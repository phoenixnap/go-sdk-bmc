regex='const SdkVersion = "(.*)"'
contents="$(cat $1)"
if [[ $contents =~ $regex ]] 
then
    version="${BASH_REMATCH[1]}"

    # Splits the matched version by dots.
    IFS='.' read -ra VER <<< "$version"

    # If the first element is greater than 1, suffix is set to
    # /v + that element.
    suffix=""
    if [ "${VER[0]}" -gt "1" ]; then
        suffix="/v${VER[0]}"
    fi
    echo $suffix
else
    # If the regex didn't match, exit with an error.
    echo "Couldn't find a match for ($regex) in $1."
    exit 1
fi