package piglatin_test

import (
	"testing"

	"github.com/igkostyuk/piglatin"
)

//nolint: funlen
func TestWords(t *testing.T) {
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
			got := piglatin.Translate(test.input)
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
			got := piglatin.Translate(test.input)
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
			got := piglatin.Translate(test.input)
			if got != test.want {
				t.Errorf("for %s got %s want %s", test.input, got, test.want)
			}
		}
	})
}

func TestPhrases(t *testing.T) {
	tt := []struct {
		input string
		want  string
	}{
		{input: "How are you?", want: "Owhay areyay ouyay?"},
		{input: "Long time no see", want: "Onglay imetay onay eesay"},
		{input: "Where are you from?", want: "Erewhay areyay ouyay omfray?"},
		{input: "I am from ...", want: "Iyay amyay omfray ..."},
		{input: "Good morning", want: "Oodgay orningmay"},
		{input: "Good afternoon", want: "Oodgay afternoonyay"},
		{input: "Good evening", want: "Oodgay eveningyay"},
		{input: "Good night", want: "Oodgay ightnay"},
		{input: "Good luck!", want: "Oodgay ucklay!"},
		{input: "Have a nice day", want: "Avehay ayay icenay ayday"},
		{input: "I understand", want: "Iyay understandyay"},
		{input: "Please speak more slowly", want: "Easeplay eakspay oremay owlyslay"},
		{input: "Please say that again", want: "Easeplay aysay atthay againyay"},
		{input: "Please write it down", want: "Easeplay itewray ityay ownday"},
		{input: "Do you speak English?", want: "Oday ouyay eakspay Englishyay?"},
		{input: "Do you speak Pig Latin?", want: "Oday ouyay eakspay Igpay Atinlay?"},
		{input: "Yes, a little", want: "Esyay, ayay ittlelay"},
		{input: "Speak to me in Pig Latin", want: "Eakspay otay emay inyay Igpay Atinlay"},
		{input: "Excuse me", want: "Excuseyay emay"},
		{input: "Sorry", want: "Orrysay"},
		{input: "I miss you", want: "Iyay issmay ouyay"},
		{input: "Get well soon", want: "Etgay ellway oonsay"},
		{input: "Go away! ", want: "Ogay awayyay!"},
		{input: "Leave me alone!", want: "Eavelay emay aloneyay!"},
		{input: "Help!", want: "Elphay!"},
		{input: "Fire!", want: "Irefay!"},
		{input: "Call the police!", want: "Allcay ethay olicepay!"},
		{input: "Merry Christmas", want: "Errymay Istmaschray"},
		{input: "Happy New Year", want: "Appyhay Ewnay Earyay"},
		{input: "Happy Birthday", want: "Appyhay Irthdaybay"},
	}
	for _, test := range tt {
		got := piglatin.Translate(test.input)
		if got != test.want {
			t.Errorf("for %s got %s want %s", test.input, got, test.want)
		}
	}
}
