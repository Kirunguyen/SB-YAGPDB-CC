{{/*"Banner S" (0.5% / 1% / 2% / 96.5%)*/}}
{{/*SSR (0.5%) = List 10 roles ; UPRATE 70% for 1 role / 30% other roles || Pity 80 => Guaranteed SSR (0.5%)*/}}
{{/*SR (1%) = Role1*/}}
{{/*UC (2%) = Role2*/}}
{{/*C (96.5%) = Fail (Nothing)*/}}

{{/*Condition 1: Required Role*/}}
{{if not (hasRoleID 1260613756690497660)}}
    {{print "<:SBw_ualasao:1095366986252439613> Ủa chưa mua gói Roll này mà" .User.Mention}}
    {{return}}
  {{end}}

{{/*Condition 2: Required NOT Role*/}}
{{if or (hasRoleID 1261697133485101097) (hasRoleID 1262419286576595014)}}
    {{print "<:SBw_susphera:1102023844996325487> " .User.Mention " còn gói Reroll chưa dùng kìa, `.reroll` ngay và luôn "}}
    {{return}}
  {{end}}

{{deleteTrigger (0)}} {{/*Cleans user input*/}}

{{/*DATABASE*/}}
{{$listRoles := cslice "1215574466797633598" "1215575230437916672" "1215575453830479892" "1215575761847713843" "1215575898460524564" "1215576214350077952" "1215576506340741130" "1215576857148129310" "1215577085771255808" "1215577466534367322"}}

{{$listSSR := cslice
"<:Sephera00i:1262227258748178434> Role <@&1215574466797633598>" 
"<:Sephera01i:1262227261021491210> Role <@&1215575230437916672>"
"<:Sephera02i:1262227262800138341> Role <@&1215575453830479892>"
"<:Sephera03i:1262227264523735100> Role <@&1215575761847713843>"
"<:Sephera04i:1262227266499514379> Role <@&1215575898460524564>"
"<:Sephera05i:1262227268311191572> Role <@&1215576214350077952>"
"<:Sephera06i:1262227270987153428> Role <@&1215576506340741130>"
"<:Sephera07i:1262227273264660490> Role <@&1215576857148129310>"
"<:Sephera08i:1262227274971746348> Role <@&1215577085771255808>"
"<:Sephera09i:1262227276943327375> Role <@&1215577466534367322>"}}

{{$listFail := cslice
  "\n<:SBc_jzpa:1059378086170333195> xịt "
  "\n<:SBc_jztroi:1059377794305495121> xịt "
  "\n<:SBc_troioi:1059377704199278733> xịt "
  "\n<:SBw_buon:1261711059090145281> xịt "
  "\n<:SBw_chet:1106617687385317467> xịt "
  "\n<:SBw_chetdi:1102184551779930142> xịt "
  "\n<:SBw_deplao:1106824983956705380> xịt "
  "\n<:SBw_khoc:1098257413985026099> xịt "
  "\n<:SBw_khongthaygihet:1107710529927127040> xịt "
  "\n<:SBw_lasaoma:1104029826119106581> xịt "
  "\n<:SBw_maytoisoroi:1145406544758984744> xịt "
  "\n<:SBw_nailtancong:1156261340172795964> xịt "
  "\n<:SBw_nemgach:1109822847687327766> xịt "
  "\n<:SBw_nemgach2:1234544771780050967> xịt "
  "\n<:SBw_phangso:1059377630559879178> xịt "
  "\n<:SBw_quatday:1113025864867795004> xịt "
  "\n<:SBw_susphera:1102023844996325487> xịt "
}}

{{/*PROCESS 1 - Rolls*/}}
{{$rate := 0}} {{$gotSSR := false}} {{$gotSR := false}} {{$gotR := false}}

{{$output := print "> <:SBw_chao500ae:1107693029395009546> " .User.Mention " đã sử dụng **Gói x10 Roll Banner S**, bây giờ <@204255221017214977> roll ngay nè"}}
{{$mID := sendMessageRetID nil $output}} {{/*Intro - Retrieve ID to edit later*/}}

  {{/*PITY - YAGPDB's "DATABASE"*/}}
  {{if not (dbGet .User.ID "counterS")}} {{dbSet .User.ID "counterS" "0"}} {{end}} {{/*Generates DB for user if N/A*/}}

    {{/*Adds 10 to PITY*/}}
  {{$counter := toInt64 (dbGet .User.ID "counterS").Value}} 
  {{$counter = add $counter 10}}
  {{dbSet .User.ID "counterS" (toString $counter)}}
    
    {{/*If reaches PITY: GetSSR + Fix PITY value*/}}
  {{if ge $counter 80}}
    {{$gotSSR = true}}
    {{$output = print $output "\n ### Sephera Bảo Hộ - Nhận 01 Role Sephera <:SB_SepheraDot:1074619339904405504>"}}
    {{$counter = sub $counter 80}}
	  {{dbSet .User.ID "counterS" (toString $counter)}}
  {{end}}

  {{/*Loops 10x*/}}
{{- range 10 -}} 
  {{sleep 1}} {{/*Delay between messages*/}}
  {{$rate = randInt 1000}} {{/*MAIN CHARACTER*/}}

  {{if lt $rate 5}} {{/*GetSSR = Sephera*/}}
    {{$gotSSR = true}} {{$output = print $output "\n<:SB_SepheraDot:1074619339904405504> **Cơ Hội** trúng phần quà XỊN!"}}
    {{editMessage nil $mID $output}}
  
  {{else if lt $rate 15}} {{/*GetSR = Reroll S*/}}
    {{$gotSR = true}} {{$output = print $output "\n<:SB_SepheraDot:1074619339904405504> **Cơ Hội** trúng phần quà XỊN!"}}
    {{editMessage nil $mID $output}}

  {{else if lt $rate 35}} {{/*GetR = Reroll A*/}}
    {{$gotR = true}} {{$output = print $output "\n<:SB_SepheraDot:1074619339904405504> **Cơ Hội** trúng phần quà XỊN!"}}
    {{editMessage nil $mID $output}}
    
  {{else}} {{/*FAIL*/}}
    {{$output = print $output (index $listFail (randInt (len $listFail)))}}
    {{editMessage nil $mID $output}}
  {{end}}

{{- end -}}

{{/*OUTTRO / OUTPUT1*/}}
{{if or $gotSSR $gotSR $gotR}} {{/*If any prizes*/}}
  {{$output = print $output "\n\n> <:SBw_cuoi:1100415938031009792> Và phần quà XỊN của " .User.Mention " là..."}} {{sleep 2}}
{{else}}
  {{$output = print $output "\n> Chúc " .User.Mention " may mắn lần sau <:SBw_tangsone:1129646754984382504>"}} {{editMessage nil $mID $output}}
{{end}}

{{/*PROCESS 2 (Pick SSR) - OUTPUT2 (SSR/SR/R/UC)*/}}
{{$picker := 0}}{{$pickID := ""}}{{$pickName := ""}} 
{{$uprateID := 2}} {{/*Position of $listRoles to UPRATE (0 - 10)*/}}    


{{if $gotSSR}}
  {{$rate = randInt 100}} {{/*UP / DOWN rate*/}}

  {{if lt $rate 70}} {{/*UPRATE = 70%*/}}
    {{$pickID = index $listRoles ($uprateID)}} {{addRoleID $pickID}}
    {{$pickName = index $listSSR ($uprateID)}} 
    
    {{$output = print $output "\n" $pickName}} {{editMessage nil $mID $output}}

  {{else}} {{/*NOT UPRATE = 30%*/}}
    {{$picker = randInt (len $listSSR)}} {{while not (eq $picker $uprateID)}} {{$picker = randInt (len $listSSR)}} {{end}}
    {{$pickID = index $listRoles ($picker)}} {{addRoleID $pickID}}
    {{$pickName = index $listSSR ($picker)}}
    
    {{$output = print $output "\n" $pickName}} {{editMessage nil $mID $output}}
  {{end}}

{{end}} {{/*SSR Ends*/}}

{{if $gotSR}}
  {{addRoleID 1261697133485101097}}
      
  {{$output = print $output "\n:regional_indicator_s: <@&1261697133485101097>, sử dụng: `.reroll`"}} {{editMessage nil $mID $output}}
{{end}}

{{if $gotR}}
  {{addRoleID 1262419286576595014}}
		  
  {{$output = print $output "\n:regional_indicator_a: <@&1262419286576595014>, sử dụng: `.reroll`"}} {{editMessage nil $mID $output}}
{{end}}
 
{{$output = print $output "\n-# Tiến độ nhận bảo hộ: [" $counter "/80]"}} {{editMessage nil $mID $output}}

{{removeRoleID 1260613756690497660}} {{/*Condition 1*/}}
