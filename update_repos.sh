#!/bin/bash

# Khai báo giá trị mới cho REPOSITORY và PULL_MERGE_BRANCH

# Đường dẫn đến tệp repository.conf
CONFIG_FILE="./config/repositories.conf"

if [ -n "$REPOSITORY"  ] && [ -n "$PULL_MERGE_BRANCH" ]; then 
    echo "Exist to update"
     # Thay đổi giá trị trong tệp repository.conf
    sed -i "s/$REPOSITORY=\(.*\)/${REPOSITORY}=$PULL_MERGE_BRANCH/" "$CONFIG_FILE"
    echo "Updated $REPOSITORY in $CONFIG_FILE with value $PULL_MERGE_BRANCH"
fi



# sed -i 's/microservice-golang=\(.*\)/microservice-golang=test/' repositories.conf

# # Kiểm tra xem tệp tồn tại trước khi thay đổi
# if [ -f "$CONFIG_FILE" ]; then
#     # Thay đổi giá trị trong tệp repository.conf
#     sed -i "s/$REPOSITORY=\(.*\)/${REPOSITORY}=$PULL_MERGE_BRANCH/" "$CONFIG_FILE"
#     echo "Updated $REPOSITORY in $CONFIG_FILE with value $PULL_MERGE_BRANCH"
# else
#     echo "$CONFIG_FILE does not exist."
# fi