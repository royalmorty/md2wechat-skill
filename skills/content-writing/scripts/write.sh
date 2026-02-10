#!/bin/bash

# Content Writing Script for Claude Code Skill
# This script provides content writing functionality for WeChat articles

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
APP_DIR="$(cd "$SCRIPT_DIR/../../../app" && pwd)"
WRITER_BINARY="$(cd "$SCRIPT_DIR/../../../scripts" && pwd)/wechatwriter"

# Function to handle write command
write_content() {
    local style="${1:-dan-koe}"
    local input_file="$2"
    local output_file="$3"
    local additional_args=""
    
    if [[ -n "$input_file" ]]; then
        additional_args="--input-type fragment $input_file"
    fi
    
    if [[ -n "$output_file" ]]; then
        additional_args="$additional_args -o $output_file"
    fi
    
    "$WRITER_BINARY" write --style "$style" $additional_args
}

# Function to handle convert command
convert_content() {
    local input_file="$1"
    local theme="${2:-default}"
    local output_file="$3"
    local additional_args=""
    
    if [[ -n "$output_file" ]]; then
        additional_args="-o $output_file"
    fi
    
    "$WRITER_BINARY" convert "$input_file" --theme "$theme" $additional_args
}

# Function to handle humanize command
humanize_content() {
    local input_file="$1"
    local intensity="${2:-medium}"
    local output_file="$3"
    local additional_args=""
    
    if [[ -n "$output_file" ]]; then
        additional_args="-o $output_file"
    fi
    
    "$WRITER_BINARY" humanize "$input_file" --intensity "$intensity" $additional_args
}

# Function to handle outline command
outline_content() {
    local template="${1:-viral}"
    local topic="$2"
    local domain="${3:-general}"
    local output_file="$4"
    local additional_args=""
    
    if [[ -n "$output_file" ]]; then
        additional_args="-o $output_file"
    fi
    
    "$WRITER_BINARY" outline --template "$template" --topic "$topic" --domain "$domain" $additional_args
}

# Main function
main() {
    if [[ $# -eq 0 ]]; then
        "$WRITER_BINARY" content-writing --help
        exit 0
    fi

    local command="$1"
    shift

    case "$command" in
        write)
            local style="dan-koe"
            local input_file=""
            local output_file=""
            
            while [[ $# -gt 0 ]]; do
                case $1 in
                    --style)
                        style="$2"
                        shift 2
                        ;;
                    --input)
                        input_file="$2"
                        shift 2
                        ;;
                    --output)
                        output_file="$2"
                        shift 2
                        ;;
                    *)
                        echo "Unknown option: $1"
                        exit 1
                        ;;
                esac
            done
            
            write_content "$style" "$input_file" "$output_file"
            ;;
            
        convert)
            local input_file=""
            local theme="default"
            local output_file=""
            
            while [[ $# -gt 0 ]]; do
                case $1 in
                    --input)
                        input_file="$2"
                        shift 2
                        ;;
                    --theme)
                        theme="$2"
                        shift 2
                        ;;
                    --output)
                        output_file="$2"
                        shift 2
                        ;;
                    *)
                        echo "Unknown option: $1"
                        exit 1
                        ;;
                esac
            done
            
            if [[ -z "$input_file" ]]; then
                echo "Error: --input is required for convert command"
                exit 1
            fi
            
            convert_content "$input_file" "$theme" "$output_file"
            ;;
            
        humanize)
            local input_file=""
            local intensity="medium"
            local output_file=""
            
            while [[ $# -gt 0 ]]; do
                case $1 in
                    --input)
                        input_file="$2"
                        shift 2
                        ;;
                    --intensity)
                        intensity="$2"
                        shift 2
                        ;;
                    --output)
                        output_file="$2"
                        shift 2
                        ;;
                    *)
                        echo "Unknown option: $1"
                        exit 1
                        ;;
                esac
            done
            
            if [[ -z "$input_file" ]]; then
                echo "Error: --input is required for humanize command"
                exit 1
            fi
            
            humanize_content "$input_file" "$intensity" "$output_file"
            ;;
            
        outline)
            local template="viral"
            local topic=""
            local domain="general"
            local output_file=""
            
            while [[ $# -gt 0 ]]; do
                case $1 in
                    --template)
                        template="$2"
                        shift 2
                        ;;
                    --topic)
                        topic="$2"
                        shift 2
                        ;;
                    --domain)
                        domain="$2"
                        shift 2
                        ;;
                    --output)
                        output_file="$2"
                        shift 2
                        ;;
                    *)
                        echo "Unknown option: $1"
                        exit 1
                        ;;
                esac
            done
            
            if [[ -z "$topic" ]]; then
                echo "Error: --topic is required for outline command"
                exit 1
            fi
            
            outline_content "$template" "$topic" "$domain" "$output_file"
            ;;
            
        help)
            "$WRITER_BINARY" content-writing --help
            ;;
            
        *)
            echo "Unknown command: $command"
            "$WRITER_BINARY" content-writing --help
            exit 1
            ;;
    esac
}

# Execute main function
main "$@"