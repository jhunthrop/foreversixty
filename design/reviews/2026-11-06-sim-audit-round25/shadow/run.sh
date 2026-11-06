#!/bin/bash
# order level band faction seed dps err
R=/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/r25; S=$R/shadow
cd /Users/jh/code/forever/.claude/worktrees/r25-shadow/sim
OUT=${OUT:-$S/results.tsv}
FROM=${FROM:-10}; TO=${TO:-41}
grep "^L " $S/ladders.txt | while read _ lvl a b d; do
  [ $lvl -lt $FROM -o $lvl -gt $TO ] && continue
  if [ $lvl -lt 30 ]; then band=20; elif [ $lvl -lt 40 ]; then band=30; elif [ $lvl -lt 50 ]; then band=40; elif [ $lvl -lt 60 ]; then band=50; else band=60; fi
  for seed in 1 50001; do for fac in alliance horde; do for o in a b d; do
    case $o in a) t=$a;; b) t=$b;; d) t=$d;; esac
    line=$($R/sim.sh $S/sb -spec priest-shadow -faction $fac -preset bare -band $band -level $lvl -talents $t -bis-dir $R/bis-frozen -iterations 10000 -seed $seed $EXTRA 2>/dev/null </dev/null | grep -m1 "DPS")
    v=$(echo "$line" | awk '{print $3"\t"$5}')
    printf "%s\t%s\t%s\t%s\t%s\t%s\n" $o $lvl $band $fac $seed "$v" >> $OUT
  done; done; done
done
echo finished >> $S/run.done
