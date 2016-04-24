// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package store

import (
	"time"

	l4g "github.com/alecthomas/log4go"
	"github.com/mdl-react-golang/model"
)

type StoreResult struct {
	Data interface{}
	Err  *model.AppError
}

type StoreChannel chan StoreResult

func Must(sc StoreChannel) interface{} {
	r := <-sc
	if r.Err != nil {
		l4g.Close()
		time.Sleep(time.Second)
		panic(r.Err)
	}

	return r.Data
}

type Store interface {
	User() UserStore
	Session() SessionStore
	OAuth() OAuthStore
	System() SystemStore
	Command() CommandStore
	Preference() PreferenceStore
	MarkSystemRanUnitTests()
	Close()
}

type UserStore interface {
	Save(user *model.User) StoreChannel
	Update(user *model.User, allowRoleUpdate bool) StoreChannel
	UpdateLastPictureUpdate(userId string) StoreChannel
	UpdateLastPingAt(userId string, time int64) StoreChannel
	UpdateLastActivityAt(userId string, time int64) StoreChannel
	UpdateUserAndSessionActivity(userId string, sessionId string, time int64) StoreChannel
	UpdatePassword(userId, newPassword string) StoreChannel
	UpdateAuthData(userId, service, authData, email string) StoreChannel
	UpdateMfaSecret(userId, secret string) StoreChannel
	UpdateMfaActive(userId string, active bool) StoreChannel
	Get(id string) StoreChannel
	GetProfiles(teamId string) StoreChannel
	GetByEmail(teamId string, email string) StoreChannel
	GetByAuth(teamId string, authData string, authService string) StoreChannel
	GetByUsername(teamId string, username string) StoreChannel
	VerifyEmail(userId string) StoreChannel
	GetEtagForProfiles(teamId string) StoreChannel
	UpdateFailedPasswordAttempts(userId string, attempts int) StoreChannel
	GetForExport(teamId string) StoreChannel
	GetTotalUsersCount() StoreChannel
	GetTotalActiveUsersCount() StoreChannel
	GetSystemAdminProfiles() StoreChannel
	PermanentDelete(userId string) StoreChannel
	AnalyticsUniqueUserCount(teamId string) StoreChannel

	GetUnreadCount(userId string) StoreChannel
}

type SessionStore interface {
	Save(session *model.Session) StoreChannel
	Get(sessionIdOrToken string) StoreChannel
	GetSessions(userId string) StoreChannel
	Remove(sessionIdOrToken string) StoreChannel
	RemoveAllSessionsForTeam(teamId string) StoreChannel
	PermanentDeleteSessionsByUser(teamId string) StoreChannel
	UpdateLastActivityAt(sessionId string, time int64) StoreChannel
	UpdateRoles(userId string, roles string) StoreChannel
	UpdateDeviceId(id string, deviceId string) StoreChannel
	AnalyticsSessionCount(teamId string) StoreChannel
}

type OAuthStore interface {
	SaveApp(app *model.OAuthApp) StoreChannel
	UpdateApp(app *model.OAuthApp) StoreChannel
	GetApp(id string) StoreChannel
	GetAppByUser(userId string) StoreChannel
	SaveAuthData(authData *model.AuthData) StoreChannel
	GetAuthData(code string) StoreChannel
	RemoveAuthData(code string) StoreChannel
	PermanentDeleteAuthDataByUser(userId string) StoreChannel
	SaveAccessData(accessData *model.AccessData) StoreChannel
	GetAccessData(token string) StoreChannel
	GetAccessDataByAuthCode(authCode string) StoreChannel
	RemoveAccessData(token string) StoreChannel
}

type SystemStore interface {
	Save(system *model.System) StoreChannel
	SaveOrUpdate(system *model.System) StoreChannel
	Update(system *model.System) StoreChannel
	Get() StoreChannel
	GetByName(name string) StoreChannel
}

type CommandStore interface {
	Save(webhook *model.Command) StoreChannel
	Get(id string) StoreChannel
	GetByTeam(teamId string) StoreChannel
	Delete(commandId string, time int64) StoreChannel
	PermanentDeleteByUser(userId string) StoreChannel
	Update(hook *model.Command) StoreChannel
	AnalyticsCommandCount(teamId string) StoreChannel
}

type PreferenceStore interface {
	Save(preferences *model.Preferences) StoreChannel
	Get(userId string, category string, name string) StoreChannel
	GetCategory(userId string, category string) StoreChannel
	GetAll(userId string) StoreChannel
	PermanentDeleteByUser(userId string) StoreChannel
	IsFeatureEnabled(feature, userId string) StoreChannel
}
