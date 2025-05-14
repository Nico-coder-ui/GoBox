#!/bin/bash

echo "Script lancé avec $# argument(s)"

for i in "$@"
do
  echo "🔹 Argument : $i"
done