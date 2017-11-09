var L = parseInt(readline()),
    H = parseInt(readline()),
    T = readline(),
    F = [];

for (var i = 0; i < H; i++) {
    var ROW = readline();
    for (var ii = 0; ii <= 26; ii++) {
        if (!F[ii]) F[ii] = [];
        (F[ii]).push(ROW.substr(ii * L, L));
    }
}

for (var i = 0; i < H; i++) {
    line = '';
    for (var C of T.split('')) {
        var idx = C.toUpperCase().charCodeAt() - 'A'.charCodeAt();
        if (!F[idx]) idx = 26;
        line += F[idx][i];
    }
    print(line);
}