var M = readline(),
    bin = '';

for (var i in M)
    bin += ("0000000" + M[i].charCodeAt(0).toString(2)).slice(-7);
print(bin2Chuck(bin));

function bin2Chuck(b) {
    var ret = '',
        i = 0;

    while (i < b.length) {
        var next = parseInt(i) + 1,
            cpt = 1;

        if (i > 0 &&
            i < b.length) ret += ' ';
        if (b[i] === '0') ret += '00 0';
        else ret += '0 0';
        while (b[next] && b[next] === b[i]) {
            i = next;
            next++;
            ret += '0';
        }
        i++;
    }

    return ret;
}