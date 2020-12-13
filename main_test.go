package main

import "testing"

func TestPigLatin(t *testing.T) {
	t.Run("word that begin with vowel sounds", func(t *testing.T) {
		tt := []struct {
			input string
			want  string
		}{
			{input: "eat", want: "eatyay"},
			{input: "omelet", want: "omeletyay"},
			{input: "are", want: "areyay"},
			{input: "egg", want: "eggyay"},
			{input: "explain", want: "explainyay"},
			{input: "always", want: "alwaysyay"},
			{input: "ends", want: "endsyay"},
			{input: "I", want: "Iyay"},
		}
		for _, test := range tt {
			got := TranslateWord(test.input)
			if got != test.want {
				t.Errorf("for %s got %s want %s", test.input, got, test.want)
			}
		}
	})
	t.Run("word that begin with consonant sounds", func(t *testing.T) {
		tt := []struct {
			input string
			want  string
		}{
			{input: "pig", want: "igpay"},
			{input: "latin", want: "atinlay"},
			{input: "banana", want: "ananabay"},
			{input: "will", want: "illway"},
			{input: "butler", want: "utlerbay"},
			{input: "happy", want: "appyhay"},
			{input: "duck", want: "uckday"},
			{input: "me", want: "emay"},
			{input: "bagel", want: "agelbay"},
		}
		for _, test := range tt {
			got := TranslateWord(test.input)
			if got != test.want {
				t.Errorf("for %s got %s want %s", test.input, got, test.want)
			}
		}
	})
	t.Run("word that begin with consonant clusters", func(t *testing.T) {
		tt := []struct {
			input string
			want  string
		}{
			{input: "smile", want: "ilesmay"},
			{input: "string", want: "ingstray"},
			{input: "stupid", want: "upidstay"},
			{input: "glove", want: "oveglay"},
			{input: "trash", want: "ashtray"},
			{input: "floor", want: "oorflay"},
			{input: "store", want: "orestay"},
		}
		for _, test := range tt {
			got := TranslateWord(test.input)
			if got != test.want {
				t.Errorf("for %s got %s want %s", test.input, got, test.want)
			}
		}
	})

}

func TestTranslatePhrase(t *testing.T) {
	tt := []struct {
		input string
		want  string
	}{
		{input: "How are you?", want: "Owhay arehay ouyay?"},
		{input: "Long time no see", want: "Onglay imetay onay eesay"},
		{input: "What's your name?", want: "   Atwhay isway ouryay amenay?"},
		{input: "My name is ...", want: " Ymay amenay isway ..."},
		{input: "Where are you from?", want: " Erewhay areway ouyay romfay?"},
		{input: "I'm from ...", want: "  Iway amway romfay ..."},
		{input: "Pleased to meet you", want: " Leasedpay otay eetmay ouyay"},
		{input: "Good morning ", want: " Oodgay orningmay"},
		{input: "Good afternoon ", want: "  Oodgay afternoonway"},
		{input: "Good evening", want: "  Oodgay eveningway"},
		{input: "Good night ", want: " Oodgay ightgay"},
		{input: "Good luck!", want: "  Oodgay ucklay"},
		{input: "Have a nice day", want: " Avehay away icenay ayday"},
		{input: "I understand ", want: "Iway understandway"},
		{input: "I don't understand  Iway on'tday understandway"},
		{input: "I don't know", want: "Iway on'tday nowkay"},
		{input: "Please speak more slowly  ", want: "Leasepay peaksay oremay lowlysay"},
		{input: "Please say that again", want: "Leasepay aysay atthay againway"},
		{input: "Please write it down", want: "Leasepay riteway itway ownday"},
		{input: "Do you speak English?  ", want: " Oday ouyay peaksay Englishway?"},
		{input: "Do you speak Pig Latin?>  ", want: "  Oday ouyay peaksay Igpay Atinlay?"},
		{input: "Yes, a little", want: "Esyay, away ittlelay"},
		{input: "Speak to me in Pig Latin", want: "Peaksay otay emay inway Igpay Atinlay"},
		{input: "Excuse me", want: "Excuseway emay"},
		{input: "How much is this?", want: "Owhay uchmay?"},
		{input: "Sorry  ", want: "Orrysay"},
		{input: "Please", want: "Leasepay"},
		{input: "I miss you ", want: "Iway issmay ouyou"},
		{input: "I love you ", want: "Iway ovelay ouyou"},
		{input: "Get well soon  ", want: "Etgay ellway oonsay"},
		{input: "Go away! ", want: "Ogay awayway"},
		{input: "Leave me alone!", want: "Eavelay emay aloneway"},
		{input: "Help!", want: "Elphay!"},
		{input: "Fire!", want: "Irefay!"},
		{input: "Stop!", want: "Topsay!"},
		{input: "Call the police! ", want: "Allcay ethay olicepay!"},
		{input: "Christmas greetings", want: " Errymay Ristmaschay"},
		{input: "New Year greetings ", want: " Appyhay Ewnay Eearyay"},
		{input: "Easter greetings   ", want: " Appyhay Easterway"},
		{input: "Birthday greetings ", want: " Appyhay Irthdaybay"},
	}
	for _, test := range tt {
		got := TranslatePhrase(test.input)
		if got != test.want {
			t.Errorf("for %s got %s want %s", test.input, got, test.want)
		}
	}
}
