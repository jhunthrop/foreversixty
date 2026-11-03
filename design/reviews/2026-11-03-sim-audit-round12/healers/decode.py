"""decode.py <class> <site talent string>: name every point."""
import json,sys
S='/Users/jh/code/forever/.claude/worktrees/r12-healers'
d=json.load(open(f'{S}/data/builds/1.60.1.70753/talents/{sys.argv[1]}.json'))
trees=d['trees'] if 'trees' in d else d
for tree,part in zip(trees,sys.argv[2].split('-')):
    tal=sorted(tree['talents'],key=lambda t:(t['tier'],t['column']))
    out=[f"{t['name']} {c}/{t['max_rank']} [{t['ranks'][0]['spell_id']}]" for t,c in zip(tal,part) if c!='0']
    print(tree.get('name'),':',', '.join(out))
