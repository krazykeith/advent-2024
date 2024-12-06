package main

func day4Input() [][]rune {
  return {'M','S','M','S','M','A','X','X','A','X','X','X','X','A','X','X','A','X','M','A','S','M','X','S','X','M','A','S','M','X','M','X','M','A','S','M','S','S','X','A','S','A','M','X','S','A','M','X','X','S','A','M','X','X','M','A','M','M','S','S','M','S','S','S','M','X','S','A','M','X','X','X','X','X','S','X','S','X','S','M','S','M','M','M','M','S','X','M','A','S','M','M','M','S','M','S','S','S','S','M','M','M','S','A','X','S','S','S','S','S','S','X','M','A','S','A','M','X','M','S','S','M','M','S','M','M','M','S','A','M','S','A','M','X','A','M','A','M','X','S'},
{'S','A','A','M','M','A','M','S','S','M','M','M','S','M','M','M','S','M','X','M','X','M','A','M','A','M','A','M','A','A','M','X','M','A','M','X','A','M','X','X','M','M','S','M','S','A','S','X','M','S','A','S','A','X','X','M','X','M','A','A','X','X','A','A','A','A','M','M','X','S','M','X','S','A','A','A','S','A','M','A','A','M','X','S','A','M','X','X','X','A','A','A','A','A','A','S','A','M','M','A','A','A','M','M','M','M','A','A','A','X','M','A','S','A','M','X','M','A','X','A','A','A','X','A','X','A','A','S','M','S','M','X','M','X','S','A','M','S','A','A'},
{'X','M','X','M','M','A','X','M','A','A','X','A','A','A','M','A','A','A','A','M','X','M','A','S','A','M','M','X','S','A','M','X','S','S','S','M','M','S','S','X','M','X','A','A','S','A','M','X','A','X','A','A','X','M','X','S','A','M','S','S','M','M','A','M','M','M','X','M','M','M','S','A','A','M','M','M','M','A','M','S','M','S','A','M','A','S','X','M','M','S','S','M','M','S','S','M','A','M','S','S','M','S','M','S','A','M','M','M','M','M','M','M','M','M','S','S','X','M','M','S','S','M','S','A','S','M','X','S','X','A','A','A','S','A','X','M','M','M','X','S'},
{'S','S','M','S','S','S','S','S','S','M','M','S','S','M','M','S','S','S','X','S','X','S','X','M','X','S','M','A','X','A','A','A','X','A','A','X','M','A','M','A','M','X','M','X','M','M','M','A','S','M','M','M','X','M','A','M','A','X','A','M','A','M','S','M','M','S','A','M','S','A','M','X','S','X','S','X','M','M','M','M','A','M','X','S','S','M','X','X','A','A','A','A','A','M','A','M','X','M','A','M','A','X','A','S','X','X','A','M','X','M','X','A','M','X','A','A','A','X','X','A','A','M','A','M','X','M','A','M','M','M','S','M','M','A','S','X','M','S','M','M'},
SAAMAAAXAMXMAXAXXAAMAXMMSMMMSMSSMSMMMASXMASAMAXSXMXSAMMSXMASMMXAAMASMAMAMAMASAMXSMMMMXMASXXMMSSXMAMMSMMMAXSSMSSMSMMSMMMMXAXMASAXAXXXMXMAXXAA
SMMMMMMMMSXSAMXSMMSMAMXXAAAAAXXAXMXASMSASASASMXXAMASMSAAXSAMXAMSMMXMMAMSSMXAMXMAXAXXMASMAMXSXMAAAAMMAMXMSMAAAMAMAAMMXAXXAXMSXSMSMXMMMASXMSSM
MXAASMMXXAXXXAMXXAXAASXSSSMSSSSMMMSMSASXMXSMMMAXAMXSAMMSXMASMSMMSSXMSXSAAMSMSMMASMMASASXSXASASXSMMXSAMXAAMMMAMAXXSMASMMMSMASAXMAXAMASASAMAXM
SSSXSAASMMSSSSXSMSMMASAAAXAMXMASAXASMXMASAMXAMAXAMAMXMAMASAMXMAAAMAAAXMMSMAAAAMASMMXMASAMMXXAMAMAAXAMXSSXMMSXMMSMXMAMSXAAMXMAMSMMASASXSAMMSM
MAXASXMMAXSAXMASAAXSXXMMMMMMMSAMSSXSMMXAMASXMXSAAMASMMAMMAASAMMMMSSMASXAAXMMMMMXSXXAXAMXMAXMAMSMMMSAMAMMMSAMAMSAMXAAAMMSMSSMSMAXSAMASAXAMAAA
MAMXMASMXMMMMMAMMMXAXSXSMXAAXMASAMAXASMMSMMASAXXXMXSMMSSMMXMASXXXAAMXMMSSSSMSXMASXSMSASXXAMSXMAAMASMMXSAMMAMAMSAMXAMXMAMAXAAAMXMAAMMMMMAMMSM
MSSMSAMXAXSXXMASXMASMMMAMSSSMSXMASMSMMAXXASMMXSMMSAMAAAAAMMMMMMXMSSMXMAXAAXAXXMAMAXAAAMMMXXSASXSMMSASAMASMXMXMXAMAXMXMSSMSMMMMSASXMXAXMXSXMA
XAAXMASMMSAMXMAMXXAAAAAAMMAXASMMMAXAMMSAMXMAXMAAAMASMMSSMMSASAAMAAMMSSMMMMMSMMXMMAMMMSMXXXAXAMAMXXMASMSMMXAMXXSAMXSSMMXAXXXSXAAMASAMXSAXSAMX
MSSMSMSMMAASAMXMSMMSSMSSXMAMXSAASXXASAAMXSSMMAMSMXMMXMXMXAMASXSXMASAMXAAMMAMMMAMXAMXAAMMSSSMSMMSMMMAMMAXXXXMMXMXMAMAAXSAMMSXMSMSMMXMAMMMMAMS
XAAAAAMAXMMXMMAAMAXAMXMAASXMAMMMMAMSMMXXMMAXMSXAXXSAMXAMMSMAMAXAMXMASMSMMMAXAMMSAAMMSSSMAAXSAXAAAAMMXSMSSMASAAMXMASXMMMAXXAAMAXSXSAMXSXXMXMX
MXSSMSSMMXXMSXSASMMMSAMXMMAMASASXMMAASMSSSSMSXSMSMMAMSASAMMSMAMMXXXAMXMAXSMSASAAMMMXAAAMMSMSMMSSSMXSAASAASAMSASXSXSXMAAMAMXASMMXASXSMSASMSSS
XMAMXAAXMXSAAAXAAMAXSASXSXXMASAMMASMSMMXAAXXSAXXAASAMSXMMSAAMSSMSSXMMAMMMSASAMXMMXSMMSMMMMASXAAMAXAMMSMSXMMXMMMMMXMASMSXSSMMAXSMMMXXAMAMMAMX
SAMXMSSMSAMMMMMSMSSMSAMAMXXMXMAMSAMMAAXMMMMMMMMSSMMAXXXMXMMMMXAAAXAASMMSAMXMXMSMSXSAXMASXMAMMMMSAMXXMAMXXAMSXMAAXASAMAMMAMASAMSASXSMAMXMMAMM
SXMAXXAMMASMXAXXMMAMMMMSMMSMMMXMMAMSSMMAASAAXMAAXXSXMXMSAMXXXSMMMSSMMSAMMXSMMXSAAASAMSAMAMAXXXAMXMSSSMSASAMMASMMSASXSAMSMSAMSASAMAMMXSAMMSSM
XAMSMSSMSAMAMMSMXXAMASXMAXAMXSMASAMMAMXSXXXSXMSSSMMMSXASXMXMAXAAXAMMMMMXSASAMAMMMAMAXMASASMSMMXMAXSASAXASMAMMMAAMAMAMXXXXMXSXXMXMAMASMMSAMMM
MSMMAMXASXMSXMAXMSXSASXSMMMSAASAMAMSAMMMSMXMAMAAAMAAXAXXMSSMASMMMASXXMAMMAMAMXSXSASXXMASAMXAMXSMSMMAMMMAMMSMXMMMMAMXMSMSMSMMXXSMSMMMXAAMASAS
AAAXSAMXMAAXAXXMXAAMXSAAAAAMMXMXXAMSAMSAAMASMMMAMXMSSXXAMAXMAMXXAAAMAMMXSAXXMMXAXXMXXSAMXMSXSAXXXAMXMXAXMAMAAXXASMSMMSAAAAMAMXAAMAASXMMSXMAS
SMMMXASXSXMSSMSMMMXMAXMSMMXSMMMMMSMMAMMSSSMMXASAXAMXAXMAMASMSSSXSXAMAMAXMAXSAMMMMAXMXMSSMASAMXSASXMXSSSMMMXSAMXXXMAMAMXMSMMSAAMSMSMSAMXXMMAM
AXAASAMASMAXMAAAXAMSMXAXMXASAMXSAXMMAMXAXMAXXASAMSSMAMSSMMSAMAXMXMXXASMAAAAXMMAMMSMXMAMXAMMXMXAAMXAAMMAMMAAMAMSAMXSMXXMXMXAMXSXXXAMXAMXXMMAS
MMMMMAMAMMSMMSMSMXMAAMSMXAXSASAMMSASXMMMMSMMMXMXMAAMXMAMAXMAMSMSAASXMMMXMMMXMMXSMASAMXMASXMASXMSMXMAMSAMMMMSASAMXSXSMMSAMXXSXMASMMSSSMMSSMAS
MSMSSSMMSAMXXXAMXSMMSMAAXMASAMXSSMMMSAASAAAMXAAAMSMMSMSSSMSAMXMXAXSAMXSAXAAMSMAXXXMSMAMAMAMXAXXMXXAAAMASASAMXMXMMSAMXASAXSAMXMAMSXAMAAAAAMAX
MAAXAMAAMASMSMAMAMXMXXSSXMAMAMAXMAMAXSAMMSSMSMSMXXAXXAAAAASXSAMSSMSXMASMSMSAAASAMXXXXAMASXSMMXSAMSSMMSASXMAMXMXMXMXMMASAMMSMAMAMMMASMMMXSMSM
MMSMASXMSAMMAMSMSXAXSAXXAMXSSXMASAMSXMMSAMMAAAAMASMMMMMSMXMASAMAAAXAMAMXAMXMXSMSXMXSXMSASXXASAMAAAXAMMAMASMMAMAMAMASXMMMXMAMAMAXAAXMASMXMAXM
SAXMXMAAMASMSMXAMAMSMMXSSMAMXAAAXAXXXAMSASMSMSMSASAAAAXAMXMMMXMAXASXMASMMSAMXMAXSMASAAMMSMSXMXSMMMSXMAMXXXSAASASASASAAMAASMSMSSSXSMSAMXSMAMX
MASMSSSMMAMAXAMAMAXAAMAMAMASXMMXXSMSMSMSAMAAAXAMASXMXMSAMXMMMMMMSAMASMSAAMMSXMAMAMASMMMMXAMXMAMASXSAMXSSMMMSMXMSASASMMMXXSAMAAAMAMXMAMXAMMMA
XMMAAAXMMXMXMMSSSMSXSMAXAMSAMXSMXXASAXAMMMXMAMSMXMMSAAMMMAXMASAAMMMASXMAMXMAMMXSMMXSMMSSSSSXMXMAMXSASAAXAAMASXXMMMASXMXXMMAMMMXMAXMMSMMMMMSX
MXMXMMMXMXXXAMAMAXAXMXSMSSXMAMMMMMAMAMSMSSXMASAMXSASMSMAXXXSAMMXXAMASXMAMSASASAMXMAMXXAAXXAXSAMSSMSAMMSMSXSASXMASMSMMSSSXAMMSMSSMXSAAAAXAAXX
MASAMAXMXASXXMAXMMSMSAMAXXAMMSAAAMAMAXXAASAXXMAMAMASXMMSXSASASAXSMMMMAMXXXAMAMXSAMMSMMMSMMMMMASAAAMXMXAXMAMXSMSMMMXASAAMSXMAXMASXAMASXMSMXSA
SASASXMAMXMAASASMSMAMAMSMSMMASMSSSXMXSMMMSMMSSSMMMMMMMAMAMASAMAXAXAMSMMSSMSMXMASAXSAMAXAXAXXSXMXMMMAXMAMSAMAMXXAAASMMMSMAXMXSMMSMSSXXXAAXAMM
MMSXMASXXSMSMMAMXAMMMAMAAAXMASAMAMMSMMXMAXAAAAXXASXSAMMMAMMMMMXAMXMMAAXASXMAMMXSXMSMSMSSSSSXMAMXMAXXSAMXSAMSSSMMMMMMXXMXMXMMXAXXMXAAMMXMMSSM
SXXXMMMMAAAXAMXMSSMMXXSMSMSMASMMXMAAAXSXXSSSMAMSASASMSSMAMAAXMXMSASXSSMMSMMAAAXMMMMXMMXAAAAAMSXMMSSSMMSAMXMAAAASXMSSMMMAXSXSMMSSMMSMAAAXAMAM
XXMMXAAMMMSMMMAMAAAAMMMMAMAMMSXSMMSSSMMMXAAXXMMSAMMMXAASASMSMXSMMAMAMXAXMASMSMXSAAMSSMMXMAMMMAAAAAMAAAMXMMMMMMMMAAAAAAMMMAAASMAMAAMXXSAMSSSM
MSMASMSMSAAAXSXMXSMMXAASXMXSMXAAAXMMAMAAMMXMMMAMAMXSMMMSMSXXMASAMAMXMMSXSXMMAXASMXMAAXSSSSSSXSSMMMSAMXSMMAMXAAASMMMSSMXSAMXMSMAXMMXSMMXXXAAS
AAAXSAAXMSSSMMASAMAXMSMMAXSXMMSMSAMXAMMXSMASAMAXASMSASASXSMSMASMMASAAXAASXSXXMMSAMMSMMSAAAAMMMMAMXAASXXAXXXSSSXXAMAAAXAAXSAMXMSMMAXXAXMMMSMM
SASAMMXMMAXAAXAMAMSMMMAMAMMAMMAAXMSSSMXAAMASASMSMSASAMASASAXMAMXSAXMAMMAMASXMXAMMMAAAXMMMMMMAXMXMMMMMAMSXSAAXMASMXMASMMMSAMXXAMAMMMSAMXAAAAX
MAMAMXAAMXXMSMMMMMMXASAMSASAMSMSMXAAAAMSSMMSAMXAXMMMAMMMAMXMMXMMMMSSMAXAMAMMAXMASMMMXXMAXSXSSSMXSAXSMXMAAMMMMMXAAAXXXMXXAAASMASAMAAXMASMXMSM
MSMMSSMSXMMMAXSMAAMAMSMMXAXMMSAMXSMMMMMMAAAMASXMMSXSAMXMMMSMSXMXAMXAASMMMASASAMAMAMSSMXAMSAMAAAASMSASMMMSMXXMMMAMMMAXMXMSMMMMASMMMSSMXAMMSMM
XXAXXMAXXAASMMMMMSSSMMMSMMMSMMAXAXMAMMXMSMMMAXAAAXMASAMXSAAAAMMASXSSMMAMSMSXMXSAXXMAMXAMMSAMSMMMSXSAMXAAMXMMAXSSMAXXMMMXMSAXSASMAMMAXMXXMAAA
MSMMAMAMMSMSAXMAMXMAXXAAAAAMASMMSMMXXXMXXSSMASXMMMSAXXSMSXMMMSMAMAMAXXAMMASAMAMXSMMMSSXSASAMXXAASMMMMXSXSASAAMAXSXMMSMAAXMMMSAMXAXSAMXMXSMXM
AAAAXMMXAAXXAMSSMASXMMSSSMSSXMMAXAMSXSAAXAXMMMMAMAMASXMASXXMAXMAMXMAMXAXSAXXMAMAAAASXMAMASAMXSMXSAAMMXMASASASAXXAMAMAMSSSMSAMAMMMXAMXAAMXSXX
SSSMSASXSXSAMXAAMXMXMAMXAAMMMAMAMAMAASMSMMMSAASAMAMAMMMAMAMMMSMXMSMAMMAMMAMSSMSSSSMSMMSMMSAMMXMAXXMSMAMAMAMAXXMMSXSSMMAXAAMXSAMSSMAXSMSAASXS
MXMAMAXXAXMAMMMSMMMMMMXSMMMMSSMXSSMMXMXMAXAMXXSXSASAMSMMSMXASXXAAXMASAMSMMMMAXAAMXAXXAAAASAMXAMXSXMAMASAMXMXMMSAMAXXXMAXMMMAMASAMMMMXAMMAMXX
MAMMMAMXMSSSMSAAXAAAAMMSAMSXAXMAMAAXSSMSSMSSMXMAMMMXAAAAAXXMSAMMMMMAMMAAAASXSMMMMSXMMSSMMXAMMMSXAAMXSASMSMMSAXASMXMASXSSXAMXSMMMSAMMXMMMSSSM
MSMXMMXAMXAXAMSMSSMSMXAXAMSMSSMSSSMMXXAAXXAAXMSASXSSSSSMXMMAMXMSMSMXXXSMMMSAXAXAAMAXMAXAMSSMAASXSXMAMXSMSAAASXMXMAMSMAAXXXSAAXSASAXXAAXAXAAX
MXXXSASXSMSMMMXXAMXXAMSSSMXAXXAAMXMASMMMSMSSMAMAXAAAAAMXAMXSAMXAAAASMMMMSMXAXXMMAMSXMAXXMAAMMMSAAXSMSAMASMMMAMXASXSAMMMMMMMMMMMASAMSAMMSSSXM
MSMAXMASMXMAMAXAMXAMAMXAAXMMMMMMSAMSSXXAMAAAASMSMMMMMMMMXXAMAMSMSMSMAAXSAAXSMSMSMSMMMSMMMSMXAXMXMMAAAMMMMAXXMASXMMSXSASAAAAAAAMXMXMXAXAAAAMA
SAMSSXMMXXSAMASMXMMSSMSMSMXXAXMAMAXAXMMMSMSMSMAMAMAMXXAMXMXSAMAAXMAXMMMSMSAXAXMAXAAAAAAXXAAXSMMAXXMMMXSASAMXXMMASMXAMASMSSSMSMSAMXSSMMMSSMAM
SASXMAAXMMSAMXAMAMXAXMMAAAASMSXMSSMAMSAMAMAXAMAMSSSSMXSAAXASXMMXMMMSXAAMXMMMSMSSSMSMSMSMMMSMMASMMSMSAMMAMSSXAXSAMMMMMAMMXXMAXMSASAMASAXAMXSM
SXMASAMXAASAMXMSSSMSSMMSMSMAAAAMAXXMMAMMAMXXMMMMMXMAMAMXSMASXSAXSAASMMSSXAAAXAAAAAXAXAAAMAAMXMMAAMMMAMMXMAXMMMMAMXAAMASAMMMSMASXMASAXMMAMAAX
XAMAMXXMMXSXMMMAAAAXAAAXAAMMMMMMASMMXXAMXSSMSASASASAMXSAXMMMAAMAMMXMASAMXXMSSMMMMMMAMSSMMMXSAASMMSXMXMSAMXSMMAMMMSSMMASMXMAMMXMASXMMSMSAMXSM
MSSMMASXXAMAMXMMSSMXSMMSMMSSXSXMXMAXMSXSAAXASXSASASXSAMXSMAMXMMMXSMSSMMSSXAXAMASAAMAMMAMXAAMMXMMXMXMSMMASAAMSMSAAMASMMMXSXMXMASAMXXSAMXMASAA
AMAASAXXMAXSMSXMAXXAMXXSAMXMAMMMMSSMXAAMMMMAXAMXMAMAMMSSMXAASXXMASMMXAXAMMMMMSASXSSMXXAMMMSSSSXMAXMXAAXAMMSMAASMMMAMXMSAMAMMSXMAMMMMAMSAMXXS
SMSMMXSMSAMMXSAMXMASXSAMXMAMAMAAMAAAMMSMXXMSMMMMMSMXMAXAXMXMXASMAXAASAMXSAMAXMXSXXMMMMSMMXXAXAASMSSSSSMSSXMMMMMMAMMSMAMASXMAMASAMSMSAMXMAMXX
AMMAXXAAXMASAXXMXXXXASXMASASMSSSMSSMMAMAMXAAAAXAAAAAMXSMMMSXAAMMXXXMSAMXSASXSXSXMSSSXAMAMMMMMSMMAAAXAAAXXXMAMXAXXSXAAMSAMXMAXAMXSAAMXSSSSSSM
MMSAMXMASAAMXMSMMSAMMMASASXSAAAXAMXAMXSAMXMSMMXMSSSXMAAXAMMMMMAXSASXMASMMXMMSMMAXMAMMASAMAXXAMAMMMMMSMMMMXSAMSSMXMMXXXMASXSMMSSXMMSMMMAAAAMA
XAAAXXXXSMXMAXSAAXXXXSMMXMXMMMMMSMSSMMSAMSXMAAXXAMAXMSMMXXAAASAMXMAAMASASXSASASMMMAMSMXASXSMSSSSSSMAXAMAMAMAMAAXSMMSSXSXMMMMAAAXMAXXAMAMMMMS
MMSSMMMXMXAXMSMMMSXSXSXSMMSXSASMXAXMAASAMXASMMMMASMAXAASMXSXXXXXAXSXMASMAAMASAAXASAAAMMMMMXAAAAAXXMXSAMAMASMMSSMAAAAXAASAXAMMSSXMMSMSMAMSXAA
MMMMMAMAMSMSMMAAASMSAMASAASASASAMSMSMMXMMMMMMXXSMMMMMMSMAAXXSXSSMMXMMXSXMXMAMMMSXSMSXSAMXSMSMSMMMXAAMMSAXAMAAAAASMMMSAMXXMAMXAMXSAAAMXSMMMMS
XAAAXMMAXAMAAMSMXSAMXMASMMMAMAMMXXAMXSXXAXMAAMMXMAMXSXMMMSMAMAMXSAXMXAXAXMSSMSASMXAMASASXSAAAAXXMXMASAXMSSSMMMMMXXMXSXXXXMASMAMASXMXMAMAMMSX
SSSSSSSMSASXSMXSAMAMMMXSMXMXMXMSAMMMAMMSSSSMMXSAMSMAXSMAXXMAMXMAMSSMSMSSMMAMXMASXMAMAMAAAMSMSMSMXSXMXMXAAMXXAAXMMXMAXSAMXMASXMMASMSSMASAMXAX
AAXAAAAXSXMAMXSMMSSMMAAXMAMXXAMMASXMAMXAXMASMAAXAMMMSXSMSMMMXAMMMAXMAMAAAMAMAMAMASXMXMXMSMMMMAMMAMXMASMMSSSSSMSAMXMAMAXSAMASAMMXSXAAMXSMMMSS
MSMXMMMMMMSMXSAMXAAASMSSSSSSMSAXAMXMSMMMSXAXMMMMAMAXMMMXSAMSSMSSMSSSSMSSSMMSSMSSMMASAMXXMASAMAMMXSXMAXAAAAAAAXMAMXMSSSMSXSXSXMXMMMSSMMMMXSAS
MMMMMAMAAAAXMXMSMASMMXAMXXAMXXAMSSSXXASXMMMMMAXMAXMMMXSAMXMAAAAAXXMAXAAAXXXAMXAAASAMASMMSAMAMMSMAMAMSSSMMSMSMMSMMMAMAAAXAXMSXMASAAMAXMMAMAMM
MAAAMASMMXSSMAXXMAMASMSMSMSMAMAMAAXAXXMASXXASXSMSAAXSMMSXMMSSMSSMSMAMMMMMXMAMMSSMMASAMMAMASXMXXMAXAMAAAXXMMAMMAAAXAMSMMMMSASASASMSSMMMSASMXA
XSSMMASXSAMXMAXXXAXXMAAAXAMMXSAAMXMXMMSAMXMXSAAAMMMXSASAMSAAAXAXMXMAMMXAXSMSMXMXXXMMXSXASXMMXXSSMMSSSMSMXASASMSSMSXXAMXAXAXSAMASAAMAAXXAMXSX
AMAMSASAMXSXSSSSMMSAMXMAMMMAMXXSSSMAXMASASAMXMMMMAMAMXMAAMMMSXXAMAXAMMMSAAAMSAMMMSSMMXMMSAMXAAXXXXXAXAAXMMMXSAXXMAMSAMSMSMAMXMAMMMSMMMMAMASM
XMAMSAMASASMAAAXXAAAASXXMAMASXAXAAXAMMXMMMAXXXAXMAMMXSMXMAAAXMSXSASXXMAXMMSMSXSAXAAASMMXSMMSSMMSMSMAMSMMAASXMXMAXXAMAMMXMXAAXMASAAXMASMSMAMA
SMSMMASXMAXMMMMMMMSSMMXSXMMAMMMMSMMSXSAXAMXMSSMSSSMSAAXSXXMASAAXMAMAAMSMXXMAMASMXMSMMXMAMMAMAXXAAAMAMAMXSMSMSXSXMMMMAMMASMMSXSASMXSXXAAMMMSS
MAXXSAMMMSMSXAAAAAAXAMAAMSSXXASAMXXMASASMMMAXAAXAAAMMSMXASMXMMMSMSMMMMAAXMMAMMMXAXXXMMMSMMASMMXMSMSMSMSAMAXAXMSAMASMMMSMSAMXAMXSXASMSMSMAXAM
MAMAMASAAAMXXSMSSSMSAMMSAAAMSMMMMMMMAMXSAASMSMMMSMMMMMXSMXMAMAXAAMSMSXMMSASXSMAMXXSAXAMXASASASAMXMAMAXXAMAMSMASXMASASAAXSMMMXMAMMMXAAAMMMMMM
MMSXMSSMSSXSAMMAMAXSXXXMMMMMXAAAXAAMMXAXXMMAAXMAXXXSXMAMMAMAMMXMSMSXMASASMMXAMXMMASXMSASXMASAMXSASMXXSMMMAMMAMSMMASAMMXXMXSMXSMXSSSSMSMAMSSS
MAMXXXXAMAASAMMASXMMXSXAXSSMMSSSSSSMSMMSXXMAMXMAMSAMAMAXXASMSMSMMASXXAMAMXSSMMASMAMMAMXMMMXMAMAXXMXSMMASAMXXMXMASAMMSSMSMAXMASXMXAAAXMMMMAAM
SSSMSMMMMMXMAXSASMAMAMMSXMAAAAAXMAMAMAASMSMMMXMASMASXMAMSASAAAAAMAMMSSMSMMXAXSAAMXXSSXMMASMSSMSMMSXAASAMASMMSMSAMXSAAAMMMSSMMSASMSMMXXAXMMSM
XAAAMAAXAXXMAMSAMXAMXSAMASXMMMMMMAMMMMMMAXAAXMMSXMAMXXAMXAMMMSMSMSMXAMAMMMMMMSMMXSAAXAASXMXAAMAASASMXMASXMAAAXMXSAMMSSMAAAXAMSAMAMXXXSSSXSXM
MSMMXMMXAXXMAMXAASXSSMASAMXXSSMXMSSXMXSMSSSMSMXMASMMAMSASXSSXMMAAXAMMMSMMSASAMAAAMMMMXMMASMMSMSSMAXXASMMAMMSMAXSMMSAAAMMSSXMMMMMMMMXXAAXMASX
AXAXXMSXMSMAMMXXMXXMAMMMMMSMSASXAXAMXAXAAAMAXMSMAXAMAAAASAXMAASMMMAMXAAAASAMMSMMMMAAXMASAMXAXAMAMAMSXSMXMMAMXSMXMXSMSXMAMAMSAMXSXAXAMXXAXMMS
MSMMSMAASAMAMXSSMSMSMSSXXAAXMAMMSMAMMASMMMMSMXXMAMAMMSMSMXMSMMAAMMAMXSSSMMAMAXMASXSMSAMMASMAXAMAMXXMAXMASMMSAXAMSAMXMAMMSAAXAMAMASXSAMMXMXSA
MAXAMMAMSASXSASAAMAXXXAAMSSXMAMAMAAXXXSAAMAMMMSXXAAXXMXXAXXASMMSXSASAMXAASAMSSSMSAXASXXSXMMSSSMSMSAMMMXXMAXMMMMAMAXAMXMXSMSMXMASAMAXAXSAAMSM
SSMMSSSMXAMXMMMMMMXMMMMMMAAMSAMXXMSAMXXXMMASAMAMXMXSXMASXMSMAXAAXMAMMSAMXMAMAAXMMAMAMMMMAXAMXXAAAMSASMSMSMMAAAMSMSMXSASXXXAAASXMASASMMSXSASX
AAAAAAMMMSMXMAXSAXSAMAMAMSSMSAXASXAASASXSSXSASMMMSMMMMASAMXXAMMSSMMMSXAXSXMMMMMMMMMAMAASMMSSMXMMSMXXXAAMASASMXXAAAMXMASAMSXSMSASAMXAMXXAMAMA
MMMMMSMAAAASAMXSASXMSASXMMAMXMMSAMXXMXSAAXASMMAAAAAMSMXSASAXXSAAAAAXMAAXMASMMXSAAASMSMMXAXAAASMAXXXXMSMSMAXMSSXMSMSXSAXAMMMMXSAMXSAXSAMXMAMA
SXMXAMMMSMSAAMAMXMAXSAXXXSXMAMXMAXMMMXMMMMMMMSXMXXXMAMXMAMXSMMMSSMMSSXMSMAMAAAXMMXXAMXSSXMSSMXMAMMSMMXXAXMXMASMXAMXMMMSAMXAXMMXMXMMXSAXXXAMA
AMXMAMAMXMXMSMAMAMXXXMXSXMASASXSXMAAMAXAMAAAAXASASXSXSAMAMXXAXXAAAAAXAAXXASMMMSSSSMMMAXMAMXAAXMMMAAAMAMXMXXMASMSAXXXAAMASMSSMSAMXMSAXSAMSSSM
MAMASMMSMMSAXXMSMSAMXMASASAMXSAMXAMSSMSSMSMSSMMMASAAXSXMAXSMMMMSSMMSXMMASASXMAMXAAAXSMMMMMXMMSMAMSMSMMSAAMXMXSASAMSSMXSAMXXSMAAXXSMMSAXMAAAM
SSMAXMASAAMMMMXAMXAMAMMMAMXSASASMSMAMAAXAMXAXAXMAMMMMMASMXSAAAXAXMAXXSAMXAXMMAXMMMMMAXAAAMXSAMXAMXAMAXSXSMASAMMMAMXAAMMMXAMXSSSMXAMXAMMMMSMX
AAMMMMASMMMSAMSASMSMSSSMSMMMASMMAAMXMMASMMMSSSMMMSXMAMMMMAMSMMSSSMAMASMSMSXXSSXSASMSXXSMMXAMASXSSXASAMXAAMXMAMAMMMSMMMXSMSMAMMAXMMSXMXSAAXXM
SXMXAMXMXSXMAMSAMAXXMAXXXAXMMMASXMMMSXAAXAXMAMXAAXMSMSMSMSMXMMXMSMASASAXAXMMMMASASASAXMASMASMMMMXMXMAMXXMXMASAXSSXMMSSMSAAMASMMMMXAAMASMSMSA
MASXMSSXAXXMXMMAMAMAMXMASMMMMSAMMSXAXMSXSASMSMSMMSAASAXAAAXMAMSAXSMAAMAMAMAMAMMMXMAMSAMAMSAMXAAXXXXMASXSMXMAMMAMXAAAAMAMSMSMSAASMAXSMASXMAMS
MXMAMAMMMMMMSSSSMASXMMSMSAMAAMASAMMXSXMAMAXAAAXAAMMMSXMMMMSAAAMXMASMSMSSSSSMMSSMXMXMMXMAMXMASMSSSMXMASMAMAMASAMXSXMMSMSMMXSASXMSSSXXMASMMAMX
MSMSMASMXAASAAAXSASAMASXSAMMXSAMASMXSAMAMMMMMSMAMXAMXXMXMAMMMXSAAASXXAXSAAMAMAAMMSAXSMMSXMAXAXXAMXAMAXMASXSASXSMMXXXXAMASAMMMXMXAMASMXSXSSMM
AAAMSXMXSMMXMMMMMSSXMASXMAMXMMMMMMMAMASXMMMSMMMAXSMSAMXAMASXSASXMAXXMMMXMASAMSXMAMAXXAAXXSSMSSMSMXXMAXXAMXMASMMAXXMMMAMAMMMMSAMMXMSMAXMASASX
SMMMAMXMMSSSMMSXMAMAMSSXMSMMMAMAMAMXSXMMASXAAAXXMAMMASMMSXSAMXSAMXSSMMAXXMXAXAAMMSSMSMMSXSAAMAAAAXMSMSMXSAMMSASXMMSASAMSSMSASMSAMSASXMXXSAMX
MMMSXAMMAAMAAMMMMASXMMMXAAMAMSMMMASAXMASAMSSXMXAXMXSAMXXSAMXMXXXMAMAAXASMSSSMMMSMAMMXAAMXSMMMMSMSMAAAXASXXSASMMXSAXMXAMXAAMASXMAXSAMMMXXMXMM
SAMXASAMMSSSMMAASMSMAAAXSAMMMXAMSMMMSSMMAMAXXMASMMAMASMMMMSAMXSMSSSSMMASAAAAXAAMMAMMSSMSMXAAXMAXAAMASMMMMMMASMSAMMSSSSMSMMMASAMSMMSMAAMSMAMA
SAMSMMXSAXMMXMSXSAMXSMAMXMXMAMMMAMAAMAXSXMASMXMAAMASAMMAASXMMXMAAAAAXXAMMMMMSMSSMMMMMXAAMXSMSSSXSSXXXXXAMSMSMMMXSASXAAAXAXMASAMXMAMSSMAASAMS
MMMSAAXMXSAMSMMXMMMAMMXSASXMASXSAXMSSXMSAMXMMAXSXMXXMAXSXSAXSAMXMSSMMMSSMXAXSAMAASXSSMSMSAXAXAAXXAXSAMXSMAXMMSAMMXSMSMMMMSMAXMMAMAXXXMSXSMSX
MAMMMMSAAMAMSASXMXMSMSASAMASASMXMASXMASMMMAAMAMXMSMXXSXMASAMSXSAAAAAXXMAMMAMMAMSMMAXAAAAMMSAMMMMMAMMAMAXXMXSAMAMSAMXMXSXAAMMSMSSSMSAXMAMXMAM
SMSMSASMSMMAMAMASAMMAMASMXXMASAMXMXAMXSAAMSSSMSAXAAXAXAMXMSXXASMSSSMSASAMSSMSXMXMMMMXMMAMAMAMXMAMMMSAMXSAXXMASAMMAMMSASMMXSXAMAMAAMMMMMMSMAX
AXAAMASAAMXAMAMMMASMAMXMAMSMMMMSSMSSMXMMMXMAAAXMMMSMMSAMXAXMMAXXAXMASAMAXAAMXXMMSAXMSMXMMMXSXMXXXMAXAMAMMMMXMMXXSSMAMASXAAXMMMASMMMMMAAASXMX
MMMSMAMXMASXSMSASXMXSMAMSXMAAMXAAXAASXXXMSMMMMMSAMXXAAAXMSAXMXMXSMMMMXMMMSSMSMSASXXMASAMASMMASMSSMMXSMASXXSAXASXAAMAMSMMMXMAASXSAXXAAMMMSASM
AAAAMXXAXMXAAASASXMAXMXSAAXSMMMSSMSAXSAMXAAMXMAMAXAMMMMXAMMXMASAMXASAMXMAXMXAAMXSMMMAMMMAXASAMAAAASAMXMMAAXAMXMMMMMAMAAAAXMSMSASXMSXSSSXSAMS
SSSXSXSMSMMMMMMXMMMMMMXMMSMAAXAMXXMASASMSXSMAMSSSMMSMSMMAXAAXAMMSXXXAMMXSXMSMSMXXAAMXSXMASMMMMXMSAMXSXSMMMMSMMMAXXSSSMSSSXAAAMAMAMXMMXMAMAMA
XAAMXMAMAMAXXXMAMAMMAMXAAXXSMMSSMXMXMAXXAAXMASMAMAXAAAXMASMMSMMMSASMXMASMAMXXAMSMMMSAAMSXMMAXAMXAMXAMXXXXAAAAASMSXSXAAAAAMMMSMSXXMAMXAMMMSMM
MXMMSAXSASAMXSSMSMXSASAMXXXMAXMMMASAMAMMMSMAXXMASXSMMXXXMAXXAMXAXXMAXMSASAMXSXXAAAAAMMMSASMSMMSAMXSAMXSMSMSSSMSASMSSMMMMMXXSAMXMMXSSSXSAAAAX
XMXMAXMMASMSAAAAAXXMASXMASMSSMSAMMSXSASMMMMMMXMXSMXAMSMMXAXSXSMSSSSSMMXMSMSXSMSSSMSSXSASAMAMSMXAXMAAXAAXXXMAMMMXMAXXXAXMAMSASMMSAAMAMXSMMXMM
XMAXMXSMMMAMMSMSMXXMXSAMXAMAAASASAMXSMMMAAAMMXMAMAMAMAAXMXASAXAMAXAAMXSAMASASAAMAMXXAAXMAMAMMSSSMASMMSMMSMMMXMASMSMXMASMAMMAMXAMMSMAMASXSSSS
XXMMSAAASMXMAXMAXMXSMSMMMXMMMXMXAMAAXAASMSMMXAMASASXSMXSAMAMMMSMMMSMMXMAXXMAMXXXSMSMMMMSXMXSMAMXMMAXAAAMMAMMAMSMMMMMMAMMMMMAMMSMMXXAMXMAXAAA
SXSAMXSAMXAMMXSASMMMAXXXSMSXXXAMMAMMSMMSAXMASXSMXAXAMAAXXMXAXAXAMXMASMXSMAMXMMXSMMSAAXAXASXMMASAMASMSMSMSAMMASAAASAAMSSXMASASAXAAXSXSAMSMMMM
SMMASAMAMXXXAXMAMAAMSMMMMASAMSAMXXSAMMAMXMMASAAMMSMMMMMSXMSMMMMMMAXSXMAMMSMSMSXAAASXMMXSXMAASMSXSAMMAMMMSMSSXSXSMSXMXAAASASASMSMMMMXMASAAXAX
MASAMXSXXXMMSSXASMMMAXAMMAMAXXMXXXMAXXMSMSMASMSMXMASXMXSAMAMASAMXSMMAMASAMASAMSSMMXAXSASASXMSAMXMASMAMXASAMXAMAMMSXSMMSMMASXMXSAMMSASMMMSMSA
MMMASXSASMAMAMXMAMMSMSXSMSSXMASXMSMSMMXAAXMASAXAAASMASAMASXMASASMMASAMASAMAMAMAXXXSMMMAXAMMXMAMXMMMMMSMMMMMMMMAMAXAMMAMMMAXXSXXAMAAMASAXXAAX
SMSMMMMAMMAMXMAXAXXAXAXSAMXASAMAAAAAMAXMSMSXMMMXXMAXAMASXMXXASXMASAMASASAMXSAMSSMXSXAMAMSMMASAMSMASXMAXXAMASXMSMSMAMMAMXMMSAMMSSMMSXMXXMASAM
AXAAAXMAMSMSSSMSAMSAMXSMAMSMMMSMSMSXSMMXMAMMAAAXSSSMMSXMAAXMXMAXMMXSXMXSAXASAMXAMASMMMMSMASAMAXAXXSASMXMASASMAMAASXMMASAAAMAMMAMAAAMXSAMXMXA
MSSSMSXMMSMAMAMMMMMMXSAXAMAMAXXXXAXXAASAMAMXSMSXMAXAAXMSMMMSMMXMSMMSMSXMMSMXXXXAMAMMXAXAMMMSSMSSSMMAMAASXMASMAMMMMSXXMAXMXXAXMASAMMAASXMSMMS
XXAMASXMASMSMXMASXSAMSASXSXSXSXSMMSXMAMAXMMXMAMAMAMMMMMXMAAMAMAMXAAXAMASAAXSXSMMMXSASMSMSMAXMXAAAXMAMXMAAMXMMSXSAMXSMMMXMASAMMAMAAXMAMAXASAM
XMMMAMXMXSAXXAMXSAMSXMMMXAMXXMXSAAAXXASXMASMSASAMASXSXMASMXSAMASXSMMMMAMXMXMAMXSAMXAAAAXAMSSSMMXMMMXXAXXSMSAXAASMMAMXAMXXXAAMMMSAMXSMSMSAMXS
SMXMASASXMAMSASAMAMMSMSXXSMXASAMMMSSXMAXSMAAMAXXMXSAMMSMSAXSXSASXAAASMSSSMSMXMAMAMXMMSMMMXXAMXXAAMSSSMSXMAMXMMXMAMASMMMMAAMMMXMAMMXSAAXXXAMA
XMASASMMAMXMMXMASMMAAMXMAXAMXMMXSXAXMXAXSXMMMSMMSAMAMXAAMMMMXMASMMMSAXMAMAAXSMXMSMSMXXMASMMAMASXSMAXXAMMMAMMXXAMXMMSAXSAMXXXSXMASXAMSMSMMMXS
SXMXASXSAMMAXAXMXXMSMSAAXMAXSASAXMAMMMSSMXXXAAAXMAMMMSMSMSAMMAAXASMXMXMAMMMMXAAAAAXMASXSMASMMXSMAMMSMMMASASMMSXSASASMMAAXXSAMASAMMMMMMAAMSAX
XASMMMASMSASXSSXSAMAASASXMAMMAMMSSMMSAMAMMMMMSMMMMMAAXAMMSASAMXSMMAAXMMMSAAMSSMSMSMMAMMMMAMXMAXXASXSAXSASMSAMAXSAMXMXSSMMXSAXAMASMXAMXSSMMAM
SAMAASXSAAXMAMAXSAMMXMAXAMSMMXMXAAMAMMSASAAAAMASASXMMXXMASAMXMXSXMSSMAXAMMSMAXMXMAAMASASASXMXSXSXSAMXMMAMMSAMMMMXMSXAMASMASXMXSMXXMSSMAAXMAM
MMMAMSAMXSMMAMSMSMMMMMMMMMMAXMSMSSMASMSAMMSMXSAMMXAAASMMMMMMASMMAMAMXSMMSAMXMMSAMMSMMSASAXXAMMAMAMAMMSMSMAXMMMAXAXAMXSAXMAXSMXAMAXAAAMSSMSAS
SASXSMAMAMMSSSMASAMXAAMXSASMMAXXMAMXMAMXMAMAAMAXXSSMMAMAMXXSAMASXMASXXAXMASXXASASAXXAMXMXMXSAXAMAMXMXAAAXSMSSSMSMSMMXMASMXMAXMAMXMSMMMXMASXS
AXASXMAMAXMAMAMAMMMSSMSAMXMXAXSMSSMMSXMAXMMMMSMMAAAXXASAXSAMAXMMSMSMAMMMSMMXMASAMAXMXSAMMXMMMSMSMXAMXMMMMMAXAAXAAAAXAAXMASMMASMMXMAAXMAMXMXM
MXMSXSXMAMMXSSMMSAMXAAMASMSMSXSASMSAAAAMMSXMXAMMSMMSSMXMSMAMAMMAMXAAMXAAAMMXMAXXAMXSAMAXXAMAASMAMSASXSASAMSMMMSMSMSMMMMMAMMXAAXMXSAMMXAMAAAA
XSXSASAMMXSAAAAXSMXXXMSSMAAAXMMSMAAMMSXMAMAMMMXMAAXAASXSAMAMMSMASMXSMMSSSXMAMMSSMMXSASAMXXXASXXAXXXMASASMXMASAXXMAAAXXAMXSSMSSMSAXXMMSASXSMS
XMAMAMAMAAMMSMMMSASMSMSXMSMMMAXXMXMXXXXMASAMXAASMMMSMMXMAMXMAXMXMMSAAXMAXAXSSMAAAMMSXMASMXSMMMMSSMSMMMMMXAAXAXSAMSSXMMSXMMAAXXAMASMMXAXXAAXA
XMASASAMMXSAAAAAXAAAAXMAMAASMSMMMMXSMMASAXMSMSXSAMXMXMAMXMSMSSSMAXASXMMMMXMMAMSSMMAXMXAMMAAAMXAAAXAAXMASXSSMMSMMMAAXSXAMAXMMMMSMAMAASXMXMMSM
MSXSAMMXSAMMSSMMMMMSMXXMMSXMAXMASAAAAMMMSSMXMASMXMASAMASAAXAAAMSMXMAMXMASMMSAMXXAMXSAXSAMXMXMMMSXSMSMMMAAAAAXXAAMMSMSAMXXAXAXAAMSSSMSAMXAAXX
AXAMAMSAMMSAMXMMSMMAMSSSMXXMXMSSMMMSSMAAMXMMMXMSAMASASASMSMMMSMSSSXAMSXMMAAMMMMSXMASXMXAAMXXXMAMMXAMAAXMMSMMMMSMXAMAXSMAAXSSXMXSMXXASAMSAMSA
SMMMMXMASAXMSASAAASXSXAAASXMXMXMXAXMXXMSMAMASAASXSMSAMMMMMMAMMAMASAMXSAXSSMSXAXSXMMSASXSMMXMASAMXMXXAMXXAMMAXXMMSMMSMAMXSAAAASXMMAMMMAMXASXM
AXMAMASMMMSMSAMMSMMMMMMMMMAMAXAMMSSMAMMAXMSAXMMMXXAMMMXAAASMSMAMMXXMAMAMAXAMMSSMASASMMXAASMMMSAAXAAXSMMMXSXSXSAAAXMASMXAXMMMMMAAMXMASXMSXMAM
SXMASAXAAXAAMAMAMXMASMMSXSAMXMAXAAAMSMSMSXMMXSAASMXMMMSSSMSAAXXMXXSMSSMAMMSMAXAXXMASXMSXMMAAAXMMMSMMXAXSXMMMAXMXSMSASXMMXAMMAXMMMXSXMMMMXXAM
MASASXSSMSSSSSMASASXMAAXASAMMSSMMSSMAMAASMMAAMMSMSAXAXMAXAMXMMSSSMXAMAAXMAAXXSAMXMMMAMASMSSMSSMXAASAMSMMASAMSMSAMMMSSMASXMMSSSMMSAMAXMASMSSS
SAMMSAMAAXAAAXMASMMMSMMMMSAMAAAAXMAMSSMMMAMMSSXMASXSXSMXMXMAXAAAASMXMASMMSSSXMXMASASXMAMXAAAAXAMSSSSXXASAMMXMAMXSMMAXAMSAAAAXAMAMASMMSAXMAAM
MMSAMXSMMMMMMMAMSMXMASAMXSXMMSSMSAXMXXXXSXMMMAMMXMMAMXMAMXSXSMMSMMXAMXMAXAMXMXXSXMASXMXMMSMMMMSAMXMASXXMSSXSMSMMMMMMSSXSMMMSSXMASXMAXMXSMMMM
}