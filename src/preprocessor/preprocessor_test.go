package preprocessor

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) {
	meta := NewMetadata("../../example/features.metadata.example.1")
	if meta == nil {
		t.Errorf("Failed to read metadata.\n")
	}

	if meta.INPUT_FILE_DIR != "../../example/example.1.input" {
		t.Errorf("Incorrect filepath.\n")
	}

	if meta.DELIMETER != " " {
		t.Errorf("Incorrect delimeter.\n")
	}

	if len(meta.FEATURE_TYPES.([]interface{})) != 8 {
		t.Errorf("Incorrect feature type vector length.\n")
	}

	if meta.HAS_FEATURE_INDEX {
		t.Errorf("Incorrect has_feature_index config.\n")
	}

	if !meta.FIXED_FEATURE_NUM {
		t.Errorf("Incorrect fixed_feature_num config.\n")
	}

	r := ReadData(meta)
	if (*r)[0] != "0 0:12 1:0.05 2:0 3:-0.15 4:32.3 5:1 6:2 7:3\n" ||
		(*r)[1] != "1 0:543 1:32.1234 2:4 3:99.99 4:0.3 5:5 6:6 7:3\n" ||
		(*r)[2] != "0 0:97 1:3.1415 2:0 3:23.90 4:1.31 5:1 6:6 7:7\n" {
		t.Errorf("Error case conversion: \n%s\n", strings.Join(*r, ""))
	}
}

func Test2(t *testing.T) {
	meta := NewMetadata("../../example/features.metadata.example.2")
	if meta == nil {
		t.Errorf("Failed to read metadata\n")
	}

	if meta.INPUT_FILE_DIR != "../../example/example.2.input" {
		t.Errorf("Incorrect filepath.\n")
	}

	if meta.DELIMETER != "|" {
		t.Errorf("Incorrect delimeter.\n")
	}

	if meta.HAS_FEATURE_INDEX {
		t.Errorf("Incorrect has_feature_index config.\n")
	}

	if meta.FIXED_FEATURE_NUM {
		t.Errorf("Incorrect fixed_feature_num config.\n")
	}

	r := ReadData(meta)
	if (*r)[0] != "0 0:5 1:0.02 2:1 3:1 4:2 5:1 6:1 7:1 8:1 9:1 10:2\n" ||
		(*r)[1] != "2 0:2 1:4.5612 4:3 7:2 11:1 12:1 13:1 14:1 15:1 16:1 17:1 18:1 19:2 20:1 21:1 22:1 23:1 24:1\n" {
		t.Errorf("Error case conversion: \n%s\n", strings.Join(*r, ""))
	}
}
