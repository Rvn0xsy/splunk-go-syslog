## splunk-go completion

Generate shell completions

### Synopsis

To load completions:

	Bash:
	
	$ source <(splunk-golang completion bash)
	
	# To load completions for each session, execute once:
	Linux:
		$ splunk-golang completion bash > /etc/bash_completion.d/splunk-golang
	MacOS:
		$ splunk-golang completion bash > /usr/local/etc/bash_completion.d/splunk-golang
	
	Zsh:
	
	$ source <(splunk-golang completion zsh)
	
	# To load completions for each session, execute once:
	$ splunk-golang completion zsh > "${fpath[1]}/_splunk-golang"
	

```
splunk-go completion
```

### Options

```
  -h, --help   help for completion
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.splunk-go.yaml)
```

### SEE ALSO

* [splunk-go](splunk-go.md)	 - A Splunk REST API client written in GO.

###### Auto generated by spf13/cobra on 12-Jul-2020