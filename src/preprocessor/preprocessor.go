package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

var OUTPUT_SUFFIX = "_out"

type FeatureMetadata struct {
	INPUT_FILE_DIR     string
	OUTPUT_FILE_DIR    string
	FIXED_FEATURE_NUM  bool
	FEATURE_NUM        int
	DELIMETER          string
	FEATURE_TYPES      interface{}
	LABEL_START        int
	LABEL_COLUMN_INDEX int
	HAS_FEATURE_INDEX  bool
}

func ReadMetadata(metadata_file_path string) *FeatureMetadata {
	file, err := os.Open(metadata_file_path)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	feature_meta := FeatureMetadata{}
	err = decoder.Decode(&feature_meta)
	if err != nil {
		glog.Errorf("error: %v\n", err)
	}
	return &feature_meta
}

// parseField converts raw feature to featureId:featureValue.
func parseField(index int, field string, field_type string,
	string_id_lookup *map[string]int, frequency_lookup *map[int]int,
	fixed_feature_num bool, initial_numeric_feature_number int) string {
	fmt.Println("input: index = ", index, " field = ", field, " field_type = ", field_type)
	if field_type == "int" || field_type == "float" {
		_, err := strconv.ParseFloat(field, 64)
		if err != nil {
			glog.Errorf("Error when converting a numeric feature value: %v\n", err)
			return ""
		}
		return strconv.Itoa(index) + ":" + field
	} else if field_type == "string" {
		if fixed_feature_num {
			if value, ok := (*string_id_lookup)[field]; ok {
				return strconv.Itoa(index) + ":" + strconv.Itoa(value)
			} else {
				(*string_id_lookup)[field] = len(*string_id_lookup)
				return strconv.Itoa(index) + ":" + strconv.Itoa((*string_id_lookup)[field])
			}
		} else {
			var id string
			var frequency string
			if value, ok := (*string_id_lookup)[field]; ok {
				(*frequency_lookup)[value]++
				id = strconv.Itoa(value)
				frequency = strconv.Itoa((*frequency_lookup)[value])
			} else {
				(*string_id_lookup)[field] = len(*string_id_lookup) + initial_numeric_feature_number
				(*frequency_lookup)[value] = 1
				id = strconv.Itoa(value)
				frequency = "1"
			}
			return id + ":" + frequency
		}
	} else {
		glog.Errorf("Error unknown feature type: %s\n", field_type)
		return ""
	}
	return ""
}

func ReadData(meta *FeatureMetadata) {
	fileList := []string{}
	err := filepath.Walk(meta.INPUT_FILE_DIR,
		func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return nil
		})
	if err != nil {
		panic(err)
	}

	// Generate feature types vector here.
	// Notice meta.FIXED_FEATURE_NUM can cause huge difference in result
	// feature format.
	// If the number of features are fixed, we consider every string type
	// features as ENUM, the result featureId will be the column index, the
	// value is the enum value.
	// If the number of features are not fixed, we consider:
	// 1. default type is string;
	// 2. string type features as text STRING, the result featureId will be
	// string dictionary Id, the value will be the string frequency.
	feature_types := map[string]string{}
	var initial_numeric_feature_number int
	if meta.FIXED_FEATURE_NUM {
		for i, feature_type := range meta.FEATURE_TYPES.([]interface{}) {
			feature_types[strconv.Itoa(i)] = feature_type.(string)
		}
	} else {
		feature_type_config := meta.FEATURE_TYPES.(map[string]interface{})
		feature_types["default_type"] = feature_type_config["default_type"].(string)
		for _, exception_type := range feature_type_config["exception_types"].([]interface{}) {
			feature_index := int(exception_type.(map[string]interface{})["feature_index"].(float64))
			feature_type := exception_type.(map[string]interface{})["feature_type"].(string)
			feature_types[strconv.Itoa(feature_index)] = feature_type
		}
		if feature_types["default_type"] == "string" {
			initial_numeric_feature_number = len(feature_types) - 1
		} else {
			for _, feature_type := range feature_types {
				if feature_type {
				}
			}
		}
	}
	fmt.Println("XXXX: ", feature_types)

	for _, filePath := range fileList {
		file, _ := os.Open(filePath)
		defer file.Close()

		string_type_feature_dictionary := map[string]int{}
		string_type_frequency_lookup := map[int]int{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line_str := strings.Trim(scanner.Text(), " ")
			line := strings.Split(line_str, meta.DELIMETER)
			feature_vec := []string{}

			for i, field := range line {
				// If feature number is fixed but number of
				// features read does not equal with the config
				// in metadata, throw an error.
				if meta.FIXED_FEATURE_NUM &&
					len(line) != meta.FEATURE_NUM+1 {
					glog.Errorf("Error case format: %v\n",
						line)
					return
				}

				// Deal with the label field.
				if i == meta.LABEL_COLUMN_INDEX {
					label, e := strconv.Atoi(field)
					if e != nil {
						glog.Errorf("Error when converting label: %v\n", err)
						return
					}
					// We want labels always starts from 0.
					// XXX should be leftmost!
					feature_vec = append(feature_vec,
						strconv.Itoa(label-meta.LABEL_START))
					continue
				}

				// Deal with the feature fields.
				feature_idx := i
				if feature_idx > meta.LABEL_COLUMN_INDEX {
					feature_idx--
				}
				feature_type := ""
				if _, ok := feature_types[strconv.Itoa(feature_idx)]; !ok {
					if !meta.FIXED_FEATURE_NUM {
						feature_type = feature_types["default_type"]
					}
				} else {
					feature_type = feature_types[strconv.Itoa(feature_idx)]
				}
				feature_str := parseField(feature_idx, field,
					feature_type,
					&string_type_feature_dictionary, &string_type_frequency_lookup,
					meta.FIXED_FEATURE_NUM, initial_numeric_feature_number)
				fmt.Println(feature_str)
			}
		}
	}
}

func main() {
	meta := ReadMetadata("./features.metadata.example.1")
	fmt.Println(meta)
	fmt.Println("file_path: ", meta.INPUT_FILE_DIR)
	fmt.Println("delimeter: ", meta.DELIMETER)
	fmt.Println("feature_types: ", meta.FEATURE_TYPES)
	fmt.Println("has_feature_index: ", meta.HAS_FEATURE_INDEX)
	ReadData(meta)

	meta = ReadMetadata("./features.metadata.example.2")
	fmt.Println(meta)
	fmt.Println("delimeter: ", meta.DELIMETER)
	fmt.Println("feature_types: ", meta.FEATURE_TYPES)
	fmt.Println("has_feature_index: ", meta.HAS_FEATURE_INDEX)
	ReadData(meta)

}
