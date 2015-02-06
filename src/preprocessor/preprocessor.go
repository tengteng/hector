package preprocessor

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

type FeatureMetadata struct {
	INPUT_FILE_DIR     string
	OUTPUT_FILE        string
	FIXED_FEATURE_NUM  bool
	FEATURE_NUM        int
	DELIMETER          string
	FEATURE_TYPES      interface{}
	LABEL_START        int
	LABEL_COLUMN_INDEX int
	HAS_FEATURE_INDEX  bool
}

func NewMetadata(metadata_file_path string) *FeatureMetadata {
	file, err := os.Open(metadata_file_path)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	meta := FeatureMetadata{}
	err = decoder.Decode(&meta)
	if err != nil {
		glog.Errorf("error: %v\n", err)
	}
	return &meta
}

// parseField converts raw feature to featureId:featureValue.
func parseField(index int, field string, field_type string,
	string_id_lookup *map[string]int, frequency_lookup *map[int]int,
	fixed_feature_num bool, initial_numeric_feature_number int) string {
	if field_type == "int" || field_type == "float" {
		_, err := strconv.ParseFloat(field, 64)
		if err != nil {
			glog.Errorf(
				"Error when converting a numeric feature value: %v\n", err)
			return ""
		}

		if !fixed_feature_num {
			// Do mapping, reassign featureId to numerical features.
		}

		return fmt.Sprintf("%d:%s", index, field)
	} else if field_type == "string" {
		if fixed_feature_num {
			if value, ok := (*string_id_lookup)[field]; ok {
				return fmt.Sprintf("%d:%d", index, value)
			} else {
				(*string_id_lookup)[field] =
					len(*string_id_lookup)
				return fmt.Sprintf("%d:%d", index,
					(*string_id_lookup)[field])
			}
		} else {
			// Generate frequency dictionary and return.
			if id, ok := (*string_id_lookup)[field]; ok {
				(*frequency_lookup)[id]++
			} else {
				id = len(*string_id_lookup) +
					initial_numeric_feature_number
				(*string_id_lookup)[field] = id
				(*frequency_lookup)[id] = 1
			}
			return ""
		}
	} else {
		glog.Errorf("Error unknown feature type: %s\n", field_type)
		return ""
	}
	return ""
}

func ReadData(meta *FeatureMetadata) *[]string {
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
	feature_types := map[int]string{}
	var initial_numeric_feature_number int
	if meta.FIXED_FEATURE_NUM {
		for i, feature_type := range meta.FEATURE_TYPES.([]interface{}) {
			feature_types[i] = feature_type.(string)
		}
	} else {
		feature_type_config :=
			meta.FEATURE_TYPES.(map[string]interface{})
		exception_types := feature_type_config["exception_types"].([]interface{})
		for _, exception_type := range exception_types {
			feature_index := int(exception_type.(map[string]interface{})["feature_index"].(float64))
			feature_type := exception_type.(map[string]interface{})["feature_type"].(string)
			feature_types[feature_index] = feature_type
		}
		initial_numeric_feature_number = len(feature_types)
	}

	result_date := []string{}

	for _, filePath := range fileList {
		file, _ := os.Open(filePath)
		defer file.Close()

		string_id_lookup := map[string]int{}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			frequency_lookup := map[int]int{}
			line_str := strings.Trim(scanner.Text(), " ")
			line := strings.Split(line_str, meta.DELIMETER)
			feature_vec := []string{}
			case_label := ""
			current_feature_id := -1

			for i, field := range line {
				// If feature number is fixed but number of
				// features read does not equal with the config
				// in metadata, throw an error.
				if meta.FIXED_FEATURE_NUM &&
					len(line) != meta.FEATURE_NUM+1 {
					glog.Errorf("Error case format: %v\n",
						line)
					return nil
				}

				// Deal with the label field.
				if i == meta.LABEL_COLUMN_INDEX {
					label, e := strconv.Atoi(field)
					if e != nil {
						glog.Errorf(
							"Error when converting label: %v\n", err)
						return nil
					}
					// We want labels always starts from 0.
					case_label = strconv.Itoa(label -
						meta.LABEL_START)
					continue
				}

				// Deal with the feature fields.
				feature_idx := i
				if feature_idx > meta.LABEL_COLUMN_INDEX {
					feature_idx--
				}
				feature_type := ""
				if _, ok := feature_types[feature_idx]; !ok {
					if !meta.FIXED_FEATURE_NUM {
						feature_type = "string"
					}
				} else {
					feature_type =
						feature_types[feature_idx]
					current_feature_id++
				}
				feature_str := parseField(current_feature_id,
					field, feature_type, &string_id_lookup,
					&frequency_lookup,
					meta.FIXED_FEATURE_NUM,
					initial_numeric_feature_number)
				if feature_str != "" {
					feature_vec = append(feature_vec,
						feature_str)
				}
			}
			// Append frequency_lookup str to feature_str
			if len(frequency_lookup) != 0 {
				keys := make([]int, len(frequency_lookup))
				i := 0
				for idx := range frequency_lookup {
					keys[i] = idx
					i++
				}
				sort.Ints(keys)
				for _, idx := range keys {
					feature_vec = append(feature_vec,
						fmt.Sprintf("%d:%d", idx,
							frequency_lookup[idx]))
				}
			}
			result_date = append(result_date,
				fmt.Sprintf("%s %s\n", case_label,
					strings.Join(feature_vec, " ")))
		}
	}

	// Write to disk if output exists.
	if meta.OUTPUT_FILE != "" {
		output_file, err := os.Create(meta.OUTPUT_FILE)
		if err != nil {
			panic(err)
		}
		defer output_file.Close()

		for _, line := range result_date {
			output_file.WriteString(line)
		}
	}
	return &result_date
}
