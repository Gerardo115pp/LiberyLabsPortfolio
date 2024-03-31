#! /usr/bin/bash

function link_dev_wp_files() {
    # wp_dev_path could be relative
    local wp_dev_path=$1
    local app_path='/var/www/libery-labs.com'
    local user_name=$(whoami)
    local wp_path="$app_path/services/wp_od/wp-content"

    if [ ! -d $wp_dev_path ]; then
        echo "Directory not found: $wp_dev_path"
        return 1
    fi

    if [ ! -d $wp_path ]; then
        echo "Directory not found: $wp_path\nDid the local app directory changed?"
        return 1
    fi

    # Link the theme

    local theme_name='libery_labs_theme'
    local theme_dir=$(realpath "$wp_dev_path/$theme_name")

    echo "Linking theme: $theme_dir" 

    if [ ! -d $theme_dir ]; then
        echo "The theme directory was not found: $theme_dir"
        return 1
    fi

    local theme_mount_target="$wp_path/themes/$theme_name"

    if [ ! -d $theme_mount_target ]; then
        echo "Linking theme to $theme_mount_target"
        sudo ln -s $theme_dir $theme_mount_target
        sudo chown -h $user_name:$user_name $theme_mount_target
    fi

    if [ -f "$app_path/services/docker-compose.yaml" ] || [ -f "$app_path/services/docker-compose.yml" ]; then
        # we check if the user(probably meaning me) has mounted the directory with the theme and the plugins in the docker container, because if he didn't, then the symblinks
        # will be broken inside the container.
        local wp_src_path=$(realpath "$wp_dev_path")
        local line_exists=false
        for posible_docker_file in \
            "$app_path/services/docker-compose.yaml",\
            "$app_path/services/docker-compose.yml" \
        ; do
            if [ -f $posible_docker_file ]; then
                if grep -q "$wp_src_path:$wp_src_path" $posible_docker_file; then
                    line_exists=true
                    break
                fi
            fi
        done
        
        if [ ! $line_exists ]; then
            echo "WARNING: Can't find a line in the docker-compose file that mounts the wp source directory. if indeed this was missed, the symlinks will be broken inside the container."
        else 
            echo "Found a line in the docker-compose file that mounts the wp source directory. GOOD!"
        fi

    fi

    local plugins_dir="$wp_dev_path/wp_plugins"

    # Link the plugins
    if [ -d $plugins_dir ]; then
        local plugins_dir=$(realpath $plugins_dir)
        for plugin_dir in "$plugins_dir"/*; do
            if [ -d $plugin_dir ]; then
                local plugin_name=$(basename $plugin_dir)
                local plugin_mount_target="$wp_path/plugins/$plugin_name"

                if [ ! -d $plugin_mount_target ]; then
                    echo "Linking plugin: $plugin_dir to $plugin_mount_target"
                    sudo ln -s $plugin_dir $plugin_mount_target
                    sudo chown -h $user_name:$user_name $plugin_mount_target
                fi
            fi
        done
    fi 

    echo 'All operations completed successfully'

    return 0
}