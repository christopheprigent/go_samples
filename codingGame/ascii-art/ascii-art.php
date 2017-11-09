<?php
fscanf(STDIN, "%d",  $L);
fscanf(STDIN, "%d",  $H);
$Text       = strtoupper(stream_get_line(STDIN, 256 + 1, "\n"));
$Char2ACSII = Array();
/*
Array 
(
    [A] => Array
        (
            [0] =>  #  
            [1] => # # 
            [2] => ### 
            [3] => # # 
            [4] => # # 
            )

    [...]

    [?] => Array
        (
            [0] => ### 
            [1] =>   # 
            [2] =>  ## 
            [3] =>     
            [4] =>  #  
        )
)
*/

// LOAD $Char2ACSII
for ($i = 0; $i < $H; $i++)
{
    $ROW = stream_get_line(STDIN, 1024 + 1, "\n");
    foreach (range(0, 26) as $number) {
        $curChar = chr($number+ord('A'));
        if ($curChar == '[')$curChar = '?';
        if (!is_array($Char2ACSII[$curChar])) $Char2ACSII[$curChar]=Array();
        $Char2ACSII[$curChar][$i]=substr($ROW, $L * $number, $L);
    }
}

//error_log(print_r($Char2ACSII,1));
printText2ACSII($Text,$Char2ACSII);


/*
 * translate and print '$str' using ref '$Char2ACSII'
 */
function printText2ACSII($str, $Char2ACSII){
    $lenT   = strlen($str);
    $H      = count($Char2ACSII['A']);
    
    for ($l=0;$l<$H;$l++){
        $line='';
        for ($i=0;$i<$lenT;$i++){
            $idx=$str[$i];
            if (!$Char2ACSII[$idx]) $idx = '?';
            $line.=$Char2ACSII[$idx][$l];
        }
       echo $line . "\n";
    }
}