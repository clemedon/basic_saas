#!/bin/bash

# Number of requests to send
num_requests=10000

# Counter for the loop
counter=0

# Infinite loop to send requests
while true; do
    # Increment counter
    ((counter++))

    # Send POST request using curl in the background
    curl -X POST 195.154.72.99/v1/run &

    # Check if we've reached the desired number of requests
    if [ $counter -eq $num_requests ]; then
        break
    fi
done

# Wait for all background processes to finish
wait

echo "All requests sent!"

