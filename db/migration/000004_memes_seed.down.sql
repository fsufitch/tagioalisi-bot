-- Auto-generated by ./memes-seed/build_sql_down.go
-- Using memes from https://old.reddit.com/r/image_linker_bot/comments/2znbrg/image_suggestion_thread_20/

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('truestory')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('lennyface')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('aliens', 'aliensguy')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('spidermanneat', 'neat')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('soon')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('tmyk', 'themoreyouknow')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('thatsracist', 'datsracist', 'dasracist')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('mybodyisready')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('thisisfine')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('itsatrap', 'admiralackbaritsatrap', 'ackbar', 'admiralakbar', 'ackbaritsatrap', 'trap', 'admiralackbar', 'atrap', 'akbar', 'ackbartrap')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('thatsapenisreverse', 'reversethatsapenis', 'sinepastaht', 'thatsapenis2')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('thatsmyfetish', 'fetish', 'myfetish', 'thisismyfetish')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('idontbelieveyou', 'idontbeleiveyou', 'dontbelieveyou', 'anchormanidontbelieveyou')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('costanza')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('disgonbgud', 'thisgonnabegood', 'disgunbegoodguywithfoldingchair', 'disgunbegoof', 'disgunbegood', 'disgonbgood', 'disgunbegud', 'disgonbegud', 'disgonbegood', 'thisgunbegood', 'thisgonbegood')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('absolutelydisgusting')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('deaddove', 'idontknowwhatiexpected', 'dontknowwhatiexpected')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('doubt')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('slowclap', 'claps', 'slowcap')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('whoosh', 'woosh')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('shutupandtakemymoney', 'frytakemymoney', 'takemymoney', 'fryshutupandtakemymoney')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('both')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('notbad', 'obamanotbad')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('ohyou')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('abandonthread', 'fuckthisshitimgoinhome', 'exitthread')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('feels', 'myfeels')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('michaelscottno', 'godno', 'nogodno')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('idontgiveashit', 'dontgiveashit')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('sadtennant', 'saddoctor', 'tennantrain', 'saddrwho', 'saddoctorwho', 'doctorwhorain', 'drwhocrying', 'doctorwhocrying')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('redditsilver')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('colbertpopcorn')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('iknowsomeofthesewords')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('thanksobama')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('foreveralone')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('nathanfillion', 'firefly', 'nevermind', 'speechless', 'nathanfilionspeechless', 'nathanfilion', 'nathanfillionspeechless', 'fireflyguy')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('ohgodwhy')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('lowqualitybait')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('dealwithit')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('dickbutt')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('sadkenau')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('youtried', 'utried', 'tried', 'youtriedstar')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('thatsapenis')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('trollface')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('whynotboth', 'porquenodos', 'porquenolasdos', 'porquenolosdos')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('oneartplease', 'artplease', 'zoidbergart', 'zoidbergartplease', '1artplease')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('itssomething')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('thatescalatedquickly')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('okay', 'ok')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('notthebees')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('feelsbadman', 'feelsbad', 'sadface')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('youregoddamnright', 'breakingbadyouregoddamright', 'goddamnedright', 'youregoddamnedright', 'goddamnright', 'yourgoddamnright', 'heisenberggoddamnright', 'yourgoddamnedright', 'breakingbadgoddamnright', 'breakingbadyouregoddamnright', 'heisenbergyouregoddamnright')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('iwanttobelieve')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('badjokecena', 'johncenastandup', 'cenajokeface', 'cenareaction')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('goodfellas', 'goodfellows')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('ducklaugh', 'duck')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('notsureifserious', 'notsureifsrs')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('ihavenoideawhatimdoing', 'chemistrydog', 'ihavenoideawhatiamdoing', 'whatamidoing', 'ivenoideawhatimdoing')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('iunderstoodthatreference', 'captainamericaiunderstoodthatreference', 'capunderstoodreference', 'reference', 'iunderstoodthatreferencecaptamerica', 'iknowthatreference', 'gotthatreference', 'inuderstoodthatreference', 'igerthatreference', 'iunderstandthisreference', 'iunderstoodthisreference', 'captainamericareference', 'iunderstandthatreference', 'igotthatreference', 'iknowthisreference', 'igetthatreference', 'igetthisreference', 'captainamerica', 'referenceunderstood')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('ivemadeahugemistake', 'hugemistake', 'imadeahugemistake')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('jlawyeahright', 'jlawrenceokay', 'jenniferlawrenceokthumbsup', 'unimpressedjenniferlawrence', 'jenniferlawrenceyeahokay', 'jlawok', 'jenniferlawrenceokayyeah', 'jenniferlawrenceok', 'jlawnodding', 'jlawsarcasticohyeah', 'jlawokay', 'jlawyeahok', 'jlawrenceok', 'jenniferlawrencethumbsup', 'jlawokthumbsup', 'jenniferlawrenceokay', 'jlawohyeahsure')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('butwhy')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('michaeljacksonpopcorn', 'michaeljackson', 'popcornmj', 'mjpopcorn', 'mjpcorn')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('rollsafe')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('youdarealmvp', 'youtherealmvp', 'realmvp', 'therealmvp', 'darealmvp')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('nope', 'no')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('2spooky4me', '2spookyforme', 'toospooky', '2spooky', 'toospookyforme', 'toospooky4me')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('ifyouknowwhatimean')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('dozensofus', 'therearedozensofus', 'tobiasdozens')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('itshappening', 'ronpaul', 'ronpaulitshappening', 'happening')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('coolstorybro', 'coolstory')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('facepalm')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('slowpoke')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('mindblown', 'brainexploding', 'headexploding')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('whatyearisit', 'whatyearisthis')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('thatsthejoke', 'thatwasthejoke', 'thatsthepoint', 'thejoke')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('sensiblechuckle')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('yeahrightsure', 'yarightsure', 'yearightsure', 'yeahsure', 'yeahright')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('sickreferencebro', 'greatreference', 'sickrefrence', 'goodreference', 'referencesoutofcontrol', 'sweetreferencebro', 'sickreference', 'sweetreference', 'sickrefrencebro', 'yourreferencesareoutofcontrol', 'broyourreferencesareoutofcontrol')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('youdontsay', 'udontsay', 'yadontsay')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('ayylmao')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('heavybreathing')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('areyouawizard')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('rekt')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('feelsgoodman', 'feelsgood')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('iseewhatyoudidthere', 'seewhatyoudidthere', 'iseewhatudidthere', 'iseewatudidthere', 'iseewatyoudidthere', 'seewhatudidthere')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('nowkiss', 'nowkith')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('picardfacepalm')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('popcorn')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('motherofgod')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);

WITH meme_name AS (
	SELECT * FROM meme_names WHERE name IN ('conspiracykeanu')
)
DELETE FROM memes WHERE id IN (select meme_id from meme_name);
