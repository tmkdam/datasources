package lib

var coreSchema =
`Data Set Name: MD_SIDC_2011_CORE               
Number of Observations: 721317
Total Record Length: 857
Total Number of Variables: 249


Columns   Description:
=======   ============
 1-  3    Database name
 5-  8    Discharge year of data
10- 25    File name
27- 29    Variable number
31- 56    Variable name
58- 61    Starting column of variable in ASCII file
63- 66    Ending column of variable in ASCII file
    68    Non-zero number of digits after decimal point for numeric variable
70- 73    Variable type (Num=numeric; Char=character)
75-174    Variable label


MD  2011 CORE               1 AGE                           1    3   Num  Age in years at admission
MD  2011 CORE               2 AGEDAY                        4    6   Num  Age in days (when age < 1 year)
MD  2011 CORE               3 AGEMONTH                      7    9   Num  Age in months (when age < 11 years)
MD  2011 CORE               4 AMONTH                       10   11   Num  Admission month
MD  2011 CORE               5 ASOURCE                      12   13   Num  Admission source (uniform)
MD  2011 CORE               6 ASOURCE_X                    14   15   Char Admission source (as received from source)
MD  2011 CORE               7 ATYPE                        16   17   Num  Admission type
MD  2011 CORE               8 AWEEKEND                     18   19   Num  Admission day is a weekend
MD  2011 CORE               9 BWT                          20   23   Num  Birth weight in grams
MD  2011 CORE              10 DIED                         24   25   Num  Died during hospitalization
MD  2011 CORE              11 DISPUNIFORM                  26   27   Num  Disposition of patient (uniform)
MD  2011 CORE              12 DISP_X                       28   29   Char Disposition of patient (as received from source)
MD  2011 CORE              13 DQTR                         30   31   Num  Discharge quarter
MD  2011 CORE              14 DRG                          32   34   Num  DRG in effect on discharge date
MD  2011 CORE              15 DRG24                        35   37   Num  DRG, version 24
MD  2011 CORE              16 DRGVER                       38   39   Num  DRG grouper version used on discharge date
MD  2011 CORE              17 DRG_NoPOA                    40   42   Num  DRG in use on discharge date, calculated without POA
MD  2011 CORE              18 DSHOSPID                     43   59   Char Data source hospital identifier
MD  2011 CORE              19 DX1                          60   64   Char Diagnosis 1
MD  2011 CORE              20 DX2                          65   69   Char Diagnosis 2
MD  2011 CORE              21 DX3                          70   74   Char Diagnosis 3
MD  2011 CORE              22 DX4                          75   79   Char Diagnosis 4
MD  2011 CORE              23 DX5                          80   84   Char Diagnosis 5
MD  2011 CORE              24 DX6                          85   89   Char Diagnosis 6
MD  2011 CORE              25 DX7                          90   94   Char Diagnosis 7
MD  2011 CORE              26 DX8                          95   99   Char Diagnosis 8
MD  2011 CORE              27 DX9                         100  104   Char Diagnosis 9
MD  2011 CORE              28 DX10                        105  109   Char Diagnosis 10
MD  2011 CORE              29 DX11                        110  114   Char Diagnosis 11
MD  2011 CORE              30 DX12                        115  119   Char Diagnosis 12
MD  2011 CORE              31 DX13                        120  124   Char Diagnosis 13
MD  2011 CORE              32 DX14                        125  129   Char Diagnosis 14
MD  2011 CORE              33 DX15                        130  134   Char Diagnosis 15
MD  2011 CORE              34 DX16                        135  139   Char Diagnosis 16
MD  2011 CORE              35 DX17                        140  144   Char Diagnosis 17
MD  2011 CORE              36 DX18                        145  149   Char Diagnosis 18
MD  2011 CORE              37 DX19                        150  154   Char Diagnosis 19
MD  2011 CORE              38 DX20                        155  159   Char Diagnosis 20
MD  2011 CORE              39 DX21                        160  164   Char Diagnosis 21
MD  2011 CORE              40 DX22                        165  169   Char Diagnosis 22
MD  2011 CORE              41 DX23                        170  174   Char Diagnosis 23
MD  2011 CORE              42 DX24                        175  179   Char Diagnosis 24
MD  2011 CORE              43 DX25                        180  184   Char Diagnosis 25
MD  2011 CORE              44 DX26                        185  189   Char Diagnosis 26
MD  2011 CORE              45 DX27                        190  194   Char Diagnosis 27
MD  2011 CORE              46 DX28                        195  199   Char Diagnosis 28
MD  2011 CORE              47 DX29                        200  204   Char Diagnosis 29
MD  2011 CORE              48 DX30                        205  209   Char Diagnosis 30
MD  2011 CORE              49 DXCCS1                      210  213   Num  CCS: diagnosis 1
MD  2011 CORE              50 DXCCS2                      214  217   Num  CCS: diagnosis 2
MD  2011 CORE              51 DXCCS3                      218  221   Num  CCS: diagnosis 3
MD  2011 CORE              52 DXCCS4                      222  225   Num  CCS: diagnosis 4
MD  2011 CORE              53 DXCCS5                      226  229   Num  CCS: diagnosis 5
MD  2011 CORE              54 DXCCS6                      230  233   Num  CCS: diagnosis 6
MD  2011 CORE              55 DXCCS7                      234  237   Num  CCS: diagnosis 7
MD  2011 CORE              56 DXCCS8                      238  241   Num  CCS: diagnosis 8
MD  2011 CORE              57 DXCCS9                      242  245   Num  CCS: diagnosis 9
MD  2011 CORE              58 DXCCS10                     246  249   Num  CCS: diagnosis 10
MD  2011 CORE              59 DXCCS11                     250  253   Num  CCS: diagnosis 11
MD  2011 CORE              60 DXCCS12                     254  257   Num  CCS: diagnosis 12
MD  2011 CORE              61 DXCCS13                     258  261   Num  CCS: diagnosis 13
MD  2011 CORE              62 DXCCS14                     262  265   Num  CCS: diagnosis 14
MD  2011 CORE              63 DXCCS15                     266  269   Num  CCS: diagnosis 15
MD  2011 CORE              64 DXCCS16                     270  273   Num  CCS: diagnosis 16
MD  2011 CORE              65 DXCCS17                     274  277   Num  CCS: diagnosis 17
MD  2011 CORE              66 DXCCS18                     278  281   Num  CCS: diagnosis 18
MD  2011 CORE              67 DXCCS19                     282  285   Num  CCS: diagnosis 19
MD  2011 CORE              68 DXCCS20                     286  289   Num  CCS: diagnosis 20
MD  2011 CORE              69 DXCCS21                     290  293   Num  CCS: diagnosis 21
MD  2011 CORE              70 DXCCS22                     294  297   Num  CCS: diagnosis 22
MD  2011 CORE              71 DXCCS23                     298  301   Num  CCS: diagnosis 23
MD  2011 CORE              72 DXCCS24                     302  305   Num  CCS: diagnosis 24
MD  2011 CORE              73 DXCCS25                     306  309   Num  CCS: diagnosis 25
MD  2011 CORE              74 DXCCS26                     310  313   Num  CCS: diagnosis 26
MD  2011 CORE              75 DXCCS27                     314  317   Num  CCS: diagnosis 27
MD  2011 CORE              76 DXCCS28                     318  321   Num  CCS: diagnosis 28
MD  2011 CORE              77 DXCCS29                     322  325   Num  CCS: diagnosis 29
MD  2011 CORE              78 DXCCS30                     326  329   Num  CCS: diagnosis 30
MD  2011 CORE              79 DXPOA1                      330  330   Char Diagnosis 1, present on admission indicator
MD  2011 CORE              80 DXPOA2                      331  331   Char Diagnosis 2, present on admission indicator
MD  2011 CORE              81 DXPOA3                      332  332   Char Diagnosis 3, present on admission indicator
MD  2011 CORE              82 DXPOA4                      333  333   Char Diagnosis 4, present on admission indicator
MD  2011 CORE              83 DXPOA5                      334  334   Char Diagnosis 5, present on admission indicator
MD  2011 CORE              84 DXPOA6                      335  335   Char Diagnosis 6, present on admission indicator
MD  2011 CORE              85 DXPOA7                      336  336   Char Diagnosis 7, present on admission indicator
MD  2011 CORE              86 DXPOA8                      337  337   Char Diagnosis 8, present on admission indicator
MD  2011 CORE              87 DXPOA9                      338  338   Char Diagnosis 9, present on admission indicator
MD  2011 CORE              88 DXPOA10                     339  339   Char Diagnosis 10, present on admission indicator
MD  2011 CORE              89 DXPOA11                     340  340   Char Diagnosis 11, present on admission indicator
MD  2011 CORE              90 DXPOA12                     341  341   Char Diagnosis 12, present on admission indicator
MD  2011 CORE              91 DXPOA13                     342  342   Char Diagnosis 13, present on admission indicator
MD  2011 CORE              92 DXPOA14                     343  343   Char Diagnosis 14, present on admission indicator
MD  2011 CORE              93 DXPOA15                     344  344   Char Diagnosis 15, present on admission indicator
MD  2011 CORE              94 DXPOA16                     345  345   Char Diagnosis 16, present on admission indicator
MD  2011 CORE              95 DXPOA17                     346  346   Char Diagnosis 17, present on admission indicator
MD  2011 CORE              96 DXPOA18                     347  347   Char Diagnosis 18, present on admission indicator
MD  2011 CORE              97 DXPOA19                     348  348   Char Diagnosis 19, present on admission indicator
MD  2011 CORE              98 DXPOA20                     349  349   Char Diagnosis 20, present on admission indicator
MD  2011 CORE              99 DXPOA21                     350  350   Char Diagnosis 21, present on admission indicator
MD  2011 CORE             100 DXPOA22                     351  351   Char Diagnosis 22, present on admission indicator
MD  2011 CORE             101 DXPOA23                     352  352   Char Diagnosis 23, present on admission indicator
MD  2011 CORE             102 DXPOA24                     353  353   Char Diagnosis 24, present on admission indicator
MD  2011 CORE             103 DXPOA25                     354  354   Char Diagnosis 25, present on admission indicator
MD  2011 CORE             104 DXPOA26                     355  355   Char Diagnosis 26, present on admission indicator
MD  2011 CORE             105 DXPOA27                     356  356   Char Diagnosis 27, present on admission indicator
MD  2011 CORE             106 DXPOA28                     357  357   Char Diagnosis 28, present on admission indicator
MD  2011 CORE             107 DXPOA29                     358  358   Char Diagnosis 29, present on admission indicator
MD  2011 CORE             108 DXPOA30                     359  359   Char Diagnosis 30, present on admission indicator
MD  2011 CORE             109 DaysBurnUnit                360  362   Num  Days in burn care unit (as received from source)
MD  2011 CORE             110 DaysCCU                     363  365   Num  Days in coronary care unit (as received from source)
MD  2011 CORE             111 DaysICU                     366  368   Num  Days in medical/surgical intensive care unit (as received from source)
MD  2011 CORE             112 DaysNICU                    369  371   Num  Days in neonetal care unit (as received from source)
MD  2011 CORE             113 DaysPICU                    372  374   Num  Days in pediatric care unit (as received from source)
MD  2011 CORE             114 DaysShockUnit               375  377   Num  Days in shock trauma unit (as received from source
MD  2011 CORE             115 ECODE1                      378  382   Char E code 1
MD  2011 CORE             116 ECODE2                      383  387   Char E code 2
MD  2011 CORE             117 ECODE3                      388  392   Char E code 3
MD  2011 CORE             118 ECODE4                      393  397   Char E code 4
MD  2011 CORE             119 ECODE5                      398  402   Char E code 5
MD  2011 CORE             120 ECODE6                      403  407   Char E code 6
MD  2011 CORE             121 ECODE7                      408  412   Char E code 7
MD  2011 CORE             122 ECODE8                      413  417   Char E code 8
MD  2011 CORE             123 ECODE9                      418  422   Char E code 9
MD  2011 CORE             124 ECODE10                     423  427   Char E code 10
MD  2011 CORE             125 ECODE11                     428  432   Char E code 11
MD  2011 CORE             126 E_CCS1                      433  436   Num  CCS: E Code 1
MD  2011 CORE             127 E_CCS2                      437  440   Num  CCS: E Code 2
MD  2011 CORE             128 E_CCS3                      441  444   Num  CCS: E Code 3
MD  2011 CORE             129 E_CCS4                      445  448   Num  CCS: E Code 4
MD  2011 CORE             130 E_CCS5                      449  452   Num  CCS: E Code 5
MD  2011 CORE             131 E_CCS6                      453  456   Num  CCS: E Code 6
MD  2011 CORE             132 E_CCS7                      457  460   Num  CCS: E Code 7
MD  2011 CORE             133 E_CCS8                      461  464   Num  CCS: E Code 8
MD  2011 CORE             134 E_CCS9                      465  468   Num  CCS: E Code 9
MD  2011 CORE             135 E_CCS10                     469  472   Num  CCS: E Code 10
MD  2011 CORE             136 E_CCS11                     473  476   Num  CCS: E Code 11
MD  2011 CORE             137 E_POA1                      477  477   Char E Code 1, present on admission indicator
MD  2011 CORE             138 E_POA2                      478  478   Char E Code 2, present on admission indicator
MD  2011 CORE             139 E_POA3                      479  479   Char E Code 3, present on admission indicator
MD  2011 CORE             140 E_POA4                      480  480   Char E Code 4, present on admission indicator
MD  2011 CORE             141 E_POA5                      481  481   Char E Code 5, present on admission indicator
MD  2011 CORE             142 E_POA6                      482  482   Char E Code 6, present on admission indicator
MD  2011 CORE             143 E_POA7                      483  483   Char E Code 7, present on admission indicator
MD  2011 CORE             144 E_POA8                      484  484   Char E Code 8, present on admission indicator
MD  2011 CORE             145 E_POA9                      485  485   Char E Code 9, present on admission indicator
MD  2011 CORE             146 E_POA10                     486  486   Char E Code 10, present on admission indicator
MD  2011 CORE             147 E_POA11                     487  487   Char E Code 11, present on admission indicator
MD  2011 CORE             148 FEMALE                      488  489   Num  Indicator of sex
MD  2011 CORE             149 HCUP_ED                     490  491   Num  HCUP Emergency Department service indicator
MD  2011 CORE             150 HCUP_OS                     492  493   Num  HCUP Observation Stay service indicator
MD  2011 CORE             151 HCUP_SURGERY_BROAD          494  495   Num  Revised HCUP_AS (PCLASSn=3 or 4, or broad definition)
MD  2011 CORE             152 HCUP_SURGERY_NARROW         496  497   Num  Revised HCUP_AS (PCLASSn=4, or narrow definition)
MD  2011 CORE             153 HISPANIC_X                  498  498   Char Hispanic ethnicity (as received from source)
MD  2011 CORE             154 HOSPBRTH                    499  501   Num  Indicator of birth in this hospital
MD  2011 CORE             155 HOSPST                      502  503   Char Hospital state postal code
MD  2011 CORE             156 HospitalUnit                504  505   Num  Indicator that patient was discharged from a special unit within an acute care hospital (reported by source)
MD  2011 CORE             157 KEY                         506  523   Num  HCUP record identifier
MD  2011 CORE             158 LOS                         524  528   Num  Length of stay (cleaned)
MD  2011 CORE             159 LOS_X                       529  534   Num  Length of stay (as received from source)
MD  2011 CORE             160 MARITALSTATUSUB04           535  535   Char Patient's marital status, UB-04 standard coding
MD  2011 CORE             161 MARITALSTATUS_X             536  536   Char Patient's marital status (as received from source)
MD  2011 CORE             162 MDC                         537  538   Num  MDC in effect on discharge date
MD  2011 CORE             163 MDC24                       539  540   Num  MDC, version 24
MD  2011 CORE             164 MDC_NoPOA                   541  542   Num  MDC in use on discharge date, calculated without POA
MD  2011 CORE             165 MDNUM1_R                    543  551   Num  Physician 1 number (re-identified)
MD  2011 CORE             166 MDNUM2_R                    552  560   Num  Physician 2 number (re-identified)
MD  2011 CORE             167 MEDINCSTQ                   561  562   Num  Median household income state quartile for patient ZIP Code
MD  2011 CORE             168 MRN_R                       563  571   Num  Medical record number (re-identified)
MD  2011 CORE             169 NCHRONIC                    572  574   Num  Number of chronic conditions
MD  2011 CORE             170 NDX                         575  576   Num  Number of diagnoses on this record
MD  2011 CORE             171 NECODE                      577  578   Num  Number of E codes on this record
MD  2011 CORE             172 NEOMAT                      579  580   Num  Neonatal and/or maternal DX and/or PR
MD  2011 CORE             173 NPR                         581  582   Num  Number of procedures on this record
MD  2011 CORE             174 ORPROC                      583  584   Num  Major operating room procedure indicator
MD  2011 CORE             175 OS_TIME                     585  592   Num  Observation stay time summed from UNITS
MD  2011 CORE             176 P7EDSRC_X                   593  594   Char Condition Code P7, Direct Inpatient Admission from Emergency Room
MD  2011 CORE             177 PAY1                        595  596   Num  Primary expected payer (uniform)
MD  2011 CORE             178 PAY1_X                      597  598   Char Primary expected payer (as received from source)
MD  2011 CORE             179 PAY2                        599  600   Num  Secondary expected payer (uniform)
MD  2011 CORE             180 PAY2_X                      601  602   Char Secondary expected payer (as received from source)
MD  2011 CORE             181 PAYER1_X                    603  604   Char Primary expected payer plan identifier (as received from source)
MD  2011 CORE             182 PAYER2_X                    605  606   Char Secondary expected payer plan identifier (as received from source)
MD  2011 CORE             183 PL_CBSA                     607  609   Num  Patient location: Core Based Statistical Area (CBSA)
MD  2011 CORE             184 PL_MSA1993                  610  612   Num  Patient location: Metropolitan Statistical Area (MSA), 1993
MD  2011 CORE             185 PL_NCHS2006                 613  614   Num  Patient Location: NCHS Urban-Rural Code (V2006)
MD  2011 CORE             186 PL_RUCA10_2005              615  616   Num  Patient location: Rural-Urban Commuting Area (RUCA) Codes, ten levels
MD  2011 CORE             187 PL_RUCA2005                 617  620 1 Num  Patient location: Rural-Urban Commuting Area (RUCA) Codes
MD  2011 CORE             188 PL_RUCA4_2005               621  622   Num  Patient location: Rural-Urban Commuting Area (RUCA) Codes, four levels
MD  2011 CORE             189 PL_RUCC2003                 623  624   Num  Patient location: Rural-Urban Continuum Codes(RUCC), 2003
MD  2011 CORE             190 PL_UIC2003                  625  626   Num  Patient location: Urban Influence Codes, 2003
MD  2011 CORE             191 PL_UR_CAT4                  627  628   Num  Patient Location: Urban-Rural 4 Categories
MD  2011 CORE             192 PR1                         629  632   Char Procedure 1
MD  2011 CORE             193 PR2                         633  636   Char Procedure 2
MD  2011 CORE             194 PR3                         637  640   Char Procedure 3
MD  2011 CORE             195 PR4                         641  644   Char Procedure 4
MD  2011 CORE             196 PR5                         645  648   Char Procedure 5
MD  2011 CORE             197 PR6                         649  652   Char Procedure 6
MD  2011 CORE             198 PR7                         653  656   Char Procedure 7
MD  2011 CORE             199 PR8                         657  660   Char Procedure 8
MD  2011 CORE             200 PR9                         661  664   Char Procedure 9
MD  2011 CORE             201 PR10                        665  668   Char Procedure 10
MD  2011 CORE             202 PR11                        669  672   Char Procedure 11
MD  2011 CORE             203 PR12                        673  676   Char Procedure 12
MD  2011 CORE             204 PR13                        677  680   Char Procedure 13
MD  2011 CORE             205 PR14                        681  684   Char Procedure 14
MD  2011 CORE             206 PR15                        685  688   Char Procedure 15
MD  2011 CORE             207 PRCCS1                      689  691   Num  CCS: procedure 1
MD  2011 CORE             208 PRCCS2                      692  694   Num  CCS: procedure 2
MD  2011 CORE             209 PRCCS3                      695  697   Num  CCS: procedure 3
MD  2011 CORE             210 PRCCS4                      698  700   Num  CCS: procedure 4
MD  2011 CORE             211 PRCCS5                      701  703   Num  CCS: procedure 5
MD  2011 CORE             212 PRCCS6                      704  706   Num  CCS: procedure 6
MD  2011 CORE             213 PRCCS7                      707  709   Num  CCS: procedure 7
MD  2011 CORE             214 PRCCS8                      710  712   Num  CCS: procedure 8
MD  2011 CORE             215 PRCCS9                      713  715   Num  CCS: procedure 9
MD  2011 CORE             216 PRCCS10                     716  718   Num  CCS: procedure 10
MD  2011 CORE             217 PRCCS11                     719  721   Num  CCS: procedure 11
MD  2011 CORE             218 PRCCS12                     722  724   Num  CCS: procedure 12
MD  2011 CORE             219 PRCCS13                     725  727   Num  CCS: procedure 13
MD  2011 CORE             220 PRCCS14                     728  730   Num  CCS: procedure 14
MD  2011 CORE             221 PRCCS15                     731  733   Num  CCS: procedure 15
MD  2011 CORE             222 PRDAY1                      734  738   Num  Number of days from admission to PR1
MD  2011 CORE             223 PRDAY2                      739  743   Num  Number of days from admission to PR2
MD  2011 CORE             224 PRDAY3                      744  748   Num  Number of days from admission to PR3
MD  2011 CORE             225 PRDAY4                      749  753   Num  Number of days from admission to PR4
MD  2011 CORE             226 PRDAY5                      754  758   Num  Number of days from admission to PR5
MD  2011 CORE             227 PRDAY6                      759  763   Num  Number of days from admission to PR6
MD  2011 CORE             228 PRDAY7                      764  768   Num  Number of days from admission to PR7
MD  2011 CORE             229 PRDAY8                      769  773   Num  Number of days from admission to PR8
MD  2011 CORE             230 PRDAY9                      774  778   Num  Number of days from admission to PR9
MD  2011 CORE             231 PRDAY10                     779  783   Num  Number of days from admission to PR10
MD  2011 CORE             232 PRDAY11                     784  788   Num  Number of days from admission to PR11
MD  2011 CORE             233 PROCTYPE                    789  791   Num  Procedure type indicator
MD  2011 CORE             234 PSTATE                      792  793   Char Patient State postal code
MD  2011 CORE             235 PSTCO                       794  798   Num  Patient state/county FIPS code
MD  2011 CORE             236 PSTCO2                      799  803   Num  Patient state/county FIPS code, possibly derived from ZIP Code
MD  2011 CORE             237 RACE                        804  805   Num  Race (uniform)
MD  2011 CORE             238 RACE_X                      806  806   Char Race (as received from source)
MD  2011 CORE             239 READMIT                     807  808   Num  Readmission
MD  2011 CORE             240 TOTCHG                      809  818   Num  Total charges (cleaned)
MD  2011 CORE             241 TOTCHG_X                    819  833 2 Num  Total charges (as received from source)
MD  2011 CORE             242 TRAN_IN                     834  835   Num  Transfer in indicator
MD  2011 CORE             243 TRAN_OUT                    836  837   Num  Transfer out indicator
MD  2011 CORE             244 YEAR                        838  841   Num  Calendar year
MD  2011 CORE             245 ZIP3                        842  844   Char Patient ZIP Code, first 3 digits
MD  2011 CORE             246 ZIPINC_QRTL                 845  847   Num  Median household income national quartile for patient ZIP Code
MD  2011 CORE             247 AYEAR                       848  851   Num  Admission year
MD  2011 CORE             248 BMONTH                      852  853   Num  Birth month
MD  2011 CORE             249 BYEAR                       854  857   Num  Birth year
`