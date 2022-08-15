# reverse-proxy
A simple tool created to bypass cors issues during development
## Usage

```
usage: reverse-proxy <url> <port> <print request>
       required <url>, string
       optional <port>, integer, default=1338
       optional <print request>. bool [true, false]. requires port
```
examples
```
reverse-proxy http://foxq-staging-api.cdex.club 1338 true
reverse-proxy http://foxq-staging-api.cdex.club 4000 false
reverse-proxy http://foxq-staging-api.cdex.club
```
### Build
if you wish to change the source code, building process is a simple golang command
```
go build .
```
this will output a reverse-proxy executable that you can copy to your binary folder where $PATH can find it