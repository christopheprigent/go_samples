<?php

$M      = stream_get_line(STDIN, 100 + 1, "\n");
$bin    = '';
error_log($M);

foreach(str_split($M) as $c){
    $bin .= sprintf( "%07d", decbin(ord($c)));
}
echo(bin2chuck($bin)."\n");

function bin2chuck($b){
    $chuck  = '';
    $i      = 0;
    $l      = strlen($b);

    error_log($b);
    while($i<$l){
        $c = $b[$i];
        if ($c=='0')
            $chuck .= '00 ';
        else
            $chuck .= '0 ';

        do {
            $i++;
            $chuck .= '0';
        } while($i<$l && $c==$b[$i]);
        if ($i<$l)
        $chuck .= ' ';
    }

    error_log("'$b'=>'$chuck'");
    return $chuck;
    }
?>