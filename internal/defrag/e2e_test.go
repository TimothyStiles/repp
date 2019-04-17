// +build e2e

<<<<<<< HEAD
package rvec
=======
package defrag
>>>>>>> 6f3450c4125f17d3ff1136ff8c157a24a9b4d467

import (
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

<<<<<<< HEAD
	"github.com/jjtimmons/rvec/config"
=======
	"github.com/jjtimmons/defrag/config"
>>>>>>> 6f3450c4125f17d3ff1136ff8c157a24a9b4d467
)

func Test_sequence_e2e(test *testing.T) {
	c := config.New()

	type testFlags struct {
		in       string
		out      string
		backbone string
		enzyme   string
		filters  string
		dbs      []string
		addgene  bool
		igem     bool
	}

	tests := []testFlags{
		testFlags{
			path.Join("..", "..", "test", "input", "backbone.fa"),
			path.Join("..", "..", "test", "output", "backbone.json"),
			"pSB1A3",
			"PstI",
			"2018,2019",
			[]string{},
			false,
			true,
		},
		testFlags{
			path.Join("..", "..", "test", "input", "110056.fa"),
			path.Join("..", "..", "test", "output", "110056.json"),
			"",
			"",
			"2019,2018",
			[]string{},
			true,
			true,
		},
		testFlags{
			path.Join("..", "..", "test", "input", "BBa_K2602025.fa"),
			path.Join("..", "..", "test", "output", "BBa_K2602025.json"),
			"pSB1A3",
			"PstI",
			"",
			[]string{},
			false,
			true,
		},
		testFlags{
			path.Join("..", "..", "test", "input", "BBa_K2779020.fa"),
			path.Join("..", "..", "test", "output", "BBa_K2779020.json"),
			"pSB1A3",
			"PstI",
			"BBa_K277",
			[]string{},
			false,
			true,
		},
		testFlags{
			path.Join("..", "..", "test", "input", "BBa_E0610.fa"),
			path.Join("..", "..", "test", "output", "BBa_E0610.json"),
			"pSB1C3",
			"EcoRI",
			"2005,2006,2007,2008,2009,2010,2011,2012,2013,2014,2015,2016,2017,2018,BBa_E061",
			[]string{},
			false,
			true,
		},
		testFlags{
			path.Join("..", "..", "test", "input", "BBa_I5310.fa"),
			path.Join("..", "..", "test", "output", "BBa_I5310.json"),
			"pSB1C3",
			"EcoRI",
			"",
			[]string{},
			true,
			true,
		},
		testFlags{
			path.Join("..", "..", "test", "input", "BBa_K2651001.fa"),
			path.Join("..", "..", "test", "output", "BBa_K2651001.json"),
			"pSB1C3",
			"EcoRI",
			"BBa_K265",
			[]string{},
			true,
			true,
		},
		testFlags{
			path.Join("..", "..", "test", "input", "BBa_K2779020.fa"),
			path.Join("..", "..", "test", "output", "BBa_K2779020.json"),
			"pSB1A3",
			"PstI",
			"BBa_K277", // no year filters needed
			[]string{},
			true,
			true,
		},
		testFlags{
			path.Join("..", "..", "test", "input", "BBa_K1085023.fa"),
			path.Join("..", "..", "test", "output", "BBa_K1085023.json"),
			"pSB1C3",
			"EcoRI",
			"2009,2010,2011,2012,BBa_K108",
			[]string{},
			true,
			true,
		},
	}

	for _, t := range tests {
		sols := Sequence(NewFlags(t.in, t.out, t.backbone, t.enzyme, t.filters, t.dbs, t.addgene, t.igem, false))

		if len(sols) < 1 {
			test.Errorf("no solutions for %s", t.in)
		}

		for _, s := range sols {
			validateJunctions(t.in, s, c, test)
		}
	}
}

func Test_features(t *testing.T) {
	test1, conf := NewFlags(
		"p10 promoter, mEGFP, T7 terminator",
		filepath.Join("..", "..", "test", "output", "features.json"),
		"pSB1A3",
		"EcoRI",
		"",
		[]string{},
		true,
		true,
		false,
	)

	type args struct {
		flags *Flags
		conf  *config.Config
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test end to end features creation",
			args{
				flags: test1,
				conf:  conf,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sols := Features(tt.args.flags, tt.args.conf)

			for _, s := range sols {
				validateJunctions(tt.name, s, conf, t)
			}
		})
	}
}

func Test_fragments(t *testing.T) {
	c := config.New()
	c.PCRMinLength = 10
	c.FragmentsMinHomology = 8
	c.FragmentsMaxHomology = 20

	type args struct {
		inputFragments []*Frag
		conf           *config.Config
	}
	tests := []struct {
		name             string
		args             args
		wantTargetVector *Frag
		wantFragments    []*Frag
	}{
		{
			"fragments with linear overlap",
			args{
				[]*Frag{
					&Frag{
						Seq:  "ACGTGCTAGCTACATCGATCGTAGCTAGCTAGCATCG",
						conf: c,
					},
					&Frag{
						Seq:  "AGCTAGCATCGACTGATCACTAGCATCGACTAGCTAG",
						conf: c,
					},
					&Frag{
						Seq:  "TCGACTAGCTAGAACTGATCTAGACGTGCTAGCTACA",
						conf: c,
					},
				},
				c,
			},
			&Frag{
				Seq: "ACGTGCTAGCTACATCGATCGTAGCTAGCTAGCATCGACTGATCACTAGCATCGACTAGCTAGAACTGATCTAG",
			},
			[]*Frag{
				&Frag{
					Seq:      "ACGTGCTAGCTACATCGATCGTAGCTAGCTAGCATCG",
					fragType: linear,
				},
				&Frag{
					Seq:      "AGCTAGCATCGACTGATCACTAGCATCGACTAGCTAG",
					fragType: linear,
				},
				&Frag{
					Seq:      "TCGACTAGCTAGAACTGATCTAGACGTGCTAGCTACA",
					fragType: linear,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTargetVector, gotFragments := fragments(tt.args.inputFragments, tt.args.conf)

			if !reflect.DeepEqual(gotTargetVector.Seq, tt.wantTargetVector.Seq) {
				t.Errorf("assembleFWD() gotTargetVector = %v, want %v", gotTargetVector, tt.wantTargetVector)
			}

			for i, wantF := range tt.wantFragments {
				if wantF.Seq != gotFragments[i].Seq {
					t.Errorf("assembleFWD() gotFragment.Seq = %v, want %v", gotFragments[i].Seq, wantF.Seq)
				}

				if wantF.fragType != gotFragments[i].fragType {
					t.Errorf("assembleFWD() gotFragment.Type = %v, want %v", gotFragments[i].fragType, wantF.fragType)
				}
			}
		})
	}
}

// if an input fragment being built is exactly the same as one in a DB, it should be used
// as is and without PCR or any preparation
func Test_vector_single_vector(t *testing.T) {
	fs, c := NewFlags(
		path.Join("..", "..", "test", "input", "109049.addgene.fa"),
		path.Join("..", "..", "test", "output", "109049.output.json"),
		"",
		"",
		"",
		[]string{},
		true,
		false,
		false,
	)

	assemblies := Sequence(fs, c) // use addgene database

	if !strings.Contains(assemblies[0][0].URL, "109049") {
		t.Fatal("failed to use 109049 to build the vector")
	}
}
