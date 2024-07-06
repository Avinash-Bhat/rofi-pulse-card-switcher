package main

import (
	"encoding/json"
	"fmt"
	"github.com/Avinash-Bhat/rofi-pulse-card-switcher/rofi"
	pulse "github.com/mafik/pulseaudio"
	"log"
	"slices"
	"strconv"
	"strings"
)

var (
	cardPropsNames = []string{
		"device.api",
		"device.bus",
		"device.bus_path",
		"device.description",
		"device.enum.api",
		"device.form_factor",
		"device.icon_name",
		"device.name",
		"device.nick",
		"device.plugged.usec",
		"device.product.id",
		"device.product.name",
		"device.string",
		"device.subsystem",
		"device.vendor.id",
		"device.vendor.name",
	}
)

type response struct {
	mode      rofi.Mode
	responses []rofi.Row
}

func (r response) String() string {
	sb := strings.Builder{}
	for _, row := range r.responses {
		sb.WriteString(row.String())
		sb.WriteByte('\n')
	}
	sb.WriteString(r.mode.String())
	return sb.String()
}
func main() {

	client, err := pulse.NewClient()
	if err != nil {
		log.Fatal("cannot connect to pulseaudio")
	}
	info, err := client.ServerInfo()
	if err != nil {
		log.Fatal("cannot connect to pulseaudio")
	}
	env := rofi.GetEnv()
	var resp response
	switch env.CurrentState {
	case 1:
		cardProps := &map[string]string{}
		if err := json.Unmarshal([]byte(env.Info), cardProps); err != nil {
			log.Fatalf("cannot parse the info: %s", env.Info)
		}
		log.Printf("[INFO] entry selected: %v, props: %s\n", env.Arg, cardProps)
		outputs, activeIndex, err := client.Outputs()
		if err != nil {
			log.Fatal("cannot get the outputs of pulseaudio")
		}
		for i, output := range outputs {
			if !output.Available {
				continue
			}
			log.Printf("output: %s:%s, active: %v, selected: %v\n", output.CardName, output.CardID, i == activeIndex, output.CardName == env.Arg)
			if i != activeIndex && output.CardName == env.Arg {
				if err := output.Activate(); err != nil {
					log.Fatalf("Cannot activate card: %s\n", env.Arg)
				}
				break
			}
		}
	default:
		resp = listCards(info, err, client)
	}

	fmt.Print(resp)
}

func listCards(info *pulse.Server, err error, client *pulse.Client) response {
	defaultSink := info.DefaultSink
	defaultSource := info.DefaultSource

	cards, err := client.Cards()
	if err != nil {
		log.Fatal("unable to list cards")
	}
	mode := rofi.Mode{
		NoCustom: true,
	}
	rows := make([]rofi.Row, len(cards))
	for i := 0; i < len(cards); i++ {
		card := cards[i]
		name := card.Name
		var busId string
		if card.Driver == "alsa" {
			name = card.PropList["device.nick"]
			busId = card.PropList["device.bus-id"]
		}
		isDefaultSink := containsAny(defaultSink, name, busId)
		isDefaultSource := containsAny(defaultSource, name, busId)
		cardInfo, err := json.Marshal(extractInfo(card))
		if err != nil {
			log.Fatalf("unable to convert cardInfo: %v", extractInfo(card))
		}
		row := rofi.Row{
			Value:   name,
			Active:  isDefaultSink || isDefaultSource,
			Meta:    card.Name,
			Urgent:  isDefaultSink || isDefaultSource,
			Display: name,
		}
		if cardInfo != nil {
			row.Info = string(cardInfo)
		}
		if isDefaultSource || isDefaultSink {
			mode.Active = strconv.Itoa(i)
			mode.NewSelection = mode.Active
			mode.KeepSelection = true
			log.Println("[INFO]", "default:", name)
		}
		rows[i] = row
	}
	return response{
		mode,
		rows,
	}
}

func extractInfo(card pulse.Card) map[string]string {
	keys := map[string]string{}
	for k, v := range card.PropList {
		if slices.Contains(cardPropsNames, k) {
			keys[k] = v
		}
	}
	return keys
}

func containsAny(s string, substrings ...string) bool {
	for _, substring := range substrings {
		if substring == "" {
			continue
		}
		if strings.Contains(s, substring) {
			return true
		}
	}
	return false
}
