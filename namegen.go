/*-
* Copyright (c) 2013, Saggi Mizrahi
* All rights reserved.
*
* Redistribution and use in source and binary forms, with or without
* modification, are permitted provided that the following conditions are met:
* 1. Redistributions of source code must retain the above copyright
*    notice, this list of conditions and the following disclaimer.
* 2. Redistributions in binary form must reproduce the above copyright
*    notice, this list of conditions and the following disclaimer in the
*    documentation and/or other materials provided with the distribution.
* 3. Neither the name of Saggi Mizrahi nor the
*    names of its contributors may be used to endorse or promote products
*    derived from this software without specific prior written permission.
*
* THIS SOFTWARE IS PROVIDED BY SAGGI MIZRAHI ''AS IS'' AND ANY
* EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
* WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
* DISCLAIMED. IN NO EVENT SHALL SAGGI MIZRAHI BE LIABLE FOR ANY
* DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
* (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
* LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
* ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
* (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
* SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

// Library for generating random names.
package namegen

import (
	"fmt"
	"math/rand"
	"time"
)

var _rand = rand.NewSource(time.Now().UnixNano())

type Gender string

const (
	Male   = Gender("Male")
	Female = Gender("Female")
)

type Name struct {
	First  string
	Last   string
	Gender Gender
}

func (n Name) String() string {
	return fmt.Sprintf("%s %s [%s]", n.First, n.Last, n.Gender)
}

// Precalculated
var (
	malen   = len(male)
	femalen = len(female)
	lastn   = len(last)
	genders = []Gender{Male, Female}
)

// Fills up empty field in a name struct.
// Useful for generating families or names of a specific gender.
func FillName(n *Name) {
	if n.Gender != Male && n.Gender != Female {
		n.Gender = genders[_rand.Int63()%2]
	}

	if n.First == "" {
		switch n.Gender {
		case Male:
			n.First = male[int(_rand.Int63())%malen]
		case Female:
			n.First = female[int(_rand.Int63())%femalen]
		}
	}

	if n.Last == "" {
		n.Last = last[int(_rand.Int63())%lastn]
	}
}

// Creates a new random name
func NewName() Name {
	var name Name
	FillName(&name)
	return name
}
