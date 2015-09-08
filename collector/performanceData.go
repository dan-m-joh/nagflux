package collector

import (
	"fmt"
	"strings"
	"github.com/griesbacher/nagflux/helper"
)

type PerformanceData struct {
	hostname         string
	service          string
	command          string
	performanceLabel string
	performanceType  string
	unit             string
	time             string
	value            string
	fieldseperator   string
	tags             []string
}

func (p *PerformanceData) String() string {
	p.tags = helper.RemoveDuplicateStrings(p.tags)
	tableName := fmt.Sprintf(`%s%s%s%s%s%s%s%s%s`,
		p.hostname, p.fieldseperator,
		p.service, p.fieldseperator,
		p.command, p.fieldseperator,
		p.performanceLabel, p.fieldseperator,
		p.performanceType)
	if p.unit != "" {
		tableName += fmt.Sprintf(`,unit=%s`, p.unit)
	}

	if len(p.tags) > 0 {
		tableName += fmt.Sprintf(`,%s`, strings.Replace(strings.Trim(fmt.Sprintf("%s",p.tags),"[]")," ", ",", -1))
	}

	tableName += fmt.Sprintf(` value=%s %s`, p.value, p.time)
	return tableName
}