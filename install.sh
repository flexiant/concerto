#!/bin/bash


cli_url=http://get.concerto.io/concerto.x64
cli_command=concerto
cli_fullpath=/usr/local/bin/$cli_command
conf_path=$HOME/.concerto
cli_conf=$conf_path/client.xml
cli_conf_exists=false
cli_fullpath_exists=false
cacert_exists=false
cert_exists=false
key_exists=false
update=false
verbose=false
LOGO_H_SIZE=127

parseArgs(){


	printf "Parse arguments ..."
	for arg in "$@"
	do
    [ "$arg" = "-u" ] && update=true
		[ "$arg" = "-v" ] && verbose=true
	done
	printf " OK\n"
}

concertoInitialize(){
	printf "Initializing ..."
	case "$(uname -m)" in
		*64)
			;;
		*)
			echo >&2 -e "Concerto CLI is available for 64 bit systems only\n"
			exit 1
			;;
	esac

	case "$(uname -s)" in
		Darwin)
			cli_url="$cli_url.darwin"
			;;
		Linux)
			cli_url="$cli_url.linux"
			;;
		*)
			$verbose && printf " (OS could not be detected. Assuming linux) "
			cli_url="$cli_url.linux"
			;;
	esac

	printf " OK\n"

	getInstallationState
}

getInstallationState(){
	[ -f $cli_conf ] && cli_conf_exists=true || cli_conf_exists=false
	[ -f $cli_fullpath ] && cli_fullpath_exists=true || cli_fullpath_exists=false
	[ -f $conf_path/ssl/ca_cert.pem ] && cacert_exists=true || cacert_exists=false
	[ -f $conf_path/ssl/cert.crt ] && cert_exists=true || cert_exists=false
	[ -f $conf_path/ssl/private/cert.key ] && key_exists=true || key_exists=false
}

writeDefaultConfig(){

	printf "Writing concerto configuration ..."
	if $cli_conf_exists;
	then
		$verbose && printf " (configuration found at '$cli_conf')."
		printf " Skipped\n"
		return
	fi

	mkdir -p "${conf_path}"
	cat <<EOF > $cli_conf
<concerto version="1.0" server="https://clients.concerto.io:886/" log_file="/var/log/concerto-client.log" log_level="info">
	<ssl cert="$cli_conf/ssl/cert.crt" key="$cli_conf/ssl/private/cert.key" server_ca="$cli_conf/ssl/ca_cert.pem" />
</concerto>
EOF

	printf " OK\n"
}

installConcertoCLI(){
	printf "Installing Concerto CLI ..."
	if ! $update && $cli_fullpath_exists;
	then
		$verbose && printf " (concerto CLI exists. Use '-u' to force update)"
		printf " Skipped\n"
		return
	fi

	command -v curl > /dev/null &&  dwld="curl -sSL -o" || \
	{	command -v wget > /dev/null && dwld="wget -qO"; } || \
	{ echo ' (curl or wget are needed to install Concerto CLI.) Failed'; exit 1; }
	printf " (you might be asked for your password to sudo now)\n"
	if ! sudo $dwld  $cli_fullpath $cli_url;
	then
		echo "(Concerto CLI Binary download failed). Failed"
		exit 1
	fi

	if ! sudo chmod +x $cli_fullpath;
	then
		echo "(Concerto CLI Binary execution flag assigment failed). Failed"
		exit 1
	fi

	echo "Binary has been installed. OK"

	current_concerto=$(command -v $cli_fullpath)
	[ $current_concerto != $cli_fullpath ] && echo "WARNING: concerto is being run from '$current_concerto'. Please, update your path to execute from $cli_fullpath"

}

installAPIKeys(){
	printf "Installing API keys ..."

	if $cacert_exists && $cert_exists && $key_exists;
	then
	 	$verbose && printf " (Concerto keys already exists)."
		printf " Skipped\n"
	else
		if ! concerto setup api_keys;
		then
			printf " (error downloading Concerto keys. Try downloading manually). Failed\n"
			certsInstructions
			exit 1
		fi
	fi

	# # if certs not there
	#  ! $cacert_exists || ! $cert_exists || ! $key_exists ] && ! $cert_exists && concerto setup api_keys
	#  getInstallationState
	#  ! $cacert_exists || ! $cert_exists || ! $cert_exists ] && ! $cert_exists && certsInstructions || echo "Concerto installed. Type 'concerto' to access CLI help"

}

certsInstructions(){
cat <<EOF
Concerto CLI uses an API Key that you can download from Concerto's Web through 'Settings' > 'User Detail' > 'New API Key'
Uncompress the downloaded file and copy as follows:
$cli_conf
└── ssl
    ├── ca_cert.pem
    ├── cert.crt
    └── private
        └── cert.key
EOF
}

installedMessage(){
	printf "\n Concerto CLI is installed.\n Type 'concerto' to access Concerto commands\n\n"
}
showLogo(){
	[ $LOGO_H_SIZE -lt $(tput cols) ] && logoFull || logoSimple
	echo "Executing Flexiant Concerto CLI install"
}

logoSimple(){
cat <<EOF

            ╔                       
            ╠▒╕                     
            ╠╢╢▒                    
            ╠╢╢╢▒╕               ╓⌐ 
            ╠╢╢╢╢╢▒        ╓╗╪╬▒Θ   
            ╠╢╢╢╢╢╢▒╦╓╤╗╬▒╢╢╢╢╢Å    
            ╠╢╢╢╢╢╢╢╢╢╬ÅÅÅ╫╢╢╢┘  ,,,
       ,╓╗@▒▒╢╢╢╢╢╢╢╢╢╢╤   ▒▒   ╬╢╢┘
,╓╗╗╬▒╬╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢▒╕    ╓▒╢Θ    
'▀Å▒╫╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢Å    ╙▒╢▒  
      '"▀Å▒╢╢╢╢╢╢╢╢╢╢╢╢▒   ╬▒   Å╢▒╦
            ╠╢╢╢╢╢╢╢╢╢╬╤╤╤╬╢╢▒╕   """
            ╠╢╢╢╢╢╢╬Ö╙ÅΘ╫╢╢╢╢╢╢▒     
            ╠╢╢╢╢╢Θ        "▀Å▒╢▒╕   
            ╠╢╢╢╢▀               '"  
            ╠╢╢▒                     
            ╠╢▀    Flexiant Concerto       
            ╠      https://start.concerto.io                   

EOF
}

logoFull(){
cat <<EOF

            ╔                        88888888888 88                        88 
            ╠▒╕                      88          88                        ""                         ,d
            ╠╢╢▒                     88          88                                                   88
            ╠╢╢╢▒╕               ╓⌐  88aaaaa     88  ,adPPYba, 8b,     ,d8 88 ,adPPYYba, 8b,dPPYba, MM88MMM
            ╠╢╢╢╢╢▒        ╓╗╪╬▒Θ    88"""""     88 a8P_____88  'Y8, ,8P'  88 ""     'Y8 88P'   '"8a  88
            ╠╢╢╢╢╢╢▒╦╓╤╗╬▒╢╢╢╢╢Å     88          88 8PP"""""""    )888(    88 ,adPPPPP88 88       88  88 
            ╠╢╢╢╢╢╢╢╢╢╬ÅÅÅ╫╢╢╢┘  ,,, 88          88 "8b,   ,aa  ,d8" "8b,  88 88,    ,88 88       88  88, 
       ,╓╗@▒▒╢╢╢╢╢╢╢╢╢╢╤   ▒▒   ╬╢╢┘ 88          88  '"Ybbd8"' 8P'     'Y8 88 '"8bbdP"Y8 88       88  "Y888 
,╓╗╗╬▒╬╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢▒╕    ╓▒╢Θ    
'▀Å▒╫╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢╢Å    ╙▒╢▒      ,ad8888ba, 
      '"▀Å▒╢╢╢╢╢╢╢╢╢╢╢╢▒   ╬▒   Å╢▒╦   d8"'    '"8b                                                            ,d
            ╠╢╢╢╢╢╢╢╢╢╬╤╤╤╬╢╢▒╕   """ d8'                                                                      88
            ╠╢╢╢╢╢╢╬Ö╙ÅΘ╫╢╢╢╢╢╢▒      88             ,adPPYba,  8b,dPPYba,   ,adPPYba,  ,adPPYba, 8b,dPPYba, MM88MMM ,adPPYba,
            ╠╢╢╢╢╢Θ        "▀Å▒╢▒╕    88            a8"     "8a 88P'   '"8a a8"     "" a8P_____88 88P'   "Y8   88   a8"     "8a
            ╠╢╢╢╢▀               '"   Y8,           8b       d8 88       88 8b         8PP""""""" 88           88   8b       d8
            ╠╢╢▒                       Y8a.    .a8P "8a,   ,a8" 88       88 "8a,   ,aa "8b,   ,aa 88           88,  "8a,   ,a8"
            ╠╢▀                         '"Y8888Y"'   '"YbbdP"'  88       88  '"Ybbd8"'  '"Ybbd8"' 88           "Y888 '"YbbdP"'
            ╠                         

EOF
}






showLogo
parseArgs $@
concertoInitialize
installConcertoCLI
writeDefaultConfig
installAPIKeys
installedMessage
