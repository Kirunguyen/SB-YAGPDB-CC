{{/*DATABASE*/}}
{{$listHero := (cslice
	(cslice "<:Airi:1183236796293849139> Airi" "<:Aoi:1183236816510390292> Aoi" "<:Astrid:1183236832306143374> Astrid" "<:Bright:1183236858059178054> Bright" "<:Butterfly:1183236863952166963> Butterfly" "<:Enzo:1183236904771125358> Enzo" "<:Kaine:1183236958793777222> Kaine" "<:Keera:1183236964518989945> Keera" "<:Kriknak:1183236972790157342> Kriknak" "<:Liliana:1183238064437481522> Liliana" "<:Murad:1183238112416104560> Murad" "<:Nakroth:1183238114345496577> Nakroth" "<:Paine:1183238131781206056> Paine" "<:Qi:1183238138345304138> Qi" "<:Quillen:1183238142518628362> Quillen" "<:Raz:1183238147103002634> Raz" "<:Ryoma:1183238163444015205> Ryoma" "<:Sinestrea:1183238170494631987> Sinestrea" "<:TheFlash:1183238203298304030> The Flash" "<:Volkath:1183238235955142716> Volkath" "<:Wukong:1183238470089584720> Ngộ Không" "<:Yan:1183238475236003894> Yan" "<:Zata:1183238500431188068> Zata" "<:Zephys:1183238504034082916> Zephys" "<:Zill:1183238505871183962> Zill" "<:Zuka:1183238513676779541> Zuka")
	
	(cslice "<:Airi:1183236796293849139> Airi" "<:Allain:1183236807014486127> Allain" "<:Amily:1183236811070390272> Amily" "<:Arduin:1183236819484164107> Arduin" "<:Arthur:1183236823829463040> Arthur" "<:Bijan:1183236847741177926> Bijan" "<:Butterfly:1183236863952166963> Butterfly" "<:Charlotte:1249726113651884114> Charlotte" "<:Dextra:1183236886404288522> Dextra" "<:Errol:1183236908889935962> Errol" "<:Florentino:1183236918754938900> Florentino" "<:KilGroth:1183236968839131226> Kil'Groth" "<:LữBố:1183238077829886056> Lữ Bố" "<:Maloch:1183238087430635550> Maloch" "<:Omen:1183238124600561724> Omen" "<:Qi:1183238138345304138> Qi" "<:Richter:1183238150793998357> Richter" "<:Rourke:1183238157869776916> Rourke" "<:Roxie:1183238161502052392> Roxie" "<:Skud:1183238172491137045> Skud" "<:Superman:1183238180556771348> Superman" "<:Taara:1183238184948220056> Taara" "<:Tachi:1183238186693042288> Tachi" "<:Veres:1183238227205820436> Veres" "<:Volkath:1183238235955142716> Volkath" "<:Wiro:1183238460153278494> Wiro" "<:WonderWoman:1183238467690434590> Wonder Woman" "<:Yena:1183238485633671188> Yena" "<:Zanis:1183238495913906206> Triệu Vân" "<:Zephys:1183238504034082916> Zephys" "<:Zuka:1183238513676779541> Zuka")
	
	(cslice "<:Arduin:1183236819484164107> Arduin" "<:Ata:1183236834604613632> Ata" "<:Baldum:1183236844155052053> Baldum" "<:Chaugnar:1183236874144337970> Chaugnar" "<:Cresht:1183236880427384832> Cresht"  "<:Dextra:1183236886404288522> Dextra" "<:Gildur:1183236921007276042> Gildur" "<:Grakk:1183236924731838554> Grakk" "<:Lumburr:1183238081957085264> Lumburr" "<:Maloch:1183238087430635550> Maloch" "<:Max:1183238093667586108> Max" "<:Mina:1183238102446260304> Mina" "<:Omega:1183238122155278376> Omega" "<:Ormarr:1183238127838560326> Ormarr" "<:Roxie:1183238161502052392> Roxie" "<:Skud:1183238172491137045> Skud" "<:Taara:1183238184948220056> Taara" "<:Thane:1183238200009957487> Thane" "<:Toro:1183238209224839258> Toro" "<:Wiro:1183238460153278494> Wiro" "<:YBneth:1183238479711305748> Y'Bneth")

	(cslice  "<:Alice:1183236804858626058> Alice" "<:Arum:1183236828090863626> Arum" "<:Aya:1183236838203342858> Aya" "<:Baldum:1183236844155052053> Baldum" "<:Chaugnar:1183236874144337970> Chaugnar" "<:Cresht:1183236880427384832> Cresht" "<:Gildur:1183236921007276042> Gildur" "<:Grakk:1183236924731838554> Grakk" "<:Helen:1183236931975393363> Helen" "<:Ishar:1183236944168243250> Ishar" "<:Krizzix:1183236979878547546> Krizzix" "<:Lumburr:1183238081957085264> Lumburr" "<:Mina:1183238102446260304> Mina" "<:Rouie:1183238154019426355> Rouie" "<:Sephera:1183238166522630257> Sephera" "<:Teemee:1183238190757314621> Teemee" "<:Toro:1183238209224839258> Toro" "<:Xeniel:1183238472291590256> Xeniel" "<:Zip:1183238509268574349> Zip")

	(cslice "<:Capheny:1183236867760603237> Capheny" "<:Celica:1183236870554005514> Celica" "<:Elandorr:1183236896827121674> Eland'orr" "<:Elsu:1183236901533139135> Elsu" "<:Erin:1249726252382818375> Erin" "<:Fennik:1183236914187345931> Fennik"  "<:Hayate:1183236928267620463> Hayate" "<:Laville:1183238061216256152> Laville" "<:Lindis:1183238069638402058> Lindis" "<:Stuart:1249728135293964340> Stuart" "<:TelAnnas:1183238194062438460> Tel'Annas" "<:Thorne:1183238206842490900> Thorne" "<:Valhein:1183238215499517982> Valhein" "<:Violet:1183238230737424434> Violet" "<:Wisp:1183238463693267024> Wisp" "<:Yorn:1183238489198837760> Yorn")

	(cslice "<:Aleister:1183236801591259236> Aleister" "<:Annette:1183236814534885396> Annette" "<:AzzenKa:1183236841403580437> Azzen'Ka" "<:Bonnie:1183236851616731226> Bonnie" "<:DArcy:1183236882700705872> D'Arcy" "<:Diaochan:1183236890019758190> Điêu Thuyền" "<:Dirak:1183236894717395055> Dirak" "<:Iggy:1183236933904781362> Iggy" "<:Ignis:1183236937864183829> Ignis" "<:Ilumia:1183236940317863936> Ilumia" "<:Ishar:1183236944168243250> Ishar" "<:Jinna:1183236947955687424> Jinna" "<:Kahlii:1183236953953550507> Kahlii" "<:Krixi:1183236975583576094> Krixi" "<:Lauriel:1183238058196349049> Lauriel" "<:Liliana:1183238064437481522> Liliana" "<:Lorion:1183238075166498897> Lorion" "<:Marja:1183238091104866345> Marja" "<:Mganga:1183238097316614214> Mganga" "<:Natalya:1183238118573342720> Natalya" "<:Preyta:1183238135874859058> Preyta" "<:Raz:1183238147103002634> Raz" "<:Sephera:1183238166522630257> Sephera" "<:Tulen:1183238213356228738> Tulen" "<:Veera:1183238221736456263> Veera" "<:Yue:1183238492109672609> Yue" "<:Zata:1183238500431188068> Zata")
)}}

{{$listClass := (cslice "Sát Thủ" "Đấu Sĩ" "Đỡ Đòn" "Trợ Thủ" "Xạ Thủ" "Pháp Sư")}}
{{$listURL := (cslice "https://cdn.discordapp.com/emojis/1114474554119688222.png" "https://cdn.discordapp.com/emojis/1114478367887065159.png" "https://cdn.discordapp.com/emojis/1114537536153464893.png" "https://cdn.discordapp.com/emojis/1114478670946500718.png" "https://cdn.discordapp.com/emojis/1114478764680826940.png" "https://cdn.discordapp.com/emojis/1114478533608214578.png")}}

{{/*PROCESS*/}}
{{$duelist := parseArgs 2 "> Cú pháp đúng là **.solo [@người_1] [@người_2]**"  
    (carg "user" "duelist A")
    (carg "user" "duelist B")}}

{{$rand := randInt 6}} {{/*"Main Character*/}}

{{$pClass := index $listHero ($rand)}} {{/*Pick a class*/}}

{{$pick1 := index $pClass (randInt (len $pClass))}} {{/*Pick 2 random heroes from the selected class*/}}
{{$pick2 := index $pClass (randInt (len $pClass))}}

{{/*OUTPUT*/}}
{{$sideA := (($duelist.Get 0).Mention)}}
{{$sideB := (($duelist.Get 1).Mention)}}

{{$fields:= (cslice 
	(sdict "name" $pick1 "value" $sideA "inline" true) 
	(sdict "name" ":crossed_swords:" "value" "**vs**" "inline" true) 
	(sdict "name" $pick2 "value" $sideB "inline" true) 
)}}

{{$embed := cembed 
	"title" (print "QUYẾT CHIẾN! Trận SOLO gay cấn của hai **" (index $listClass ($rand)) "**!")
	"fields" $fields
	"thumbnail" (sdict "url" (index $listURL ($rand)))
	"footer" (sdict 
		"text" "Sephera's Biển" 
		"icon_url" "https://media.discordapp.net/attachments/1270931389256044606/1271475546576130058/Sephera_Cat.png") 
}}
{{sendMessage nil $embed}}
