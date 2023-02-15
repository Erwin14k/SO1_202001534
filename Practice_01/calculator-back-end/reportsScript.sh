#!/usr/bin/env bash

# Get today's date in the format used in logs.txt
today_date=$(date "+%Y-%m-%d")

# Count the total number of lines in the file
total_lines=$(wc -l logs.txt | cut -d' ' -f1)

# Count the number of occurrences of -1499 in the file
total_errors=$(grep -c -- "-,-1499;" logs.txt)

# Count the number of occurrences of + in the file
total_sums=$(grep -c -- "+," logs.txt)

# Count the number of occurrences of - in the file
total_subtractions=$(grep -c -- "-," logs.txt)

# Count the number of occurrences of * in the file
total_multiplications=$(grep -c -- "\*," logs.txt)

# Count the number of occurrences of / in the file
total_divisions=$(grep -c -- "/," logs.txt)

# Print the results
echo "Total logs registered: $total_lines"
echo "Total Errors: $total_errors"
echo "Total Sums: $total_sums"
echo "Total Subtractions: $total_subtractions"
echo "Total Multiplications: $total_multiplications"
echo "Total Divisions: $total_divisions"
echo "Total logs registered today: $total_today"
