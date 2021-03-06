// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api

import (
	"github.com/mdl-react-golang/model"
	"github.com/mdl-react-golang/store"
	"github.com/mdl-react-golang/utils"
)

var Client *model.Client

func Setup() {
	if Srv == nil {
		utils.LoadConfig("config.json")
		utils.InitTranslations()
		utils.Cfg.TeamSettings.MaxUsersPerTeam = 50
		NewServer()
		StartServer()
		InitApi()
		Client = model.NewClient("http://localhost" + utils.Cfg.ServiceSettings.ListenAddress)

		Srv.Store.MarkSystemRanUnitTests()
	}
}

func SetupBenchmark() (*model.Team, *model.User, *model.Channel) {
	Setup()

	team := &model.Team{DisplayName: "Benchmark Team", Name: "z-z-" + model.NewId() + "a", Email: "benchmark@nowhere.com", Type: model.TEAM_OPEN}
	team = Client.Must(Client.CreateTeam(team)).Data.(*model.Team)
	user := &model.User{TeamId: team.Id, Email: model.NewId() + "success+test@simulator.amazonses.com", Nickname: "Mr. Benchmarker", Password: "pwd"}
	user = Client.Must(Client.CreateUser(user, "")).Data.(*model.User)
	store.Must(Srv.Store.User().VerifyEmail(user.Id))
	Client.LoginByEmail(team.Name, user.Email, "pwd")
	channel := &model.Channel{DisplayName: "Benchmark Channel", Name: "a" + model.NewId() + "a", Type: model.CHANNEL_OPEN, TeamId: team.Id}
	channel = Client.Must(Client.CreateChannel(channel)).Data.(*model.Channel)

	return team, user, channel
}

func TearDown() {
	if Srv != nil {
		StopServer()
	}
}
