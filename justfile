# Run check-file against /opt/file and /opt/sub/file.
test:
    #!/usr/bin/env bash
    if just check /opt/file; then
        echo 'OK: Found file at /opt/file'
    else
        echo '{{ style('error') }}ERROR:{{ NORMAL }} Did not find file at /opt/file'
    fi
    if just check /opt/sub/file; then
        echo '{{ style('error') }}ERROR:{{ NORMAL }} Found file at /opt/sub/file'
    else
        echo 'OK: No file found at /opt/sub/file'
    fi

check file:
    ${DAGGER:-dagger} --progress plain call -m . check-file --file {{ quote(file) }}
