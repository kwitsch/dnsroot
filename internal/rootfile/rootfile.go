package rootfile

import (
	"bufio"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	_namedRootUrl = "https://www.internic.net/domain/named.root"
)

var (
	_lastUpdateReg       = regexp.MustCompile(`^;\s+last update:\s+(\w+ \d{2}, \d{4})`)
	_zoneVersionReg      = regexp.MustCompile(`^;\s+related version of root zone:\s+(\d+)`)
	_serverFormerlyReg   = regexp.MustCompile(`^;\s+FORMERLY\s+([\w\.]+)`)
	_serverOperatedByReg = regexp.MustCompile(`^;\s+OPERATED BY\s+([\w\.,\s]+)`)
	_serverNameReg       = regexp.MustCompile(`^\.\s+\d+\s+NS\s+(\w\.ROOT-SERVERS\.NET)`)
	_serverAReg          = regexp.MustCompile(`^[A-Z]\.ROOT-SERVERS\.NET\.\s+\d+\s+A\s+(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
	_serverAAAAReg       = regexp.MustCompile(`^[A-Z]\.ROOT-SERVERS\.NET\.\s+\d+\s+AAAA\s+([\d\w:]+)`)
)

func Get() (*RootFile, error) {
	req, err := http.NewRequest("GET", _namedRootUrl, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fileScanner := bufio.NewScanner(resp.Body)
	fileScanner.Split(bufio.ScanLines)

	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return processLines(lines), nil
}

func processLines(lines []string) *RootFile {
	res := RootFile{
		Servers: make([]RootServer, 0),
	}

	var curSrv RootServer
	for _, line := range lines {
		if lastUpdate, found := getRegString(line, _lastUpdateReg); found {
			if t, err := time.Parse("January 2, 2006", lastUpdate); err == nil {
				res.LastUpdate = t.Format(time.DateOnly)
			}
		} else if version, found := getRegString(line, _zoneVersionReg); found {
			res.Version = version
		} else if formerly, found := getRegString(line, _serverFormerlyReg); found {
			res.AppendServer(&curSrv)
			curSrv = RootServer{
				Formerly: strings.TrimSpace(formerly),
			}
		} else if operatedBy, found := getRegString(line, _serverOperatedByReg); found {
			res.AppendServer(&curSrv)
			curSrv = RootServer{
				OperatedBy: strings.TrimSpace(operatedBy),
			}
		} else if name, found := getRegString(line, _serverNameReg); found {
			curSrv.Name = name
		} else if a, found := getRegString(line, _serverAReg); found {
			curSrv.A = a
		} else if aaaa, found := getRegString(line, _serverAAAAReg); found {
			curSrv.AAAA = aaaa
		}
	}
	res.AppendServer(&curSrv)

	return &res
}

func getRegString(line string, reg *regexp.Regexp) (string, bool) {
	res := reg.FindStringSubmatch(line)
	if len(res) > 1 {
		return res[1], true
	}

	return "", false
}
