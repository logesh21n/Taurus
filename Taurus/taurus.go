package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running command: %s, output: %s", err, output)
	}
	fmt.Printf("Output:\n%s\n", output)
	return nil
}

func main() {
	domainsFile := flag.String("d", "domains.txt", "File containing the list of domains")
	help := flag.Bool("h", false, "Show help")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	commands := []string{
		fmt.Sprintf("subfinder -dL %s -all -recursive > subdomains1.txt", *domainsFile),


		fmt.Sprintf("interlace -tL %s -c \"curl -s 'https://rapiddns.io/subdomain/_target_?full=1'\" | grep -Eo '([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,}' | sort -u >> subdomains2.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL %s -c \"curl -s 'https://rapiddns.io/subdomain/_target_?full=2'\" | grep -Eo '([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,}' | sort -u >> subdomains2.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL %s -c \"curl -s 'https://rapiddns.io/subdomain/_target_?full=3'\" | grep -Eo '([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,}' | sort -u >> subdomains2.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL %s -c \"sublist3r -d '_target_'\" | grep -Eo '([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,}' | sed 's/^92m//' | sort -u >> subdomains3.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL %s -c \"true | openssl s_client -connect _target_:443 2>/dev/null | openssl x509 -noout -text | grep 'DNS' | tr ',' '\\n' | cut -d ':' -f2\" | grep -Eo '([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,}' | sort -u > subdomains5.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL %s -c \"echo '_target_' | gau --subs | unfurl -u domains\" | grep -Eo '([a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,}' | sort -u > subdomains6.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL %s -c \"curl -I -s '_target_' | grep -iE 'content-security-policy|CSP' | tr ' ' '\\n' | grep '\\.' | tr -d ';' | sed 's/\\*\\.//g' | sort -u\" > subdomains9.txt", *domainsFile),
		
		
		fmt.Sprintf("cat subdomains1.txt subdomains2.txt subdomains3.txt subdomains4.txt subdomains5.txt subdomains6.txt subdomains7.txt subdomains8.txt subdomains9.txt | sort -u > sortedSubdomains.txt", *domainsFile),
		
		
		fmt.Sprintf("cat domains.txt | sed 's/\\(\\.[^.]*\\)*$//' > firstWords.txt", *domainsFile),
		
		
		fmt.Sprintf("grep -F -f firstWords.txt sortedSubdomains.txt >> correctedSubdomains.txt", *domainsFile),
		
		
		fmt.Sprintf("cat correctedSubdomains.txt | httpx -sc -cl --title -o aliveSubdomain.txt", *domainsFile),
		
		
		fmt.Sprintf("cat correctedSubdomains.txt | gau --threads 10 --o gauUrls.txt", *domainsFile),
		
		
		fmt.Sprintf("cat correctedSubdomains.txt | waybackurls > waybackurls.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL correctedSubdomains.txt -c \"curl -s 'http://web.archive.org/cdx/search/cdx?url=_target_/*&output=text&fl=original&collapse=urlkey'\" | sort -u | uro > allUrls.txt", *domainsFile),
		
		
		fmt.Sprintf("cat allUrls.txt gauUrls.txt waybackurls.txt | grep '.js$' > jsFiles.txt", *domainsFile),
		
		
		fmt.Sprintf("cat jsFiles.txt | httpx -threads 20 > aliveJs.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL aliveJs.txt -c \"curl -s '_target_' | grep -Po '((http|https):\\/\\/)?(([\\w\\.-]*)\\.([\\w]*)\\.([A-z]))\\w+' | sort -u\" >> jsSubdomains.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL aliveJs.txt -c \"curl -s '_target_' | sed -nE 's/.*((http|https):\\/\\/)?(([\\w.-]*)\\.([\\w]*)\\.([A-Za-z]))\\w+.*/\\1/p'\" | sort -u >> jsSubdomains.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL aliveJs.txt -c \"curl -s -I '_target_' | grep -E '(?i)((access_key|access_token|admin_pass|admin_user|algolia_admin_key|algolia_api_key|alias_pass|alicloud_access_key|amazon_secret_access_key|amazonaws|ansible_vault_password|aos_key|api_key|api_key_secret|api_key_sid|api_secret|api.googlemaps AIza|apidocs|apikey|apiSecret|app_debug|app_id|app_key|app_log_level|app_secret|appkey|appkeysecret|application_key|appsecret|appspot|auth_token|authorizationToken|authsecret|aws_access|aws_access_key_id|aws_bucket|aws_key|aws_secret|aws_secret_key|aws_token|AWSSecretKey|b2_app_key|bashrc password|bintray_apikey|bintray_gpg_password|bintray_key|bintraykey|bluemix_api_key|bluemix_pass|browserstack_access_key|bucket_password|bucketeer_aws_access_key_id|bucketeer_aws_secret_access_key|built_branch_deploy_key|bx_password|cache_driver|cache_s3_secret_key|cattle_access_key|cattle_secret_key|certificate_password|ci_deploy_password|client_secret|client_zpk_secret_key|clojars_password|cloud_api_key|cloud_watch_aws_access_key|cloudant_password|cloudflare_api_key|cloudflare_auth_key|cloudinary_api_secret|cloudinary_name|codecov_token|config|conn.login|connectionstring|consumer_key|consumer_secret|credentials|cypress_record_key|database_password|database_schema_test|datadog_api_key|datadog_app_key|db_password|db_server|db_username|dbpasswd|dbpassword|dbuser|deploy_password|digitalocean_ssh_key_body|digitalocean_ssh_key_ids|docker_hub_password|docker_key|docker_pass|docker_passwd|docker_password|dockerhub_password|dockerhubpassword|dot-files|dotfiles|droplet_travis_password|dynamoaccesskeyid|dynamosecretaccesskey|elastica_host|elastica_port|elasticsearch_password|encryption_key|encryption_password|env.heroku_api_key|env.sonatype_password|eureka.awssecretkey)[a-z0-9_ .\\-,]{0,25})(=|>|:=|\\|\\|:|<=|=>|:).{0,5}[\\'\\\"]([0-9a-zA-Z\\-_=]{8,64})[\\'\\\"]\" > secrets.txt", *domainsFile),
		
		
		fmt.Sprintf("interlace -tL correctedSubdomains.txt -c \"dig _target_ CNAME\" > dig_CNAME.txt", *domainsFile),
	}

	for _, command := range commands {
		fmt.Printf("Running command: %s\n", command)
		if err := runCommand(command); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to continue with the Nmap scan? (yes or no): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response) 

	if strings.ToLower(response) == "yes" {
		nmapCommand := fmt.Sprintf("interlace -tL correctedSubdomains.txt -c \"nmap -sC -sV '_target_' -T4\" > nmap")
		fmt.Printf("Running Nmap command: %s\n", nmapCommand)
		if err := runCommand(nmapCommand); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Nmap scan skipped.")
	}

	fmt.Println("All commands executed successfully!")
}

