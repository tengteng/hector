# Use ann and generated model@model/a6a.model.ann to test data@./data/test/a6a.t
./bin/hector-run --method ann --action test --test ./data/test/a6a.t --model model/a6a.model.ann

# Use svm to train data@data/train/a6a then test data@./data/test/a6a.t
./bin/hector-run --method svm --train data/train/a6a --test ./data/test/a6a.t

# Use random forest to train data@train/a1a then save the model@model/a1a.model
./bin/hector-run --method rf --action train --train train/a1a --model model/a1a.model
