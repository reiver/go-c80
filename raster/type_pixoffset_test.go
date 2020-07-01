package c80raster_test

import (
	"github.com/reiver/go-c80/raster"

	"testing"
)

func TestTypePixOffset(t *testing.T) {

	tests := []struct{
		X int
		Y int
		Expected int
	}{
		{
			X:        0,
			Y:        0,
			Expected: 0,
		},
		{
			X:        1,
			Y:        0,
			Expected: 1,
		},
		{
			X:        2,
			Y:        0,
			Expected: 2,
		},
		{
			X:        3,
			Y:        0,
			Expected: 3,
		},
		{
			X:        4,
			Y:        0,
			Expected: 4,
		},
		{
			X:        5,
			Y:        0,
			Expected: 5,
		},
		{
			X:          6,
			Y:          0,
			Expected:   6,
		},
		{
			X:          7,
			Y:          0,
			Expected:   7,
		},
		{
			X:          8,
			Y:          0,
			Expected:   8,
		},
		{
			X:          9,
			Y:          0,
			Expected:   9,
		},
		{
			X:         10,
			Y:          0,
			Expected:  10,
		},
		{
			X:         11,
			Y:          0,
			Expected:  11,
		},
		{
			X:         12,
			Y:          0,
			Expected:  12,
		},
		{
			X:         13,
			Y:          0,
			Expected:  13,
		},
		{
			X:         14,
			Y:          0,
			Expected:  14,
		},
		{
			X:         15,
			Y:          0,
			Expected:  15,
		},
		{
			X:         16,
			Y:          0,
			Expected:  16,
		},
		{
			X:         17,
			Y:          0,
			Expected:  17,
		},
		{
			X:         18,
			Y:          0,
			Expected:  18,
		},
		{
			X:         19,
			Y:          0,
			Expected:  19,
		},
		{
			X:         20,
			Y:          0,
			Expected:  20,
		},
		{
			X:         21,
			Y:          0,
			Expected:  21,
		},
		{
			X:         22,
			Y:          0,
			Expected:  22,
		},
		{
			X:         23,
			Y:          0,
			Expected:  23,
		},
		{
			X:         24,
			Y:          0,
			Expected:  24,
		},
		{
			X:         25,
			Y:          0,
			Expected:  25,
		},
		{
			X:         26,
			Y:          0,
			Expected:  26,
		},
		{
			X:         27,
			Y:          0,
			Expected:  27,
		},
		{
			X:         28,
			Y:          0,
			Expected:  28,
		},
		{
			X:         29,
			Y:          0,
			Expected:  29,
		},
		{
			X:         30,
			Y:          0,
			Expected:  30,
		},
		{
			X:         31,
			Y:          0,
			Expected:  31,
		},
		{
			X:         32,
			Y:          0,
			Expected:  32,
		},
		{
			X:         33,
			Y:          0,
			Expected:  33,
		},
		{
			X:         34,
			Y:          0,
			Expected:  34,
		},
		{
			X:         35,
			Y:          0,
			Expected:  35,
		},
		{
			X:         36,
			Y:          0,
			Expected:  36,
		},
		{
			X:         37,
			Y:          0,
			Expected:  37,
		},
		{
			X:         38,
			Y:          0,
			Expected:  38,
		},
		{
			X:         39,
			Y:          0,
			Expected:  39,
		},
		{
			X:         40,
			Y:          0,
			Expected:  40,
		},
		{
			X:         41,
			Y:          0,
			Expected:  41,
		},
		{
			X:         42,
			Y:          0,
			Expected:  42,
		},
		{
			X:         43,
			Y:          0,
			Expected:  43,
		},
		{
			X:         44,
			Y:          0,
			Expected:  44,
		},
		{
			X:         45,
			Y:          0,
			Expected:  45,
		},
		{
			X:         46,
			Y:          0,
			Expected:  46,
		},
		{
			X:         47,
			Y:          0,
			Expected:  47,
		},
		{
			X:         48,
			Y:          0,
			Expected:  48,
		},
		{
			X:         49,
			Y:          0,
			Expected:  49,
		},
		{
			X:         50,
			Y:          0,
			Expected:  50,
		},
		{
			X:         51,
			Y:          0,
			Expected:  51,
		},
		{
			X:         52,
			Y:          0,
			Expected:  52,
		},
		{
			X:         53,
			Y:          0,
			Expected:  53,
		},
		{
			X:         54,
			Y:          0,
			Expected:  54,
		},
		{
			X:         55,
			Y:          0,
			Expected:  55,
		},
		{
			X:         56,
			Y:          0,
			Expected:  56,
		},
		{
			X:         57,
			Y:          0,
			Expected:  57,
		},
		{
			X:         58,
			Y:          0,
			Expected:  58,
		},
		{
			X:         59,
			Y:          0,
			Expected:  59,
		},
		{
			X:         60,
			Y:          0,
			Expected:  60,
		},
		{
			X:         61,
			Y:          0,
			Expected:  61,
		},
		{
			X:         62,
			Y:          0,
			Expected:  62,
		},
		{
			X:         63,
			Y:          0,
			Expected:  63,
		},
		{
			X:         64,
			Y:          0,
			Expected:  64,
		},
		{
			X:         65,
			Y:          0,
			Expected:  65,
		},
		{
			X:         66,
			Y:          0,
			Expected:  66,
		},
		{
			X:         67,
			Y:          0,
			Expected:  67,
		},
		{
			X:         68,
			Y:          0,
			Expected:  68,
		},
		{
			X:         69,
			Y:          0,
			Expected:  69,
		},
		{
			X:         70,
			Y:          0,
			Expected:  70,
		},
		{
			X:         71,
			Y:          0,
			Expected:  71,
		},
		{
			X:         72,
			Y:          0,
			Expected:  72,
		},
		{
			X:         73,
			Y:          0,
			Expected:  73,
		},
		{
			X:         74,
			Y:          0,
			Expected:  74,
		},
		{
			X:         75,
			Y:          0,
			Expected:  75,
		},
		{
			X:         76,
			Y:          0,
			Expected:  76,
		},
		{
			X:         77,
			Y:          0,
			Expected:  77,
		},
		{
			X:         78,
			Y:          0,
			Expected:  78,
		},
		{
			X:         79,
			Y:          0,
			Expected:  79,
		},
		{
			X:         80,
			Y:          0,
			Expected:  80,
		},
		{
			X:         81,
			Y:          0,
			Expected:  81,
		},
		{
			X:         82,
			Y:          0,
			Expected:  82,
		},
		{
			X:         83,
			Y:          0,
			Expected:  83,
		},
		{
			X:         84,
			Y:          0,
			Expected:  84,
		},
		{
			X:         85,
			Y:          0,
			Expected:  85,
		},
		{
			X:         86,
			Y:          0,
			Expected:  86,
		},
		{
			X:         87,
			Y:          0,
			Expected:  87,
		},
		{
			X:         88,
			Y:          0,
			Expected:  88,
		},
		{
			X:         89,
			Y:          0,
			Expected:  89,
		},
		{
			X:         90,
			Y:          0,
			Expected:  90,
		},
		{
			X:         91,
			Y:          0,
			Expected:  91,
		},
		{
			X:         92,
			Y:          0,
			Expected:  92,
		},
		{
			X:         93,
			Y:          0,
			Expected:  93,
		},
		{
			X:         94,
			Y:          0,
			Expected:  94,
		},
		{
			X:         95,
			Y:          0,
			Expected:  95,
		},
		{
			X:         96,
			Y:          0,
			Expected:  96,
		},
		{
			X:         97,
			Y:          0,
			Expected:  97,
		},
		{
			X:         98,
			Y:          0,
			Expected:  98,
		},
		{
			X:         99,
			Y:          0,
			Expected:  99,
		},
		{
			X:        100,
			Y:          0,
			Expected: 100,
		},
		{
			X:        101,
			Y:          0,
			Expected: 101,
		},
		{
			X:        102,
			Y:          0,
			Expected: 102,
		},
		{
			X:        103,
			Y:          0,
			Expected: 103,
		},
		{
			X:        104,
			Y:          0,
			Expected: 104,
		},
		{
			X:        105,
			Y:          0,
			Expected: 105,
		},
		{
			X:        106,
			Y:          0,
			Expected: 106,
		},
		{
			X:        107,
			Y:          0,
			Expected: 107,
		},
		{
			X:        108,
			Y:          0,
			Expected: 108,
		},
		{
			X:        109,
			Y:          0,
			Expected: 109,
		},
		{
			X:        110,
			Y:          0,
			Expected: 110,
		},
		{
			X:        111,
			Y:          0,
			Expected: 111,
		},
		{
			X:        112,
			Y:          0,
			Expected: 112,
		},
		{
			X:        113,
			Y:          0,
			Expected: 113,
		},
		{
			X:        114,
			Y:          0,
			Expected: 114,
		},
		{
			X:        115,
			Y:          0,
			Expected: 115,
		},
		{
			X:        116,
			Y:          0,
			Expected: 116,
		},
		{
			X:        117,
			Y:          0,
			Expected: 117,
		},
		{
			X:        118,
			Y:          0,
			Expected: 118,
		},
		{
			X:        119,
			Y:          0,
			Expected: 119,
		},
		{
			X:        120,
			Y:          0,
			Expected: 120,
		},
		{
			X:        121,
			Y:          0,
			Expected: 121,
		},
		{
			X:        122,
			Y:          0,
			Expected: 122,
		},
		{
			X:        123,
			Y:          0,
			Expected: 123,
		},
		{
			X:        124,
			Y:          0,
			Expected: 124,
		},
		{
			X:        125,
			Y:          0,
			Expected: 125,
		},
		{
			X:        126,
			Y:          0,
			Expected: 126,
		},
		{
			X:        127,
			Y:          0,
			Expected: 127,
		},
		{
			X:        128,
			Y:          0,
			Expected: 128,
		},
		{
			X:        129,
			Y:          0,
			Expected: 129,
		},
		{
			X:        130,
			Y:          0,
			Expected: 130,
		},
		{
			X:        131,
			Y:          0,
			Expected: 131,
		},
		{
			X:        132,
			Y:          0,
			Expected: 132,
		},
		{
			X:        133,
			Y:          0,
			Expected: 133,
		},
		{
			X:        134,
			Y:          0,
			Expected: 134,
		},
		{
			X:        135,
			Y:          0,
			Expected: 135,
		},
		{
			X:        136,
			Y:          0,
			Expected: 136,
		},
		{
			X:        137,
			Y:          0,
			Expected: 137,
		},
		{
			X:        138,
			Y:          0,
			Expected: 138,
		},
		{
			X:        139,
			Y:          0,
			Expected: 139,
		},
		{
			X:        140,
			Y:          0,
			Expected: 140,
		},
		{
			X:        140,
			Y:          0,
			Expected: 140,
		},
		{
			X:        141,
			Y:          0,
			Expected: 141,
		},
		{
			X:        142,
			Y:          0,
			Expected: 142,
		},
		{
			X:        143,
			Y:          0,
			Expected: 143,
		},
		{
			X:        144,
			Y:          0,
			Expected: 144,
		},
		{
			X:        145,
			Y:          0,
			Expected: 145,
		},
		{
			X:        146,
			Y:          0,
			Expected: 146,
		},
		{
			X:        147,
			Y:          0,
			Expected: 147,
		},
		{
			X:        148,
			Y:          0,
			Expected: 148,
		},
		{
			X:        149,
			Y:          0,
			Expected: 149,
		},
		{
			X:        150,
			Y:          0,
			Expected: 150,
		},
		{
			X:        151,
			Y:          0,
			Expected: 151,
		},
		{
			X:        152,
			Y:          0,
			Expected: 152,
		},
		{
			X:        153,
			Y:          0,
			Expected: 153,
		},
		{
			X:        154,
			Y:          0,
			Expected: 154,
		},
		{
			X:        155,
			Y:          0,
			Expected: 155,
		},
		{
			X:        156,
			Y:          0,
			Expected: 156,
		},
		{
			X:        157,
			Y:          0,
			Expected: 157,
		},
		{
			X:        158,
			Y:          0,
			Expected: 158,
		},
		{
			X:        159,
			Y:          0,
			Expected: 159,
		},
		{
			X:        160,
			Y:          0,
			Expected: 160,
		},
		{
			X:        161,
			Y:          0,
			Expected: 161,
		},
		{
			X:        162,
			Y:          0,
			Expected: 162,
		},
		{
			X:        163,
			Y:          0,
			Expected: 163,
		},
		{
			X:        164,
			Y:          0,
			Expected: 164,
		},
		{
			X:        165,
			Y:          0,
			Expected: 165,
		},
		{
			X:        166,
			Y:          0,
			Expected: 166,
		},
		{
			X:        167,
			Y:          0,
			Expected: 167,
		},
		{
			X:        168,
			Y:          0,
			Expected: 168,
		},
		{
			X:        169,
			Y:          0,
			Expected: 169,
		},
		{
			X:        170,
			Y:          0,
			Expected: 170,
		},
		{
			X:        171,
			Y:          0,
			Expected: 171,
		},
		{
			X:        172,
			Y:          0,
			Expected: 172,
		},
		{
			X:        173,
			Y:          0,
			Expected: 173,
		},
		{
			X:        174,
			Y:          0,
			Expected: 174,
		},
		{
			X:        175,
			Y:          0,
			Expected: 175,
		},
		{
			X:        176,
			Y:          0,
			Expected: 176,
		},
		{
			X:        177,
			Y:          0,
			Expected: 177,
		},
		{
			X:        178,
			Y:          0,
			Expected: 178,
		},
		{
			X:        179,
			Y:          0,
			Expected: 179,
		},
		{
			X:        180,
			Y:          0,
			Expected: 180,
		},
		{
			X:        181,
			Y:          0,
			Expected: 181,
		},
		{
			X:        182,
			Y:          0,
			Expected: 182,
		},
		{
			X:        183,
			Y:          0,
			Expected: 183,
		},
		{
			X:        184,
			Y:          0,
			Expected: 184,
		},
		{
			X:        185,
			Y:          0,
			Expected: 185,
		},
		{
			X:        186,
			Y:          0,
			Expected: 186,
		},
		{
			X:        187,
			Y:          0,
			Expected: 187,
		},
		{
			X:        188,
			Y:          0,
			Expected: 188,
		},
		{
			X:        189,
			Y:          0,
			Expected: 189,
		},
		{
			X:        190,
			Y:          0,
			Expected: 190,
		},
		{
			X:        191,
			Y:          0,
			Expected: 191,
		},
		{
			X:        192,
			Y:          0,
			Expected: 192,
		},
		{
			X:        193,
			Y:          0,
			Expected: 193,
		},
		{
			X:        194,
			Y:          0,
			Expected: 194,
		},
		{
			X:        195,
			Y:          0,
			Expected: 195,
		},
		{
			X:        196,
			Y:          0,
			Expected: 196,
		},
		{
			X:        197,
			Y:          0,
			Expected: 197,
		},
		{
			X:        198,
			Y:          0,
			Expected: 198,
		},
		{
			X:        199,
			Y:          0,
			Expected: 199,
		},
		{
			X:        200,
			Y:          0,
			Expected: 200,
		},
		{
			X:        201,
			Y:          0,
			Expected: 201,
		},
		{
			X:        202,
			Y:          0,
			Expected: 202,
		},
		{
			X:        203,
			Y:          0,
			Expected: 203,
		},
		{
			X:        204,
			Y:          0,
			Expected: 204,
		},
		{
			X:        205,
			Y:          0,
			Expected: 205,
		},
		{
			X:        206,
			Y:          0,
			Expected: 206,
		},
		{
			X:        207,
			Y:          0,
			Expected: 207,
		},
		{
			X:        208,
			Y:          0,
			Expected: 208,
		},
		{
			X:        209,
			Y:          0,
			Expected: 209,
		},
		{
			X:        210,
			Y:          0,
			Expected: 210,
		},
		{
			X:        211,
			Y:          0,
			Expected: 211,
		},
		{
			X:        212,
			Y:          0,
			Expected: 212,
		},
		{
			X:        213,
			Y:          0,
			Expected: 213,
		},
		{
			X:        214,
			Y:          0,
			Expected: 214,
		},
		{
			X:        215,
			Y:          0,
			Expected: 215,
		},
		{
			X:        216,
			Y:          0,
			Expected: 216,
		},
		{
			X:        217,
			Y:          0,
			Expected: 217,
		},
		{
			X:        218,
			Y:          0,
			Expected: 218,
		},
		{
			X:        219,
			Y:          0,
			Expected: 219,
		},
		{
			X:        220,
			Y:          0,
			Expected: 220,
		},
		{
			X:        221,
			Y:          0,
			Expected: 221,
		},
		{
			X:        222,
			Y:          0,
			Expected: 222,
		},
		{
			X:        223,
			Y:          0,
			Expected: 223,
		},
		{
			X:        224,
			Y:          0,
			Expected: 224,
		},
		{
			X:        225,
			Y:          0,
			Expected: 225,
		},
		{
			X:        226,
			Y:          0,
			Expected: 226,
		},
		{
			X:        227,
			Y:          0,
			Expected: 227,
		},
		{
			X:        228,
			Y:          0,
			Expected: 228,
		},
		{
			X:        229,
			Y:          0,
			Expected: 229,
		},
		{
			X:        230,
			Y:          0,
			Expected: 230,
		},
		{
			X:        231,
			Y:          0,
			Expected: 231,
		},
		{
			X:        232,
			Y:          0,
			Expected: 232,
		},
		{
			X:        233,
			Y:          0,
			Expected: 233,
		},
		{
			X:        234,
			Y:          0,
			Expected: 234,
		},
		{
			X:        235,
			Y:          0,
			Expected: 235,
		},
		{
			X:        236,
			Y:          0,
			Expected: 236,
		},
		{
			X:        237,
			Y:          0,
			Expected: 237,
		},
		{
			X:        238,
			Y:          0,
			Expected: 238,
		},
		{
			X:        239,
			Y:          0,
			Expected: 239,
		},
		{
			X:        240,
			Y:          0,
			Expected: 240,
		},
		{
			X:        241,
			Y:          0,
			Expected: 241,
		},
		{
			X:        242,
			Y:          0,
			Expected: 242,
		},
		{
			X:        243,
			Y:          0,
			Expected: 243,
		},
		{
			X:        244,
			Y:          0,
			Expected: 244,
		},
		{
			X:        245,
			Y:          0,
			Expected: 245,
		},
		{
			X:        246,
			Y:          0,
			Expected: 246,
		},
		{
			X:        247,
			Y:          0,
			Expected: 247,
		},
		{
			X:        248,
			Y:          0,
			Expected: 248,
		},
		{
			X:        249,
			Y:          0,
			Expected: 249,
		},
		{
			X:        250,
			Y:          0,
			Expected: 250,
		},
		{
			X:        251,
			Y:          0,
			Expected: 251,
		},
		{
			X:        252,
			Y:          0,
			Expected: 252,
		},
		{
			X:        253,
			Y:          0,
			Expected: 253,
		},
		{
			X:        254,
			Y:          0,
			Expected: 254,
		},
		{
			X:        255,
			Y:          0,
			Expected: 255,
		},









		{
			X:          0,
			Y:          1,
			Expected: 256,
		},
		{
			X:          1,
			Y:          1,
			Expected: 257,
		},
		{
			X:          2,
			Y:          1,
			Expected: 258,
		},
		{
			X:          3,
			Y:          1,
			Expected: 259,
		},
		{
			X:          4,
			Y:          1,
			Expected: 260,
		},
		{
			X:          5,
			Y:          1,
			Expected: 261,
		},
		{
			X:          6,
			Y:          1,
			Expected: 262,
		},
		{
			X:          7,
			Y:          1,
			Expected: 263,
		},
		{
			X:          8,
			Y:          1,
			Expected: 264,
		},
		{
			X:          9,
			Y:          1,
			Expected: 265,
		},
		{
			X:         10,
			Y:          1,
			Expected: 266,
		},
		{
			X:         11,
			Y:          1,
			Expected: 267,
		},
		{
			X:         12,
			Y:          1,
			Expected: 268,
		},
		{
			X:         13,
			Y:          1,
			Expected: 269,
		},
		{
			X:         14,
			Y:          1,
			Expected: 270,
		},
		{
			X:         15,
			Y:          1,
			Expected: 271,
		},
		{
			X:         16,
			Y:          1,
			Expected: 272,
		},
		{
			X:         17,
			Y:          1,
			Expected: 273,
		},
		{
			X:         18,
			Y:          1,
			Expected: 274,
		},
		{
			X:         19,
			Y:          1,
			Expected: 275,
		},
		{
			X:         20,
			Y:          1,
			Expected: 276,
		},
		{
			X:         21,
			Y:          1,
			Expected: 277,
		},
		{
			X:         22,
			Y:          1,
			Expected: 278,
		},
		{
			X:         23,
			Y:          1,
			Expected: 279,
		},
		{
			X:         24,
			Y:          1,
			Expected: 280,
		},
		{
			X:         25,
			Y:          1,
			Expected: 281,
		},
		{
			X:         26,
			Y:          1,
			Expected: 282,
		},
		{
			X:         27,
			Y:          1,
			Expected: 283,
		},
		{
			X:         28,
			Y:          1,
			Expected: 284,
		},
		{
			X:         29,
			Y:          1,
			Expected: 285,
		},
		{
			X:         30,
			Y:          1,
			Expected: 286,
		},
		{
			X:         31,
			Y:          1,
			Expected: 287,
		},
		{
			X:         32,
			Y:          1,
			Expected: 288,
		},
		{
			X:         33,
			Y:          1,
			Expected: 289,
		},
		{
			X:         34,
			Y:          1,
			Expected: 290,
		},
		{
			X:         35,
			Y:          1,
			Expected: 291,
		},
		{
			X:         36,
			Y:          1,
			Expected: 292,
		},
		{
			X:         37,
			Y:          1,
			Expected: 293,
		},
		{
			X:         38,
			Y:          1,
			Expected: 294,
		},
		{
			X:         39,
			Y:          1,
			Expected: 295,
		},
		{
			X:         40,
			Y:          1,
			Expected: 296,
		},
		{
			X:         41,
			Y:          1,
			Expected: 297,
		},
		{
			X:         42,
			Y:          1,
			Expected: 298,
		},
		{
			X:         43,
			Y:          1,
			Expected: 299,
		},
		{
			X:         44,
			Y:          1,
			Expected: 300,
		},
		{
			X:         45,
			Y:          1,
			Expected: 301,
		},
		{
			X:         46,
			Y:          1,
			Expected: 302,
		},
		{
			X:         47,
			Y:          1,
			Expected: 303,
		},
		{
			X:         48,
			Y:          1,
			Expected: 304,
		},
		{
			X:         49,
			Y:          1,
			Expected: 305,
		},
		{
			X:         50,
			Y:          1,
			Expected: 306,
		},
		{
			X:         51,
			Y:          1,
			Expected: 307,
		},
		{
			X:         52,
			Y:          1,
			Expected: 308,
		},
		{
			X:         53,
			Y:          1,
			Expected: 309,
		},
		{
			X:         54,
			Y:          1,
			Expected: 310,
		},
		{
			X:         55,
			Y:          1,
			Expected: 311,
		},
		{
			X:         56,
			Y:          1,
			Expected: 312,
		},
		{
			X:         57,
			Y:          1,
			Expected: 313,
		},
		{
			X:         58,
			Y:          1,
			Expected: 314,
		},
		{
			X:         59,
			Y:          1,
			Expected: 315,
		},
		{
			X:         60,
			Y:          1,
			Expected: 316,
		},
		{
			X:         61,
			Y:          1,
			Expected: 317,
		},
		{
			X:         62,
			Y:          1,
			Expected: 318,
		},
		{
			X:         63,
			Y:          1,
			Expected: 319,
		},
		{
			X:         64,
			Y:          1,
			Expected: 320,
		},
		{
			X:         65,
			Y:          1,
			Expected: 321,
		},
		{
			X:         66,
			Y:          1,
			Expected: 322,
		},
		{
			X:         67,
			Y:          1,
			Expected: 323,
		},
		{
			X:         68,
			Y:          1,
			Expected: 324,
		},
		{
			X:         69,
			Y:          1,
			Expected: 325,
		},
		{
			X:         70,
			Y:          1,
			Expected: 326,
		},
		{
			X:         71,
			Y:          1,
			Expected: 327,
		},
		{
			X:         72,
			Y:          1,
			Expected: 328,
		},
		{
			X:         73,
			Y:          1,
			Expected: 329,
		},
		{
			X:         74,
			Y:          1,
			Expected: 330,
		},
		{
			X:         75,
			Y:          1,
			Expected: 331,
		},
		{
			X:         76,
			Y:          1,
			Expected: 332,
		},
		{
			X:         77,
			Y:          1,
			Expected: 333,
		},
		{
			X:         78,
			Y:          1,
			Expected: 334,
		},
		{
			X:         79,
			Y:          1,
			Expected: 335,
		},
		{
			X:         80,
			Y:          1,
			Expected: 336,
		},
		{
			X:         81,
			Y:          1,
			Expected: 337,
		},
		{
			X:         82,
			Y:          1,
			Expected: 338,
		},
		{
			X:         83,
			Y:          1,
			Expected: 339,
		},
		{
			X:         84,
			Y:          1,
			Expected: 340,
		},
		{
			X:         85,
			Y:          1,
			Expected: 341,
		},
		{
			X:         86,
			Y:          1,
			Expected: 342,
		},
		{
			X:         87,
			Y:          1,
			Expected: 343,
		},
		{
			X:         88,
			Y:          1,
			Expected: 344,
		},
		{
			X:         89,
			Y:          1,
			Expected: 345,
		},
		{
			X:         90,
			Y:          1,
			Expected: 346,
		},
		{
			X:         91,
			Y:          1,
			Expected: 347,
		},
		{
			X:         92,
			Y:          1,
			Expected: 348,
		},
		{
			X:         93,
			Y:          1,
			Expected: 349,
		},
		{
			X:         94,
			Y:          1,
			Expected: 350,
		},
		{
			X:         95,
			Y:          1,
			Expected: 351,
		},
		{
			X:         96,
			Y:          1,
			Expected: 352,
		},
		{
			X:         97,
			Y:          1,
			Expected: 353,
		},
		{
			X:         98,
			Y:          1,
			Expected: 354,
		},
		{
			X:         99,
			Y:          1,
			Expected: 355,
		},
		{
			X:        100,
			Y:          1,
			Expected: 356,
		},
		{
			X:        101,
			Y:          1,
			Expected: 357,
		},
		{
			X:        102,
			Y:          1,
			Expected: 358,
		},
		{
			X:        103,
			Y:          1,
			Expected: 359,
		},
		{
			X:        104,
			Y:          1,
			Expected: 360,
		},
		{
			X:        105,
			Y:          1,
			Expected: 361,
		},
		{
			X:        106,
			Y:          1,
			Expected: 362,
		},
		{
			X:        107,
			Y:          1,
			Expected: 363,
		},
		{
			X:        108,
			Y:          1,
			Expected: 364,
		},
		{
			X:        109,
			Y:          1,
			Expected: 365,
		},
		{
			X:        110,
			Y:          1,
			Expected: 366,
		},
		{
			X:        111,
			Y:          1,
			Expected: 367,
		},
		{
			X:        112,
			Y:          1,
			Expected: 368,
		},
		{
			X:        113,
			Y:          1,
			Expected: 369,
		},
		{
			X:        114,
			Y:          1,
			Expected: 370,
		},
		{
			X:        115,
			Y:          1,
			Expected: 371,
		},
		{
			X:        116,
			Y:          1,
			Expected: 372,
		},
		{
			X:        117,
			Y:          1,
			Expected: 373,
		},
		{
			X:        118,
			Y:          1,
			Expected: 374,
		},
		{
			X:        119,
			Y:          1,
			Expected: 375,
		},
		{
			X:        120,
			Y:          1,
			Expected: 376,
		},
		{
			X:        121,
			Y:          1,
			Expected: 377,
		},
		{
			X:        122,
			Y:          1,
			Expected: 378,
		},
		{
			X:        123,
			Y:          1,
			Expected: 379,
		},
		{
			X:        124,
			Y:          1,
			Expected: 380,
		},
		{
			X:        125,
			Y:          1,
			Expected: 381,
		},
		{
			X:        126,
			Y:          1,
			Expected: 382,
		},
		{
			X:        127,
			Y:          1,
			Expected: 383,
		},
		{
			X:        128,
			Y:          1,
			Expected: 384,
		},
		{
			X:        129,
			Y:          1,
			Expected: 385,
		},
		{
			X:        130,
			Y:          1,
			Expected: 386,
		},
		{
			X:        131,
			Y:          1,
			Expected: 387,
		},
		{
			X:        132,
			Y:          1,
			Expected: 388,
		},
		{
			X:        133,
			Y:          1,
			Expected: 389,
		},
		{
			X:        134,
			Y:          1,
			Expected: 390,
		},
		{
			X:        135,
			Y:          1,
			Expected: 391,
		},
		{
			X:        136,
			Y:          1,
			Expected: 392,
		},
		{
			X:        137,
			Y:          1,
			Expected: 393,
		},
		{
			X:        138,
			Y:          1,
			Expected: 394,
		},
		{
			X:        139,
			Y:          1,
			Expected: 395,
		},
		{
			X:        140,
			Y:          1,
			Expected: 396,
		},
		{
			X:        141,
			Y:          1,
			Expected: 397,
		},
		{
			X:        142,
			Y:          1,
			Expected: 398,
		},
		{
			X:        143,
			Y:          1,
			Expected: 399,
		},
		{
			X:        144,
			Y:          1,
			Expected: 400,
		},
		{
			X:        145,
			Y:          1,
			Expected: 401,
		},
		{
			X:        146,
			Y:          1,
			Expected: 402,
		},
		{
			X:        147,
			Y:          1,
			Expected: 403,
		},
		{
			X:        148,
			Y:          1,
			Expected: 404,
		},
		{
			X:        149,
			Y:          1,
			Expected: 405,
		},
		{
			X:        150,
			Y:          1,
			Expected: 406,
		},
		{
			X:        151,
			Y:          1,
			Expected: 407,
		},
		{
			X:        152,
			Y:          1,
			Expected: 408,
		},
		{
			X:        153,
			Y:          1,
			Expected: 409,
		},
		{
			X:        154,
			Y:          1,
			Expected: 410,
		},
		{
			X:        155,
			Y:          1,
			Expected: 411,
		},
		{
			X:        156,
			Y:          1,
			Expected: 412,
		},
		{
			X:        157,
			Y:          1,
			Expected: 413,
		},
		{
			X:        158,
			Y:          1,
			Expected: 414,
		},
		{
			X:        159,
			Y:          1,
			Expected: 415,
		},
		{
			X:        160,
			Y:          1,
			Expected: 416,
		},
		{
			X:        161,
			Y:          1,
			Expected: 417,
		},
		{
			X:        162,
			Y:          1,
			Expected: 418,
		},
		{
			X:        163,
			Y:          1,
			Expected: 419,
		},
		{
			X:        164,
			Y:          1,
			Expected: 420,
		},
		{
			X:        165,
			Y:          1,
			Expected: 421,
		},
		{
			X:        166,
			Y:          1,
			Expected: 422,
		},
		{
			X:        167,
			Y:          1,
			Expected: 423,
		},
		{
			X:        168,
			Y:          1,
			Expected: 424,
		},
		{
			X:        169,
			Y:          1,
			Expected: 425,
		},
		{
			X:        170,
			Y:          1,
			Expected: 426,
		},
		{
			X:        171,
			Y:          1,
			Expected: 427,
		},
		{
			X:        172,
			Y:          1,
			Expected: 428,
		},
		{
			X:        173,
			Y:          1,
			Expected: 429,
		},
		{
			X:        174,
			Y:          1,
			Expected: 430,
		},
		{
			X:        175,
			Y:          1,
			Expected: 431,
		},
		{
			X:        176,
			Y:          1,
			Expected: 432,
		},
		{
			X:        177,
			Y:          1,
			Expected: 433,
		},
		{
			X:        178,
			Y:          1,
			Expected: 434,
		},
		{
			X:        179,
			Y:          1,
			Expected: 435,
		},
		{
			X:        180,
			Y:          1,
			Expected: 436,
		},
		{
			X:        181,
			Y:          1,
			Expected: 437,
		},
		{
			X:        182,
			Y:          1,
			Expected: 438,
		},
		{
			X:        183,
			Y:          1,
			Expected: 439,
		},
		{
			X:        184,
			Y:          1,
			Expected: 440,
		},
		{
			X:        185,
			Y:          1,
			Expected: 441,
		},
		{
			X:        186,
			Y:          1,
			Expected: 442,
		},
		{
			X:        187,
			Y:          1,
			Expected: 443,
		},
		{
			X:        188,
			Y:          1,
			Expected: 444,
		},
		{
			X:        189,
			Y:          1,
			Expected: 445,
		},
		{
			X:        190,
			Y:          1,
			Expected: 446,
		},
		{
			X:        191,
			Y:          1,
			Expected: 447,
		},
		{
			X:        192,
			Y:          1,
			Expected: 448,
		},
		{
			X:        193,
			Y:          1,
			Expected: 449,
		},
		{
			X:        194,
			Y:          1,
			Expected: 450,
		},
		{
			X:        195,
			Y:          1,
			Expected: 451,
		},
		{
			X:        196,
			Y:          1,
			Expected: 452,
		},
		{
			X:        197,
			Y:          1,
			Expected: 453,
		},
		{
			X:        198,
			Y:          1,
			Expected: 454,
		},
		{
			X:        199,
			Y:          1,
			Expected: 455,
		},
		{
			X:        200,
			Y:          1,
			Expected: 456,
		},
		{
			X:        201,
			Y:          1,
			Expected: 457,
		},
		{
			X:        202,
			Y:          1,
			Expected: 458,
		},
		{
			X:        203,
			Y:          1,
			Expected: 459,
		},
		{
			X:        204,
			Y:          1,
			Expected: 460,
		},
		{
			X:        205,
			Y:          1,
			Expected: 461,
		},
		{
			X:        206,
			Y:          1,
			Expected: 462,
		},
		{
			X:        207,
			Y:          1,
			Expected: 463,
		},
		{
			X:        208,
			Y:          1,
			Expected: 464,
		},
		{
			X:        209,
			Y:          1,
			Expected: 465,
		},
		{
			X:        210,
			Y:          1,
			Expected: 466,
		},
		{
			X:        211,
			Y:          1,
			Expected: 467,
		},
		{
			X:        212,
			Y:          1,
			Expected: 468,
		},
		{
			X:        213,
			Y:          1,
			Expected: 469,
		},
		{
			X:        214,
			Y:          1,
			Expected: 470,
		},
		{
			X:        215,
			Y:          1,
			Expected: 471,
		},
		{
			X:        216,
			Y:          1,
			Expected: 472,
		},
		{
			X:        217,
			Y:          1,
			Expected: 473,
		},
		{
			X:        218,
			Y:          1,
			Expected: 474,
		},
		{
			X:        219,
			Y:          1,
			Expected: 475,
		},
		{
			X:        220,
			Y:          1,
			Expected: 476,
		},
		{
			X:        221,
			Y:          1,
			Expected: 477,
		},
		{
			X:        222,
			Y:          1,
			Expected: 478,
		},
		{
			X:        223,
			Y:          1,
			Expected: 479,
		},
		{
			X:        224,
			Y:          1,
			Expected: 480,
		},
		{
			X:        225,
			Y:          1,
			Expected: 481,
		},
		{
			X:        226,
			Y:          1,
			Expected: 482,
		},
		{
			X:        227,
			Y:          1,
			Expected: 483,
		},
		{
			X:        228,
			Y:          1,
			Expected: 484,
		},
		{
			X:        229,
			Y:          1,
			Expected: 485,
		},
		{
			X:        230,
			Y:          1,
			Expected: 486,
		},
		{
			X:        231,
			Y:          1,
			Expected: 487,
		},
		{
			X:        232,
			Y:          1,
			Expected: 488,
		},
		{
			X:        233,
			Y:          1,
			Expected: 489,
		},
		{
			X:        234,
			Y:          1,
			Expected: 490,
		},
		{
			X:        235,
			Y:          1,
			Expected: 491,
		},
		{
			X:        236,
			Y:          1,
			Expected: 492,
		},
		{
			X:        237,
			Y:          1,
			Expected: 493,
		},
		{
			X:        238,
			Y:          1,
			Expected: 494,
		},
		{
			X:        239,
			Y:          1,
			Expected: 495,
		},
		{
			X:        240,
			Y:          1,
			Expected: 496,
		},
		{
			X:        241,
			Y:          1,
			Expected: 497,
		},
		{
			X:        242,
			Y:          1,
			Expected: 498,
		},
		{
			X:        243,
			Y:          1,
			Expected: 499,
		},
		{
			X:        244,
			Y:          1,
			Expected: 500,
		},
		{
			X:        245,
			Y:          1,
			Expected: 501,
		},
		{
			X:        246,
			Y:          1,
			Expected: 502,
		},
		{
			X:        247,
			Y:          1,
			Expected: 503,
		},
		{
			X:        248,
			Y:          1,
			Expected: 504,
		},
		{
			X:        249,
			Y:          1,
			Expected: 505,
		},
		{
			X:        250,
			Y:          1,
			Expected: 506,
		},
		{
			X:        251,
			Y:          1,
			Expected: 507,
		},
		{
			X:        252,
			Y:          1,
			Expected: 508,
		},
		{
			X:        253,
			Y:          1,
			Expected: 509,
		},
		{
			X:        254,
			Y:          1,
			Expected: 510,
		},
		{
			X:        255,
			Y:          1,
			Expected: 511,
		},









		{
			X:          0,
			Y:         70,
			Expected: 17920,
		},
		{
			X:          1,
			Y:         70,
			Expected: 17921,
		},
		{
			X:          2,
			Y:         70,
			Expected: 17922,
		},
		{
			X:          3,
			Y:         70,
			Expected: 17923,
		},
		{
			X:          4,
			Y:         70,
			Expected: 17924,
		},
		{
			X:          5,
			Y:         70,
			Expected: 17925,
		},
		{
			X:          6,
			Y:         70,
			Expected: 17926,
		},
		{
			X:          7,
			Y:         70,
			Expected: 17927,
		},
		{
			X:          8,
			Y:         70,
			Expected: 17928,
		},
		{
			X:          9,
			Y:         70,
			Expected: 17929,
		},
		{
			X:         10,
			Y:         70,
			Expected: 17930,
		},
		{
			X:         11,
			Y:         70,
			Expected: 17931,
		},
		{
			X:         12,
			Y:         70,
			Expected: 17932,
		},
		{
			X:         13,
			Y:         70,
			Expected: 17933,
		},
		{
			X:         14,
			Y:         70,
			Expected: 17934,
		},
		{
			X:         15,
			Y:         70,
			Expected: 17935,
		},
		{
			X:         16,
			Y:         70,
			Expected: 17936,
		},
		{
			X:         17,
			Y:         70,
			Expected: 17937,
		},
		{
			X:         18,
			Y:         70,
			Expected: 17938,
		},
		{
			X:         19,
			Y:         70,
			Expected: 17939,
		},
		{
			X:         20,
			Y:         70,
			Expected: 17940,
		},
		{
			X:         21,
			Y:         70,
			Expected: 17941,
		},
		{
			X:         22,
			Y:         70,
			Expected: 17942,
		},
		{
			X:         23,
			Y:         70,
			Expected: 17943,
		},
		{
			X:         24,
			Y:         70,
			Expected: 17944,
		},
		{
			X:         25,
			Y:         70,
			Expected: 17945,
		},
		{
			X:         26,
			Y:         70,
			Expected: 17946,
		},
		{
			X:         27,
			Y:         70,
			Expected: 17947,
		},
		{
			X:         28,
			Y:         70,
			Expected: 17948,
		},
		{
			X:         29,
			Y:         70,
			Expected: 17949,
		},
		{
			X:         30,
			Y:         70,
			Expected: 17950,
		},
		{
			X:         31,
			Y:         70,
			Expected: 17951,
		},
		{
			X:         32,
			Y:         70,
			Expected: 17952,
		},
		{
			X:         33,
			Y:         70,
			Expected: 17953,
		},
		{
			X:         34,
			Y:         70,
			Expected: 17954,
		},
		{
			X:         35,
			Y:         70,
			Expected: 17955,
		},
		{
			X:         36,
			Y:         70,
			Expected: 17956,
		},
		{
			X:         37,
			Y:         70,
			Expected: 17957,
		},
		{
			X:         38,
			Y:         70,
			Expected: 17958,
		},
		{
			X:         39,
			Y:         70,
			Expected: 17959,
		},
		{
			X:         40,
			Y:         70,
			Expected: 17960,
		},
		{
			X:         41,
			Y:         70,
			Expected: 17961,
		},
		{
			X:         42,
			Y:         70,
			Expected: 17962,
		},
		{
			X:         43,
			Y:         70,
			Expected: 17963,
		},
		{
			X:         44,
			Y:         70,
			Expected: 17964,
		},
		{
			X:         45,
			Y:         70,
			Expected: 17965,
		},
		{
			X:         46,
			Y:         70,
			Expected: 17966,
		},
		{
			X:         47,
			Y:         70,
			Expected: 17967,
		},
		{
			X:         48,
			Y:         70,
			Expected: 17968,
		},
		{
			X:         49,
			Y:         70,
			Expected: 17969,
		},
		{
			X:         50,
			Y:         70,
			Expected: 17970,
		},
		{
			X:         51,
			Y:         70,
			Expected: 17971,
		},
		{
			X:         52,
			Y:         70,
			Expected: 17972,
		},
		{
			X:         53,
			Y:         70,
			Expected: 17973,
		},
		{
			X:         54,
			Y:         70,
			Expected: 17974,
		},
		{
			X:         55,
			Y:         70,
			Expected: 17975,
		},
		{
			X:         56,
			Y:         70,
			Expected: 17976,
		},
		{
			X:         57,
			Y:         70,
			Expected: 17977,
		},
		{
			X:         58,
			Y:         70,
			Expected: 17978,
		},
		{
			X:         59,
			Y:         70,
			Expected: 17979,
		},
		{
			X:         60,
			Y:         70,
			Expected: 17980,
		},
		{
			X:         61,
			Y:         70,
			Expected: 17981,
		},
		{
			X:         62,
			Y:         70,
			Expected: 17982,
		},
		{
			X:         63,
			Y:         70,
			Expected: 17983,
		},
		{
			X:         64,
			Y:         70,
			Expected: 17984,
		},
		{
			X:         65,
			Y:         70,
			Expected: 17985,
		},
		{
			X:         66,
			Y:         70,
			Expected: 17986,
		},
		{
			X:         67,
			Y:         70,
			Expected: 17987,
		},
		{
			X:         68,
			Y:         70,
			Expected: 17988,
		},
		{
			X:         69,
			Y:         70,
			Expected: 17989,
		},
		{
			X:         70,
			Y:         70,
			Expected: 17990,
		},
		{
			X:         71,
			Y:         70,
			Expected: 17991,
		},
		{
			X:         72,
			Y:         70,
			Expected: 17992,
		},
		{
			X:         73,
			Y:         70,
			Expected: 17993,
		},
		{
			X:         74,
			Y:         70,
			Expected: 17994,
		},
		{
			X:         75,
			Y:         70,
			Expected: 17995,
		},
		{
			X:         76,
			Y:         70,
			Expected: 17996,
		},
		{
			X:         77,
			Y:         70,
			Expected: 17997,
		},
		{
			X:         78,
			Y:         70,
			Expected: 17998,
		},
		{
			X:         79,
			Y:         70,
			Expected: 17999,
		},
		{
			X:         80,
			Y:         70,
			Expected: 18000,
		},
		{
			X:         81,
			Y:         70,
			Expected: 18001,
		},
		{
			X:         82,
			Y:         70,
			Expected: 18002,
		},
		{
			X:         83,
			Y:         70,
			Expected: 18003,
		},
		{
			X:         84,
			Y:         70,
			Expected: 18004,
		},
		{
			X:         85,
			Y:         70,
			Expected: 18005,
		},
		{
			X:         86,
			Y:         70,
			Expected: 18006,
		},
		{
			X:         87,
			Y:         70,
			Expected: 18007,
		},
		{
			X:         88,
			Y:         70,
			Expected: 18008,
		},
		{
			X:         89,
			Y:         70,
			Expected: 18009,
		},
		{
			X:         90,
			Y:         70,
			Expected: 18010,
		},
		{
			X:         91,
			Y:         70,
			Expected: 18011,
		},
		{
			X:         92,
			Y:         70,
			Expected: 18012,
		},
		{
			X:         93,
			Y:         70,
			Expected: 18013,
		},
		{
			X:         94,
			Y:         70,
			Expected: 18014,
		},
		{
			X:         95,
			Y:         70,
			Expected: 18015,
		},
		{
			X:         96,
			Y:         70,
			Expected: 18016,
		},
		{
			X:         97,
			Y:         70,
			Expected: 18017,
		},
		{
			X:         98,
			Y:         70,
			Expected: 18018,
		},
		{
			X:         99,
			Y:         70,
			Expected: 18019,
		},
		{
			X:        100,
			Y:         70,
			Expected: 18020,
		},
		{
			X:        101,
			Y:         70,
			Expected: 18021,
		},
		{
			X:        102,
			Y:         70,
			Expected: 18022,
		},
		{
			X:        103,
			Y:         70,
			Expected: 18023,
		},
		{
			X:        104,
			Y:         70,
			Expected: 18024,
		},
		{
			X:        105,
			Y:         70,
			Expected: 18025,
		},
		{
			X:        106,
			Y:         70,
			Expected: 18026,
		},
		{
			X:        107,
			Y:         70,
			Expected: 18027,
		},
		{
			X:        108,
			Y:         70,
			Expected: 18028,
		},
		{
			X:        109,
			Y:         70,
			Expected: 18029,
		},
		{
			X:        110,
			Y:         70,
			Expected: 18030,
		},
		{
			X:        111,
			Y:         70,
			Expected: 18031,
		},
		{
			X:        112,
			Y:         70,
			Expected: 18032,
		},
		{
			X:        113,
			Y:         70,
			Expected: 18033,
		},
		{
			X:        114,
			Y:         70,
			Expected: 18034,
		},
		{
			X:        115,
			Y:         70,
			Expected: 18035,
		},
		{
			X:        116,
			Y:         70,
			Expected: 18036,
		},
		{
			X:        117,
			Y:         70,
			Expected: 18037,
		},
		{
			X:        118,
			Y:         70,
			Expected: 18038,
		},
		{
			X:        119,
			Y:         70,
			Expected: 18039,
		},
		{
			X:        120,
			Y:         70,
			Expected: 18040,
		},
		{
			X:        121,
			Y:         70,
			Expected: 18041,
		},
		{
			X:        122,
			Y:         70,
			Expected: 18042,
		},
		{
			X:        123,
			Y:         70,
			Expected: 18043,
		},
		{
			X:        124,
			Y:         70,
			Expected: 18044,
		},
		{
			X:        125,
			Y:         70,
			Expected: 18045,
		},
		{
			X:        126,
			Y:         70,
			Expected: 18046,
		},
		{
			X:        127,
			Y:         70,
			Expected: 18047,
		},
		{
			X:        128,
			Y:         70,
			Expected: 18048,
		},
		{
			X:        129,
			Y:         70,
			Expected: 18049,
		},
		{
			X:        130,
			Y:         70,
			Expected: 18050,
		},
		{
			X:        131,
			Y:         70,
			Expected: 18051,
		},
		{
			X:        132,
			Y:         70,
			Expected: 18052,
		},
		{
			X:        133,
			Y:         70,
			Expected: 18053,
		},
		{
			X:        134,
			Y:         70,
			Expected: 18054,
		},
		{
			X:        135,
			Y:         70,
			Expected: 18055,
		},
		{
			X:        136,
			Y:         70,
			Expected: 18056,
		},
		{
			X:        137,
			Y:         70,
			Expected: 18057,
		},
		{
			X:        138,
			Y:         70,
			Expected: 18058,
		},
		{
			X:        139,
			Y:         70,
			Expected: 18059,
		},
		{
			X:        140,
			Y:         70,
			Expected: 18060,
		},
		{
			X:        141,
			Y:         70,
			Expected: 18061,
		},
		{
			X:        142,
			Y:         70,
			Expected: 18062,
		},
		{
			X:        143,
			Y:         70,
			Expected: 18063,
		},
		{
			X:        144,
			Y:         70,
			Expected: 18064,
		},
		{
			X:        145,
			Y:         70,
			Expected: 18065,
		},
		{
			X:        146,
			Y:         70,
			Expected: 18066,
		},
		{
			X:        147,
			Y:         70,
			Expected: 18067,
		},
		{
			X:        148,
			Y:         70,
			Expected: 18068,
		},
		{
			X:        149,
			Y:         70,
			Expected: 18069,
		},
		{
			X:        150,
			Y:         70,
			Expected: 18070,
		},
		{
			X:        151,
			Y:         70,
			Expected: 18071,
		},
		{
			X:        152,
			Y:         70,
			Expected: 18072,
		},
		{
			X:        153,
			Y:         70,
			Expected: 18073,
		},
		{
			X:        154,
			Y:         70,
			Expected: 18074,
		},
		{
			X:        155,
			Y:         70,
			Expected: 18075,
		},
		{
			X:        156,
			Y:         70,
			Expected: 18076,
		},
		{
			X:        157,
			Y:         70,
			Expected: 18077,
		},
		{
			X:        158,
			Y:         70,
			Expected: 18078,
		},
		{
			X:        159,
			Y:         70,
			Expected: 18079,
		},
		{
			X:        160,
			Y:         70,
			Expected: 18080,
		},
		{
			X:        161,
			Y:         70,
			Expected: 18081,
		},
		{
			X:        162,
			Y:         70,
			Expected: 18082,
		},
		{
			X:        163,
			Y:         70,
			Expected: 18083,
		},
		{
			X:        164,
			Y:         70,
			Expected: 18084,
		},
		{
			X:        165,
			Y:         70,
			Expected: 18085,
		},
		{
			X:        166,
			Y:         70,
			Expected: 18086,
		},
		{
			X:        167,
			Y:         70,
			Expected: 18087,
		},
		{
			X:        168,
			Y:         70,
			Expected: 18088,
		},
		{
			X:        169,
			Y:         70,
			Expected: 18089,
		},
		{
			X:        170,
			Y:         70,
			Expected: 18090,
		},
		{
			X:        171,
			Y:         70,
			Expected: 18091,
		},
		{
			X:        172,
			Y:         70,
			Expected: 18092,
		},
		{
			X:        173,
			Y:         70,
			Expected: 18093,
		},
		{
			X:        174,
			Y:         70,
			Expected: 18094,
		},
		{
			X:        175,
			Y:         70,
			Expected: 18095,
		},
		{
			X:        176,
			Y:         70,
			Expected: 18096,
		},
		{
			X:        177,
			Y:         70,
			Expected: 18097,
		},
		{
			X:        178,
			Y:         70,
			Expected: 18098,
		},
		{
			X:        179,
			Y:         70,
			Expected: 18099,
		},
		{
			X:        180,
			Y:         70,
			Expected: 18100,
		},
		{
			X:        181,
			Y:         70,
			Expected: 18101,
		},
		{
			X:        182,
			Y:         70,
			Expected: 18102,
		},
		{
			X:        183,
			Y:         70,
			Expected: 18103,
		},
		{
			X:        184,
			Y:         70,
			Expected: 18104,
		},
		{
			X:        185,
			Y:         70,
			Expected: 18105,
		},
		{
			X:        186,
			Y:         70,
			Expected: 18106,
		},
		{
			X:        187,
			Y:         70,
			Expected: 18107,
		},
		{
			X:        188,
			Y:         70,
			Expected: 18108,
		},
		{
			X:        189,
			Y:         70,
			Expected: 18109,
		},
		{
			X:        190,
			Y:         70,
			Expected: 18110,
		},
		{
			X:        191,
			Y:         70,
			Expected: 18111,
		},
		{
			X:        192,
			Y:         70,
			Expected: 18112,
		},
		{
			X:        193,
			Y:         70,
			Expected: 18113,
		},
		{
			X:        194,
			Y:         70,
			Expected: 18114,
		},
		{
			X:        195,
			Y:         70,
			Expected: 18115,
		},
		{
			X:        196,
			Y:         70,
			Expected: 18116,
		},
		{
			X:        197,
			Y:         70,
			Expected: 18117,
		},
		{
			X:        198,
			Y:         70,
			Expected: 18118,
		},
		{
			X:        199,
			Y:         70,
			Expected: 18119,
		},
		{
			X:        200,
			Y:         70,
			Expected: 18120,
		},
		{
			X:        201,
			Y:         70,
			Expected: 18121,
		},
		{
			X:        202,
			Y:         70,
			Expected: 18122,
		},
		{
			X:        203,
			Y:         70,
			Expected: 18123,
		},
		{
			X:        204,
			Y:         70,
			Expected: 18124,
		},
		{
			X:        205,
			Y:         70,
			Expected: 18125,
		},
		{
			X:        206,
			Y:         70,
			Expected: 18126,
		},
		{
			X:        207,
			Y:         70,
			Expected: 18127,
		},
		{
			X:        208,
			Y:         70,
			Expected: 18128,
		},
		{
			X:        209,
			Y:         70,
			Expected: 18129,
		},
		{
			X:        210,
			Y:         70,
			Expected: 18130,
		},
		{
			X:        211,
			Y:         70,
			Expected: 18131,
		},
		{
			X:        212,
			Y:         70,
			Expected: 18132,
		},
		{
			X:        213,
			Y:         70,
			Expected: 18133,
		},
		{
			X:        214,
			Y:         70,
			Expected: 18134,
		},
		{
			X:        215,
			Y:         70,
			Expected: 18135,
		},
		{
			X:        216,
			Y:         70,
			Expected: 18136,
		},
		{
			X:        217,
			Y:         70,
			Expected: 18137,
		},
		{
			X:        218,
			Y:         70,
			Expected: 18138,
		},
		{
			X:        219,
			Y:         70,
			Expected: 18139,
		},
		{
			X:        220,
			Y:         70,
			Expected: 18140,
		},
		{
			X:        221,
			Y:         70,
			Expected: 18141,
		},
		{
			X:        222,
			Y:         70,
			Expected: 18142,
		},
		{
			X:        223,
			Y:         70,
			Expected: 18143,
		},
		{
			X:        224,
			Y:         70,
			Expected: 18144,
		},
		{
			X:        225,
			Y:         70,
			Expected: 18145,
		},
		{
			X:        226,
			Y:         70,
			Expected: 18146,
		},
		{
			X:        227,
			Y:         70,
			Expected: 18147,
		},
		{
			X:        228,
			Y:         70,
			Expected: 18148,
		},
		{
			X:        229,
			Y:         70,
			Expected: 18149,
		},
		{
			X:        230,
			Y:         70,
			Expected: 18150,
		},
		{
			X:        231,
			Y:         70,
			Expected: 18151,
		},
		{
			X:        232,
			Y:         70,
			Expected: 18152,
		},
		{
			X:        233,
			Y:         70,
			Expected: 18153,
		},
		{
			X:        234,
			Y:         70,
			Expected: 18154,
		},
		{
			X:        235,
			Y:         70,
			Expected: 18155,
		},
		{
			X:        236,
			Y:         70,
			Expected: 18156,
		},
		{
			X:        237,
			Y:         70,
			Expected: 18157,
		},
		{
			X:        238,
			Y:         70,
			Expected: 18158,
		},
		{
			X:        239,
			Y:         70,
			Expected: 18159,
		},
		{
			X:        240,
			Y:         70,
			Expected: 18160,
		},
		{
			X:        241,
			Y:         70,
			Expected: 18161,
		},
		{
			X:        242,
			Y:         70,
			Expected: 18162,
		},
		{
			X:        243,
			Y:         70,
			Expected: 18163,
		},
		{
			X:        244,
			Y:         70,
			Expected: 18164,
		},
		{
			X:        245,
			Y:         70,
			Expected: 18165,
		},
		{
			X:        246,
			Y:         70,
			Expected: 18166,
		},
		{
			X:        247,
			Y:         70,
			Expected: 18167,
		},
		{
			X:        248,
			Y:         70,
			Expected: 18168,
		},
		{
			X:        249,
			Y:         70,
			Expected: 18169,
		},
		{
			X:        250,
			Y:         70,
			Expected: 18170,
		},
		{
			X:        251,
			Y:         70,
			Expected: 18171,
		},
		{
			X:        252,
			Y:         70,
			Expected: 18172,
		},
		{
			X:        253,
			Y:         70,
			Expected: 18173,
		},
		{
			X:        254,
			Y:         70,
			Expected: 18174,
		},
		{
			X:        255,
			Y:         70,
			Expected: 18175,
		},









		{
			X:            0,
			Y:          101,
			Expected: 25856,
		},
		{
			X:            1,
			Y:          101,
			Expected: 25857,
		},
		{
			X:            2,
			Y:          101,
			Expected: 25858,
		},
		{
			X:            3,
			Y:          101,
			Expected: 25859,
		},
		{
			X:            4,
			Y:          101,
			Expected: 25860,
		},
		{
			X:            5,
			Y:          101,
			Expected: 25861,
		},
		{
			X:            6,
			Y:          101,
			Expected: 25862,
		},
		{
			X:            7,
			Y:          101,
			Expected: 25863,
		},
		{
			X:            8,
			Y:          101,
			Expected: 25864,
		},
		{
			X:            9,
			Y:          101,
			Expected: 25865,
		},
		{
			X:           10,
			Y:          101,
			Expected: 25866,
		},
		{
			X:           11,
			Y:          101,
			Expected: 25867,
		},
		{
			X:           12,
			Y:          101,
			Expected: 25868,
		},
		{
			X:           13,
			Y:          101,
			Expected: 25869,
		},
		{
			X:           14,
			Y:          101,
			Expected: 25870,
		},
		{
			X:           15,
			Y:          101,
			Expected: 25871,
		},
		{
			X:           16,
			Y:          101,
			Expected: 25872,
		},
		{
			X:           17,
			Y:          101,
			Expected: 25873,
		},
		{
			X:           18,
			Y:          101,
			Expected: 25874,
		},
		{
			X:           19,
			Y:          101,
			Expected: 25875,
		},
		{
			X:           20,
			Y:          101,
			Expected: 25876,
		},
		{
			X:           21,
			Y:          101,
			Expected: 25877,
		},
		{
			X:           22,
			Y:          101,
			Expected: 25878,
		},
		{
			X:           23,
			Y:          101,
			Expected: 25879,
		},
		{
			X:           24,
			Y:          101,
			Expected: 25880,
		},
		{
			X:           25,
			Y:          101,
			Expected: 25881,
		},
		{
			X:           26,
			Y:          101,
			Expected: 25882,
		},
		{
			X:           27,
			Y:          101,
			Expected: 25883,
		},
		{
			X:           28,
			Y:          101,
			Expected: 25884,
		},
		{
			X:           29,
			Y:          101,
			Expected: 25885,
		},
		{
			X:           30,
			Y:          101,
			Expected: 25886,
		},
		{
			X:           31,
			Y:          101,
			Expected: 25887,
		},
		{
			X:           32,
			Y:          101,
			Expected: 25888,
		},
		{
			X:           33,
			Y:          101,
			Expected: 25889,
		},
		{
			X:           34,
			Y:          101,
			Expected: 25890,
		},
		{
			X:           35,
			Y:          101,
			Expected: 25891,
		},
		{
			X:           36,
			Y:          101,
			Expected: 25892,
		},
		{
			X:           37,
			Y:          101,
			Expected: 25893,
		},
		{
			X:           38,
			Y:          101,
			Expected: 25894,
		},
		{
			X:           39,
			Y:          101,
			Expected: 25895,
		},
		{
			X:           40,
			Y:          101,
			Expected: 25896,
		},
		{
			X:           41,
			Y:          101,
			Expected: 25897,
		},
		{
			X:           42,
			Y:          101,
			Expected: 25898,
		},
		{
			X:           43,
			Y:          101,
			Expected: 25899,
		},
		{
			X:           44,
			Y:          101,
			Expected: 25900,
		},
		{
			X:           45,
			Y:          101,
			Expected: 25901,
		},
		{
			X:           46,
			Y:          101,
			Expected: 25902,
		},
		{
			X:           47,
			Y:          101,
			Expected: 25903,
		},
		{
			X:           48,
			Y:          101,
			Expected: 25904,
		},
		{
			X:           49,
			Y:          101,
			Expected: 25905,
		},
		{
			X:           50,
			Y:          101,
			Expected: 25906,
		},
		{
			X:           51,
			Y:          101,
			Expected: 25907,
		},
		{
			X:           52,
			Y:          101,
			Expected: 25908,
		},
		{
			X:           53,
			Y:          101,
			Expected: 25909,
		},
		{
			X:           54,
			Y:          101,
			Expected: 25910,
		},
		{
			X:           55,
			Y:          101,
			Expected: 25911,
		},
		{
			X:           56,
			Y:          101,
			Expected: 25912,
		},
		{
			X:           57,
			Y:          101,
			Expected: 25913,
		},
		{
			X:           58,
			Y:          101,
			Expected: 25914,
		},
		{
			X:           59,
			Y:          101,
			Expected: 25915,
		},
		{
			X:           60,
			Y:          101,
			Expected: 25916,
		},
		{
			X:           61,
			Y:          101,
			Expected: 25917,
		},
		{
			X:           62,
			Y:          101,
			Expected: 25918,
		},
		{
			X:           63,
			Y:          101,
			Expected: 25919,
		},
		{
			X:           64,
			Y:          101,
			Expected: 25920,
		},
		{
			X:           65,
			Y:          101,
			Expected: 25921,
		},
		{
			X:           66,
			Y:          101,
			Expected: 25922,
		},
		{
			X:           67,
			Y:          101,
			Expected: 25923,
		},
		{
			X:           68,
			Y:          101,
			Expected: 25924,
		},
		{
			X:           69,
			Y:          101,
			Expected: 25925,
		},
		{
			X:           70,
			Y:          101,
			Expected: 25926,
		},
		{
			X:           71,
			Y:          101,
			Expected: 25927,
		},
		{
			X:           72,
			Y:          101,
			Expected: 25928,
		},
		{
			X:           73,
			Y:          101,
			Expected: 25929,
		},
		{
			X:           74,
			Y:          101,
			Expected: 25930,
		},
		{
			X:           75,
			Y:          101,
			Expected: 25931,
		},
		{
			X:           76,
			Y:          101,
			Expected: 25932,
		},
		{
			X:           77,
			Y:          101,
			Expected: 25933,
		},
		{
			X:           78,
			Y:          101,
			Expected: 25934,
		},
		{
			X:           79,
			Y:          101,
			Expected: 25935,
		},
		{
			X:           80,
			Y:          101,
			Expected: 25936,
		},
		{
			X:           81,
			Y:          101,
			Expected: 25937,
		},
		{
			X:           82,
			Y:          101,
			Expected: 25938,
		},
		{
			X:           83,
			Y:          101,
			Expected: 25939,
		},
		{
			X:           84,
			Y:          101,
			Expected: 25940,
		},
		{
			X:           85,
			Y:          101,
			Expected: 25941,
		},
		{
			X:           86,
			Y:          101,
			Expected: 25942,
		},
		{
			X:           87,
			Y:          101,
			Expected: 25943,
		},
		{
			X:           88,
			Y:          101,
			Expected: 25944,
		},
		{
			X:           89,
			Y:          101,
			Expected: 25945,
		},
		{
			X:           90,
			Y:          101,
			Expected: 25946,
		},
		{
			X:           91,
			Y:          101,
			Expected: 25947,
		},
		{
			X:           92,
			Y:          101,
			Expected: 25948,
		},
		{
			X:           93,
			Y:          101,
			Expected: 25949,
		},
		{
			X:           94,
			Y:          101,
			Expected: 25950,
		},
		{
			X:           95,
			Y:          101,
			Expected: 25951,
		},
		{
			X:           96,
			Y:          101,
			Expected: 25952,
		},
		{
			X:           97,
			Y:          101,
			Expected: 25953,
		},
		{
			X:           98,
			Y:          101,
			Expected: 25954,
		},
		{
			X:           99,
			Y:          101,
			Expected: 25955,
		},
		{
			X:          100,
			Y:          101,
			Expected: 25956,
		},
		{
			X:          101,
			Y:          101,
			Expected: 25957,
		},
		{
			X:          102,
			Y:          101,
			Expected: 25958,
		},
		{
			X:          103,
			Y:          101,
			Expected: 25959,
		},
		{
			X:          104,
			Y:          101,
			Expected: 25960,
		},
		{
			X:          105,
			Y:          101,
			Expected: 25961,
		},
		{
			X:          106,
			Y:          101,
			Expected: 25962,
		},
		{
			X:          107,
			Y:          101,
			Expected: 25963,
		},
		{
			X:          108,
			Y:          101,
			Expected: 25964,
		},
		{
			X:          109,
			Y:          101,
			Expected: 25965,
		},
		{
			X:          110,
			Y:          101,
			Expected: 25966,
		},
		{
			X:          111,
			Y:          101,
			Expected: 25967,
		},
		{
			X:          112,
			Y:          101,
			Expected: 25968,
		},
		{
			X:          113,
			Y:          101,
			Expected: 25969,
		},
		{
			X:          114,
			Y:          101,
			Expected: 25970,
		},
		{
			X:          115,
			Y:          101,
			Expected: 25971,
		},
		{
			X:          116,
			Y:          101,
			Expected: 25972,
		},
		{
			X:          117,
			Y:          101,
			Expected: 25973,
		},
		{
			X:          118,
			Y:          101,
			Expected: 25974,
		},
		{
			X:          119,
			Y:          101,
			Expected: 25975,
		},
		{
			X:          120,
			Y:          101,
			Expected: 25976,
		},
		{
			X:          121,
			Y:          101,
			Expected: 25977,
		},
		{
			X:          122,
			Y:          101,
			Expected: 25978,
		},
		{
			X:          123,
			Y:          101,
			Expected: 25979,
		},
		{
			X:          124,
			Y:          101,
			Expected: 25980,
		},
		{
			X:          125,
			Y:          101,
			Expected: 25981,
		},
		{
			X:          126,
			Y:          101,
			Expected: 25982,
		},
		{
			X:          127,
			Y:          101,
			Expected: 25983,
		},
		{
			X:          128,
			Y:          101,
			Expected: 25984,
		},
		{
			X:          129,
			Y:          101,
			Expected: 25985,
		},
		{
			X:          130,
			Y:          101,
			Expected: 25986,
		},
		{
			X:          131,
			Y:          101,
			Expected: 25987,
		},
		{
			X:          132,
			Y:          101,
			Expected: 25988,
		},
		{
			X:          133,
			Y:          101,
			Expected: 25989,
		},
		{
			X:          134,
			Y:          101,
			Expected: 25990,
		},
		{
			X:          135,
			Y:          101,
			Expected: 25991,
		},
		{
			X:          136,
			Y:          101,
			Expected: 25992,
		},
		{
			X:          137,
			Y:          101,
			Expected: 25993,
		},
		{
			X:          138,
			Y:          101,
			Expected: 25994,
		},
		{
			X:          139,
			Y:          101,
			Expected: 25995,
		},
		{
			X:          140,
			Y:          101,
			Expected: 25996,
		},
		{
			X:          141,
			Y:          101,
			Expected: 25997,
		},
		{
			X:          142,
			Y:          101,
			Expected: 25998,
		},
		{
			X:          143,
			Y:          101,
			Expected: 25999,
		},
		{
			X:          144,
			Y:          101,
			Expected: 26000,
		},
		{
			X:          145,
			Y:          101,
			Expected: 26001,
		},
		{
			X:          146,
			Y:          101,
			Expected: 26002,
		},
		{
			X:          147,
			Y:          101,
			Expected: 26003,
		},
		{
			X:          148,
			Y:          101,
			Expected: 26004,
		},
		{
			X:          149,
			Y:          101,
			Expected: 26005,
		},
		{
			X:          150,
			Y:          101,
			Expected: 26006,
		},
		{
			X:          151,
			Y:          101,
			Expected: 26007,
		},
		{
			X:          152,
			Y:          101,
			Expected: 26008,
		},
		{
			X:          153,
			Y:          101,
			Expected: 26009,
		},
		{
			X:          154,
			Y:          101,
			Expected: 26010,
		},
		{
			X:          155,
			Y:          101,
			Expected: 26011,
		},
		{
			X:          156,
			Y:          101,
			Expected: 26012,
		},
		{
			X:          157,
			Y:          101,
			Expected: 26013,
		},
		{
			X:          158,
			Y:          101,
			Expected: 26014,
		},
		{
			X:          159,
			Y:          101,
			Expected: 26015,
		},
		{
			X:          160,
			Y:          101,
			Expected: 26016,
		},
		{
			X:          161,
			Y:          101,
			Expected: 26017,
		},
		{
			X:          162,
			Y:          101,
			Expected: 26018,
		},
		{
			X:          163,
			Y:          101,
			Expected: 26019,
		},
		{
			X:          164,
			Y:          101,
			Expected: 26020,
		},
		{
			X:          165,
			Y:          101,
			Expected: 26021,
		},
		{
			X:          166,
			Y:          101,
			Expected: 26022,
		},
		{
			X:          167,
			Y:          101,
			Expected: 26023,
		},
		{
			X:          168,
			Y:          101,
			Expected: 26024,
		},
		{
			X:          169,
			Y:          101,
			Expected: 26025,
		},
		{
			X:          170,
			Y:          101,
			Expected: 26026,
		},
		{
			X:          171,
			Y:          101,
			Expected: 26027,
		},
		{
			X:          172,
			Y:          101,
			Expected: 26028,
		},
		{
			X:          173,
			Y:          101,
			Expected: 26029,
		},
		{
			X:          174,
			Y:          101,
			Expected: 26030,
		},
		{
			X:          175,
			Y:          101,
			Expected: 26031,
		},
		{
			X:          176,
			Y:          101,
			Expected: 26032,
		},
		{
			X:          177,
			Y:          101,
			Expected: 26033,
		},
		{
			X:          178,
			Y:          101,
			Expected: 26034,
		},
		{
			X:          179,
			Y:          101,
			Expected: 26035,
		},
		{
			X:          180,
			Y:          101,
			Expected: 26036,
		},
		{
			X:          181,
			Y:          101,
			Expected: 26037,
		},
		{
			X:          182,
			Y:          101,
			Expected: 26038,
		},
		{
			X:          183,
			Y:          101,
			Expected: 26039,
		},
		{
			X:          184,
			Y:          101,
			Expected: 26040,
		},
		{
			X:          185,
			Y:          101,
			Expected: 26041,
		},
		{
			X:          186,
			Y:          101,
			Expected: 26042,
		},
		{
			X:          187,
			Y:          101,
			Expected: 26043,
		},
		{
			X:          188,
			Y:          101,
			Expected: 26044,
		},
		{
			X:          189,
			Y:          101,
			Expected: 26045,
		},
		{
			X:          190,
			Y:          101,
			Expected: 26046,
		},
		{
			X:          191,
			Y:          101,
			Expected: 26047,
		},
		{
			X:          192,
			Y:          101,
			Expected: 26048,
		},
		{
			X:          193,
			Y:          101,
			Expected: 26049,
		},
		{
			X:          194,
			Y:          101,
			Expected: 26050,
		},
		{
			X:          195,
			Y:          101,
			Expected: 26051,
		},
		{
			X:          196,
			Y:          101,
			Expected: 26052,
		},
		{
			X:          197,
			Y:          101,
			Expected: 26053,
		},
		{
			X:          198,
			Y:          101,
			Expected: 26054,
		},
		{
			X:          199,
			Y:          101,
			Expected: 26055,
		},
		{
			X:          200,
			Y:          101,
			Expected: 26056,
		},
		{
			X:          201,
			Y:          101,
			Expected: 26057,
		},
		{
			X:          202,
			Y:          101,
			Expected: 26058,
		},
		{
			X:          203,
			Y:          101,
			Expected: 26059,
		},
		{
			X:          204,
			Y:          101,
			Expected: 26060,
		},
		{
			X:          205,
			Y:          101,
			Expected: 26061,
		},
		{
			X:          206,
			Y:          101,
			Expected: 26062,
		},
		{
			X:          207,
			Y:          101,
			Expected: 26063,
		},
		{
			X:          208,
			Y:          101,
			Expected: 26064,
		},
		{
			X:          209,
			Y:          101,
			Expected: 26065,
		},
		{
			X:          210,
			Y:          101,
			Expected: 26066,
		},
		{
			X:          211,
			Y:          101,
			Expected: 26067,
		},
		{
			X:          212,
			Y:          101,
			Expected: 26068,
		},
		{
			X:          213,
			Y:          101,
			Expected: 26069,
		},
		{
			X:          214,
			Y:          101,
			Expected: 26070,
		},
		{
			X:          215,
			Y:          101,
			Expected: 26071,
		},
		{
			X:          216,
			Y:          101,
			Expected: 26072,
		},
		{
			X:          217,
			Y:          101,
			Expected: 26073,
		},
		{
			X:          218,
			Y:          101,
			Expected: 26074,
		},
		{
			X:          219,
			Y:          101,
			Expected: 26075,
		},
		{
			X:          220,
			Y:          101,
			Expected: 26076,
		},
		{
			X:          221,
			Y:          101,
			Expected: 26077,
		},
		{
			X:          222,
			Y:          101,
			Expected: 26078,
		},
		{
			X:          223,
			Y:          101,
			Expected: 26079,
		},
		{
			X:          224,
			Y:          101,
			Expected: 26080,
		},
		{
			X:          225,
			Y:          101,
			Expected: 26081,
		},
		{
			X:          226,
			Y:          101,
			Expected: 26082,
		},
		{
			X:          227,
			Y:          101,
			Expected: 26083,
		},
		{
			X:          228,
			Y:          101,
			Expected: 26084,
		},
		{
			X:          229,
			Y:          101,
			Expected: 26085,
		},
		{
			X:          230,
			Y:          101,
			Expected: 26086,
		},
		{
			X:          231,
			Y:          101,
			Expected: 26087,
		},
		{
			X:          232,
			Y:          101,
			Expected: 26088,
		},
		{
			X:          233,
			Y:          101,
			Expected: 26089,
		},
		{
			X:          234,
			Y:          101,
			Expected: 26090,
		},
		{
			X:          235,
			Y:          101,
			Expected: 26091,
		},
		{
			X:          236,
			Y:          101,
			Expected: 26092,
		},
		{
			X:          237,
			Y:          101,
			Expected: 26093,
		},
		{
			X:          238,
			Y:          101,
			Expected: 26094,
		},
		{
			X:          239,
			Y:          101,
			Expected: 26095,
		},
		{
			X:          240,
			Y:          101,
			Expected: 26096,
		},
		{
			X:          241,
			Y:          101,
			Expected: 26097,
		},
		{
			X:          242,
			Y:          101,
			Expected: 26098,
		},
		{
			X:          243,
			Y:          101,
			Expected: 26099,
		},
		{
			X:          244,
			Y:          101,
			Expected: 26100,
		},
		{
			X:          245,
			Y:          101,
			Expected: 26101,
		},
		{
			X:          246,
			Y:          101,
			Expected: 26102,
		},
		{
			X:          247,
			Y:          101,
			Expected: 26103,
		},
		{
			X:          248,
			Y:          101,
			Expected: 26104,
		},
		{
			X:          249,
			Y:          101,
			Expected: 26105,
		},
		{
			X:          250,
			Y:          101,
			Expected: 26106,
		},
		{
			X:          251,
			Y:          101,
			Expected: 26107,
		},
		{
			X:          252,
			Y:          101,
			Expected: 26108,
		},
		{
			X:          253,
			Y:          101,
			Expected: 26109,
		},
		{
			X:          254,
			Y:          101,
			Expected: 26110,
		},
		{
			X:          255,
			Y:          101,
			Expected: 26111,
		},









		{
			X:          0,
			Y:        255,
			Expected: 65280,
		},
		{
			X:          1,
			Y:        255,
			Expected: 65281,
		},
		{
			X:          2,
			Y:        255,
			Expected: 65282,
		},
		{
			X:          3,
			Y:        255,
			Expected: 65283,
		},
		{
			X:          4,
			Y:        255,
			Expected: 65284,
		},
		{
			X:          5,
			Y:        255,
			Expected: 65285,
		},
		{
			X:          6,
			Y:        255,
			Expected: 65286,
		},
		{
			X:          7,
			Y:        255,
			Expected: 65287,
		},
		{
			X:          8,
			Y:        255,
			Expected: 65288,
		},
		{
			X:          9,
			Y:        255,
			Expected: 65289,
		},
		{
			X:         10,
			Y:        255,
			Expected: 65290,
		},
		{
			X:         11,
			Y:        255,
			Expected: 65291,
		},
		{
			X:         12,
			Y:        255,
			Expected: 65292,
		},
		{
			X:         13,
			Y:        255,
			Expected: 65293,
		},
		{
			X:         14,
			Y:        255,
			Expected: 65294,
		},
		{
			X:         15,
			Y:        255,
			Expected: 65295,
		},
		{
			X:         16,
			Y:        255,
			Expected: 65296,
		},
		{
			X:         17,
			Y:        255,
			Expected: 65297,
		},
		{
			X:         18,
			Y:        255,
			Expected: 65298,
		},
		{
			X:         19,
			Y:        255,
			Expected: 65299,
		},
		{
			X:         20,
			Y:        255,
			Expected: 65300,
		},
		{
			X:         21,
			Y:        255,
			Expected: 65301,
		},
		{
			X:         22,
			Y:        255,
			Expected: 65302,
		},
		{
			X:         23,
			Y:        255,
			Expected: 65303,
		},
		{
			X:         24,
			Y:        255,
			Expected: 65304,
		},
		{
			X:         25,
			Y:        255,
			Expected: 65305,
		},
		{
			X:         26,
			Y:        255,
			Expected: 65306,
		},
		{
			X:         27,
			Y:        255,
			Expected: 65307,
		},
		{
			X:         28,
			Y:        255,
			Expected: 65308,
		},
		{
			X:         29,
			Y:        255,
			Expected: 65309,
		},
		{
			X:         30,
			Y:        255,
			Expected: 65310,
		},
		{
			X:         31,
			Y:        255,
			Expected: 65311,
		},
		{
			X:         32,
			Y:        255,
			Expected: 65312,
		},
		{
			X:         33,
			Y:        255,
			Expected: 65313,
		},
		{
			X:         34,
			Y:        255,
			Expected: 65314,
		},
		{
			X:         35,
			Y:        255,
			Expected: 65315,
		},
		{
			X:         36,
			Y:        255,
			Expected: 65316,
		},
		{
			X:         37,
			Y:        255,
			Expected: 65317,
		},
		{
			X:         38,
			Y:        255,
			Expected: 65318,
		},
		{
			X:         39,
			Y:        255,
			Expected: 65319,
		},
		{
			X:         40,
			Y:        255,
			Expected: 65320,
		},
		{
			X:         41,
			Y:        255,
			Expected: 65321,
		},
		{
			X:         42,
			Y:        255,
			Expected: 65322,
		},
		{
			X:         43,
			Y:        255,
			Expected: 65323,
		},
		{
			X:         44,
			Y:        255,
			Expected: 65324,
		},
		{
			X:         45,
			Y:        255,
			Expected: 65325,
		},
		{
			X:         46,
			Y:        255,
			Expected: 65326,
		},
		{
			X:         47,
			Y:        255,
			Expected: 65327,
		},
		{
			X:         48,
			Y:        255,
			Expected: 65328,
		},
		{
			X:         49,
			Y:        255,
			Expected: 65329,
		},
		{
			X:         50,
			Y:        255,
			Expected: 65330,
		},
		{
			X:         51,
			Y:        255,
			Expected: 65331,
		},
		{
			X:         52,
			Y:        255,
			Expected: 65332,
		},
		{
			X:         53,
			Y:        255,
			Expected: 65333,
		},
		{
			X:         54,
			Y:        255,
			Expected: 65334,
		},
		{
			X:         55,
			Y:        255,
			Expected: 65335,
		},
		{
			X:         56,
			Y:        255,
			Expected: 65336,
		},
		{
			X:         57,
			Y:        255,
			Expected: 65337,
		},
		{
			X:         58,
			Y:        255,
			Expected: 65338,
		},
		{
			X:         59,
			Y:        255,
			Expected: 65339,
		},
		{
			X:         60,
			Y:        255,
			Expected: 65340,
		},
		{
			X:         61,
			Y:        255,
			Expected: 65341,
		},
		{
			X:         62,
			Y:        255,
			Expected: 65342,
		},
		{
			X:         63,
			Y:        255,
			Expected: 65343,
		},
		{
			X:         64,
			Y:        255,
			Expected: 65344,
		},
		{
			X:         65,
			Y:        255,
			Expected: 65345,
		},
		{
			X:         66,
			Y:        255,
			Expected: 65346,
		},
		{
			X:         67,
			Y:        255,
			Expected: 65347,
		},
		{
			X:         68,
			Y:        255,
			Expected: 65348,
		},
		{
			X:         69,
			Y:        255,
			Expected: 65349,
		},
		{
			X:         70,
			Y:        255,
			Expected: 65350,
		},
		{
			X:         71,
			Y:        255,
			Expected: 65351,
		},
		{
			X:         72,
			Y:        255,
			Expected: 65352,
		},
		{
			X:         73,
			Y:        255,
			Expected: 65353,
		},
		{
			X:         74,
			Y:        255,
			Expected: 65354,
		},
		{
			X:         75,
			Y:        255,
			Expected: 65355,
		},
		{
			X:         76,
			Y:        255,
			Expected: 65356,
		},
		{
			X:         77,
			Y:        255,
			Expected: 65357,
		},
		{
			X:         78,
			Y:        255,
			Expected: 65358,
		},
		{
			X:         79,
			Y:        255,
			Expected: 65359,
		},
		{
			X:         80,
			Y:        255,
			Expected: 65360,
		},
		{
			X:         81,
			Y:        255,
			Expected: 65361,
		},
		{
			X:         82,
			Y:        255,
			Expected: 65362,
		},
		{
			X:         83,
			Y:        255,
			Expected: 65363,
		},
		{
			X:         84,
			Y:        255,
			Expected: 65364,
		},
		{
			X:         85,
			Y:        255,
			Expected: 65365,
		},
		{
			X:         86,
			Y:        255,
			Expected: 65366,
		},
		{
			X:         87,
			Y:        255,
			Expected: 65367,
		},
		{
			X:         88,
			Y:        255,
			Expected: 65368,
		},
		{
			X:         89,
			Y:        255,
			Expected: 65369,
		},
		{
			X:         90,
			Y:        255,
			Expected: 65370,
		},
		{
			X:         91,
			Y:        255,
			Expected: 65371,
		},
		{
			X:         92,
			Y:        255,
			Expected: 65372,
		},
		{
			X:         93,
			Y:        255,
			Expected: 65373,
		},
		{
			X:         94,
			Y:        255,
			Expected: 65374,
		},
		{
			X:         95,
			Y:        255,
			Expected: 65375,
		},
		{
			X:         96,
			Y:        255,
			Expected: 65376,
		},
		{
			X:         97,
			Y:        255,
			Expected: 65377,
		},
		{
			X:         98,
			Y:        255,
			Expected: 65378,
		},
		{
			X:         99,
			Y:        255,
			Expected: 65379,
		},
		{
			X:        100,
			Y:        255,
			Expected: 65380,
		},
		{
			X:        101,
			Y:        255,
			Expected: 65381,
		},
		{
			X:        102,
			Y:        255,
			Expected: 65382,
		},
		{
			X:        103,
			Y:        255,
			Expected: 65383,
		},
		{
			X:        104,
			Y:        255,
			Expected: 65384,
		},
		{
			X:        105,
			Y:        255,
			Expected: 65385,
		},
		{
			X:        106,
			Y:        255,
			Expected: 65386,
		},
		{
			X:        107,
			Y:        255,
			Expected: 65387,
		},
		{
			X:        108,
			Y:        255,
			Expected: 65388,
		},
		{
			X:        109,
			Y:        255,
			Expected: 65389,
		},
		{
			X:        110,
			Y:        255,
			Expected: 65390,
		},
		{
			X:        111,
			Y:        255,
			Expected: 65391,
		},
		{
			X:        112,
			Y:        255,
			Expected: 65392,
		},
		{
			X:        113,
			Y:        255,
			Expected: 65393,
		},
		{
			X:        114,
			Y:        255,
			Expected: 65394,
		},
		{
			X:        115,
			Y:        255,
			Expected: 65395,
		},
		{
			X:        116,
			Y:        255,
			Expected: 65396,
		},
		{
			X:        117,
			Y:        255,
			Expected: 65397,
		},
		{
			X:        118,
			Y:        255,
			Expected: 65398,
		},
		{
			X:        119,
			Y:        255,
			Expected: 65399,
		},
		{
			X:        120,
			Y:        255,
			Expected: 65400,
		},
		{
			X:        121,
			Y:        255,
			Expected: 65401,
		},
		{
			X:        122,
			Y:        255,
			Expected: 65402,
		},
		{
			X:        123,
			Y:        255,
			Expected: 65403,
		},
		{
			X:        124,
			Y:        255,
			Expected: 65404,
		},
		{
			X:        125,
			Y:        255,
			Expected: 65405,
		},
		{
			X:        126,
			Y:        255,
			Expected: 65406,
		},
		{
			X:        127,
			Y:        255,
			Expected: 65407,
		},
		{
			X:        128,
			Y:        255,
			Expected: 65408,
		},
		{
			X:        129,
			Y:        255,
			Expected: 65409,
		},
		{
			X:        130,
			Y:        255,
			Expected: 65410,
		},
		{
			X:        131,
			Y:        255,
			Expected: 65411,
		},
		{
			X:        132,
			Y:        255,
			Expected: 65412,
		},
		{
			X:        133,
			Y:        255,
			Expected: 65413,
		},
		{
			X:        134,
			Y:        255,
			Expected: 65414,
		},
		{
			X:        135,
			Y:        255,
			Expected: 65415,
		},
		{
			X:        136,
			Y:        255,
			Expected: 65416,
		},
		{
			X:        137,
			Y:        255,
			Expected: 65417,
		},
		{
			X:        138,
			Y:        255,
			Expected: 65418,
		},
		{
			X:        139,
			Y:        255,
			Expected: 65419,
		},
		{
			X:        140,
			Y:        255,
			Expected: 65420,
		},
		{
			X:        141,
			Y:        255,
			Expected: 65421,
		},
		{
			X:        142,
			Y:        255,
			Expected: 65422,
		},
		{
			X:        143,
			Y:        255,
			Expected: 65423,
		},
		{
			X:        144,
			Y:        255,
			Expected: 65424,
		},
		{
			X:        145,
			Y:        255,
			Expected: 65425,
		},
		{
			X:        146,
			Y:        255,
			Expected: 65426,
		},
		{
			X:        147,
			Y:        255,
			Expected: 65427,
		},
		{
			X:        148,
			Y:        255,
			Expected: 65428,
		},
		{
			X:        149,
			Y:        255,
			Expected: 65429,
		},
		{
			X:        150,
			Y:        255,
			Expected: 65430,
		},
		{
			X:        151,
			Y:        255,
			Expected: 65431,
		},
		{
			X:        152,
			Y:        255,
			Expected: 65432,
		},
		{
			X:        153,
			Y:        255,
			Expected: 65433,
		},
		{
			X:        154,
			Y:        255,
			Expected: 65434,
		},
		{
			X:        155,
			Y:        255,
			Expected: 65435,
		},
		{
			X:        156,
			Y:        255,
			Expected: 65436,
		},
		{
			X:        157,
			Y:        255,
			Expected: 65437,
		},
		{
			X:        158,
			Y:        255,
			Expected: 65438,
		},
		{
			X:        159,
			Y:        255,
			Expected: 65439,
		},
		{
			X:        160,
			Y:        255,
			Expected: 65440,
		},
		{
			X:        161,
			Y:        255,
			Expected: 65441,
		},
		{
			X:        162,
			Y:        255,
			Expected: 65442,
		},
		{
			X:        163,
			Y:        255,
			Expected: 65443,
		},
		{
			X:        164,
			Y:        255,
			Expected: 65444,
		},
		{
			X:        165,
			Y:        255,
			Expected: 65445,
		},
		{
			X:        166,
			Y:        255,
			Expected: 65446,
		},
		{
			X:        167,
			Y:        255,
			Expected: 65447,
		},
		{
			X:        168,
			Y:        255,
			Expected: 65448,
		},
		{
			X:        169,
			Y:        255,
			Expected: 65449,
		},
		{
			X:        170,
			Y:        255,
			Expected: 65450,
		},
		{
			X:        171,
			Y:        255,
			Expected: 65451,
		},
		{
			X:        172,
			Y:        255,
			Expected: 65452,
		},
		{
			X:        173,
			Y:        255,
			Expected: 65453,
		},
		{
			X:        174,
			Y:        255,
			Expected: 65454,
		},
		{
			X:        175,
			Y:        255,
			Expected: 65455,
		},
		{
			X:        176,
			Y:        255,
			Expected: 65456,
		},
		{
			X:        177,
			Y:        255,
			Expected: 65457,
		},
		{
			X:        178,
			Y:        255,
			Expected: 65458,
		},
		{
			X:        179,
			Y:        255,
			Expected: 65459,
		},
		{
			X:        180,
			Y:        255,
			Expected: 65460,
		},
		{
			X:        181,
			Y:        255,
			Expected: 65461,
		},
		{
			X:        182,
			Y:        255,
			Expected: 65462,
		},
		{
			X:        183,
			Y:        255,
			Expected: 65463,
		},
		{
			X:        184,
			Y:        255,
			Expected: 65464,
		},
		{
			X:        185,
			Y:        255,
			Expected: 65465,
		},
		{
			X:        186,
			Y:        255,
			Expected: 65466,
		},
		{
			X:        187,
			Y:        255,
			Expected: 65467,
		},
		{
			X:        188,
			Y:        255,
			Expected: 65468,
		},
		{
			X:        189,
			Y:        255,
			Expected: 65469,
		},
		{
			X:        190,
			Y:        255,
			Expected: 65470,
		},
		{
			X:        191,
			Y:        255,
			Expected: 65471,
		},
		{
			X:        192,
			Y:        255,
			Expected: 65472,
		},
		{
			X:        193,
			Y:        255,
			Expected: 65473,
		},
		{
			X:        194,
			Y:        255,
			Expected: 65474,
		},
		{
			X:        195,
			Y:        255,
			Expected: 65475,
		},
		{
			X:        196,
			Y:        255,
			Expected: 65476,
		},
		{
			X:        197,
			Y:        255,
			Expected: 65477,
		},
		{
			X:        198,
			Y:        255,
			Expected: 65478,
		},
		{
			X:        199,
			Y:        255,
			Expected: 65479,
		},
		{
			X:        200,
			Y:        255,
			Expected: 65480,
		},
		{
			X:        201,
			Y:        255,
			Expected: 65481,
		},
		{
			X:        202,
			Y:        255,
			Expected: 65482,
		},
		{
			X:        203,
			Y:        255,
			Expected: 65483,
		},
		{
			X:        204,
			Y:        255,
			Expected: 65484,
		},
		{
			X:        205,
			Y:        255,
			Expected: 65485,
		},
		{
			X:        206,
			Y:        255,
			Expected: 65486,
		},
		{
			X:        207,
			Y:        255,
			Expected: 65487,
		},
		{
			X:        208,
			Y:        255,
			Expected: 65488,
		},
		{
			X:        209,
			Y:        255,
			Expected: 65489,
		},
		{
			X:        210,
			Y:        255,
			Expected: 65490,
		},
		{
			X:        211,
			Y:        255,
			Expected: 65491,
		},
		{
			X:        212,
			Y:        255,
			Expected: 65492,
		},
		{
			X:        213,
			Y:        255,
			Expected: 65493,
		},
		{
			X:        214,
			Y:        255,
			Expected: 65494,
		},
		{
			X:        215,
			Y:        255,
			Expected: 65495,
		},
		{
			X:        216,
			Y:        255,
			Expected: 65496,
		},
		{
			X:        217,
			Y:        255,
			Expected: 65497,
		},
		{
			X:        218,
			Y:        255,
			Expected: 65498,
		},
		{
			X:        219,
			Y:        255,
			Expected: 65499,
		},
		{
			X:        220,
			Y:        255,
			Expected: 65500,
		},
		{
			X:        221,
			Y:        255,
			Expected: 65501,
		},
		{
			X:        222,
			Y:        255,
			Expected: 65502,
		},
		{
			X:        223,
			Y:        255,
			Expected: 65503,
		},
		{
			X:        224,
			Y:        255,
			Expected: 65504,
		},
		{
			X:        225,
			Y:        255,
			Expected: 65505,
		},
		{
			X:        226,
			Y:        255,
			Expected: 65506,
		},
		{
			X:        227,
			Y:        255,
			Expected: 65507,
		},
		{
			X:        228,
			Y:        255,
			Expected: 65508,
		},
		{
			X:        229,
			Y:        255,
			Expected: 65509,
		},
		{
			X:        230,
			Y:        255,
			Expected: 65510,
		},
		{
			X:        231,
			Y:        255,
			Expected: 65511,
		},
		{
			X:        232,
			Y:        255,
			Expected: 65512,
		},
		{
			X:        233,
			Y:        255,
			Expected: 65513,
		},
		{
			X:        234,
			Y:        255,
			Expected: 65514,
		},
		{
			X:        235,
			Y:        255,
			Expected: 65515,
		},
		{
			X:        236,
			Y:        255,
			Expected: 65516,
		},
		{
			X:        237,
			Y:        255,
			Expected: 65517,
		},
		{
			X:        238,
			Y:        255,
			Expected: 65518,
		},
		{
			X:        239,
			Y:        255,
			Expected: 65519,
		},
		{
			X:        240,
			Y:        255,
			Expected: 65520,
		},
		{
			X:        241,
			Y:        255,
			Expected: 65521,
		},
		{
			X:        242,
			Y:        255,
			Expected: 65522,
		},
		{
			X:        243,
			Y:        255,
			Expected: 65523,
		},
		{
			X:        244,
			Y:        255,
			Expected: 65524,
		},
		{
			X:        245,
			Y:        255,
			Expected: 65525,
		},
		{
			X:        246,
			Y:        255,
			Expected: 65526,
		},
		{
			X:        247,
			Y:        255,
			Expected: 65527,
		},
		{
			X:        248,
			Y:        255,
			Expected: 65528,
		},
		{
			X:        249,
			Y:        255,
			Expected: 65529,
		},
		{
			X:        250,
			Y:        255,
			Expected: 65530,
		},
		{
			X:        251,
			Y:        255,
			Expected: 65531,
		},
		{
			X:        252,
			Y:        255,
			Expected: 65532,
		},
		{
			X:        253,
			Y:        255,
			Expected: 65533,
		},
		{
			X:        254,
			Y:        255,
			Expected: 65534,
		},
		{
			X:        255,
			Y:        255,
			Expected: 65535,
		},
	}

	for testNumber, test := range tests {

		var raster c80raster.Type

		actual := raster.PixOffset(test.X, test.Y)

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, actual pix offset was not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
