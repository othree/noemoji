package noemoji

import (
  "unicode/utf8"
)

// Mapping from character to concrete escape code.
// Taken from github.com/kyokomi/emoji
var emojiMap = map[string]string{
	"\U00002651": ":capricorn:",
	"\U0001f51a": ":end:",
	"\U0001f4f5": ":no_mobile_phones:",
	"\U0001f46b": ":couple:",
	"\U000026c4": ":snowman:",
	"\U0001f304": ":sunrise_over_mountains:",
	"\U0001f69f": ":suspension_railway:",
	"\U0001f504": ":arrows_counterclockwise:",
	"\U0001f41b": ":bug:",
	"\U0001f615": ":confused:",
	"\U0001f457": ":dress:",
	// "\U0001f41d": ":honeybee:",
	"\U0001f318": ":waning_crescent_moon:",
	"\U0001f388": ":balloon:",
	"\U0001f68c": ":bus:",
	"\U0001f4e6": ":package:",
	"\U0000270f": ":pencil2:",
	"\U0001f621": ":rage:",
	"\U0001f47e": ":space_invader:",
	"\U000025fd": ":white_medium_small_square:",
	"\U000023e9": ":fast_forward:",
	"\U0001f358": ":rice_cracker:",
	"\U0001f4e8": ":incoming_envelope:",
	"\U0001f202": ":sa:",
	"\U0001f6ba": ":womens:",
	"\U000027a1": ":arrow_right:",
	"\U0001f477": ":construction_worker:",
	"\U0001f3b6": ":notes:",
	"\U0001f410": ":goat:",
	"\U00002754": ":grey_question:",
	"\U0001f3ee": ":lantern:",
	"\U0001f391": ":rice_scene:",
	"\U0001f3c3": ":running:",
	"\U0001f3a1": ":ferris_wheel:",
	"\U0001f3bc": ":musical_score:",
	"\U00002747": ":sparkle:",
	"\U0001f609": ":wink:",
	"\U0001f3a8": ":art:",
	"\U0001f55e": ":clock330:",
	"\U0001f4bd": ":minidisc:",
	"\U0001f6ab": ":no_entry_sign:",
	"\U0001f390": ":wind_chime:",
	"\U0001f300": ":cyclone:",
	"\U0001f33f": ":herb:",
	"\U0001f406": ":leopard:",
	"\U0001f34c": ":banana:",
	"\U0001f45c": ":handbag:",
	"\U0001f36f": ":honey_pot:",
	"\U0001f197": ":ok:",
	"\U00002665": ":hearts:",
	"\U0001f6c2": ":passport_control:",
	"\U0001f5ff": ":moyai:",
	"\U0001f604": ":smile:",
	"\U0001f405": ":tiger2:",
	"\U0001f500": ":twisted_rightwards_arrows:",
	"\U0001f6b8": ":children_crossing:",
	"\U0001f42e": ":cow:",
	"\U0000261d": ":point_up:",
	"\U0001f3e0": ":house:",
	"\U0001f473": ":man_with_turban:",
	"\U0001f69e": ":mountain_railway:",
	"\U0001f4f3": ":vibration_mode:",
	"\U0001f421": ":blowfish:",
	"\U0001f1ee\U0001f1f9": ":it:",
	"\U0001f362": ":oden:",
	"\U0001f552": ":clock3:",
	"\U0001f36d": ":lollipop:",
	"\U0001f68b": ":train:",
	"\U00002702": ":scissors:",
	"\U0001f4d0": ":triangular_ruler:",
	"\U0001f492": ":wedding:",
	"\U0001f526": ":flashlight:",
	"\U00003299": ":secret:",
	"\U0001f363": ":sushi:",
	"\U0001f699": ":blue_car:",
	"\U0001f4bf": ":cd:",
	"\U0001f30c": ":milky_way:",
	"\U0001f393": ":mortar_board:",
	"\U0001f451": ":crown:",
	"\U0001f4ac": ":speech_balloon:",
	"\U0001f371": ":bento:",
	"\U00002755": ":grey_exclamation:",
	"\U0001f3e8": ":hotel:",
	"\U0001f51f": ":keycap_ten:",
	"\U0001f4f0": ":newspaper:",
	"\U0001f4e4": ":outbox_tray:",
	"\U0001f40e": ":racehorse:",
	"\U0001f606": ":laughing:",
	"\U00002b1b": ":black_large_square:",
	"\U0001f4da": ":books:",
	"\U00002733": ":eight_spoked_asterisk:",
	"\U00002714": ":heavy_check_mark:",
	"\U000024c2": ":m:",
	"\U0001f44b": ":wave:",
	"\U0001f6b4": ":bicyclist:",
	"\U0001f378": ":cocktail:",
	"\U0001f3f0": ":european_castle:",
	"\U0001f447": ":point_down:",
	"\U0001f5fc": ":tokyo_tower:",
	"\U0001f50b": ":battery:",
	"\U0001f483": ":dancer:",
	"\U0001f501": ":repeat:",
	"\U0001f1f7\U0001f1fa": ":ru:",
	"\U0001f311": ":new_moon:",
	"\U000026ea": ":church:",
	"\U0001f4c5": ":date:",
	"\U0001f30e": ":earth_americas:",
	"\U0001f463": ":footprints:",
	"\U0000264e": ":libra:",
	"\U0001f6a0": ":mountain_cableway:",
	"\U0001f53b": ":small_red_triangle_down:",
	"\U0001f51d": ":top:",
	"\U0001f60e": ":sunglasses:",
	"\U0001f521": ":abcd:",
	"\U0001f191": ":cl:",
	"\U0001f3bf": ":ski:",
	"\U0001f4d6": ":book:",
	"\U000023f3": ":hourglass_flowing_sand:",
	"\U0001f61d": ":stuck_out_tongue_closed_eyes:",
	"\U0001f630": ":cold_sweat:",
	"\U0001f3a7": ":headphones:",
	"\U0001f38a": ":confetti_ball:",
	"\U0000264a": ":gemini:",
	"\U0001f195": ":new:",
	"\U0001f64f": ":pray:",
	"\U0000231a": ":watch:",
	"\U00002615": ":coffee:",
	"\U0001f47b": ":ghost:",
	"\U0001f51b": ":on:",
	"\U0001f45d": ":pouch:",
	"\U0001f695": ":taxi:",
	// "\U0001f52a": ":hocho:",
	"\U0001f60b": ":yum:",
	"\U00002795": ":heavy_plus_sign:",
	"\U0001f389": ":tada:",
	"\U00002935": ":arrow_heading_down:",
	"\U0001f560": ":clock530:",
	"\U0001f357": ":poultry_leg:",
	"\U0001f418": ":elephant:",
	// "\U0001f1ec\U0001f1e7": ":gb:",
	"\U0001f004": ":mahjong:",
	"\U0001f35a": ":rice:",
	"\U0001f3b5": ":musical_note:",
	"\U0001f530": ":beginner:",
	"\U0001f53a": ":small_red_triangle:",
	"\U0001f345": ":tomato:",
	"\U0001f566": ":clock1130:",
	"\U0001f3ef": ":japanese_castle:",
	"\U0001f31e": ":sun_with_face:",
	"\U00000034\U000020e3": ":four:",
	"\U0001f3a4": ":microphone:",
	"\U0001f3be": ":tennis:",
	"\U00002195": ":arrow_up_down:",
	"\U0001f1e8\U0001f1f3": ":cn:",
	"\U0001f3c7": ":horse_racing:",
	"\U0001f6b3": ":no_bicycles:",
	"\U0001f40c": ":snail:",
	"\U0001f193": ":free:",
	"\U0001f41e": ":beetle:",
	"\U000025aa": ":black_small_square:",
	"\U0001f4c1": ":file_folder:",
	"\U0001f62f": ":hushed:",
	"\U0001f480": ":skull:",
	"\U0001f18e": ":ab:",
	"\U0001f680": ":rocket:",
	"\U0001f360": ":sweet_potato:",
	"\U0001f3b8": ":guitar:",
	"\U0001f429": ":poodle:",
	"\U0001f337": ":tulip:",
	"\U0001f536": ":large_orange_diamond:",
	":-1:": "\U0001f44e",
	"\U0001f4c8": ":chart_with_upwards_trend:",
	"\U0001f1e9\U0001f1ea": ":de:",
	"\U0001f347": ":grapes:",
	"\U0001f250": ":ideograph_advantage:",
	"\U0001f479": ":japanese_ogre:",
	// "\U0000260e": ":telephone:",
	"\U0001f55d": ":clock230:",
	"\U0000231b": ":hourglass:",
	"\U000021a9": ":leftwards_arrow_with_hook:",
	"\U0001f387": ":sparkler:",
	"\U0001f0cf": ":black_joker:",
	"\U0001f562": ":clock730:",
	"\U0001f31b": ":first_quarter_moon_with_face:",
	"\U0001f468": ":man:",
	"\U0001f553": ":clock4:",
	"\U0001f3a3": ":fishing_pole_and_fish:",
	"\U0001f3a9": ":tophat:",
	"\U000025fb": ":white_medium_square:",
	"\U0001f4e3": ":mega:",
	"\U0001f35d": ":spaghetti:",
	"\U0001f3af": ":dart:",
	"\U0001f467": ":girl:",
	"\U0001f452": ":womans_hat:",
	"\U0001f685": ":bullettrain_front:",
	"\U0001f3ec": ":department_store:",
	"\U0001f493": ":heartbeat:",
	"\U0001f334": ":palm_tree:",
	"\U0001f3ca": ":swimmer:",
	"\U0001f49b": ":yellow_heart:",
	"\U00002197": ":arrow_upper_right:",
	"\U0001f551": ":clock2:",
	"\U0001f460": ":high_heel:",
	"\U000023eb": ":arrow_double_up:",
	"\U0001f622": ":cry:",
	"\U0001f4c0": ":dvd:",
	":e-mail:":                  "\U0001f4e7",
	"\U0001f37c": ":baby_bottle:",
	"\U0001f192": ":cool:",
	"\U0001f4be": ":floppy_disk:",
	"\U0001f4f1": ":iphone:",
	"\U0001f690": ":minibus:",
	"\U0001f413": ":rooster:",
	"\U00000033\U000020e3": ":three:",
	"\U000025ab": ":white_small_square:",
	"\U0000264b": ":cancer:",
	"\U00002753": ":question:",
	"\U0001f376": ":sake:",
	"\U0001f382": ":birthday:",
	"\U0001f415": ":dog2:",
	"\U0001f4e2": ":loudspeaker:",
	"\U0001f53c": ":arrow_up_small:",
	"\U0001f42b": ":camel:",
	"\U0001f428": ":koala:",
	"\U0001f50e": ":mag_right:",
	"\U000026bd": ":soccer:",
	"\U0001f6b2": ":bike:",
	"\U0001f33e": ":ear_of_rice:",
	"\U0001f4a9": ":shit:",
	"\U0001f232": ":u7981:",
	"\U0001f6c0": ":bath:",
	"\U0001f476": ":baby:",
	"\U0001f50f": ":lock_with_ink_pen:",
	"\U0001f454": ":necktie:",
	"\U0001f459": ":bikini:",
	"\U0001f60a": ":blush:",
	"\U0001f497": ":heartpulse:",
	"\U0001f43d": ":pig_nose:",
	"\U0001f4cf": ":straight_ruler:",
	"\U0001f235": ":u6e80:",
	"\U0001f381": ":gift:",
	"\U0001f6a5": ":traffic_light:",
	"\U0001f33a": ":hibiscus:",
	"\U0001f491": ":couple_with_heart:",
	"\U0001f4cc": ":pushpin:",
	"\U0001f236": ":u6709:",
	"\U0001f6b6": ":walking:",
	"\U0001f600": ":grinning:",
	"\U00000023\U000020e3": ":hash:",
	"\U0001f518": ":radio_button:",
	// "\U0000270b": ":raised_hand:",
	"\U0001f367": ":shaved_ice:",
	"\U0001f488": ":barber:",
	"\U0001f431": ":cat:",
	// "\U00002757": ":heavy_exclamation_mark:",
	"\U0001f368": ":ice_cream:",
	"\U0001f637": ":mask:",
	"\U0001f416": ":pig2:",
	"\U0001f6a9": ":triangular_flag_on_post:",
	"\U00002196": ":arrow_upper_left:",
	"\U0001f41d": ":bee:",
	"\U0001f37a": ":beer:",
	"\U00002712": ":black_nib:",
	"\U00002757": ":exclamation:",
	"\U0001f436": ":dog:",
	"\U0001f525": ":fire:",
	"\U0001f41c": ":ant:",
	"\U0001f494": ":broken_heart:",
	"\U0001f4b9": ":chart:",
	"\U0001f550": ":clock1:",
	"\U0001f4a3": ":bomb:",
	"\U0000264d": ":virgo:",
	"\U0001f170": ":a:",
	"\U0001f374": ":fork_and_knife:",
	"\U000000a9": ":copyright:",
	"\U000027b0": ":curly_loop:",
	"\U0001f315": ":full_moon:",
	"\U0001f45e": ":shoe:",
	"\U0001f3e4": ":european_post_office:",
	"\U0001f196": ":ng:",
	"\U0001f3e2": ":office:",
	"\U0001f64b": ":raising_hand:",
	"\U0001f49e": ":revolving_hearts:",
	"\U00002652": ":aquarius:",
	"\U0001f50c": ":electric_plug:",
	"\U0001f356": ":meat_on_bone:",
	"\U0001f6b9": ":mens:",
	"\U0001f4bc": ":briefcase:",
	"\U0001f6a2": ":ship:",
	"\U00002693": ":anchor:",
	"\U00002611": ":ballot_box_with_check:",
	"\U0001f43b": ":bear:",
	"\U0001f37b": ":beers:",
	"\U0001f42a": ":dromedary_camel:",
	"\U0001f529": ":nut_and_bolt:",
	"\U0001f6a7": ":construction:",
	"\U000026f3": ":golf:",
	"\U0001f6bd": ":toilet:",
	"\U0001f4d8": ":blue_book:",
	"\U0001f4a5": ":boom:",
	"\U0001f333": ":deciduous_tree:",
	"\U0001f61a": ":kissing_closed_eyes:",
	"\U0001f63a": ":smiley_cat:",
	"\U000026fd": ":fuelpump:",
	"\U0001f48b": ":kiss:",
	"\U0001f559": ":clock10:",
	"\U0001f411": ":sheep:",
	"\U0001f4ae": ":white_flower:",
	"\U0001f417": ":boar:",
	"\U0001f4b1": ":currency_exchange:",
	// "\U0001f44a": ":facepunch:",
	"\U0001f3b4": ":flower_playing_cards:",
	"\U0001f64d": ":person_frowning:",
	// "\U0001f4a9": ":poop:",
	// "\U0001f606": ":satisfied:",
	"\U0001f3b1": ":8ball:",
	"\U0001f625": ":disappointed_relieved:",
	"\U0001f43c": ":panda_face:",
	"\U0001f3ab": ":ticket:",
	"\U0001f1fa\U0001f1f8": ":us:",
	"\U0001f312": ":waxing_crescent_moon:",
	"\U0001f409": ":dragon:",
	"\U0001f52b": ":gun:",
	"\U0001f5fb": ":mount_fuji:",
	"\U0001f31a": ":new_moon_with_face:",
	"\U0001f31f": ":star2:",
	"\U0001f62c": ":grimacing:",
	"\U0001f616": ":confounded:",
	"\U00003297": ":congratulations:",
	"\U0001f36e": ":custard:",
	"\U0001f626": ":frowning:",
	"\U0001f341": ":maple_leaf:",
	"\U0001f693": ":police_car:",
	"\U00002601": ":cloud:",
	"\U0001f456": ":jeans:",
	"\U0001f41f": ":fish:",
	"\U00003030": ":wavy_dash:",
	"\U0001f554": ":clock5:",
	"\U0001f385": ":santa:",
	"\U0001f5fe": ":japan:",
	"\U0001f696": ":oncoming_taxi:",
	"\U0001f433": ":whale:",
	"\U000025b6": ":arrow_forward:",
	"\U0001f618": ":kissing_heart:",
	"\U0001f684": ":bullettrain_side:",
	"\U0001f628": ":fearful:",
	"\U0001f4b0": ":moneybag:",
	// "\U0001f3c3": ":runner:",
	"\U0001f4eb": ":mailbox:",
	"\U0001f461": ":sandal:",
	"\U0001f4a4": ":zzz:",
	"\U0001f34e": ":apple:",
	"\U00002934": ":arrow_heading_up:",
	"\U0001f46a": ":family:",
	"\U00002796": ":heavy_minus_sign:",
	"\U0001f3b7": ":saxophone:",
	"\U0001f239": ":u5272:",
	"\U0001f532": ":black_square_button:",
	"\U0001f490": ":bouquet:",
	"\U0001f48c": ":love_letter:",
	"\U0001f687": ":metro:",
	"\U0001f539": ":small_blue_diamond:",
	"\U0001f4ad": ":thought_balloon:",
	"\U00002b06": ":arrow_up:",
	"\U0001f6b7": ":no_pedestrians:",
	"\U0001f60f": ":smirk:",
	"\U0001f499": ":blue_heart:",
	"\U0001f537": ":large_blue_diamond:",
	"\U0001f19a": ":vs:",
	"\U0000270c": ":v:",
	"\U0000267f": ":wheelchair:",
	"\U0001f48f": ":couplekiss:",
	"\U000026fa": ":tent:",
	"\U0001f49c": ":purple_heart:",
	"\U0000263a": ":relaxed:",
	"\U0001f251": ":accept:",
	"\U0001f49a": ":green_heart:",
	"\U0001f63e": ":pouting_cat:",
	"\U0001f68a": ":tram:",
	"\U0000203c": ":bangbang:",
	// "\U0001f4a5": ":collision:",
	"\U0001f3ea": ":convenience_store:",
	"\U0001f471": ":person_with_blond_hair:",
	"\U0001f1ec\U0001f1e7": ":uk:",
	"\U0001f351": ":peach:",
	"\U0001f62b": ":tired_face:",
	"\U0001f35e": ":bread:",
	"\U0001f4ea": ":mailbox_closed:",
	"\U0001f62e": ":open_mouth:",
	"\U0001f437": ":pig:",
	"\U0001f6ae": ":put_litter_in_its_place:",
	"\U0001f233": ":u7a7a:",
	"\U0001f4a1": ":bulb:",
	"\U0001f558": ":clock9:",
	"\U0001f4e9": ":envelope_with_arrow:",
	"\U00002653": ":pisces:",
	"\U0001f6c4": ":baggage_claim:",
	"\U0001f373": ":egg:",
	"\U0001f605": ":sweat_smile:",
	"\U000026f5": ":boat:",
	"\U0001f1eb\U0001f1f7": ":fr:",
	"\U00002797": ":heavy_division_sign:",
	"\U0001f4aa": ":muscle:",
	"\U0001f43e": ":paw_prints:",
	"\U00002b05": ":arrow_left:",
	"\U000026ab": ":black_circle:",
	"\U0001f619": ":kissing_smiling_eyes:",
	"\U00002b50": ":star:",
	"\U0001f682": ":steam_locomotive:",
	"\U0001f522": ":1234:",
	"\U0001f55c": ":clock130:",
	"\U0001f1f0\U0001f1f7": ":kr:",
	"\U0001f69d": ":monorail:",
	"\U0001f3eb": ":school:",
	"\U00000037\U000020e3": ":seven:",
	"\U0001f424": ":baby_chick:",
	"\U0001f309": ":bridge_at_night:",
	"\U00002668": ":hotsprings:",
	"\U0001f339": ":rose:",
	"\U0001f3e9": ":love_hotel:",
	"\U0001f478": ":princess:",
	"\U0001f35c": ":ramen:",
	"\U0001f4dc": ":scroll:",
	"\U0001f420": ":tropical_fish:",
	"\U0001f63b": ":heart_eyes_cat:",
	"\U0001f481": ":information_desk_person:",
	"\U0001f42d": ":mouse:",
	"\U0001f6ad": ":no_smoking:",
	"\U0001f3e3": ":post_office:",
	"\U0001f320": ":stars:",
	"\U000023ec": ":arrow_double_down:",
	"\U0001f513": ":unlock:",
	"\U000025c0": ":arrow_backward:",
	"\U0000270b": ":hand:",
	"\U0001f3e5": ":hospital:",
	"\U0001f30a": ":ocean:",
	"\U0001f6b5": ":mountain_bicyclist:",
	"\U0001f419": ":octopus:",
	"\U0001f198": ":sos:",
	"\U0001f635": ":dizzy_face:",
	"\U0001f445": ":tongue:",
	"\U0001f686": ":train2:",
	"\U0001f3c1": ":checkered_flag:",
	"\U0001f4d9": ":orange_book:",
	"\U0001f509": ":sound:",
	"\U0001f6a1": ":aerial_tramway:",
	"\U0001f514": ":bell:",
	"\U0001f432": ":dragon_face:",
	// "\U0001f42c": ":flipper:",
	"\U0001f646": ":ok_woman:",
	"\U0001f3ad": ":performing_arts:",
	"\U0001f4ef": ":postal_horn:",
	"\U0001f565": ":clock1030:",
	"\U00002709": ":mail:", // "\U00002709": ":email:",
	"\U0001f4d7": ":green_book:",
	"\U0001f446": ":point_up_2:",
	"\U0001f506": ":high_brightness:",
	"\U0001f3bd": ":running_shirt_with_sash:",
	"\U0001f516": ":bookmark:",
	"\U0001f62d": ":sob:",
	"\U00002198": ":arrow_lower_right:",
	"\U0001f448": ":point_left:",
	"\U0001f45b": ":purse:",
	"\U00002728": ":sparkles:",
	"\U000025fe": ":black_medium_small_square:",
	"\U0001f4b7": ":pound:",
	"\U0001f430": ":rabbit:",
	"\U0001f469": ":woman:",
	"\U0000274e": ":negative_squared_cross_mark:",
	// "\U0001f4d6": ":open_book:",
	"\U0001f608": ":smiling_imp:",
	"\U00002660": ":spades:",
	"\U000026be": ":baseball:",
	"\U000026f2": ":fountain:",
	"\U0001f602": ":joy:",
	"\U0001f484": ":lipstick:",
	"\U000026c5": ":partly_sunny:",
	"\U0001f40f": ":ram:",
	"\U0001f534": ":red_circle:",
	"\U0001f46e": ":cop:",
	"\U0001f34f": ":green_apple:",
	"\U000000ae": ":registered:",
	":+1:":                          "\U0001f44d",
	"\U0001f63f": ":crying_cat_face:",
	"\U0001f607": ":innocent:",
	"\U0001f4f4": ":mobile_phone_off:",
	"\U0001f51e": ":underage:",
	"\U0001f42c": ":dolphin:",
	"\U0001f465": ":busts_in_silhouette:",
	"\U00002614": ":umbrella:",
	"\U0001f47c": ":angel:",
	"\U0001f538": ":small_orange_diamond:",
	"\U0001f33b": ":sunflower:",
	"\U0001f517": ":link:",
	"\U0001f4d3": ":notebook:",
	"\U0001f68d": ":oncoming_bus:",
	"\U0001f4d1": ":bookmark_tabs:",
	"\U0001f4c6": ":calendar:",
	// "\U0001f3ee": ":izakaya_lantern:",
	// "\U0001f45e": ":mans_shoe:",
	"\U0001f4db": ":name_badge:",
	"\U0001f510": ":closed_lock_with_key:",
	"\U0000270a": ":fist:",
	"\U0001f194": ":id:",
	"\U0001f691": ":ambulance:",
	"\U0001f3b9": ":musical_keyboard:",
	"\U0001f380": ":ribbon:",
	"\U0001f331": ":seedling:",
	"\U0001f4fa": ":tv:",
	"\U0001f3c8": ":football:",
	"\U0001f485": ":nail_care:",
	"\U0001f4ba": ":seat:",
	"\U000023f0": ":alarm_clock:",
	"\U0001f4b8": ":money_with_wings:",
	"\U0001f60c": ":relieved:",
	"\U0001f45a": ":womans_clothes:",
	"\U0001f444": ":lips:",
	"\U00002663": ":clubs:",
	"\U0001f3e1": ":house_with_garden:",
	"\U0001f305": ":sunrise:",
	"\U0001f412": ":monkey:",
	"\U00000036\U000020e3": ":six:",
	"\U0001f603": ":smiley:",
	// "\U0001f43e": ":feet:",
	"\U0001f316": ":waning_gibbous_moon:",
	"\U0001f4b4": ":yen:",
	"\U0001f6bc": ":baby_symbol:",
	"\U0001f4f6": ":signal_strength:",
	"\U0001f466": ":boy:",
	"\U0001f68f": ":busstop:",
	"\U0001f4bb": ":computer:",
	"\U0001f303": ":night_with_stars:",
	"\U0001f475": ":older_woman:",
	"\U0001f17f": ":parking:",
	"\U0001f3ba": ":trumpet:",
	"\U0001f4af": ":100:",
	"\U0001f4a6": ":sweat_drops:",
	"\U0001f6be": ":wc:",
	"\U0001f171": ":b:",
	"\U0001f498": ":cupid:",
	"\U00000035\U000020e3": ":five:",
	"\U0000303d": ":part_alternation_mark:",
	"\U0001f3c2": ":snowboarder:",
	"\U000026a0": ":warning:",
	"\U00002b1c": ":white_large_square:",
	"\U000026a1": ":zap:",
	"\U0001f53d": ":arrow_down_small:",
	"\U0001f55f": ":clock430:",
	"\U0001f611": ":expressionless:",
	"\U0000260e": ":phone:",
	"\U0001f3a2": ":roller_coaster:",
	"\U0001f34b": ":lemon:",
	"\U00000031\U000020e3": ":one:",
	"\U0001f384": ":christmas_tree:",
	// "\U0001f4a9": ":hankey:",
	"\U0001f425": ":hatched_chick:",
	"\U0001f238": ":u7533:",
	"\U0001f535": ":large_blue_circle:",
	"\U0001f199": ":up:",
	"\U0001f377": ":wine_glass:",
	"\U0000274c": ":x:",
	"\U0001f443": ":nose:",
	"\U000023ea": ":rewind:",
	"\U0001f495": ":two_hearts:",
	// "\U00002709": ":envelope:",
	"\U0001f698": ":oncoming_automobile:",
	"\U000026ce": ":ophiuchus:",
	"\U0001f48d": ":ring:",
	"\U0001f379": ":tropical_drink:",
	"\U0001f422": ":turtle:",
	"\U0001f319": ":crescent_moon:",
	"\U0001f201": ":koko:",
	"\U0001f52c": ":microscope:",
	"\U0001f3c9": ":rugby_football:",
	"\U0001f6ac": ":smoking:",
	"\U0001f4a2": ":anger:",
	"\U00002648": ":aries:",
	"\U0001f306": ":city_sunset:",
	"\U0001f567": ":clock1230:",
	"\U0001f4ed": ":mailbox_with_no_mail:",
	"\U0001f3a5": ":movie_camera:",
	"\U0001f4df": ":pager:",
	"\U00000030\U000020e3": ":zero:",
	"\U0001f3e6": ":bank:",
	"\U00002734": ":eight_pointed_black_star:",
	"\U0001f52a": ":knife:",
	"\U0001f21a": ":u7121:",
	"\U0001f6c3": ":customs:",
	"\U0001f348": ":melon:",
	"\U0001f6a3": ":rowboat:",
	"\U0001f33d": ":corn:",
	"\U0001f346": ":eggplant:",
	"\U0001f49f": ":heart_decoration:",
	"\U0001f6a8": ":rotating_light:",
	"\U0001f4cd": ":round_pushpin:",
	"\U0001f408": ":cat2:",
	"\U0001f36b": ":chocolate_bar:",
	"\U0001f515": ":no_bell:",
	"\U0001f4fb": ":radio:",
	"\U0001f4a7": ":droplet:",
	"\U0001f354": ":hamburger:",
	"\U0001f692": ":fire_engine:",
	"\U00002764": ":heart:",
	"\U0001f6b0": ":potable_water:",
	"\U0001f4de": ":telephone_receiver:",
	"\U0001f4a8": ":dash:",
	"\U0001f310": ":globe_with_meridians:",
	"\U0001f482": ":guardsman:",
	"\U00002716": ":heavy_multiplication_x:",
	"\U0001f4c9": ":chart_with_downwards_trend:",
	"\U0001f47f": ":imp:",
	"\U0001f30f": ":earth_asia:",
	"\U0001f401": ":mouse2:",
	"\U0001f4d4": ":notebook_with_decorative_cover:",
	"\U0001f52d": ":telescope:",
	"\U0001f68e": ":trolleybus:",
	"\U0001f4c7": ":card_index:",
	"\U0001f4b6": ":euro:",
	"\U0001f4b5": ":dollar:",
	"\U0001f4e0": ":fax:",
	"\U0001f4ec": ":mailbox_with_mail:",
	"\U0001f64c": ":raised_hands:",
	"\U0001f61e": ":disappointed:",
	"\U0001f301": ":foggy:",
	"\U0001f64e": ":person_with_pouting_face:",
	"\U0001f5fd": ":statue_of_liberty:",
	"\U0001f38e": ":dolls:",
	"\U0001f688": ":light_rail:",
	"\U0001f4dd": ":pencil:",
	"\U0001f64a": ":speak_no_evil:",
	"\U0001f4f2": ":calling:",
	"\U0001f563": ":clock830:",
	"\U0001f404": ":cow2:",
	"\U0001f649": ":hear_no_evil:",
	"\U0001f640": ":scream_cat:",
	"\U0001f638": ":smile_cat:",
	"\U0001f69c": ":tractor:",
	"\U0001f55a": ":clock11:",
	"\U0001f369": ":doughnut:",
	"\U0001f528": ":hammer:",
	"\U000027bf": ":loop:",
	"\U0001f314": ":moon:",
	"\U0001f51c": ":soon:",
	"\U0001f3a6": ":cinema:",
	"\U0001f3ed": ":factory:",
	"\U0001f633": ":flushed:",
	"\U0001f507": ":mute:",
	"\U0001f610": ":neutral_face:",
	"\U0000264f": ":scorpius:",
	"\U0001f43a": ":wolf:",
	"\U0001f3ac": ":clapper:",
	"\U0001f639": ":joy_cat:",
	"\U0001f614": ":pensive:",
	"\U0001f634": ":sleeping:",
	"\U0001f4b3": ":credit_card:",
	"\U0000264c": ":leo:",
	"\U0001f472": ":man_with_gua_pi_mao:",
	"\U0001f450": ":open_hands:",
	"\U0001f375": ":tea:",
	"\U00002b07": ":arrow_down:",
	"\U00000039\U000020e3": ":nine:",
	"\U0001f44a": ":punch:",
	"\U0001f3b0": ":slot_machine:",
	"\U0001f44f": ":clap:",
	"\U00002139": ":information_source:",
	"\U0001f42f": ":tiger:",
	"\U0001f307": ":city_sunrise:",
	"\U0001f361": ":dango:",
	"\U0001f44e": ":thumbsdown:",
	"\U0001f22f": ":u6307:",
	"\U0001f35b": ":curry:",
	"\U0001f352": ":cherries:",
	"\U0001f555": ":clock6:",
	"\U0001f556": ":clock7:",
	"\U0001f474": ":older_man:",
	"\U0001f694": ":oncoming_police_car:",
	"\U0001f489": ":syringe:",
	"\U0001f4b2": ":heavy_dollar_sign:",
	"\U0001f4c2": ":open_file_folder:",
	"\U000021aa": ":arrow_right_hook:",
	"\U0001f69b": ":articulated_lorry:",
	"\U0001f46f": ":dancers:",
	"\U0001f63d": ":kissing_cat:",
	"\U0001f308": ":rainbow:",
	"\U0001f234": ":u5408:",
	"\U0001f462": ":boot:",
	"\U0001f3a0": ":carousel_horse:",
	"\U0001f364": ":fried_shrimp:",
	"\U0001f512": ":lock:",
	":non-potable_water:":              "\U0001f6b1",
	"\U00002b55": ":o:",
	"\U0001f623": ":persevere:",
	"\U0001f4a0": ":diamond_shape_with_a_dot_inside:",
	"\U0001f342": ":fallen_leaf:",
	"\U0001f486": ":massage:",
	"\U0001f30b": ":volcano:",
	"\U0001f48e": ":gem:",
	"\U0001f6bf": ":shower:",
	"\U0001f508": ":speaker:",
	"\U0001f31c": ":last_quarter_moon_with_face:",
	"\U0001f50d": ":mag:",
	"\U0001f627": ":anguished:",
	"\U0001f435": ":monkey_face:",
	"\U00002600": ":sunny:",
	"\U0001f34a": ":tangerine:",
	"\U0001f449": ":point_right:",
	"\U0001f683": ":railway_car:",
	"\U0001f624": ":triumph:",
	"\U00000032\U000020e3": ":two:",
	"\U0001f49d": ":gift_heart:",
	"\U0001f4d2": ":ledger:",
	"\U00002650": ":sagittarius:",
	"\U00002744": ":snowflake:",
	"\U0001f524": ":abc:",
	"\U0001f434": ":horse:",
	"\U0001f44c": ":ok_hand:",
	"\U0001f4f9": ":video_camera:",
	"\U0001f496": ":sparkling_heart:",
	"\U00002649": ":taurus:",
	"\U0001f438": ":frog:",
	"\U0001f439": ":hamster:",
	"\U0001f681": ":helicopter:",
	"\U0001f35f": ":fries:",
	"\U0001f344": ":mushroom:",
	"\U0001f427": ":penguin:",
	"\U0001f69a": ":truck:",
	"\U0001f4ca": ":bar_chart:",
	"\U0001f332": ":evergreen_tree:",
	"\U0001f647": ":bow:",
	"\U0001f55b": ":clock12:",
	"\U0001f340": ":four_leaf_clover:",
	"\U0001f4e5": ":inbox_tray:",
	"\U0001f63c": ":smirk_cat:",
	"\U0001f46c": ":two_men_holding_hands:",
	"\U0001f403": ":water_buffalo:",
	"\U0001f47d": ":alien:",
	"\U0001f3ae": ":video_game:",
	"\U0001f36c": ":candy:",
	"\U0001f4c4": ":page_facing_up:",
	"\U0001f349": ":watermelon:",
	"\U00002705": ":white_check_mark:",
	"\U0001f33c": ":blossom:",
	"\U0001f40a": ":crocodile:",
	"\U0001f636": ":no_mouth:",
	"\U0001f17e": ":o2:",
	"\U0001f455": ":shirt:",
	"\U0001f557": ":clock8:",
	"\U0001f440": ":eyes:",
	"\U0001f407": ":rabbit2:",
	"\U0001f38b": ":tanabata_tree:",
	"\U0001f527": ":wrench:",
	"\U0001f1ea\U0001f1f8": ":es:",
	"\U0001f3c6": ":trophy:",
	"\U0001f46d": ":two_women_holding_hands:",
	"\U0001f561": ":clock630:",
	"\U0001f34d": ":pineapple:",
	"\U0001f61b": ":stuck_out_tongue:",
	"\U0001f620": ":angry:",
	"\U0001f45f": ":athletic_shoe:",
	"\U0001f36a": ":cookie:",
	"\U0001f38f": ":flags:",
	"\U0001f3b2": ":game_die:",
	"\U0001f426": ":bird:",
	"\U0001f383": ":jack_o_lantern:",
	"\U0001f402": ":ox:",
	"\U0001f4ce": ":paperclip:",
	"\U0001f62a": ":sleepy:",
	"\U0001f632": ":astonished:",
	"\U0001f519": ":back:",
	"\U0001f4d5": ":closed_book:",
	"\U0001f423": ":hatching_chick:",
	"\U0001f503": ":arrows_clockwise:",
	"\U0001f697": ":car:",
	"\U0001f442": ":ear:",
	"\U0001f487": ":haircut:",
	"\U0001f366": ":icecream:",
	"\U0001f464": ":bust_in_silhouette:",
	"\U00002666": ":diamonds:",
	"\U0001f645": ":no_good:",
	"\U0001f355": ":pizza:",
	"\U0001f414": ":chicken:",
	"\U0001f453": ":eyeglasses:",
	"\U0001f648": ":see_no_evil:",
	"\U0001f30d": ":earth_africa:",
	"\U0001f386": ":fireworks:",
	"\U0001f4c3": ":page_with_curl:",
	"\U0001f359": ":rice_ball:",
	"\U0001f533": ":white_square_button:",
	"\U0001f370": ":cake:",
	// "\U0001f697": ":red_car:",
	"\U00002122": ":tm:",
	"\U0001f612": ":unamused:",
	"\U0001f365": ":fish_cake:",
	"\U0001f511": ":key:",
	"\U0001f6a4": ":speedboat:",
	"\U0001f302": ":closed_umbrella:",
	"\U0001f350": ":pear:",
	"\U0001f4e1": ":satellite:",
	"\U0001f631": ":scream:",
	"\U0001f313": ":first_quarter_moon:",
	"\U0001f1ef\U0001f1f5": ":jp:",
	"\U0001f502": ":repeat_one:",
	"\U0001f41a": ":shell:",
	"\U00002049": ":interrobang:",
	"\U0001f531": ":trident:",
	"\U0001f23a": ":u55b6:",
	"\U0001f3e7": ":atm:",
	"\U0001f6aa": ":door:",
	"\U0001f617": ":kissing:",
	"\U0001f52f": ":six_pointed_star:",
	"\U0001f44d": ":thumbsup:",
	"\U0001f237": ":u6708:",
	"\U0001f6af": ":do_not_litter:",
	"\U0001f40b": ":whale2:",
	"\U0001f392": ":school_satchel:",
	"\U0001f335": ":cactus:",
	"\U0001f4cb": ":clipboard:",
	"\U0001f4ab": ":dizzy:",
	// "\U0001f314": ":waxing_gibbous_moon:",
	"\U0001f4f7": ":camera:",
	"\U0001f520": ":capital_abcd:",
	"\U0001f343": ":leaves:",
	"\U0001f6c5": ":left_luggage:",
	"\U0001f38d": ":bamboo:",
	"\U0001f3b3": ":bowling:",
	"\U00000038\U000020e3": ":eight:",
	"\U0001f458": ":kimono:",
	"\U00002194": ":left_right_arrow:",
	"\U0001f61c": ":stuck_out_tongue_winking_eye:",
	"\U0001f3c4": ":surfer:",
	"\U0001f613": ":sweat:",
	"\U0001f3bb": ":violin:",
	"\U0001f4ee": ":postbox:",
	"\U0001f470": ":bride_with_veil:",
	"\U0000267b": ":recycle:",
	"\U0001f689": ":station:",
	"\U0001f4fc": ":vhs:",
	"\U0001f38c": ":crossed_flags:",
	// "\U0001f4dd": ":memo:",
	"\U000026d4": ":no_entry:",
	"\U000026aa": ":white_circle:",
	"\U00002199": ":arrow_lower_left:",
	"\U0001f330": ":chestnut:",
	"\U0001f52e": ":crystal_ball:",
	"\U0001f317": ":last_quarter_moon:",
	"\U0001f50a": ":loud_sound:",
	"\U0001f353": ":strawberry:",
	"\U0001f61f": ":worried:",
	"\U0001f3aa": ":circus_tent:",
	"\U0001f629": ":weary:",
	"\U0001f6c1": ":bathtub:",
	"\U0001f40d": ":snake:",
	"\U0001f601": ":grin:",
	"\U0001f523": ":symbols:",
	"\U00002708": ":airplane:",
	"\U0001f60d": ":heart_eyes:",
	// "\U000026f5": ":sailboat:",
	"\U0001f372": ":stew:",
	// "\U0001f455": ":tshirt:",
	"\U0001f400": ":rat:",
	"\U000025fc": ":black_medium_square:",
	"\U0001f564": ":clock930:",
	"\U0001f31d": ":full_moon_with_face:",
	"\U0001f47a": ":japanese_goblin:",
	"\U0001f6bb": ":restroom:",
	"\U0001f6a6": ":vertical_traffic_light:",
	"\U0001f3c0": ":basketball:",
	"\U0001f338": ":cherry_blossom:",
	"\U0001f505": ":low_brightness:",
	"\U0001f48a": ":pill:",
}

// Emoji returns the unicode value for the given emoji. If the
// specified emoji does not exist, Emoji() returns the empty string.
func Noemoji(emoji string) string {
	val, ok := emojiMap[emoji]
	if !ok {
		return ""
	}
	return val
}

// Emojitize takes in a string with emojis specified in it, and returns
// a string with every emoji place holder replaced with it's unicode value
// (unless it could not be found, in which case it is let alone).
// http://maiyang.github.io/golang/%E5%AD%97%E7%AC%A6%E4%B8%B2/emoji/%E8%A1%A8%E6%83%85/2015/06/16/golang-character-length/
func Noemojitize(content string) string {
  new_content := ""
  for _, value := range content {
    _, size := utf8.DecodeRuneInString(string(value))
    if size <= 3 {
      new_content += string(value)
    } else {
      new_content += Noemoji(string(value))
    }
  }
  return new_content
}
