## splunk-go search

Search Splunk for events.

### Synopsis

The search command is used to perform search queries via the Splunk REST API. 
Searching splunk using the CLI requires one argument of a SPL file containing your search.

e.g. splunk-go search ~/.splunk-go/searches/my-search.spl



```
splunk-go search [string or filepath] [flags]
```

### Options

```
  -h, --help          help for search
  -i, --interactive   Runs the search command in interactive mode.
  -m, --mode string   The output mode of the search: csv, json, json_cols, json_rows, raw, xml (default "json")
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.splunk-go.yaml)
```

### SEE ALSO

* [splunk-go](splunk-go.md)	 - A Splunk REST API client written in GO.

###### Auto generated by spf13/cobra on 19-Dec-2023
