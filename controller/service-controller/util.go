package srvcont

import "strings"

// GetFilePath returns the file path from the request path.
// e.g. requestPath: /devingen/assets/icon
//      domain:      devingen
//      result:      assets/icon
func GetFilePath(requestPath, domain string, trim bool) string {

	if len(requestPath) <= len(domain)+1 {
		return ""
	}
	filePath := requestPath[len(domain)+2:]
	filePathLen := len(filePath)

	// remove trailing /
	if trim && filePathLen > 0 && filePath[filePathLen-1:] == "/" {
		return filePath[:filePathLen-1]
	}
	return filePath[:]
}

// GetDomainAndPath returns the domain and file path from the request path.
// e.g. requestPath: /devingen/assets/icon
//      domain:      devingen
//      path:        assets/icon
func GetDomainAndPath(requestPath string, trim bool) (string, string) {
	pathWithoutLeadingSlash := requestPath[1:]
	domain := pathWithoutLeadingSlash[:strings.Index(pathWithoutLeadingSlash, "/")]
	path := pathWithoutLeadingSlash[len(domain)+1:]

	// remove trailing /
	if trim {
		pathLen := len(path)
		if pathLen > 0 && path[pathLen-1:] == "/" {
			path = path[:pathLen-1]
		}
	}
	return domain, path
}
