# Use ann and generated model@model/a6a.model.ann to test data@./data/test/a6a.t
./bin/hector-run --method ann --action test --test data/test/a6a.t --model model/a6a.model.ann

# Use svm to train data@data/train/a6a then test data@./data/test/a6a.t
./bin/hector-run --method svm --train data/train/a6a --test ./data/test/a6a.t

# Use random forest to train data@train/a1a then save the model@model/a1a.model
./bin/hector-run --method rf --action train --train data/train/a1a --model model/a1a.model

# Use logistical regression model@model/a1a.model to test data@data/test/a1a.t and save prediction@data/prediction/a1a.prediction
./bin/hector-run --method lr --action test --test data/test/a1a.t --model model/a1a.model --prediction data/prediction/a1a.prediction

# Use nerual network model@model/a1a.model to test data@data/test/a1a.t then save prediction@data/prediction/a1a.prediction and plot roc@./ROC.png
./bin/hector-run --method ann --action test --test data/test/a1a.t --model model/a1a.model --prediction data/prediction/a1a.prediction --roc_path ./ROC.png

./bin/hector-run --action preprocess --execution_plan_path ./example/features.metadata.example.train

./bin/hector-cv --method lr --train data/train/a6a --cv 10
