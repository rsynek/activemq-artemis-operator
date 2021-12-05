// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisSecurity) DeepCopyInto(out *ActiveMQArtemisSecurity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisSecurity.
func (in *ActiveMQArtemisSecurity) DeepCopy() *ActiveMQArtemisSecurity {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemisSecurity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisSecurityList) DeepCopyInto(out *ActiveMQArtemisSecurityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveMQArtemisSecurity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisSecurityList.
func (in *ActiveMQArtemisSecurityList) DeepCopy() *ActiveMQArtemisSecurityList {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisSecurityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemisSecurityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisSecuritySpec) DeepCopyInto(out *ActiveMQArtemisSecuritySpec) {
	*out = *in
	in.LoginModules.DeepCopyInto(&out.LoginModules)
	in.SecurityDomains.DeepCopyInto(&out.SecurityDomains)
	in.SecuritySettings.DeepCopyInto(&out.SecuritySettings)
	if in.ApplyToCrNames != nil {
		in, out := &in.ApplyToCrNames, &out.ApplyToCrNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisSecuritySpec.
func (in *ActiveMQArtemisSecuritySpec) DeepCopy() *ActiveMQArtemisSecuritySpec {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisSecuritySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisSecurityStatus) DeepCopyInto(out *ActiveMQArtemisSecurityStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisSecurityStatus.
func (in *ActiveMQArtemisSecurityStatus) DeepCopy() *ActiveMQArtemisSecurityStatus {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisSecurityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowedListEntryType) DeepCopyInto(out *AllowedListEntryType) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowedListEntryType.
func (in *AllowedListEntryType) DeepCopy() *AllowedListEntryType {
	if in == nil {
		return nil
	}
	out := new(AllowedListEntryType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorisationConfigType) DeepCopyInto(out *AuthorisationConfigType) {
	*out = *in
	if in.AllowedList != nil {
		in, out := &in.AllowedList, &out.AllowedList
		*out = make([]AllowedListEntryType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultAccess != nil {
		in, out := &in.DefaultAccess, &out.DefaultAccess
		*out = make([]DefaultAccessType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RoleAccess != nil {
		in, out := &in.RoleAccess, &out.RoleAccess
		*out = make([]RoleAccessType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorisationConfigType.
func (in *AuthorisationConfigType) DeepCopy() *AuthorisationConfigType {
	if in == nil {
		return nil
	}
	out := new(AuthorisationConfigType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerDomainType) DeepCopyInto(out *BrokerDomainType) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.LoginModules != nil {
		in, out := &in.LoginModules, &out.LoginModules
		*out = make([]LoginModuleReferenceType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerDomainType.
func (in *BrokerDomainType) DeepCopy() *BrokerDomainType {
	if in == nil {
		return nil
	}
	out := new(BrokerDomainType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BrokerSecuritySettingType) DeepCopyInto(out *BrokerSecuritySettingType) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]PermissionType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BrokerSecuritySettingType.
func (in *BrokerSecuritySettingType) DeepCopy() *BrokerSecuritySettingType {
	if in == nil {
		return nil
	}
	out := new(BrokerSecuritySettingType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectorConfigType) DeepCopyInto(out *ConnectorConfigType) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.RmiRegistryPort != nil {
		in, out := &in.RmiRegistryPort, &out.RmiRegistryPort
		*out = new(int32)
		**out = **in
	}
	if in.JmxRealm != nil {
		in, out := &in.JmxRealm, &out.JmxRealm
		*out = new(string)
		**out = **in
	}
	if in.ObjectName != nil {
		in, out := &in.ObjectName, &out.ObjectName
		*out = new(string)
		**out = **in
	}
	if in.AuthenticatorType != nil {
		in, out := &in.AuthenticatorType, &out.AuthenticatorType
		*out = new(string)
		**out = **in
	}
	if in.Secured != nil {
		in, out := &in.Secured, &out.Secured
		*out = new(bool)
		**out = **in
	}
	if in.KeyStoreProvider != nil {
		in, out := &in.KeyStoreProvider, &out.KeyStoreProvider
		*out = new(string)
		**out = **in
	}
	if in.KeyStorePath != nil {
		in, out := &in.KeyStorePath, &out.KeyStorePath
		*out = new(string)
		**out = **in
	}
	if in.KeyStorePassword != nil {
		in, out := &in.KeyStorePassword, &out.KeyStorePassword
		*out = new(string)
		**out = **in
	}
	if in.TrustStoreProvider != nil {
		in, out := &in.TrustStoreProvider, &out.TrustStoreProvider
		*out = new(string)
		**out = **in
	}
	if in.TrustStorePath != nil {
		in, out := &in.TrustStorePath, &out.TrustStorePath
		*out = new(string)
		**out = **in
	}
	if in.TrustStorePassword != nil {
		in, out := &in.TrustStorePassword, &out.TrustStorePassword
		*out = new(string)
		**out = **in
	}
	if in.PasswordCodec != nil {
		in, out := &in.PasswordCodec, &out.PasswordCodec
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectorConfigType.
func (in *ConnectorConfigType) DeepCopy() *ConnectorConfigType {
	if in == nil {
		return nil
	}
	out := new(ConnectorConfigType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultAccessType) DeepCopyInto(out *DefaultAccessType) {
	*out = *in
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultAccessType.
func (in *DefaultAccessType) DeepCopy() *DefaultAccessType {
	if in == nil {
		return nil
	}
	out := new(DefaultAccessType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuestLoginModuleType) DeepCopyInto(out *GuestLoginModuleType) {
	*out = *in
	if in.GuestUser != nil {
		in, out := &in.GuestUser, &out.GuestUser
		*out = new(string)
		**out = **in
	}
	if in.GuestRole != nil {
		in, out := &in.GuestRole, &out.GuestRole
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuestLoginModuleType.
func (in *GuestLoginModuleType) DeepCopy() *GuestLoginModuleType {
	if in == nil {
		return nil
	}
	out := new(GuestLoginModuleType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyValueType) DeepCopyInto(out *KeyValueType) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyValueType.
func (in *KeyValueType) DeepCopy() *KeyValueType {
	if in == nil {
		return nil
	}
	out := new(KeyValueType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakLoginModuleType) DeepCopyInto(out *KeycloakLoginModuleType) {
	*out = *in
	if in.ModuleType != nil {
		in, out := &in.ModuleType, &out.ModuleType
		*out = new(string)
		**out = **in
	}
	in.Configuration.DeepCopyInto(&out.Configuration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakLoginModuleType.
func (in *KeycloakLoginModuleType) DeepCopy() *KeycloakLoginModuleType {
	if in == nil {
		return nil
	}
	out := new(KeycloakLoginModuleType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycloakModuleConfigurationType) DeepCopyInto(out *KeycloakModuleConfigurationType) {
	*out = *in
	if in.Realm != nil {
		in, out := &in.Realm, &out.Realm
		*out = new(string)
		**out = **in
	}
	if in.RealmPublicKey != nil {
		in, out := &in.RealmPublicKey, &out.RealmPublicKey
		*out = new(string)
		**out = **in
	}
	if in.AuthServerUrl != nil {
		in, out := &in.AuthServerUrl, &out.AuthServerUrl
		*out = new(string)
		**out = **in
	}
	if in.SslRequired != nil {
		in, out := &in.SslRequired, &out.SslRequired
		*out = new(string)
		**out = **in
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(string)
		**out = **in
	}
	if in.PublicClient != nil {
		in, out := &in.PublicClient, &out.PublicClient
		*out = new(bool)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]KeyValueType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UseResourceRoleMappings != nil {
		in, out := &in.UseResourceRoleMappings, &out.UseResourceRoleMappings
		*out = new(bool)
		**out = **in
	}
	if in.EnableCors != nil {
		in, out := &in.EnableCors, &out.EnableCors
		*out = new(bool)
		**out = **in
	}
	if in.CorsMaxAge != nil {
		in, out := &in.CorsMaxAge, &out.CorsMaxAge
		*out = new(int64)
		**out = **in
	}
	if in.CorsAllowedMethods != nil {
		in, out := &in.CorsAllowedMethods, &out.CorsAllowedMethods
		*out = new(string)
		**out = **in
	}
	if in.CorsAllowedHeaders != nil {
		in, out := &in.CorsAllowedHeaders, &out.CorsAllowedHeaders
		*out = new(string)
		**out = **in
	}
	if in.CorsExposedHeaders != nil {
		in, out := &in.CorsExposedHeaders, &out.CorsExposedHeaders
		*out = new(string)
		**out = **in
	}
	if in.ExposeToken != nil {
		in, out := &in.ExposeToken, &out.ExposeToken
		*out = new(bool)
		**out = **in
	}
	if in.BearerOnly != nil {
		in, out := &in.BearerOnly, &out.BearerOnly
		*out = new(bool)
		**out = **in
	}
	if in.AutoDetectBearerOnly != nil {
		in, out := &in.AutoDetectBearerOnly, &out.AutoDetectBearerOnly
		*out = new(bool)
		**out = **in
	}
	if in.ConnectionPoolSize != nil {
		in, out := &in.ConnectionPoolSize, &out.ConnectionPoolSize
		*out = new(int64)
		**out = **in
	}
	if in.AllowAnyHostName != nil {
		in, out := &in.AllowAnyHostName, &out.AllowAnyHostName
		*out = new(bool)
		**out = **in
	}
	if in.DisableTrustManager != nil {
		in, out := &in.DisableTrustManager, &out.DisableTrustManager
		*out = new(bool)
		**out = **in
	}
	if in.TrustStore != nil {
		in, out := &in.TrustStore, &out.TrustStore
		*out = new(string)
		**out = **in
	}
	if in.TrustStorePassword != nil {
		in, out := &in.TrustStorePassword, &out.TrustStorePassword
		*out = new(string)
		**out = **in
	}
	if in.ClientKeyStore != nil {
		in, out := &in.ClientKeyStore, &out.ClientKeyStore
		*out = new(string)
		**out = **in
	}
	if in.ClientKeyStorePassword != nil {
		in, out := &in.ClientKeyStorePassword, &out.ClientKeyStorePassword
		*out = new(string)
		**out = **in
	}
	if in.ClientKeyPassword != nil {
		in, out := &in.ClientKeyPassword, &out.ClientKeyPassword
		*out = new(string)
		**out = **in
	}
	if in.AlwaysRefreshToken != nil {
		in, out := &in.AlwaysRefreshToken, &out.AlwaysRefreshToken
		*out = new(bool)
		**out = **in
	}
	if in.RegisterNodeAtStartup != nil {
		in, out := &in.RegisterNodeAtStartup, &out.RegisterNodeAtStartup
		*out = new(bool)
		**out = **in
	}
	if in.RegisterNodePeriod != nil {
		in, out := &in.RegisterNodePeriod, &out.RegisterNodePeriod
		*out = new(int64)
		**out = **in
	}
	if in.TokenStore != nil {
		in, out := &in.TokenStore, &out.TokenStore
		*out = new(string)
		**out = **in
	}
	if in.TokenCookiePath != nil {
		in, out := &in.TokenCookiePath, &out.TokenCookiePath
		*out = new(string)
		**out = **in
	}
	if in.PrincipalAttribute != nil {
		in, out := &in.PrincipalAttribute, &out.PrincipalAttribute
		*out = new(string)
		**out = **in
	}
	if in.ProxyUrl != nil {
		in, out := &in.ProxyUrl, &out.ProxyUrl
		*out = new(string)
		**out = **in
	}
	if in.TurnOffChangeSessionIdOnLogin != nil {
		in, out := &in.TurnOffChangeSessionIdOnLogin, &out.TurnOffChangeSessionIdOnLogin
		*out = new(bool)
		**out = **in
	}
	if in.TokenMinimumTimeToLive != nil {
		in, out := &in.TokenMinimumTimeToLive, &out.TokenMinimumTimeToLive
		*out = new(int64)
		**out = **in
	}
	if in.MinTimeBetweenJwksRequests != nil {
		in, out := &in.MinTimeBetweenJwksRequests, &out.MinTimeBetweenJwksRequests
		*out = new(int64)
		**out = **in
	}
	if in.PublicKeyCacheTtl != nil {
		in, out := &in.PublicKeyCacheTtl, &out.PublicKeyCacheTtl
		*out = new(int64)
		**out = **in
	}
	if in.IgnoreOauthQueryParameter != nil {
		in, out := &in.IgnoreOauthQueryParameter, &out.IgnoreOauthQueryParameter
		*out = new(bool)
		**out = **in
	}
	if in.VerifyTokenAudience != nil {
		in, out := &in.VerifyTokenAudience, &out.VerifyTokenAudience
		*out = new(bool)
		**out = **in
	}
	if in.EnableBasicAuth != nil {
		in, out := &in.EnableBasicAuth, &out.EnableBasicAuth
		*out = new(bool)
		**out = **in
	}
	if in.ConfidentialPort != nil {
		in, out := &in.ConfidentialPort, &out.ConfidentialPort
		*out = new(int32)
		**out = **in
	}
	if in.RedirectRewriteRules != nil {
		in, out := &in.RedirectRewriteRules, &out.RedirectRewriteRules
		*out = make([]KeyValueType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycloakModuleConfigurationType.
func (in *KeycloakModuleConfigurationType) DeepCopy() *KeycloakModuleConfigurationType {
	if in == nil {
		return nil
	}
	out := new(KeycloakModuleConfigurationType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginModuleReferenceType) DeepCopyInto(out *LoginModuleReferenceType) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Flag != nil {
		in, out := &in.Flag, &out.Flag
		*out = new(string)
		**out = **in
	}
	if in.Debug != nil {
		in, out := &in.Debug, &out.Debug
		*out = new(bool)
		**out = **in
	}
	if in.Reload != nil {
		in, out := &in.Reload, &out.Reload
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginModuleReferenceType.
func (in *LoginModuleReferenceType) DeepCopy() *LoginModuleReferenceType {
	if in == nil {
		return nil
	}
	out := new(LoginModuleReferenceType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginModulesType) DeepCopyInto(out *LoginModulesType) {
	*out = *in
	if in.PropertiesLoginModules != nil {
		in, out := &in.PropertiesLoginModules, &out.PropertiesLoginModules
		*out = make([]PropertiesLoginModuleType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GuestLoginModules != nil {
		in, out := &in.GuestLoginModules, &out.GuestLoginModules
		*out = make([]GuestLoginModuleType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KeycloakLoginModules != nil {
		in, out := &in.KeycloakLoginModules, &out.KeycloakLoginModules
		*out = make([]KeycloakLoginModuleType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginModulesType.
func (in *LoginModulesType) DeepCopy() *LoginModulesType {
	if in == nil {
		return nil
	}
	out := new(LoginModulesType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementSecuritySettingsType) DeepCopyInto(out *ManagementSecuritySettingsType) {
	*out = *in
	if in.HawtioRoles != nil {
		in, out := &in.HawtioRoles, &out.HawtioRoles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Connector.DeepCopyInto(&out.Connector)
	in.Authorisation.DeepCopyInto(&out.Authorisation)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementSecuritySettingsType.
func (in *ManagementSecuritySettingsType) DeepCopy() *ManagementSecuritySettingsType {
	if in == nil {
		return nil
	}
	out := new(ManagementSecuritySettingsType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionType) DeepCopyInto(out *PermissionType) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionType.
func (in *PermissionType) DeepCopy() *PermissionType {
	if in == nil {
		return nil
	}
	out := new(PermissionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PropertiesLoginModuleType) DeepCopyInto(out *PropertiesLoginModuleType) {
	*out = *in
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]UserType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PropertiesLoginModuleType.
func (in *PropertiesLoginModuleType) DeepCopy() *PropertiesLoginModuleType {
	if in == nil {
		return nil
	}
	out := new(PropertiesLoginModuleType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAccessType) DeepCopyInto(out *RoleAccessType) {
	*out = *in
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.AccessList != nil {
		in, out := &in.AccessList, &out.AccessList
		*out = make([]DefaultAccessType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAccessType.
func (in *RoleAccessType) DeepCopy() *RoleAccessType {
	if in == nil {
		return nil
	}
	out := new(RoleAccessType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityDomainsType) DeepCopyInto(out *SecurityDomainsType) {
	*out = *in
	in.BrokerDomain.DeepCopyInto(&out.BrokerDomain)
	in.ConsoleDomain.DeepCopyInto(&out.ConsoleDomain)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityDomainsType.
func (in *SecurityDomainsType) DeepCopy() *SecurityDomainsType {
	if in == nil {
		return nil
	}
	out := new(SecurityDomainsType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecuritySettingsType) DeepCopyInto(out *SecuritySettingsType) {
	*out = *in
	if in.Broker != nil {
		in, out := &in.Broker, &out.Broker
		*out = make([]BrokerSecuritySettingType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Management.DeepCopyInto(&out.Management)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecuritySettingsType.
func (in *SecuritySettingsType) DeepCopy() *SecuritySettingsType {
	if in == nil {
		return nil
	}
	out := new(SecuritySettingsType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserType) DeepCopyInto(out *UserType) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserType.
func (in *UserType) DeepCopy() *UserType {
	if in == nil {
		return nil
	}
	out := new(UserType)
	in.DeepCopyInto(out)
	return out
}
