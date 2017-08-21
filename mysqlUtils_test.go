package mysqlgoutils

import (
        "testing"
)

func TestSplitHostOptionalPortAndSchema(t *testing.T) {
        assertSplit := func(addr, expectHost string, expectPort int, expectSchema string, expectErr bool) {
                host, port, schema, err := SplitHostOptionalPortAndSchema(addr)
                if host != expectHost {
                        t.Errorf("Expected SplitHostOptionalPortAndSchema(\"%s\") to return host of \"%s\", instead found \"%s\"", addr, expectHost, host)
                }
                if port != expectPort {
                        t.Errorf("Expected SplitHostOptionalPortAndSchema(\"%s\") to return port of %d, instead found %d", addr, expectPort, port)
                }
                if schema != expectSchema {
                        t.Errorf("Expected SplitHostOptionalPortAndSchema(\"%s\") to return schema of %d, instead found \"%s\"", addr, expectSchema, schema)
                }
                if expectErr && err == nil {
                        t.Errorf("Expected SplitHostOptionalPortAndSchema(\"%s\") to return an error, but instead found nil", addr)
                } else if !expectErr && err != nil {
                        t.Errorf("Expected SplitHostOptionalPortAndSchema(\"%s\") to return NOT return an error, but instead found %s", addr, err)
                }
        }

        assertSplit("", "", 0, "", true)
        assertSplit("foo", "foo", 0, "", false)
        assertSplit("1.2.3.4", "1.2.3.4", 0, "", false)
        assertSplit("some.host:1234", "some.host", 1234, "", false)
        assertSplit("some.host:text", "", 0, "", true)
        assertSplit("some.host:1234:5678", "", 0, "", true)
        assertSplit("some.host:0", "", 0, "", true)
        assertSplit("some.host:-5", "", 0, "", true)
        assertSplit("fe80::1", "", 0, "", true)
        assertSplit("[fe80::1]", "[fe80::1]", 0, "", false)
        assertSplit("[fe80::1]:3306", "[fe80::1]", 3306, "", false)
        assertSplit("[fe80::bd0f:a8bc:6480:238b%11]", "[fe80::bd0f:a8bc:6480:238b%11]", 0, "", false)
        assertSplit("[fe80::bd0f:a8bc:6480:238b%11]:443", "[fe80::bd0f:a8bc:6480:238b%11]", 443, "", false)
        assertSplit("[fe80::bd0f:a8bc:6480:238b%11]:sup", "", 0, "", true)
        assertSplit("[fe80::bd0f:a8bc:6480:238b%11]:123:456", "", 0, "", true)

	    assertSplit("|dbtest", "", 0, "", true)
        assertSplit("1.2.3.4|", "", 0, "", true)
        assertSplit("1.2.3.4|dbtest", "1.2.3.4", 0, "dbtest", false)
        assertSplit("1.2.3.4:1234|dbtest", "1.2.3.4", 1234, "dbtest", false)
        assertSplit("1.2.3.4:1234|dbtest|foo", "", 0, "", true)
        assertSplit("some.host", "some.host", 0, "", false)
        assertSplit("some.host|dbtest", "some.host", 0, "dbtest", false)
        assertSplit("some.host:1234|dbtest", "some.host", 1234, "dbtest", false)

}
