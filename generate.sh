#!/bin/sh

for i in $(seq 100); do
    name="proj${i}"
    echo "Generating ${name}..."
    rm -rf "${name}"
    cp -r template "${name}"
    sed -i'.bak' -e "s/template/${name}/" "${name}/go.mod"
    rm "${name}/go.mod.bak"
done
