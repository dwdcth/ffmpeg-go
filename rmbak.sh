
for file in $1/*; do
    if [ -f "$1_bak/$(basename $file)" ]; then
        echo "rm $1_bak/$(basename $file)"
        rm "$1_bak/$(basename $file)"
    fi
done