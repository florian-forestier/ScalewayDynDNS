# Online/Scaleway dynamic DNS

This little project offer tools to help people who want to set a DNS dynamically. For example, for people who are self-hosting their stuff at home.

## How to use
1. Build the go program : `go build -o dyndns-scaleway ./cmd`, or take it from Release page.
2. Start the program with `domain` and `key` parameter. 
   * Example : `./dyndns-scaleway -domain "my.domain.that.i.want" -key "OnlineToken"`.
   * You can use `SCALEWAY_DYNDNS_DOMAIN` and `SCALEWAY_API_KEY` environment variables to replace command-line arguments.
   * You can specify multiple domains, BUT they must belong to the same user (because you can define only one API key). Simply separate each domain by a semicolon.
     * Example : `SCALEWAY_DYNDNS_DOMAIN="my_subdomain.artheriom.fr;my_other_subdomain.artheriom.fr" ./dyndns-scaleway`
3. Use a crontab to call automatically the program on a regular basis

## Licence

Script submitted under the Do What The Fuck You Want to the Public Licence. Do what you want with it, it was just a short development for a friend. It comes with absolutely NO WARRANTY.