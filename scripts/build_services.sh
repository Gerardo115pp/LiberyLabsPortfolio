#! /usr/bin/bash 

function build_wp_database() {
    local dockerfile_path=$1
    local docker_image_tag="defalt115/libery-labs-wp-database:$2"

    echo "Building $docker_image_tag from $dockerfile"

    if [ ! -d $dockerfile_path ]; then
        echo "Directory not found: $dockerfile"
        return 1
    fi

    if [ ! -f "$dockerfile_path/Dockerfile" ]; then
        echo "Dockerfile not found in $dockerfile_path"
        return 1
    fi

    cd $dockerfile_path

    docker buildx build -t $docker_image_tag . --load 
    
    if [ $? -ne 0 ]; then
        echo "Failed to build $docker_image_tag"
        return 1
    fi

    popd

    echo "Successfully built $docker_image_tag"
    return 0
}