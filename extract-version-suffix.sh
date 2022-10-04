regex='const SdkVersion = "(.*)"'
contents="$(cat $1)"
if [[ $contents =~ $regex ]] 
then
    version="${BASH_REMATCH[1]}"
    IFS='.' read -ra VER <<< "$version"

    suffix=""
    if [ "${VER[0]}" -gt "1" ]; then
        suffix="/v${VER[0]}"
    fi
    echo $suffix
else
    echo "Couldn't find a match for ($regex) in $1."
    exit 1
fi