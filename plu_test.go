package plu

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type pluSuite struct {
	c Codes
}

var _ = Suite(&pluSuite{})

func (p *pluSuite) SetUpSuite(c *C) {
	var err error
	p.c, err = New()
	c.Assert(err, IsNil)
}

func (p *pluSuite) TestPLU(c *C) {
	c.Assert(p.c.Valid("4011"), Equals, true)
	c.Assert(p.c.Valid(4011), Equals, true)
	c.Assert(p.c.Organic(4011), Equals, false)
	c.Assert(p.c.GM(4011), Equals, false)
	c.Assert(p.c.RetailerAssigned(4011), Equals, false)
}

func (p *pluSuite) TestOrganic(c *C) {
	c.Assert(p.c.Organic(94011), Equals, true)
}

func (p *pluSuite) TestGM(c *C) {
	c.Assert(p.c.GM(84011), Equals, true)
}

func (p *pluSuite) TestInvalid(c *C) {
	c.Assert(p.c.Valid(2000), Equals, false)
	c.Assert(p.c.Valid(5000), Equals, false)
	c.Assert(p.c.Valid(400), Equals, false)
	c.Assert(p.c.Valid(40111), Equals, false)
}

func (p *pluSuite) TestInvalidModifier(c *C) {
	c.Assert(p.c.Valid(74011), Equals, false)
}

func (p *pluSuite) TestName(c *C) {
	c.Assert(p.c.Name(94011), Equals, "Bananas")
}

func (p *pluSuite) TestNoName(c *C) {
	c.Assert(p.c.Name(4000), Equals, "")
}

func (p *pluSuite) TestRetailerAssigned(c *C) {
	c.Assert(p.c.RetailerAssigned(3170), Equals, true)
}

func (p *pluSuite) TestAll(c *C) {
	m1 := map[string]string{}
	for k, v := range p.c.All() {
		m1[v] = k
	}
	m2 := p.c.All()
	for k, v := range m1 {
		c.Assert(m2[v], Equals, k)
	}
	c.Assert(len(m1), Equals, len(m2))
}
