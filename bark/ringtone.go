package bark

// RingtoneName represents name of built-in ringtones of Bark app.
type RingtoneName string

// Name of available built-in ringtones of Bark app
const (
	RingtoneDefault            RingtoneName = ""
	RingtoneAlarm              RingtoneName = "alarm"
	RingtoneAnticipate         RingtoneName = "anticipate"
	RingtoneBell               RingtoneName = "bell"
	RingtoneBirdsong           RingtoneName = "birdsong"
	RingtoneBloom              RingtoneName = "bloom"
	RingtoneCalypso            RingtoneName = "calypso"
	RingtoneChime              RingtoneName = "chime"
	RingtoneChoo               RingtoneName = "choo"
	RingtoneDescent            RingtoneName = "descent"
	RingtoneElectronic         RingtoneName = "electronic"
	RingtoneFanfare            RingtoneName = "fanfare"
	RingtoneGlass              RingtoneName = "glass"
	RingtoneGoToSleep          RingtoneName = "gotosleep"
	RingtoneHealthNotification RingtoneName = "healthnotification"
	RingtoneHorn               RingtoneName = "horn"
	RingtoneLadder             RingtoneName = "ladder"
	RingtoneMailSent           RingtoneName = "mailsent"
	RingtoneMinuet             RingtoneName = "minuet"
	RingtoneMultiwayInvitation RingtoneName = "multiwayinvitation"
	RingtoneNewMail            RingtoneName = "newmail"
	RingtoneNewsFlash          RingtoneName = "newsflash"
	RingtoneNoir               RingtoneName = "noir"
	RingtonePaymentSuccess     RingtoneName = "paymentsuccess"
	RingtoneShake              RingtoneName = "shake"
	RingtoneSherwoodForest     RingtoneName = "sherwoodforest"
	RingtoneSilence            RingtoneName = "silence"
	RingtoneSpell              RingtoneName = "spell"
	RingtoneSuspense           RingtoneName = "suspense"
	RingtoneTelegraph          RingtoneName = "telegraph"
	RingtoneTiptoes            RingtoneName = "tiptoes"
	RingtoneTypewriters        RingtoneName = "typewriters"
	RingtoneUpdate             RingtoneName = "update"
)

// AllRingtones contains all available built-in ringtones.
var AllRingtones = []RingtoneName{
	RingtoneSilence,
	RingtoneAlarm,
	RingtoneAnticipate,
	RingtoneBell,
	RingtoneBirdsong,
	RingtoneBloom,
	RingtoneCalypso,
	RingtoneChime,
	RingtoneChoo,
	RingtoneDescent,
	RingtoneElectronic,
	RingtoneFanfare,
	RingtoneGlass,
	RingtoneGoToSleep,
	RingtoneHealthNotification,
	RingtoneHorn,
	RingtoneLadder,
	RingtoneMailSent,
	RingtoneMinuet,
	RingtoneMultiwayInvitation,
	RingtoneNewMail,
	RingtoneNewsFlash,
	RingtoneNoir,
	RingtonePaymentSuccess,
	RingtoneShake,
	RingtoneSherwoodForest,
	RingtoneSpell,
	RingtoneSuspense,
	RingtoneTelegraph,
	RingtoneTiptoes,
	RingtoneTypewriters,
	RingtoneUpdate,
}
