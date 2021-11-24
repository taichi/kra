// Code generated from c:\dev\repo\go\kra\parser\Named.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Copyright 2021 taichi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 20, 305,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 62, 10, 3, 12, 3, 14, 3, 65, 11, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 76, 10, 4, 12, 4,
	14, 4, 79, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 98, 10, 8, 12, 8,
	14, 8, 101, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9,
	111, 10, 9, 12, 9, 14, 9, 114, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10,
	120, 10, 10, 3, 11, 6, 11, 123, 10, 11, 13, 11, 14, 11, 124, 3, 11, 3,
	11, 7, 11, 129, 10, 11, 12, 11, 14, 11, 132, 11, 11, 5, 11, 134, 10, 11,
	3, 11, 3, 11, 6, 11, 138, 10, 11, 13, 11, 14, 11, 139, 5, 11, 142, 10,
	11, 3, 11, 3, 11, 5, 11, 146, 10, 11, 3, 11, 6, 11, 149, 10, 11, 13, 11,
	14, 11, 150, 5, 11, 153, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 6, 11, 159,
	10, 11, 13, 11, 14, 11, 160, 5, 11, 163, 10, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 7, 12, 169, 10, 12, 12, 12, 14, 12, 172, 11, 12, 3, 13, 3, 13, 6, 13,
	176, 10, 13, 13, 13, 14, 13, 177, 3, 14, 3, 14, 5, 14, 182, 10, 14, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 304, 10, 26, 4, 63, 77, 2, 27, 3, 3,
	5, 4, 7, 5, 9, 2, 11, 2, 13, 6, 15, 2, 17, 2, 19, 7, 21, 8, 23, 9, 25,
	10, 27, 2, 29, 2, 31, 2, 33, 11, 35, 12, 37, 13, 39, 14, 41, 15, 43, 16,
	45, 17, 47, 18, 49, 19, 51, 20, 3, 2, 15, 5, 2, 11, 13, 15, 15, 34, 34,
	5, 2, 50, 59, 67, 72, 99, 104, 3, 2, 50, 59, 4, 2, 75, 75, 107, 107, 4,
	2, 80, 80, 112, 112, 4, 2, 36, 36, 94, 94, 4, 2, 41, 41, 94, 94, 4, 2,
	71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 5, 2, 40, 40, 93, 93, 95, 95, 4,
	2, 63, 63, 96, 96, 5, 2, 39, 39, 47, 47, 126, 126, 4, 2, 49, 49, 128, 128,
	4, 587, 2, 67, 2, 92, 2, 99, 2, 124, 2, 172, 2, 172, 2, 183, 2, 183, 2,
	188, 2, 188, 2, 194, 2, 216, 2, 218, 2, 248, 2, 250, 2, 707, 2, 712, 2,
	723, 2, 738, 2, 742, 2, 750, 2, 750, 2, 752, 2, 752, 2, 882, 2, 886, 2,
	888, 2, 889, 2, 892, 2, 895, 2, 897, 2, 897, 2, 904, 2, 904, 2, 906, 2,
	908, 2, 910, 2, 910, 2, 912, 2, 931, 2, 933, 2, 1015, 2, 1017, 2, 1155,
	2, 1164, 2, 1329, 2, 1331, 2, 1368, 2, 1371, 2, 1371, 2, 1379, 2, 1417,
	2, 1490, 2, 1516, 2, 1522, 2, 1524, 2, 1570, 2, 1612, 2, 1648, 2, 1649,
	2, 1651, 2, 1749, 2, 1751, 2, 1751, 2, 1767, 2, 1768, 2, 1776, 2, 1777,
	2, 1788, 2, 1790, 2, 1793, 2, 1793, 2, 1810, 2, 1810, 2, 1812, 2, 1841,
	2, 1871, 2, 1959, 2, 1971, 2, 1971, 2, 1996, 2, 2028, 2, 2038, 2, 2039,
	2, 2044, 2, 2044, 2, 2050, 2, 2071, 2, 2076, 2, 2076, 2, 2086, 2, 2086,
	2, 2090, 2, 2090, 2, 2114, 2, 2138, 2, 2146, 2, 2156, 2, 2210, 2, 2230,
	2, 2232, 2, 2239, 2, 2310, 2, 2363, 2, 2367, 2, 2367, 2, 2386, 2, 2386,
	2, 2394, 2, 2403, 2, 2419, 2, 2434, 2, 2439, 2, 2446, 2, 2449, 2, 2450,
	2, 2453, 2, 2474, 2, 2476, 2, 2482, 2, 2484, 2, 2484, 2, 2488, 2, 2491,
	2, 2495, 2, 2495, 2, 2512, 2, 2512, 2, 2526, 2, 2527, 2, 2529, 2, 2531,
	2, 2546, 2, 2547, 2, 2558, 2, 2558, 2, 2567, 2, 2572, 2, 2577, 2, 2578,
	2, 2581, 2, 2602, 2, 2604, 2, 2610, 2, 2612, 2, 2613, 2, 2615, 2, 2616,
	2, 2618, 2, 2619, 2, 2651, 2, 2654, 2, 2656, 2, 2656, 2, 2676, 2, 2678,
	2, 2695, 2, 2703, 2, 2705, 2, 2707, 2, 2709, 2, 2730, 2, 2732, 2, 2738,
	2, 2740, 2, 2741, 2, 2743, 2, 2747, 2, 2751, 2, 2751, 2, 2770, 2, 2770,
	2, 2786, 2, 2787, 2, 2811, 2, 2811, 2, 2823, 2, 2830, 2, 2833, 2, 2834,
	2, 2837, 2, 2858, 2, 2860, 2, 2866, 2, 2868, 2, 2869, 2, 2871, 2, 2875,
	2, 2879, 2, 2879, 2, 2910, 2, 2911, 2, 2913, 2, 2915, 2, 2931, 2, 2931,
	2, 2949, 2, 2949, 2, 2951, 2, 2956, 2, 2960, 2, 2962, 2, 2964, 2, 2967,
	2, 2971, 2, 2972, 2, 2974, 2, 2974, 2, 2976, 2, 2977, 2, 2981, 2, 2982,
	2, 2986, 2, 2988, 2, 2992, 2, 3003, 2, 3026, 2, 3026, 2, 3079, 2, 3086,
	2, 3088, 2, 3090, 2, 3092, 2, 3114, 2, 3116, 2, 3131, 2, 3135, 2, 3135,
	2, 3162, 2, 3164, 2, 3170, 2, 3171, 2, 3202, 2, 3202, 2, 3207, 2, 3214,
	2, 3216, 2, 3218, 2, 3220, 2, 3242, 2, 3244, 2, 3253, 2, 3255, 2, 3259,
	2, 3263, 2, 3263, 2, 3296, 2, 3296, 2, 3298, 2, 3299, 2, 3315, 2, 3316,
	2, 3335, 2, 3342, 2, 3344, 2, 3346, 2, 3348, 2, 3388, 2, 3391, 2, 3391,
	2, 3408, 2, 3408, 2, 3414, 2, 3416, 2, 3425, 2, 3427, 2, 3452, 2, 3457,
	2, 3463, 2, 3480, 2, 3484, 2, 3507, 2, 3509, 2, 3517, 2, 3519, 2, 3519,
	2, 3522, 2, 3528, 2, 3587, 2, 3634, 2, 3636, 2, 3637, 2, 3650, 2, 3656,
	2, 3715, 2, 3716, 2, 3718, 2, 3718, 2, 3721, 2, 3722, 2, 3724, 2, 3724,
	2, 3727, 2, 3727, 2, 3734, 2, 3737, 2, 3739, 2, 3745, 2, 3747, 2, 3749,
	2, 3751, 2, 3751, 2, 3753, 2, 3753, 2, 3756, 2, 3757, 2, 3759, 2, 3762,
	2, 3764, 2, 3765, 2, 3775, 2, 3775, 2, 3778, 2, 3782, 2, 3784, 2, 3784,
	2, 3806, 2, 3809, 2, 3842, 2, 3842, 2, 3906, 2, 3913, 2, 3915, 2, 3950,
	2, 3978, 2, 3982, 2, 4098, 2, 4140, 2, 4161, 2, 4161, 2, 4178, 2, 4183,
	2, 4188, 2, 4191, 2, 4195, 2, 4195, 2, 4199, 2, 4200, 2, 4208, 2, 4210,
	2, 4215, 2, 4227, 2, 4240, 2, 4240, 2, 4258, 2, 4295, 2, 4297, 2, 4297,
	2, 4303, 2, 4303, 2, 4306, 2, 4348, 2, 4350, 2, 4682, 2, 4684, 2, 4687,
	2, 4690, 2, 4696, 2, 4698, 2, 4698, 2, 4700, 2, 4703, 2, 4706, 2, 4746,
	2, 4748, 2, 4751, 2, 4754, 2, 4786, 2, 4788, 2, 4791, 2, 4794, 2, 4800,
	2, 4802, 2, 4802, 2, 4804, 2, 4807, 2, 4810, 2, 4824, 2, 4826, 2, 4882,
	2, 4884, 2, 4887, 2, 4890, 2, 4956, 2, 4994, 2, 5009, 2, 5026, 2, 5111,
	2, 5114, 2, 5119, 2, 5123, 2, 5742, 2, 5745, 2, 5761, 2, 5763, 2, 5788,
	2, 5794, 2, 5868, 2, 5875, 2, 5882, 2, 5890, 2, 5902, 2, 5904, 2, 5907,
	2, 5922, 2, 5939, 2, 5954, 2, 5971, 2, 5986, 2, 5998, 2, 6000, 2, 6002,
	2, 6018, 2, 6069, 2, 6105, 2, 6105, 2, 6110, 2, 6110, 2, 6178, 2, 6265,
	2, 6274, 2, 6278, 2, 6281, 2, 6314, 2, 6316, 2, 6316, 2, 6322, 2, 6391,
	2, 6402, 2, 6432, 2, 6482, 2, 6511, 2, 6514, 2, 6518, 2, 6530, 2, 6573,
	2, 6578, 2, 6603, 2, 6658, 2, 6680, 2, 6690, 2, 6742, 2, 6825, 2, 6825,
	2, 6919, 2, 6965, 2, 6983, 2, 6989, 2, 7045, 2, 7074, 2, 7088, 2, 7089,
	2, 7100, 2, 7143, 2, 7170, 2, 7205, 2, 7247, 2, 7249, 2, 7260, 2, 7295,
	2, 7298, 2, 7306, 2, 7403, 2, 7406, 2, 7408, 2, 7411, 2, 7415, 2, 7416,
	2, 7426, 2, 7617, 2, 7682, 2, 7959, 2, 7962, 2, 7967, 2, 7970, 2, 8007,
	2, 8010, 2, 8015, 2, 8018, 2, 8025, 2, 8027, 2, 8027, 2, 8029, 2, 8029,
	2, 8031, 2, 8031, 2, 8033, 2, 8063, 2, 8066, 2, 8118, 2, 8120, 2, 8126,
	2, 8128, 2, 8128, 2, 8132, 2, 8134, 2, 8136, 2, 8142, 2, 8146, 2, 8149,
	2, 8152, 2, 8157, 2, 8162, 2, 8174, 2, 8180, 2, 8182, 2, 8184, 2, 8190,
	2, 8307, 2, 8307, 2, 8321, 2, 8321, 2, 8338, 2, 8350, 2, 8452, 2, 8452,
	2, 8457, 2, 8457, 2, 8460, 2, 8469, 2, 8471, 2, 8471, 2, 8475, 2, 8479,
	2, 8486, 2, 8486, 2, 8488, 2, 8488, 2, 8490, 2, 8490, 2, 8492, 2, 8495,
	2, 8497, 2, 8507, 2, 8510, 2, 8513, 2, 8519, 2, 8523, 2, 8528, 2, 8528,
	2, 8581, 2, 8582, 2, 11266, 2, 11312, 2, 11314, 2, 11360, 2, 11362, 2,
	11494, 2, 11501, 2, 11504, 2, 11508, 2, 11509, 2, 11522, 2, 11559, 2, 11561,
	2, 11561, 2, 11567, 2, 11567, 2, 11570, 2, 11625, 2, 11633, 2, 11633, 2,
	11650, 2, 11672, 2, 11682, 2, 11688, 2, 11690, 2, 11696, 2, 11698, 2, 11704,
	2, 11706, 2, 11712, 2, 11714, 2, 11720, 2, 11722, 2, 11728, 2, 11730, 2,
	11736, 2, 11738, 2, 11744, 2, 11825, 2, 11825, 2, 12295, 2, 12296, 2, 12339,
	2, 12343, 2, 12349, 2, 12350, 2, 12355, 2, 12440, 2, 12447, 2, 12449, 2,
	12451, 2, 12540, 2, 12542, 2, 12545, 2, 12551, 2, 12592, 2, 12595, 2, 12688,
	2, 12706, 2, 12732, 2, 12786, 2, 12801, 2, 13314, 2, 19895, 2, 19970, 2,
	40940, 2, 40962, 2, 42126, 2, 42194, 2, 42239, 2, 42242, 2, 42510, 2, 42514,
	2, 42529, 2, 42540, 2, 42541, 2, 42562, 2, 42608, 2, 42625, 2, 42655, 2,
	42658, 2, 42727, 2, 42777, 2, 42785, 2, 42788, 2, 42890, 2, 42893, 2, 42928,
	2, 42930, 2, 42937, 2, 43001, 2, 43011, 2, 43013, 2, 43015, 2, 43017, 2,
	43020, 2, 43022, 2, 43044, 2, 43074, 2, 43125, 2, 43140, 2, 43189, 2, 43252,
	2, 43257, 2, 43261, 2, 43261, 2, 43263, 2, 43263, 2, 43276, 2, 43303, 2,
	43314, 2, 43336, 2, 43362, 2, 43390, 2, 43398, 2, 43444, 2, 43473, 2, 43473,
	2, 43490, 2, 43494, 2, 43496, 2, 43505, 2, 43516, 2, 43520, 2, 43522, 2,
	43562, 2, 43586, 2, 43588, 2, 43590, 2, 43597, 2, 43618, 2, 43640, 2, 43644,
	2, 43644, 2, 43648, 2, 43697, 2, 43699, 2, 43699, 2, 43703, 2, 43704, 2,
	43707, 2, 43711, 2, 43714, 2, 43714, 2, 43716, 2, 43716, 2, 43741, 2, 43743,
	2, 43746, 2, 43756, 2, 43764, 2, 43766, 2, 43779, 2, 43784, 2, 43787, 2,
	43792, 2, 43795, 2, 43800, 2, 43810, 2, 43816, 2, 43818, 2, 43824, 2, 43826,
	2, 43868, 2, 43870, 2, 43879, 2, 43890, 2, 44004, 2, 44034, 2, 55205, 2,
	55218, 2, 55240, 2, 55245, 2, 55293, 2, 63746, 2, 64111, 2, 64114, 2, 64219,
	2, 64258, 2, 64264, 2, 64277, 2, 64281, 2, 64287, 2, 64287, 2, 64289, 2,
	64298, 2, 64300, 2, 64312, 2, 64314, 2, 64318, 2, 64320, 2, 64320, 2, 64322,
	2, 64323, 2, 64325, 2, 64326, 2, 64328, 2, 64435, 2, 64469, 2, 64831, 2,
	64850, 2, 64913, 2, 64916, 2, 64969, 2, 65010, 2, 65021, 2, 65138, 2, 65142,
	2, 65144, 2, 65278, 2, 65315, 2, 65340, 2, 65347, 2, 65372, 2, 65384, 2,
	65472, 2, 65476, 2, 65481, 2, 65484, 2, 65489, 2, 65492, 2, 65497, 2, 65500,
	2, 65502, 2, 2, 3, 13, 3, 15, 3, 40, 3, 42, 3, 60, 3, 62, 3, 63, 3, 65,
	3, 79, 3, 82, 3, 95, 3, 130, 3, 252, 3, 642, 3, 670, 3, 674, 3, 722, 3,
	770, 3, 801, 3, 815, 3, 834, 3, 836, 3, 843, 3, 850, 3, 887, 3, 898, 3,
	927, 3, 930, 3, 965, 3, 970, 3, 977, 3, 1026, 3, 1183, 3, 1202, 3, 1237,
	3, 1242, 3, 1277, 3, 1282, 3, 1321, 3, 1330, 3, 1381, 3, 1538, 3, 1848,
	3, 1858, 3, 1879, 3, 1890, 3, 1897, 3, 2050, 3, 2055, 3, 2058, 3, 2058,
	3, 2060, 3, 2103, 3, 2105, 3, 2106, 3, 2110, 3, 2110, 3, 2113, 3, 2135,
	3, 2146, 3, 2168, 3, 2178, 3, 2208, 3, 2274, 3, 2292, 3, 2294, 3, 2295,
	3, 2306, 3, 2327, 3, 2338, 3, 2363, 3, 2434, 3, 2489, 3, 2496, 3, 2497,
	3, 2562, 3, 2562, 3, 2578, 3, 2581, 3, 2583, 3, 2585, 3, 2587, 3, 2613,
	3, 2658, 3, 2686, 3, 2690, 3, 2718, 3, 2754, 3, 2761, 3, 2763, 3, 2790,
	3, 2818, 3, 2871, 3, 2882, 3, 2903, 3, 2914, 3, 2932, 3, 2946, 3, 2963,
	3, 3074, 3, 3146, 3, 3202, 3, 3252, 3, 3266, 3, 3316, 3, 4101, 3, 4153,
	3, 4229, 3, 4273, 3, 4306, 3, 4330, 3, 4357, 3, 4392, 3, 4434, 3, 4468,
	3, 4472, 3, 4472, 3, 4485, 3, 4532, 3, 4547, 3, 4550, 3, 4572, 3, 4572,
	3, 4574, 3, 4574, 3, 4610, 3, 4627, 3, 4629, 3, 4653, 3, 4738, 3, 4744,
	3, 4746, 3, 4746, 3, 4748, 3, 4751, 3, 4753, 3, 4767, 3, 4769, 3, 4778,
	3, 4786, 3, 4832, 3, 4871, 3, 4878, 3, 4881, 3, 4882, 3, 4885, 3, 4906,
	3, 4908, 3, 4914, 3, 4916, 3, 4917, 3, 4919, 3, 4923, 3, 4927, 3, 4927,
	3, 4946, 3, 4946, 3, 4959, 3, 4963, 3, 5122, 3, 5174, 3, 5193, 3, 5196,
	3, 5250, 3, 5297, 3, 5318, 3, 5319, 3, 5321, 3, 5321, 3, 5506, 3, 5552,
	3, 5594, 3, 5597, 3, 5634, 3, 5681, 3, 5702, 3, 5702, 3, 5762, 3, 5804,
	3, 5890, 3, 5915, 3, 6306, 3, 6369, 3, 6401, 3, 6401, 3, 6658, 3, 6658,
	3, 6669, 3, 6708, 3, 6716, 3, 6716, 3, 6738, 3, 6738, 3, 6750, 3, 6789,
	3, 6792, 3, 6795, 3, 6850, 3, 6906, 3, 7170, 3, 7178, 3, 7180, 3, 7216,
	3, 7234, 3, 7234, 3, 7284, 3, 7313, 3, 7426, 3, 7432, 3, 7434, 3, 7435,
	3, 7437, 3, 7474, 3, 7496, 3, 7496, 3, 8194, 3, 9115, 3, 9346, 3, 9541,
	3, 12290, 3, 13360, 3, 17410, 3, 17992, 3, 26626, 3, 27194, 3, 27202, 3,
	27232, 3, 27346, 3, 27375, 3, 27394, 3, 27441, 3, 27458, 3, 27461, 3, 27493,
	3, 27513, 3, 27519, 3, 27537, 3, 28418, 3, 28486, 3, 28498, 3, 28498, 3,
	28565, 3, 28577, 3, 28642, 3, 28643, 3, 28674, 3, 34798, 3, 34818, 3, 35572,
	3, 45058, 3, 45344, 3, 45426, 3, 45821, 3, 48130, 3, 48236, 3, 48242, 3,
	48254, 3, 48258, 3, 48266, 3, 48274, 3, 48283, 3, 54274, 3, 54358, 3, 54360,
	3, 54430, 3, 54432, 3, 54433, 3, 54436, 3, 54436, 3, 54439, 3, 54440, 3,
	54443, 3, 54446, 3, 54448, 3, 54459, 3, 54461, 3, 54461, 3, 54463, 3, 54469,
	3, 54471, 3, 54535, 3, 54537, 3, 54540, 3, 54543, 3, 54550, 3, 54552, 3,
	54558, 3, 54560, 3, 54587, 3, 54589, 3, 54592, 3, 54594, 3, 54598, 3, 54600,
	3, 54600, 3, 54604, 3, 54610, 3, 54612, 3, 54951, 3, 54954, 3, 54978, 3,
	54980, 3, 55004, 3, 55006, 3, 55036, 3, 55038, 3, 55062, 3, 55064, 3, 55094,
	3, 55096, 3, 55120, 3, 55122, 3, 55152, 3, 55154, 3, 55178, 3, 55180, 3,
	55210, 3, 55212, 3, 55236, 3, 55238, 3, 55245, 3, 59394, 3, 59590, 3, 59650,
	3, 59717, 3, 60930, 3, 60933, 3, 60935, 3, 60961, 3, 60963, 3, 60964, 3,
	60966, 3, 60966, 3, 60969, 3, 60969, 3, 60971, 3, 60980, 3, 60982, 3, 60985,
	3, 60987, 3, 60987, 3, 60989, 3, 60989, 3, 60996, 3, 60996, 3, 61001, 3,
	61001, 3, 61003, 3, 61003, 3, 61005, 3, 61005, 3, 61007, 3, 61009, 3, 61011,
	3, 61012, 3, 61014, 3, 61014, 3, 61017, 3, 61017, 3, 61019, 3, 61019, 3,
	61021, 3, 61021, 3, 61023, 3, 61023, 3, 61025, 3, 61025, 3, 61027, 3, 61028,
	3, 61030, 3, 61030, 3, 61033, 3, 61036, 3, 61038, 3, 61044, 3, 61046, 3,
	61049, 3, 61051, 3, 61054, 3, 61056, 3, 61056, 3, 61058, 3, 61067, 3, 61069,
	3, 61085, 3, 61091, 3, 61093, 3, 61095, 3, 61099, 3, 61101, 3, 61117, 3,
	2, 4, 42712, 4, 42754, 4, 46902, 4, 46914, 4, 47135, 4, 47138, 4, 52899,
	4, 52914, 4, 60386, 4, 63490, 4, 64031, 4, 57, 2, 50, 2, 59, 2, 1634, 2,
	1643, 2, 1778, 2, 1787, 2, 1986, 2, 1995, 2, 2408, 2, 2417, 2, 2536, 2,
	2545, 2, 2664, 2, 2673, 2, 2792, 2, 2801, 2, 2920, 2, 2929, 2, 3048, 2,
	3057, 2, 3176, 2, 3185, 2, 3304, 2, 3313, 2, 3432, 2, 3441, 2, 3560, 2,
	3569, 2, 3666, 2, 3675, 2, 3794, 2, 3803, 2, 3874, 2, 3883, 2, 4162, 2,
	4171, 2, 4242, 2, 4251, 2, 6114, 2, 6123, 2, 6162, 2, 6171, 2, 6472, 2,
	6481, 2, 6610, 2, 6619, 2, 6786, 2, 6795, 2, 6802, 2, 6811, 2, 6994, 2,
	7003, 2, 7090, 2, 7099, 2, 7234, 2, 7243, 2, 7250, 2, 7259, 2, 42530, 2,
	42539, 2, 43218, 2, 43227, 2, 43266, 2, 43275, 2, 43474, 2, 43483, 2, 43506,
	2, 43515, 2, 43602, 2, 43611, 2, 44018, 2, 44027, 2, 65298, 2, 65307, 2,
	1186, 3, 1195, 3, 4200, 3, 4209, 3, 4338, 3, 4347, 3, 4408, 3, 4417, 3,
	4562, 3, 4571, 3, 4850, 3, 4859, 3, 5202, 3, 5211, 3, 5330, 3, 5339, 3,
	5714, 3, 5723, 3, 5826, 3, 5835, 3, 5938, 3, 5947, 3, 6370, 3, 6379, 3,
	7250, 3, 7259, 3, 7506, 3, 7515, 3, 27234, 3, 27243, 3, 27474, 3, 27483,
	3, 55248, 3, 55297, 3, 59730, 3, 59739, 3, 367, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2,
	2, 2, 2, 51, 3, 2, 2, 2, 3, 53, 3, 2, 2, 2, 5, 57, 3, 2, 2, 2, 7, 71, 3,
	2, 2, 2, 9, 84, 3, 2, 2, 2, 11, 86, 3, 2, 2, 2, 13, 88, 3, 2, 2, 2, 15,
	91, 3, 2, 2, 2, 17, 104, 3, 2, 2, 2, 19, 119, 3, 2, 2, 2, 21, 162, 3, 2,
	2, 2, 23, 164, 3, 2, 2, 2, 25, 173, 3, 2, 2, 2, 27, 181, 3, 2, 2, 2, 29,
	183, 3, 2, 2, 2, 31, 185, 3, 2, 2, 2, 33, 187, 3, 2, 2, 2, 35, 189, 3,
	2, 2, 2, 37, 191, 3, 2, 2, 2, 39, 193, 3, 2, 2, 2, 41, 195, 3, 2, 2, 2,
	43, 197, 3, 2, 2, 2, 45, 199, 3, 2, 2, 2, 47, 201, 3, 2, 2, 2, 49, 203,
	3, 2, 2, 2, 51, 303, 3, 2, 2, 2, 53, 54, 9, 2, 2, 2, 54, 55, 3, 2, 2, 2,
	55, 56, 8, 2, 2, 2, 56, 4, 3, 2, 2, 2, 57, 58, 7, 49, 2, 2, 58, 59, 7,
	44, 2, 2, 59, 63, 3, 2, 2, 2, 60, 62, 11, 2, 2, 2, 61, 60, 3, 2, 2, 2,
	62, 65, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 66, 3,
	2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 67, 7, 44, 2, 2, 67, 68, 7, 49, 2, 2,
	68, 69, 3, 2, 2, 2, 69, 70, 8, 3, 2, 2, 70, 6, 3, 2, 2, 2, 71, 72, 7, 47,
	2, 2, 72, 73, 7, 47, 2, 2, 73, 77, 3, 2, 2, 2, 74, 76, 11, 2, 2, 2, 75,
	74, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 77, 75, 3, 2, 2,
	2, 78, 80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 7, 12, 2, 2, 81, 82,
	3, 2, 2, 2, 82, 83, 8, 4, 2, 2, 83, 8, 3, 2, 2, 2, 84, 85, 9, 3, 2, 2,
	85, 10, 3, 2, 2, 2, 86, 87, 9, 4, 2, 2, 87, 12, 3, 2, 2, 2, 88, 89, 9,
	5, 2, 2, 89, 90, 9, 6, 2, 2, 90, 14, 3, 2, 2, 2, 91, 99, 7, 36, 2, 2, 92,
	93, 7, 94, 2, 2, 93, 98, 11, 2, 2, 2, 94, 95, 7, 36, 2, 2, 95, 98, 7, 36,
	2, 2, 96, 98, 10, 7, 2, 2, 97, 92, 3, 2, 2, 2, 97, 94, 3, 2, 2, 2, 97,
	96, 3, 2, 2, 2, 98, 101, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2,
	2, 2, 100, 102, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 102, 103, 7, 36, 2, 2,
	103, 16, 3, 2, 2, 2, 104, 112, 7, 41, 2, 2, 105, 106, 7, 94, 2, 2, 106,
	111, 11, 2, 2, 2, 107, 108, 7, 41, 2, 2, 108, 111, 7, 41, 2, 2, 109, 111,
	10, 8, 2, 2, 110, 105, 3, 2, 2, 2, 110, 107, 3, 2, 2, 2, 110, 109, 3, 2,
	2, 2, 111, 114, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2,
	113, 115, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 116, 7, 41, 2, 2, 116,
	18, 3, 2, 2, 2, 117, 120, 5, 15, 8, 2, 118, 120, 5, 17, 9, 2, 119, 117,
	3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 20, 3, 2, 2, 2, 121, 123, 5, 11,
	6, 2, 122, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2,
	124, 125, 3, 2, 2, 2, 125, 133, 3, 2, 2, 2, 126, 130, 5, 47, 24, 2, 127,
	129, 5, 11, 6, 2, 128, 127, 3, 2, 2, 2, 129, 132, 3, 2, 2, 2, 130, 128,
	3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2,
	2, 2, 133, 126, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 142, 3, 2, 2, 2,
	135, 137, 5, 47, 24, 2, 136, 138, 5, 11, 6, 2, 137, 136, 3, 2, 2, 2, 138,
	139, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 142,
	3, 2, 2, 2, 141, 122, 3, 2, 2, 2, 141, 135, 3, 2, 2, 2, 142, 152, 3, 2,
	2, 2, 143, 145, 9, 9, 2, 2, 144, 146, 9, 10, 2, 2, 145, 144, 3, 2, 2, 2,
	145, 146, 3, 2, 2, 2, 146, 148, 3, 2, 2, 2, 147, 149, 5, 11, 6, 2, 148,
	147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150, 151,
	3, 2, 2, 2, 151, 153, 3, 2, 2, 2, 152, 143, 3, 2, 2, 2, 152, 153, 3, 2,
	2, 2, 153, 163, 3, 2, 2, 2, 154, 155, 7, 50, 2, 2, 155, 156, 7, 122, 2,
	2, 156, 158, 3, 2, 2, 2, 157, 159, 5, 9, 5, 2, 158, 157, 3, 2, 2, 2, 159,
	160, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 163,
	3, 2, 2, 2, 162, 141, 3, 2, 2, 2, 162, 154, 3, 2, 2, 2, 163, 22, 3, 2,
	2, 2, 164, 170, 5, 27, 14, 2, 165, 169, 5, 27, 14, 2, 166, 169, 5, 31,
	16, 2, 167, 169, 7, 38, 2, 2, 168, 165, 3, 2, 2, 2, 168, 166, 3, 2, 2,
	2, 168, 167, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170,
	171, 3, 2, 2, 2, 171, 24, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 175, 7,
	38, 2, 2, 174, 176, 5, 11, 6, 2, 175, 174, 3, 2, 2, 2, 176, 177, 3, 2,
	2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 26, 3, 2, 2, 2,
	179, 182, 5, 29, 15, 2, 180, 182, 7, 97, 2, 2, 181, 179, 3, 2, 2, 2, 181,
	180, 3, 2, 2, 2, 182, 28, 3, 2, 2, 2, 183, 184, 9, 15, 2, 2, 184, 30, 3,
	2, 2, 2, 185, 186, 9, 16, 2, 2, 186, 32, 3, 2, 2, 2, 187, 188, 7, 42, 2,
	2, 188, 34, 3, 2, 2, 2, 189, 190, 7, 43, 2, 2, 190, 36, 3, 2, 2, 2, 191,
	192, 7, 65, 2, 2, 192, 38, 3, 2, 2, 2, 193, 194, 7, 46, 2, 2, 194, 40,
	3, 2, 2, 2, 195, 196, 7, 66, 2, 2, 196, 42, 3, 2, 2, 2, 197, 198, 7, 60,
	2, 2, 198, 44, 3, 2, 2, 2, 199, 200, 7, 61, 2, 2, 200, 46, 3, 2, 2, 2,
	201, 202, 7, 48, 2, 2, 202, 48, 3, 2, 2, 2, 203, 204, 7, 44, 2, 2, 204,
	50, 3, 2, 2, 2, 205, 304, 9, 11, 2, 2, 206, 207, 7, 40, 2, 2, 207, 304,
	7, 40, 2, 2, 208, 209, 7, 40, 2, 2, 209, 304, 7, 62, 2, 2, 210, 211, 7,
	66, 2, 2, 211, 304, 7, 66, 2, 2, 212, 213, 7, 66, 2, 2, 213, 304, 7, 64,
	2, 2, 214, 304, 7, 35, 2, 2, 215, 216, 7, 35, 2, 2, 216, 304, 7, 35, 2,
	2, 217, 218, 7, 35, 2, 2, 218, 304, 7, 63, 2, 2, 219, 304, 9, 12, 2, 2,
	220, 221, 7, 63, 2, 2, 221, 304, 7, 64, 2, 2, 222, 304, 7, 64, 2, 2, 223,
	224, 7, 64, 2, 2, 224, 304, 7, 63, 2, 2, 225, 226, 7, 64, 2, 2, 226, 304,
	7, 64, 2, 2, 227, 304, 7, 37, 2, 2, 228, 229, 7, 37, 2, 2, 229, 304, 7,
	63, 2, 2, 230, 231, 7, 37, 2, 2, 231, 304, 7, 64, 2, 2, 232, 233, 7, 37,
	2, 2, 233, 234, 7, 64, 2, 2, 234, 304, 7, 64, 2, 2, 235, 236, 7, 37, 2,
	2, 236, 304, 7, 37, 2, 2, 237, 238, 7, 47, 2, 2, 238, 304, 7, 64, 2, 2,
	239, 240, 7, 47, 2, 2, 240, 241, 7, 64, 2, 2, 241, 304, 7, 64, 2, 2, 242,
	243, 7, 47, 2, 2, 243, 244, 7, 126, 2, 2, 244, 304, 7, 47, 2, 2, 245, 304,
	7, 62, 2, 2, 246, 247, 7, 62, 2, 2, 247, 304, 7, 63, 2, 2, 248, 249, 7,
	62, 2, 2, 249, 304, 7, 66, 2, 2, 250, 251, 7, 62, 2, 2, 251, 304, 7, 96,
	2, 2, 252, 253, 7, 62, 2, 2, 253, 304, 7, 64, 2, 2, 254, 255, 7, 62, 2,
	2, 255, 256, 7, 47, 2, 2, 256, 304, 7, 64, 2, 2, 257, 258, 7, 62, 2, 2,
	258, 304, 7, 62, 2, 2, 259, 260, 7, 62, 2, 2, 260, 261, 7, 62, 2, 2, 261,
	304, 7, 63, 2, 2, 262, 263, 7, 62, 2, 2, 263, 264, 7, 65, 2, 2, 264, 304,
	7, 64, 2, 2, 265, 304, 9, 13, 2, 2, 266, 267, 7, 126, 2, 2, 267, 304, 7,
	126, 2, 2, 268, 269, 7, 126, 2, 2, 269, 270, 7, 126, 2, 2, 270, 304, 7,
	49, 2, 2, 271, 272, 7, 126, 2, 2, 272, 304, 7, 49, 2, 2, 273, 304, 7, 45,
	2, 2, 274, 275, 7, 65, 2, 2, 275, 304, 7, 40, 2, 2, 276, 277, 7, 65, 2,
	2, 277, 304, 7, 37, 2, 2, 278, 279, 7, 65, 2, 2, 279, 304, 7, 47, 2, 2,
	280, 281, 7, 65, 2, 2, 281, 304, 7, 126, 2, 2, 282, 304, 9, 14, 2, 2, 283,
	284, 7, 128, 2, 2, 284, 304, 7, 63, 2, 2, 285, 286, 7, 128, 2, 2, 286,
	287, 7, 64, 2, 2, 287, 288, 7, 63, 2, 2, 288, 304, 7, 128, 2, 2, 289, 290,
	7, 128, 2, 2, 290, 291, 7, 64, 2, 2, 291, 304, 7, 128, 2, 2, 292, 293,
	7, 128, 2, 2, 293, 294, 7, 62, 2, 2, 294, 295, 7, 63, 2, 2, 295, 304, 7,
	128, 2, 2, 296, 297, 7, 128, 2, 2, 297, 298, 7, 62, 2, 2, 298, 304, 7,
	128, 2, 2, 299, 300, 7, 128, 2, 2, 300, 304, 7, 44, 2, 2, 301, 302, 7,
	128, 2, 2, 302, 304, 7, 128, 2, 2, 303, 205, 3, 2, 2, 2, 303, 206, 3, 2,
	2, 2, 303, 208, 3, 2, 2, 2, 303, 210, 3, 2, 2, 2, 303, 212, 3, 2, 2, 2,
	303, 214, 3, 2, 2, 2, 303, 215, 3, 2, 2, 2, 303, 217, 3, 2, 2, 2, 303,
	219, 3, 2, 2, 2, 303, 220, 3, 2, 2, 2, 303, 222, 3, 2, 2, 2, 303, 223,
	3, 2, 2, 2, 303, 225, 3, 2, 2, 2, 303, 227, 3, 2, 2, 2, 303, 228, 3, 2,
	2, 2, 303, 230, 3, 2, 2, 2, 303, 232, 3, 2, 2, 2, 303, 235, 3, 2, 2, 2,
	303, 237, 3, 2, 2, 2, 303, 239, 3, 2, 2, 2, 303, 242, 3, 2, 2, 2, 303,
	245, 3, 2, 2, 2, 303, 246, 3, 2, 2, 2, 303, 248, 3, 2, 2, 2, 303, 250,
	3, 2, 2, 2, 303, 252, 3, 2, 2, 2, 303, 254, 3, 2, 2, 2, 303, 257, 3, 2,
	2, 2, 303, 259, 3, 2, 2, 2, 303, 262, 3, 2, 2, 2, 303, 265, 3, 2, 2, 2,
	303, 266, 3, 2, 2, 2, 303, 268, 3, 2, 2, 2, 303, 271, 3, 2, 2, 2, 303,
	273, 3, 2, 2, 2, 303, 274, 3, 2, 2, 2, 303, 276, 3, 2, 2, 2, 303, 278,
	3, 2, 2, 2, 303, 280, 3, 2, 2, 2, 303, 282, 3, 2, 2, 2, 303, 283, 3, 2,
	2, 2, 303, 285, 3, 2, 2, 2, 303, 289, 3, 2, 2, 2, 303, 292, 3, 2, 2, 2,
	303, 296, 3, 2, 2, 2, 303, 299, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 304,
	52, 3, 2, 2, 2, 25, 2, 63, 77, 97, 99, 110, 112, 119, 124, 130, 133, 139,
	141, 145, 150, 152, 160, 162, 168, 170, 177, 181, 303, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "", "", "", "", "'('", "')'", "'?'", "','", "'@'",
	"':'", "';'", "'.'", "'*'",
}

var lexerSymbolicNames = []string{
	"", "SPACES", "BLOCK_COMMENT", "LINE_COMMENT", "IN", "STRING", "NUMBER",
	"IDENTIFIER", "DECPARAM", "OPEN_PAREN", "CLOSE_PAREN", "QMARK", "COMMA",
	"AT", "COLON", "SEMI", "DOT", "STAR", "ANY_SYMBOL",
}

var lexerRuleNames = []string{
	"SPACES", "BLOCK_COMMENT", "LINE_COMMENT", "HEX_DIGIT", "DIGIT", "IN",
	"DQUOTA_STRING", "SQUOTA_STRING", "STRING", "NUMBER", "IDENTIFIER", "DECPARAM",
	"LETTER", "UNICODE_LETTER", "UNICODE_DIGIT", "OPEN_PAREN", "CLOSE_PAREN",
	"QMARK", "COMMA", "AT", "COLON", "SEMI", "DOT", "STAR", "ANY_SYMBOL",
}

type NamedLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewNamedLexer(input antlr.CharStream) *NamedLexer {

	l := new(NamedLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Named.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NamedLexer tokens.
const (
	NamedLexerSPACES        = 1
	NamedLexerBLOCK_COMMENT = 2
	NamedLexerLINE_COMMENT  = 3
	NamedLexerIN            = 4
	NamedLexerSTRING        = 5
	NamedLexerNUMBER        = 6
	NamedLexerIDENTIFIER    = 7
	NamedLexerDECPARAM      = 8
	NamedLexerOPEN_PAREN    = 9
	NamedLexerCLOSE_PAREN   = 10
	NamedLexerQMARK         = 11
	NamedLexerCOMMA         = 12
	NamedLexerAT            = 13
	NamedLexerCOLON         = 14
	NamedLexerSEMI          = 15
	NamedLexerDOT           = 16
	NamedLexerSTAR          = 17
	NamedLexerANY_SYMBOL    = 18
)
