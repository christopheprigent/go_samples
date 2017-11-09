var N = parseInt(readline()),
    pi = [],
    prev = false,
    minDiff;
for (var i = 0; i < N; i++) {
    pi.push(parseInt(readline()));
}

for (var p of pi.sort(function(a, b) { return a - b; })) {
    if (!prev || minDiff > p - prev)
        minDiff = p - prev;

    prev = p;
}
print(minDiff);