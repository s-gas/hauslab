#!/bin/bash

cpu=$(top -bn1 | grep "%Cpu(s)" | awk '{print $2}');
ram=$(free | grep "Mem:" | awk '${printf "%.2f\n", $3/$2*100}');

echo $cpu;
echo $ram;
