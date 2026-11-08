#!/bin/bash
start=$(date +%s%3N)

for i in {1..100}; do
    echo "hello deep $i" | nc localhost 8000
done

end=$(date +%s%3N)

echo "Runtime: $((end - start)) ms"