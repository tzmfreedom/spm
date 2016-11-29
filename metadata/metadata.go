package metadata

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type DeployProblemType string

const (
	DeployProblemTypeWarning DeployProblemType = "Warning"

	DeployProblemTypeError DeployProblemType = "Error"
)

type ManageableState string

const (
	ManageableStateReleased ManageableState = "released"

	ManageableStateDeleted ManageableState = "deleted"

	ManageableStateDeprecated ManageableState = "deprecated"

	ManageableStateInstalled ManageableState = "installed"

	ManageableStateBeta ManageableState = "beta"

	ManageableStateUnmanaged ManageableState = "unmanaged"
)

type RetrieveStatus string

const (
	RetrieveStatusPending RetrieveStatus = "Pending"

	RetrieveStatusInProgress RetrieveStatus = "InProgress"

	RetrieveStatusSucceeded RetrieveStatus = "Succeeded"

	RetrieveStatusFailed RetrieveStatus = "Failed"
)

type DeployStatus string

const (
	DeployStatusPending DeployStatus = "Pending"

	DeployStatusInProgress DeployStatus = "InProgress"

	DeployStatusSucceeded DeployStatus = "Succeeded"

	DeployStatusSucceededPartial DeployStatus = "SucceededPartial"

	DeployStatusFailed DeployStatus = "Failed"

	DeployStatusCanceling DeployStatus = "Canceling"

	DeployStatusCanceled DeployStatus = "Canceled"
)

type ActionLinkType string

const (
	ActionLinkTypeAPI ActionLinkType = "API"

	ActionLinkTypeAPIAsync ActionLinkType = "APIAsync"

	ActionLinkTypeDownload ActionLinkType = "Download"

	ActionLinkTypeUI ActionLinkType = "UI"
)

type ActionLinkHttpMethod string

const (
	ActionLinkHttpMethodHttpDelete ActionLinkHttpMethod = "HttpDelete"

	ActionLinkHttpMethodHttpHead ActionLinkHttpMethod = "HttpHead"

	ActionLinkHttpMethodHttpGet ActionLinkHttpMethod = "HttpGet"

	ActionLinkHttpMethodHttpPatch ActionLinkHttpMethod = "HttpPatch"

	ActionLinkHttpMethodHttpPost ActionLinkHttpMethod = "HttpPost"

	ActionLinkHttpMethodHttpPut ActionLinkHttpMethod = "HttpPut"
)

type ActionLinkUserVisibility string

const (
	ActionLinkUserVisibilityCreator ActionLinkUserVisibility = "Creator"

	ActionLinkUserVisibilityEveryone ActionLinkUserVisibility = "Everyone"

	ActionLinkUserVisibilityEveryoneButCreator ActionLinkUserVisibility = "EveryoneButCreator"

	ActionLinkUserVisibilityManager ActionLinkUserVisibility = "Manager"

	ActionLinkUserVisibilityCustomUser ActionLinkUserVisibility = "CustomUser"

	ActionLinkUserVisibilityCustomExcludedUser ActionLinkUserVisibility = "CustomExcludedUser"
)

type PlatformActionGroupCategory string

const (
	PlatformActionGroupCategoryPrimary PlatformActionGroupCategory = "Primary"

	PlatformActionGroupCategoryOverflow PlatformActionGroupCategory = "Overflow"
)

type ActionLinkExecutionsAllowed string

const (
	ActionLinkExecutionsAllowedOnce ActionLinkExecutionsAllowed = "Once"

	ActionLinkExecutionsAllowedOncePerUser ActionLinkExecutionsAllowed = "OncePerUser"

	ActionLinkExecutionsAllowedUnlimited ActionLinkExecutionsAllowed = "Unlimited"
)

type ReportSummaryType string

const (
	ReportSummaryTypeSum ReportSummaryType = "Sum"

	ReportSummaryTypeAverage ReportSummaryType = "Average"

	ReportSummaryTypeMaximum ReportSummaryType = "Maximum"

	ReportSummaryTypeMinimum ReportSummaryType = "Minimum"

	ReportSummaryTypeNone ReportSummaryType = "None"
)

type ReportJobSourceTypes string

const (
	ReportJobSourceTypesTabular ReportJobSourceTypes = "tabular"

	ReportJobSourceTypesSummary ReportJobSourceTypes = "summary"

	ReportJobSourceTypesSnapshot ReportJobSourceTypes = "snapshot"
)

type ProcessSubmitterType string

const (
	ProcessSubmitterTypeGroup ProcessSubmitterType = "group"

	ProcessSubmitterTypeRole ProcessSubmitterType = "role"

	ProcessSubmitterTypeUser ProcessSubmitterType = "user"

	ProcessSubmitterTypeRoleSubordinates ProcessSubmitterType = "roleSubordinates"

	ProcessSubmitterTypeRoleSubordinatesInternal ProcessSubmitterType = "roleSubordinatesInternal"

	ProcessSubmitterTypeOwner ProcessSubmitterType = "owner"

	ProcessSubmitterTypeCreator ProcessSubmitterType = "creator"

	ProcessSubmitterTypePartnerUser ProcessSubmitterType = "partnerUser"

	ProcessSubmitterTypeCustomerPortalUser ProcessSubmitterType = "customerPortalUser"

	ProcessSubmitterTypePortalRole ProcessSubmitterType = "portalRole"

	ProcessSubmitterTypePortalRoleSubordinates ProcessSubmitterType = "portalRoleSubordinates"

	ProcessSubmitterTypeAllInternalUsers ProcessSubmitterType = "allInternalUsers"
)

type WorkflowActionType string

const (
	WorkflowActionTypeFieldUpdate WorkflowActionType = "FieldUpdate"

	WorkflowActionTypeKnowledgePublish WorkflowActionType = "KnowledgePublish"

	WorkflowActionTypeTask WorkflowActionType = "Task"

	WorkflowActionTypeAlert WorkflowActionType = "Alert"

	WorkflowActionTypeSend WorkflowActionType = "Send"

	WorkflowActionTypeOutboundMessage WorkflowActionType = "OutboundMessage"

	WorkflowActionTypeFlowAction WorkflowActionType = "FlowAction"
)

type NextOwnerType string

const (
	NextOwnerTypeAdhoc NextOwnerType = "adhoc"

	NextOwnerTypeUser NextOwnerType = "user"

	NextOwnerTypeUserHierarchyField NextOwnerType = "userHierarchyField"

	NextOwnerTypeRelatedUserField NextOwnerType = "relatedUserField"

	NextOwnerTypeQueue NextOwnerType = "queue"
)

type RoutingType string

const (
	RoutingTypeUnanimous RoutingType = "Unanimous"

	RoutingTypeFirstResponse RoutingType = "FirstResponse"
)

type FilterOperation string

const (
	FilterOperationEquals FilterOperation = "equals"

	FilterOperationNotEqual FilterOperation = "notEqual"

	FilterOperationLessThan FilterOperation = "lessThan"

	FilterOperationGreaterThan FilterOperation = "greaterThan"

	FilterOperationLessOrEqual FilterOperation = "lessOrEqual"

	FilterOperationGreaterOrEqual FilterOperation = "greaterOrEqual"

	FilterOperationContains FilterOperation = "contains"

	FilterOperationNotContain FilterOperation = "notContain"

	FilterOperationStartsWith FilterOperation = "startsWith"

	FilterOperationIncludes FilterOperation = "includes"

	FilterOperationExcludes FilterOperation = "excludes"

	FilterOperationWithin FilterOperation = "within"
)

type StepCriteriaNotMetType string

const (
	StepCriteriaNotMetTypeApproveRecord StepCriteriaNotMetType = "ApproveRecord"

	StepCriteriaNotMetTypeRejectRecord StepCriteriaNotMetType = "RejectRecord"

	StepCriteriaNotMetTypeGotoNextStep StepCriteriaNotMetType = "GotoNextStep"
)

type StepRejectBehaviorType string

const (
	StepRejectBehaviorTypeRejectRequest StepRejectBehaviorType = "RejectRequest"

	StepRejectBehaviorTypeBackToPrevious StepRejectBehaviorType = "BackToPrevious"
)

type RecordEditabilityType string

const (
	RecordEditabilityTypeAdminOnly RecordEditabilityType = "AdminOnly"

	RecordEditabilityTypeAdminOrCurrentApprover RecordEditabilityType = "AdminOrCurrentApprover"
)

type AssignToLookupValueType string

const (
	AssignToLookupValueTypeUser AssignToLookupValueType = "User"

	AssignToLookupValueTypeQueue AssignToLookupValueType = "Queue"
)

type BusinessHoursSourceType string

const (
	BusinessHoursSourceTypeNone BusinessHoursSourceType = "None"

	BusinessHoursSourceTypeCase BusinessHoursSourceType = "Case"

	BusinessHoursSourceTypeStatic BusinessHoursSourceType = "Static"
)

type EscalationStartTimeType string

const (
	EscalationStartTimeTypeCaseCreation EscalationStartTimeType = "CaseCreation"

	EscalationStartTimeTypeCaseLastModified EscalationStartTimeType = "CaseLastModified"
)

type PlatformActionListContext string

const (
	PlatformActionListContextListView PlatformActionListContext = "ListView"

	PlatformActionListContextRelatedList PlatformActionListContext = "RelatedList"

	PlatformActionListContextListViewRecord PlatformActionListContext = "ListViewRecord"

	PlatformActionListContextRelatedListRecord PlatformActionListContext = "RelatedListRecord"

	PlatformActionListContextRecord PlatformActionListContext = "Record"

	PlatformActionListContextFeedElement PlatformActionListContext = "FeedElement"

	PlatformActionListContextChatter PlatformActionListContext = "Chatter"

	PlatformActionListContextGlobal PlatformActionListContext = "Global"

	PlatformActionListContextFlexipage PlatformActionListContext = "Flexipage"

	PlatformActionListContextMruList PlatformActionListContext = "MruList"

	PlatformActionListContextMruRow PlatformActionListContext = "MruRow"

	PlatformActionListContextRecordEdit PlatformActionListContext = "RecordEdit"

	PlatformActionListContextPhoto PlatformActionListContext = "Photo"

	PlatformActionListContextBannerPhoto PlatformActionListContext = "BannerPhoto"

	PlatformActionListContextObjectHomeChart PlatformActionListContext = "ObjectHomeChart"

	PlatformActionListContextListViewDefinition PlatformActionListContext = "ListViewDefinition"

	PlatformActionListContextDockable PlatformActionListContext = "Dockable"

	PlatformActionListContextLookup PlatformActionListContext = "Lookup"

	PlatformActionListContextAssistant PlatformActionListContext = "Assistant"
)

type PlatformActionType string

const (
	PlatformActionTypeQuickAction PlatformActionType = "QuickAction"

	PlatformActionTypeStandardButton PlatformActionType = "StandardButton"

	PlatformActionTypeCustomButton PlatformActionType = "CustomButton"

	PlatformActionTypeProductivityAction PlatformActionType = "ProductivityAction"

	PlatformActionTypeActionLink PlatformActionType = "ActionLink"

	PlatformActionTypeInvocableAction PlatformActionType = "InvocableAction"
)

type AuraBundleType string

const (
	AuraBundleTypeApplication AuraBundleType = "Application"

	AuraBundleTypeComponent AuraBundleType = "Component"

	AuraBundleTypeEvent AuraBundleType = "Event"

	AuraBundleTypeInterface AuraBundleType = "Interface"

	AuraBundleTypeTokens AuraBundleType = "Tokens"
)

type AuthProviderType string

const (
	AuthProviderTypeFacebook AuthProviderType = "Facebook"

	AuthProviderTypeJanrain AuthProviderType = "Janrain"

	AuthProviderTypeSalesforce AuthProviderType = "Salesforce"

	AuthProviderTypeOpenIdConnect AuthProviderType = "OpenIdConnect"

	AuthProviderTypeMicrosoftACS AuthProviderType = "MicrosoftACS"

	AuthProviderTypeLinkedIn AuthProviderType = "LinkedIn"

	AuthProviderTypeTwitter AuthProviderType = "Twitter"

	AuthProviderTypeGoogle AuthProviderType = "Google"

	AuthProviderTypeGitHub AuthProviderType = "GitHub"

	AuthProviderTypeCustom AuthProviderType = "Custom"
)

type ForecastCategories string

const (
	ForecastCategoriesOmitted ForecastCategories = "Omitted"

	ForecastCategoriesPipeline ForecastCategories = "Pipeline"

	ForecastCategoriesBestCase ForecastCategories = "BestCase"

	ForecastCategoriesForecast ForecastCategories = "Forecast"

	ForecastCategoriesClosed ForecastCategories = "Closed"
)

type FeedItemDisplayFormat string

const (
	FeedItemDisplayFormatDefault FeedItemDisplayFormat = "Default"

	FeedItemDisplayFormatHideBlankLines FeedItemDisplayFormat = "HideBlankLines"
)

type FeedItemType string

const (
	FeedItemTypeTrackedChange FeedItemType = "TrackedChange"

	FeedItemTypeUserStatus FeedItemType = "UserStatus"

	FeedItemTypeTextPost FeedItemType = "TextPost"

	FeedItemTypeAdvancedTextPost FeedItemType = "AdvancedTextPost"

	FeedItemTypeLinkPost FeedItemType = "LinkPost"

	FeedItemTypeContentPost FeedItemType = "ContentPost"

	FeedItemTypePollPost FeedItemType = "PollPost"

	FeedItemTypeRypplePost FeedItemType = "RypplePost"

	FeedItemTypeProfileSkillPost FeedItemType = "ProfileSkillPost"

	FeedItemTypeDashboardComponentSnapshot FeedItemType = "DashboardComponentSnapshot"

	FeedItemTypeApprovalPost FeedItemType = "ApprovalPost"

	FeedItemTypeCaseCommentPost FeedItemType = "CaseCommentPost"

	FeedItemTypeReplyPost FeedItemType = "ReplyPost"

	FeedItemTypeEmailMessageEvent FeedItemType = "EmailMessageEvent"

	FeedItemTypeCallLogPost FeedItemType = "CallLogPost"

	FeedItemTypeChangeStatusPost FeedItemType = "ChangeStatusPost"

	FeedItemTypeAttachArticleEvent FeedItemType = "AttachArticleEvent"

	FeedItemTypeMilestoneEvent FeedItemType = "MilestoneEvent"

	FeedItemTypeActivityEvent FeedItemType = "ActivityEvent"

	FeedItemTypeChatTranscriptPost FeedItemType = "ChatTranscriptPost"

	FeedItemTypeCollaborationGroupCreated FeedItemType = "CollaborationGroupCreated"

	FeedItemTypeCollaborationGroupUnarchived FeedItemType = "CollaborationGroupUnarchived"

	FeedItemTypeSocialPost FeedItemType = "SocialPost"

	FeedItemTypeQuestionPost FeedItemType = "QuestionPost"

	FeedItemTypeFacebookPost FeedItemType = "FacebookPost"

	FeedItemTypeBasicTemplateFeedItem FeedItemType = "BasicTemplateFeedItem"

	FeedItemTypeCreateRecordEvent FeedItemType = "CreateRecordEvent"

	FeedItemTypeCanvasPost FeedItemType = "CanvasPost"

	FeedItemTypeAnnouncementPost FeedItemType = "AnnouncementPost"
)

type EmailToCaseOnFailureActionType string

const (
	EmailToCaseOnFailureActionTypeBounce EmailToCaseOnFailureActionType = "Bounce"

	EmailToCaseOnFailureActionTypeDiscard EmailToCaseOnFailureActionType = "Discard"

	EmailToCaseOnFailureActionTypeRequeue EmailToCaseOnFailureActionType = "Requeue"
)

type EmailToCaseRoutingAddressType string

const (
	EmailToCaseRoutingAddressTypeEmailToCase EmailToCaseRoutingAddressType = "EmailToCase"

	EmailToCaseRoutingAddressTypeOutlook EmailToCaseRoutingAddressType = "Outlook"
)

type MappingOperation string

const (
	MappingOperationAutofill MappingOperation = "Autofill"

	MappingOperationOverwrite MappingOperation = "Overwrite"
)

type CleanRuleStatus string

const (
	CleanRuleStatusInactive CleanRuleStatus = "Inactive"

	CleanRuleStatusActive CleanRuleStatus = "Active"
)

type CommunityTemplateBundleInfoType string

const (
	CommunityTemplateBundleInfoTypeHighlight CommunityTemplateBundleInfoType = "Highlight"

	CommunityTemplateBundleInfoTypePreviewImage CommunityTemplateBundleInfoType = "PreviewImage"
)

type CommunityTemplateCategory string

const (
	CommunityTemplateCategoryIT CommunityTemplateCategory = "IT"

	CommunityTemplateCategoryMarketing CommunityTemplateCategory = "Marketing"

	CommunityTemplateCategorySales CommunityTemplateCategory = "Sales"

	CommunityTemplateCategoryService CommunityTemplateCategory = "Service"
)

type CommunityThemeLayoutType string

const (
	CommunityThemeLayoutTypeInner CommunityThemeLayoutType = "Inner"
)

type AccessMethod string

const (
	AccessMethodGet AccessMethod = "Get"

	AccessMethodPost AccessMethod = "Post"
)

type CanvasLocationOptions string

const (
	CanvasLocationOptionsNone CanvasLocationOptions = "None"

	CanvasLocationOptionsChatter CanvasLocationOptions = "Chatter"

	CanvasLocationOptionsUserProfile CanvasLocationOptions = "UserProfile"

	CanvasLocationOptionsVisualforce CanvasLocationOptions = "Visualforce"

	CanvasLocationOptionsAura CanvasLocationOptions = "Aura"

	CanvasLocationOptionsPublisher CanvasLocationOptions = "Publisher"

	CanvasLocationOptionsChatterFeed CanvasLocationOptions = "ChatterFeed"

	CanvasLocationOptionsServiceDesk CanvasLocationOptions = "ServiceDesk"

	CanvasLocationOptionsOpenCTI CanvasLocationOptions = "OpenCTI"

	CanvasLocationOptionsAppLauncher CanvasLocationOptions = "AppLauncher"

	CanvasLocationOptionsMobileNav CanvasLocationOptions = "MobileNav"

	CanvasLocationOptionsPageLayout CanvasLocationOptions = "PageLayout"
)

type CanvasOptions string

const (
	CanvasOptionsHideShare CanvasOptions = "HideShare"

	CanvasOptionsHideHeader CanvasOptions = "HideHeader"

	CanvasOptionsPersonalEnabled CanvasOptions = "PersonalEnabled"
)

type SamlInitiationMethod string

const (
	SamlInitiationMethodNone SamlInitiationMethod = "None"

	SamlInitiationMethodIdpInitiated SamlInitiationMethod = "IdpInitiated"

	SamlInitiationMethodSpInitiated SamlInitiationMethod = "SpInitiated"
)

type DevicePlatformType string

const (
	DevicePlatformTypeIos DevicePlatformType = "ios"

	DevicePlatformTypeAndroid DevicePlatformType = "android"
)

type DeviceType string

const (
	DeviceTypePhone DeviceType = "phone"

	DeviceTypeTablet DeviceType = "tablet"

	DeviceTypeMinitablet DeviceType = "minitablet"
)

type ConnectedAppOauthAccessScope string

const (
	ConnectedAppOauthAccessScopeBasic ConnectedAppOauthAccessScope = "Basic"

	ConnectedAppOauthAccessScopeApi ConnectedAppOauthAccessScope = "Api"

	ConnectedAppOauthAccessScopeWeb ConnectedAppOauthAccessScope = "Web"

	ConnectedAppOauthAccessScopeFull ConnectedAppOauthAccessScope = "Full"

	ConnectedAppOauthAccessScopeChatter ConnectedAppOauthAccessScope = "Chatter"

	ConnectedAppOauthAccessScopeCustomApplications ConnectedAppOauthAccessScope = "CustomApplications"

	ConnectedAppOauthAccessScopeRefreshToken ConnectedAppOauthAccessScope = "RefreshToken"

	ConnectedAppOauthAccessScopeOpenID ConnectedAppOauthAccessScope = "OpenID"

	ConnectedAppOauthAccessScopeProfile ConnectedAppOauthAccessScope = "Profile"

	ConnectedAppOauthAccessScopeEmail ConnectedAppOauthAccessScope = "Email"

	ConnectedAppOauthAccessScopeAddress ConnectedAppOauthAccessScope = "Address"

	ConnectedAppOauthAccessScopePhone ConnectedAppOauthAccessScope = "Phone"

	ConnectedAppOauthAccessScopeOfflineAccess ConnectedAppOauthAccessScope = "OfflineAccess"

	ConnectedAppOauthAccessScopeCustomPermissions ConnectedAppOauthAccessScope = "CustomPermissions"

	ConnectedAppOauthAccessScopeWave ConnectedAppOauthAccessScope = "Wave"
)

type SamlEncryptionType string

const (
	SamlEncryptionTypeAES128 SamlEncryptionType = "AES128"

	SamlEncryptionTypeAES256 SamlEncryptionType = "AES256"

	SamlEncryptionTypeTripleDes SamlEncryptionType = "TripleDes"
)

type SamlNameIdFormatType string

const (
	SamlNameIdFormatTypeUnspecified SamlNameIdFormatType = "Unspecified"

	SamlNameIdFormatTypeEmailAddress SamlNameIdFormatType = "EmailAddress"

	SamlNameIdFormatTypePersistent SamlNameIdFormatType = "Persistent"

	SamlNameIdFormatTypeTransient SamlNameIdFormatType = "Transient"
)

type SamlSubjectType string

const (
	SamlSubjectTypeUsername SamlSubjectType = "Username"

	SamlSubjectTypeFederationId SamlSubjectType = "FederationId"

	SamlSubjectTypeUserId SamlSubjectType = "UserId"

	SamlSubjectTypeSpokeId SamlSubjectType = "SpokeId"

	SamlSubjectTypeCustomAttribute SamlSubjectType = "CustomAttribute"

	SamlSubjectTypePersistentId SamlSubjectType = "PersistentId"
)

type FormFactor string

const (
	FormFactorSmall FormFactor = "Small"

	FormFactorMedium FormFactor = "Medium"

	FormFactorLarge FormFactor = "Large"
)

type ActionOverrideType string

const (
	ActionOverrideTypeDefault ActionOverrideType = "Default"

	ActionOverrideTypeStandard ActionOverrideType = "Standard"

	ActionOverrideTypeScontrol ActionOverrideType = "Scontrol"

	ActionOverrideTypeVisualforce ActionOverrideType = "Visualforce"

	ActionOverrideTypeFlexipage ActionOverrideType = "Flexipage"
)

type NavType string

const (
	NavTypeStandard NavType = "Standard"

	NavTypeConsole NavType = "Console"
)

type UiType string

const (
	UiTypeAloha UiType = "Aloha"

	UiTypeLightning UiType = "Lightning"
)

type SortOrder string

const (
	SortOrderAsc SortOrder = "Asc"

	SortOrderDesc SortOrder = "Desc"
)

type FieldType string

const (
	FieldTypeAutoNumber FieldType = "AutoNumber"

	FieldTypeLookup FieldType = "Lookup"

	FieldTypeMasterDetail FieldType = "MasterDetail"

	FieldTypeCheckbox FieldType = "Checkbox"

	FieldTypeCurrency FieldType = "Currency"

	FieldTypeDate FieldType = "Date"

	FieldTypeDateTime FieldType = "DateTime"

	FieldTypeEmail FieldType = "Email"

	FieldTypeNumber FieldType = "Number"

	FieldTypePercent FieldType = "Percent"

	FieldTypePhone FieldType = "Phone"

	FieldTypePicklist FieldType = "Picklist"

	FieldTypeMultiselectPicklist FieldType = "MultiselectPicklist"

	FieldTypeText FieldType = "Text"

	FieldTypeTextArea FieldType = "TextArea"

	FieldTypeLongTextArea FieldType = "LongTextArea"

	FieldTypeHtml FieldType = "Html"

	FieldTypeUrl FieldType = "Url"

	FieldTypeEncryptedText FieldType = "EncryptedText"

	FieldTypeSummary FieldType = "Summary"

	FieldTypeHierarchy FieldType = "Hierarchy"

	FieldTypeFile FieldType = "File"

	FieldTypeMetadataRelationship FieldType = "MetadataRelationship"

	FieldTypeExternalLookup FieldType = "ExternalLookup"

	FieldTypeIndirectLookup FieldType = "IndirectLookup"

	FieldTypeCustomDataType FieldType = "CustomDataType"
)

type FeedItemVisibility string

const (
	FeedItemVisibilityAllUsers FeedItemVisibility = "AllUsers"

	FeedItemVisibilityInternalUsers FeedItemVisibility = "InternalUsers"
)

type DeleteConstraint string

const (
	DeleteConstraintCascade DeleteConstraint = "Cascade"

	DeleteConstraintRestrict DeleteConstraint = "Restrict"

	DeleteConstraintSetNull DeleteConstraint = "SetNull"
)

type FieldManageability string

const (
	FieldManageabilityDeveloperControlled FieldManageability = "DeveloperControlled"

	FieldManageabilitySubscriberControlled FieldManageability = "SubscriberControlled"

	FieldManageabilityLocked FieldManageability = "Locked"
)

type TreatBlanksAs string

const (
	TreatBlanksAsBlankAsBlank TreatBlanksAs = "BlankAsBlank"

	TreatBlanksAsBlankAsZero TreatBlanksAs = "BlankAsZero"
)

type EncryptedFieldMaskChar string

const (
	EncryptedFieldMaskCharAsterisk EncryptedFieldMaskChar = "asterisk"

	EncryptedFieldMaskCharX EncryptedFieldMaskChar = "X"
)

type EncryptedFieldMaskType string

const (
	EncryptedFieldMaskTypeAll EncryptedFieldMaskType = "all"

	EncryptedFieldMaskTypeCreditCard EncryptedFieldMaskType = "creditCard"

	EncryptedFieldMaskTypeSsn EncryptedFieldMaskType = "ssn"

	EncryptedFieldMaskTypeLastFour EncryptedFieldMaskType = "lastFour"

	EncryptedFieldMaskTypeSin EncryptedFieldMaskType = "sin"

	EncryptedFieldMaskTypeNino EncryptedFieldMaskType = "nino"
)

type SummaryOperations string

const (
	SummaryOperationsCount SummaryOperations = "count"

	SummaryOperationsSum SummaryOperations = "sum"

	SummaryOperationsMin SummaryOperations = "min"

	SummaryOperationsMax SummaryOperations = "max"
)

type Channel string

const (
	ChannelAllChannels Channel = "AllChannels"

	ChannelApp Channel = "App"

	ChannelPkb Channel = "Pkb"

	ChannelCsp Channel = "Csp"

	ChannelPrm Channel = "Prm"
)

type Template string

const (
	TemplatePage Template = "Page"

	TemplateTab Template = "Tab"

	TemplateToc Template = "Toc"
)

type CustomSettingsType string

const (
	CustomSettingsTypeList CustomSettingsType = "List"

	CustomSettingsTypeHierarchy CustomSettingsType = "Hierarchy"
)

type DeploymentStatus string

const (
	DeploymentStatusInDevelopment DeploymentStatus = "InDevelopment"

	DeploymentStatusDeployed DeploymentStatus = "Deployed"
)

type SharingModel string

const (
	SharingModelPrivate SharingModel = "Private"

	SharingModelRead SharingModel = "Read"

	SharingModelReadSelect SharingModel = "ReadSelect"

	SharingModelReadWrite SharingModel = "ReadWrite"

	SharingModelReadWriteTransfer SharingModel = "ReadWriteTransfer"

	SharingModelFullAccess SharingModel = "FullAccess"

	SharingModelControlledByParent SharingModel = "ControlledByParent"
)

type Gender string

const (
	GenderNeuter Gender = "Neuter"

	GenderMasculine Gender = "Masculine"

	GenderFeminine Gender = "Feminine"

	GenderAnimateMasculine Gender = "AnimateMasculine"
)

type FilterScope string

const (
	FilterScopeEverything FilterScope = "Everything"

	FilterScopeMine FilterScope = "Mine"

	FilterScopeQueue FilterScope = "Queue"

	FilterScopeDelegated FilterScope = "Delegated"

	FilterScopeMyTerritory FilterScope = "MyTerritory"

	FilterScopeMyTeamTerritory FilterScope = "MyTeamTerritory"

	FilterScopeTeam FilterScope = "Team"
)

type Language string

const (
	LanguageEnUS Language = "enUS"

	LanguageDe Language = "de"

	LanguageEs Language = "es"

	LanguageFr Language = "fr"

	LanguageIt Language = "it"

	LanguageJa Language = "ja"

	LanguageSv Language = "sv"

	LanguageKo Language = "ko"

	LanguageZhTW Language = "zhTW"

	LanguageZhCN Language = "zhCN"

	LanguagePtBR Language = "ptBR"

	LanguageNlNL Language = "nlNL"

	LanguageDa Language = "da"

	LanguageTh Language = "th"

	LanguageFi Language = "fi"

	LanguageRu Language = "ru"

	LanguageEsMX Language = "esMX"

	LanguageNo Language = "no"

	LanguageHu Language = "hu"

	LanguagePl Language = "pl"

	LanguageCs Language = "cs"

	LanguageTr Language = "tr"

	LanguageIn Language = "in"

	LanguageRo Language = "ro"

	LanguageVi Language = "vi"

	LanguageUk Language = "uk"

	LanguageIw Language = "iw"

	LanguageEl Language = "el"

	LanguageBg Language = "bg"

	LanguageEnGB Language = "enGB"

	LanguageAr Language = "ar"

	LanguageSk Language = "sk"

	LanguagePtPT Language = "ptPT"

	LanguageHr Language = "hr"

	LanguageSl Language = "sl"

	LanguageFrCA Language = "frCA"

	LanguageKa Language = "ka"

	LanguageSr Language = "sr"

	LanguageSh Language = "sh"

	LanguageEnAU Language = "enAU"

	LanguageEnMY Language = "enMY"

	LanguageEnIN Language = "enIN"

	LanguageEnPH Language = "enPH"

	LanguageEnCA Language = "enCA"

	LanguageRoMD Language = "roMD"

	LanguageBs Language = "bs"

	LanguageMk Language = "mk"

	LanguageLv Language = "lv"

	LanguageLt Language = "lt"

	LanguageEt Language = "et"

	LanguageSq Language = "sq"

	LanguageShME Language = "shME"

	LanguageMt Language = "mt"

	LanguageGa Language = "ga"

	LanguageEu Language = "eu"

	LanguageCy Language = "cy"

	LanguageIs Language = "is"

	LanguageMs Language = "ms"

	LanguageTl Language = "tl"

	LanguageLb Language = "lb"

	LanguageRm Language = "rm"

	LanguageHy Language = "hy"

	LanguageHi Language = "hi"

	LanguageUr Language = "ur"

	LanguageBn Language = "bn"

	LanguageDeAT Language = "deAT"

	LanguageDeCH Language = "deCH"

	LanguageTa Language = "ta"

	LanguageArDZ Language = "arDZ"

	LanguageArBH Language = "arBH"

	LanguageArEG Language = "arEG"

	LanguageArIQ Language = "arIQ"

	LanguageArJO Language = "arJO"

	LanguageArKW Language = "arKW"

	LanguageArLB Language = "arLB"

	LanguageArLY Language = "arLY"

	LanguageArMA Language = "arMA"

	LanguageArOM Language = "arOM"

	LanguageArQA Language = "arQA"

	LanguageArSA Language = "arSA"

	LanguageArSD Language = "arSD"

	LanguageArSY Language = "arSY"

	LanguageArTN Language = "arTN"

	LanguageArAE Language = "arAE"

	LanguageArYE Language = "arYE"

	LanguageZhSG Language = "zhSG"

	LanguageZhHK Language = "zhHK"

	LanguageEnHK Language = "enHK"

	LanguageEnIE Language = "enIE"

	LanguageEnSG Language = "enSG"

	LanguageEnZA Language = "enZA"

	LanguageFrBE Language = "frBE"

	LanguageFrLU Language = "frLU"

	LanguageFrCH Language = "frCH"

	LanguageDeLU Language = "deLU"

	LanguageItCH Language = "itCH"

	LanguageEsAR Language = "esAR"

	LanguageEsBO Language = "esBO"

	LanguageEsCL Language = "esCL"

	LanguageEsCO Language = "esCO"

	LanguageEsCR Language = "esCR"

	LanguageEsDO Language = "esDO"

	LanguageEsEC Language = "esEC"

	LanguageEsSV Language = "esSV"

	LanguageEsGT Language = "esGT"

	LanguageEsHN Language = "esHN"

	LanguageEsNI Language = "esNI"

	LanguageEsPA Language = "esPA"

	LanguageEsPY Language = "esPY"

	LanguageEsPE Language = "esPE"

	LanguageEsPR Language = "esPR"

	LanguageEsUS Language = "esUS"

	LanguageEsUY Language = "esUY"

	LanguageEsVE Language = "esVE"

	LanguageEo Language = "eo"

	LanguageIwEO Language = "iwEO"
)

type StartsWith string

const (
	StartsWithConsonant StartsWith = "Consonant"

	StartsWithVowel StartsWith = "Vowel"

	StartsWithSpecial StartsWith = "Special"
)

type SetupObjectVisibility string

const (
	SetupObjectVisibilityProtected SetupObjectVisibility = "Protected"

	SetupObjectVisibilityPublic SetupObjectVisibility = "Public"
)

type WebLinkAvailability string

const (
	WebLinkAvailabilityOnline WebLinkAvailability = "online"

	WebLinkAvailabilityOffline WebLinkAvailability = "offline"
)

type WebLinkDisplayType string

const (
	WebLinkDisplayTypeLink WebLinkDisplayType = "link"

	WebLinkDisplayTypeButton WebLinkDisplayType = "button"

	WebLinkDisplayTypeMassActionButton WebLinkDisplayType = "massActionButton"
)

type Encoding string

const (
	EncodingUTF8 Encoding = "UTF8"

	EncodingISO88591 Encoding = "ISO88591"

	EncodingShiftJIS Encoding = "ShiftJIS"

	EncodingISO2022JP Encoding = "ISO2022JP"

	EncodingEUCJP Encoding = "EUCJP"

	EncodingKsc56011987 Encoding = "ksc56011987"

	EncodingBig5 Encoding = "Big5"

	EncodingGB2312 Encoding = "GB2312"

	EncodingBig5HKSCS Encoding = "Big5HKSCS"

	EncodingXSJIS0213 Encoding = "xSJIS0213"
)

type WebLinkType string

const (
	WebLinkTypeUrl WebLinkType = "url"

	WebLinkTypeSControl WebLinkType = "sControl"

	WebLinkTypeJavascript WebLinkType = "javascript"

	WebLinkTypePage WebLinkType = "page"

	WebLinkTypeFlow WebLinkType = "flow"
)

type WebLinkWindowType string

const (
	WebLinkWindowTypeNewWindow WebLinkWindowType = "newWindow"

	WebLinkWindowTypeSidebar WebLinkWindowType = "sidebar"

	WebLinkWindowTypeNoSidebar WebLinkWindowType = "noSidebar"

	WebLinkWindowTypeReplace WebLinkWindowType = "replace"

	WebLinkWindowTypeOnClickJavaScript WebLinkWindowType = "onClickJavaScript"
)

type WebLinkPosition string

const (
	WebLinkPositionFullScreen WebLinkPosition = "fullScreen"

	WebLinkPositionNone WebLinkPosition = "none"

	WebLinkPositionTopLeft WebLinkPosition = "topLeft"
)

type Article string

const (
	ArticleNone Article = "None"

	ArticleIndefinite Article = "Indefinite"

	ArticleDefinite Article = "Definite"
)

type CaseType string

const (
	CaseTypeNominative CaseType = "Nominative"

	CaseTypeAccusative CaseType = "Accusative"

	CaseTypeGenitive CaseType = "Genitive"

	CaseTypeDative CaseType = "Dative"

	CaseTypeInessive CaseType = "Inessive"

	CaseTypeElative CaseType = "Elative"

	CaseTypeIllative CaseType = "Illative"

	CaseTypeAdessive CaseType = "Adessive"

	CaseTypeAblative CaseType = "Ablative"

	CaseTypeAllative CaseType = "Allative"

	CaseTypeEssive CaseType = "Essive"

	CaseTypeTranslative CaseType = "Translative"

	CaseTypePartitive CaseType = "Partitive"

	CaseTypeObjective CaseType = "Objective"

	CaseTypeSubjective CaseType = "Subjective"

	CaseTypeInstrumental CaseType = "Instrumental"

	CaseTypePrepositional CaseType = "Prepositional"

	CaseTypeLocative CaseType = "Locative"

	CaseTypeVocative CaseType = "Vocative"

	CaseTypeSublative CaseType = "Sublative"

	CaseTypeSuperessive CaseType = "Superessive"

	CaseTypeDelative CaseType = "Delative"

	CaseTypeCausalfinal CaseType = "Causalfinal"

	CaseTypeEssiveformal CaseType = "Essiveformal"

	CaseTypeTermanative CaseType = "Termanative"

	CaseTypeDistributive CaseType = "Distributive"

	CaseTypeErgative CaseType = "Ergative"

	CaseTypeAdverbial CaseType = "Adverbial"

	CaseTypeAbessive CaseType = "Abessive"

	CaseTypeComitative CaseType = "Comitative"
)

type Possessive string

const (
	PossessiveNone Possessive = "None"

	PossessiveFirst Possessive = "First"

	PossessiveSecond Possessive = "Second"
)

type SiteClickjackProtectionLevel string

const (
	SiteClickjackProtectionLevelAllowAllFraming SiteClickjackProtectionLevel = "AllowAllFraming"

	SiteClickjackProtectionLevelSameOriginOnly SiteClickjackProtectionLevel = "SameOriginOnly"

	SiteClickjackProtectionLevelNoFraming SiteClickjackProtectionLevel = "NoFraming"
)

type SiteRedirect string

const (
	SiteRedirectPermanent SiteRedirect = "Permanent"

	SiteRedirectTemporary SiteRedirect = "Temporary"
)

type SiteType string

const (
	SiteTypeSiteforce SiteType = "Siteforce"

	SiteTypeVisualforce SiteType = "Visualforce"

	SiteTypeChatterNetwork SiteType = "ChatterNetwork"

	SiteTypeChatterNetworkPicasso SiteType = "ChatterNetworkPicasso"

	SiteTypeUser SiteType = "User"
)

type ChartBackgroundDirection string

const (
	ChartBackgroundDirectionTopToBottom ChartBackgroundDirection = "TopToBottom"

	ChartBackgroundDirectionLeftToRight ChartBackgroundDirection = "LeftToRight"

	ChartBackgroundDirectionDiagonal ChartBackgroundDirection = "Diagonal"
)

type DashboardFilterOperation string

const (
	DashboardFilterOperationEquals DashboardFilterOperation = "equals"

	DashboardFilterOperationNotEqual DashboardFilterOperation = "notEqual"

	DashboardFilterOperationLessThan DashboardFilterOperation = "lessThan"

	DashboardFilterOperationGreaterThan DashboardFilterOperation = "greaterThan"

	DashboardFilterOperationLessOrEqual DashboardFilterOperation = "lessOrEqual"

	DashboardFilterOperationGreaterOrEqual DashboardFilterOperation = "greaterOrEqual"

	DashboardFilterOperationContains DashboardFilterOperation = "contains"

	DashboardFilterOperationNotContain DashboardFilterOperation = "notContain"

	DashboardFilterOperationStartsWith DashboardFilterOperation = "startsWith"

	DashboardFilterOperationIncludes DashboardFilterOperation = "includes"

	DashboardFilterOperationExcludes DashboardFilterOperation = "excludes"

	DashboardFilterOperationBetween DashboardFilterOperation = "between"
)

type ChartRangeType string

const (
	ChartRangeTypeAuto ChartRangeType = "Auto"

	ChartRangeTypeManual ChartRangeType = "Manual"
)

type ChartAxis string

const (
	ChartAxisX ChartAxis = "x"

	ChartAxisY ChartAxis = "y"

	ChartAxisY2 ChartAxis = "y2"

	ChartAxisR ChartAxis = "r"
)

type DashboardComponentType string

const (
	DashboardComponentTypeBar DashboardComponentType = "Bar"

	DashboardComponentTypeBarGrouped DashboardComponentType = "BarGrouped"

	DashboardComponentTypeBarStacked DashboardComponentType = "BarStacked"

	DashboardComponentTypeBarStacked100 DashboardComponentType = "BarStacked100"

	DashboardComponentTypeColumn DashboardComponentType = "Column"

	DashboardComponentTypeColumnGrouped DashboardComponentType = "ColumnGrouped"

	DashboardComponentTypeColumnStacked DashboardComponentType = "ColumnStacked"

	DashboardComponentTypeColumnStacked100 DashboardComponentType = "ColumnStacked100"

	DashboardComponentTypeLine DashboardComponentType = "Line"

	DashboardComponentTypeLineGrouped DashboardComponentType = "LineGrouped"

	DashboardComponentTypePie DashboardComponentType = "Pie"

	DashboardComponentTypeTable DashboardComponentType = "Table"

	DashboardComponentTypeMetric DashboardComponentType = "Metric"

	DashboardComponentTypeGauge DashboardComponentType = "Gauge"

	DashboardComponentTypeLineCumulative DashboardComponentType = "LineCumulative"

	DashboardComponentTypeLineGroupedCumulative DashboardComponentType = "LineGroupedCumulative"

	DashboardComponentTypeScontrol DashboardComponentType = "Scontrol"

	DashboardComponentTypeVisualforcePage DashboardComponentType = "VisualforcePage"

	DashboardComponentTypeDonut DashboardComponentType = "Donut"

	DashboardComponentTypeFunnel DashboardComponentType = "Funnel"

	DashboardComponentTypeColumnLine DashboardComponentType = "ColumnLine"

	DashboardComponentTypeColumnLineGrouped DashboardComponentType = "ColumnLineGrouped"

	DashboardComponentTypeColumnLineStacked DashboardComponentType = "ColumnLineStacked"

	DashboardComponentTypeColumnLineStacked100 DashboardComponentType = "ColumnLineStacked100"

	DashboardComponentTypeScatter DashboardComponentType = "Scatter"

	DashboardComponentTypeScatterGrouped DashboardComponentType = "ScatterGrouped"
)

type DashboardComponentFilter string

const (
	DashboardComponentFilterRowLabelAscending DashboardComponentFilter = "RowLabelAscending"

	DashboardComponentFilterRowLabelDescending DashboardComponentFilter = "RowLabelDescending"

	DashboardComponentFilterRowValueAscending DashboardComponentFilter = "RowValueAscending"

	DashboardComponentFilterRowValueDescending DashboardComponentFilter = "RowValueDescending"
)

type ChartUnits string

const (
	ChartUnitsAuto ChartUnits = "Auto"

	ChartUnitsInteger ChartUnits = "Integer"

	ChartUnitsHundreds ChartUnits = "Hundreds"

	ChartUnitsThousands ChartUnits = "Thousands"

	ChartUnitsMillions ChartUnits = "Millions"

	ChartUnitsBillions ChartUnits = "Billions"

	ChartUnitsTrillions ChartUnits = "Trillions"
)

type ChartLegendPosition string

const (
	ChartLegendPositionRight ChartLegendPosition = "Right"

	ChartLegendPositionBottom ChartLegendPosition = "Bottom"

	ChartLegendPositionOnChart ChartLegendPosition = "OnChart"
)

type DashboardType string

const (
	DashboardTypeSpecifiedUser DashboardType = "SpecifiedUser"

	DashboardTypeLoggedInUser DashboardType = "LoggedInUser"

	DashboardTypeMyTeamUser DashboardType = "MyTeamUser"
)

type DashboardComponentSize string

const (
	DashboardComponentSizeNarrow DashboardComponentSize = "Narrow"

	DashboardComponentSizeMedium DashboardComponentSize = "Medium"

	DashboardComponentSizeWide DashboardComponentSize = "Wide"
)

type DupeActionType string

const (
	DupeActionTypeAllow DupeActionType = "Allow"

	DupeActionTypeBlock DupeActionType = "Block"
)

type DupeSecurityOptionType string

const (
	DupeSecurityOptionTypeEnforceSharingRules DupeSecurityOptionType = "EnforceSharingRules"

	DupeSecurityOptionTypeBypassSharingRules DupeSecurityOptionType = "BypassSharingRules"
)

type MilestoneTimeUnits string

const (
	MilestoneTimeUnitsMinutes MilestoneTimeUnits = "Minutes"

	MilestoneTimeUnitsHours MilestoneTimeUnits = "Hours"

	MilestoneTimeUnitsDays MilestoneTimeUnits = "Days"
)

type EventDeliveryType string

const (
	EventDeliveryTypeStartFlow EventDeliveryType = "StartFlow"

	EventDeliveryTypeResumeFlow EventDeliveryType = "ResumeFlow"
)

type ExternalPrincipalType string

const (
	ExternalPrincipalTypeAnonymous ExternalPrincipalType = "Anonymous"

	ExternalPrincipalTypePerUser ExternalPrincipalType = "PerUser"

	ExternalPrincipalTypeNamedUser ExternalPrincipalType = "NamedUser"
)

type AuthenticationProtocol string

const (
	AuthenticationProtocolNoAuthentication AuthenticationProtocol = "NoAuthentication"

	AuthenticationProtocolOauth AuthenticationProtocol = "Oauth"

	AuthenticationProtocolPassword AuthenticationProtocol = "Password"
)

type ExternalDataSourceType string

const (
	ExternalDataSourceTypeDatajourney ExternalDataSourceType = "Datajourney"

	ExternalDataSourceTypeIdentity ExternalDataSourceType = "Identity"

	ExternalDataSourceTypeOData ExternalDataSourceType = "OData"

	ExternalDataSourceTypeOData4 ExternalDataSourceType = "OData4"

	ExternalDataSourceTypeBoxDataSourceProvider ExternalDataSourceType = "BoxDataSourceProvider"

	ExternalDataSourceTypeSfdcOrg ExternalDataSourceType = "SfdcOrg"

	ExternalDataSourceTypeSimpleURL ExternalDataSourceType = "SimpleURL"

	ExternalDataSourceTypeWrapper ExternalDataSourceType = "Wrapper"
)

type RegionFlagStatus string

const (
	RegionFlagStatusDisabled RegionFlagStatus = "disabled"

	RegionFlagStatusEnabled RegionFlagStatus = "enabled"
)

type ComponentInstancePropertyTypeEnum string

const (
	ComponentInstancePropertyTypeEnumDecorator ComponentInstancePropertyTypeEnum = "decorator"
)

type FlexiPageRegionMode string

const (
	FlexiPageRegionModeAppend FlexiPageRegionMode = "Append"

	FlexiPageRegionModePrepend FlexiPageRegionMode = "Prepend"

	FlexiPageRegionModeReplace FlexiPageRegionMode = "Replace"
)

type FlexiPageRegionType string

const (
	FlexiPageRegionTypeRegion FlexiPageRegionType = "Region"

	FlexiPageRegionTypeFacet FlexiPageRegionType = "Facet"
)

type FlexiPageType string

const (
	FlexiPageTypeAppPage FlexiPageType = "AppPage"

	FlexiPageTypeObjectPage FlexiPageType = "ObjectPage"

	FlexiPageTypeRecordPage FlexiPageType = "RecordPage"

	FlexiPageTypeHomePage FlexiPageType = "HomePage"

	FlexiPageTypeMailAppAppPage FlexiPageType = "MailAppAppPage"

	FlexiPageTypeCommAppPage FlexiPageType = "CommAppPage"

	FlexiPageTypeCommObjectPage FlexiPageType = "CommObjectPage"

	FlexiPageTypeCommQuickActionCreatePage FlexiPageType = "CommQuickActionCreatePage"

	FlexiPageTypeCommRecordPage FlexiPageType = "CommRecordPage"

	FlexiPageTypeCommRelatedListPage FlexiPageType = "CommRelatedListPage"

	FlexiPageTypeCommSearchResultPage FlexiPageType = "CommSearchResultPage"

	FlexiPageTypeCommThemeLayoutPage FlexiPageType = "CommThemeLayoutPage"

	FlexiPageTypeUtilityBar FlexiPageType = "UtilityBar"
)

type FlowAssignmentOperator string

const (
	FlowAssignmentOperatorAssign FlowAssignmentOperator = "Assign"

	FlowAssignmentOperatorAdd FlowAssignmentOperator = "Add"

	FlowAssignmentOperatorSubtract FlowAssignmentOperator = "Subtract"

	FlowAssignmentOperatorAddItem FlowAssignmentOperator = "AddItem"
)

type FlowComparisonOperator string

const (
	FlowComparisonOperatorEqualTo FlowComparisonOperator = "EqualTo"

	FlowComparisonOperatorNotEqualTo FlowComparisonOperator = "NotEqualTo"

	FlowComparisonOperatorGreaterThan FlowComparisonOperator = "GreaterThan"

	FlowComparisonOperatorLessThan FlowComparisonOperator = "LessThan"

	FlowComparisonOperatorGreaterThanOrEqualTo FlowComparisonOperator = "GreaterThanOrEqualTo"

	FlowComparisonOperatorLessThanOrEqualTo FlowComparisonOperator = "LessThanOrEqualTo"

	FlowComparisonOperatorStartsWith FlowComparisonOperator = "StartsWith"

	FlowComparisonOperatorEndsWith FlowComparisonOperator = "EndsWith"

	FlowComparisonOperatorContains FlowComparisonOperator = "Contains"

	FlowComparisonOperatorIsNull FlowComparisonOperator = "IsNull"

	FlowComparisonOperatorWasSet FlowComparisonOperator = "WasSet"

	FlowComparisonOperatorWasSelected FlowComparisonOperator = "WasSelected"

	FlowComparisonOperatorWasVisited FlowComparisonOperator = "WasVisited"
)

type FlowRecordFilterOperator string

const (
	FlowRecordFilterOperatorEqualTo FlowRecordFilterOperator = "EqualTo"

	FlowRecordFilterOperatorNotEqualTo FlowRecordFilterOperator = "NotEqualTo"

	FlowRecordFilterOperatorGreaterThan FlowRecordFilterOperator = "GreaterThan"

	FlowRecordFilterOperatorLessThan FlowRecordFilterOperator = "LessThan"

	FlowRecordFilterOperatorGreaterThanOrEqualTo FlowRecordFilterOperator = "GreaterThanOrEqualTo"

	FlowRecordFilterOperatorLessThanOrEqualTo FlowRecordFilterOperator = "LessThanOrEqualTo"

	FlowRecordFilterOperatorStartsWith FlowRecordFilterOperator = "StartsWith"

	FlowRecordFilterOperatorEndsWith FlowRecordFilterOperator = "EndsWith"

	FlowRecordFilterOperatorContains FlowRecordFilterOperator = "Contains"

	FlowRecordFilterOperatorIsNull FlowRecordFilterOperator = "IsNull"
)

type FlowDataType string

const (
	FlowDataTypeCurrency FlowDataType = "Currency"

	FlowDataTypeDate FlowDataType = "Date"

	FlowDataTypeNumber FlowDataType = "Number"

	FlowDataTypeString FlowDataType = "String"

	FlowDataTypeBoolean FlowDataType = "Boolean"

	FlowDataTypeSObject FlowDataType = "SObject"

	FlowDataTypeDateTime FlowDataType = "DateTime"

	FlowDataTypePicklist FlowDataType = "Picklist"

	FlowDataTypeMultipicklist FlowDataType = "Multipicklist"
)

type FlowScreenFieldType string

const (
	FlowScreenFieldTypeDisplayText FlowScreenFieldType = "DisplayText"

	FlowScreenFieldTypeInputField FlowScreenFieldType = "InputField"

	FlowScreenFieldTypeLargeTextArea FlowScreenFieldType = "LargeTextArea"

	FlowScreenFieldTypePasswordField FlowScreenFieldType = "PasswordField"

	FlowScreenFieldTypeRadioButtons FlowScreenFieldType = "RadioButtons"

	FlowScreenFieldTypeDropdownBox FlowScreenFieldType = "DropdownBox"

	FlowScreenFieldTypeMultiSelectCheckboxes FlowScreenFieldType = "MultiSelectCheckboxes"

	FlowScreenFieldTypeMultiSelectPicklist FlowScreenFieldType = "MultiSelectPicklist"
)

type IterationOrder string

const (
	IterationOrderAsc IterationOrder = "Asc"

	IterationOrderDesc IterationOrder = "Desc"
)

type InvocableActionType string

const (
	InvocableActionTypeApex InvocableActionType = "apex"

	InvocableActionTypeChatterPost InvocableActionType = "chatterPost"

	InvocableActionTypeContentWorkspaceEnableFolders InvocableActionType = "contentWorkspaceEnableFolders"

	InvocableActionTypeEmailAlert InvocableActionType = "emailAlert"

	InvocableActionTypeEmailSimple InvocableActionType = "emailSimple"

	InvocableActionTypeFlow InvocableActionType = "flow"

	InvocableActionTypeMetricRefresh InvocableActionType = "metricRefresh"

	InvocableActionTypeQuickAction InvocableActionType = "quickAction"

	InvocableActionTypeSubmit InvocableActionType = "submit"

	InvocableActionTypeThanks InvocableActionType = "thanks"

	InvocableActionTypeThunderResponse InvocableActionType = "thunderResponse"
)

type FlowProcessType string

const (
	FlowProcessTypeAutoLaunchedFlow FlowProcessType = "AutoLaunchedFlow"

	FlowProcessTypeFlow FlowProcessType = "Flow"

	FlowProcessTypeWorkflow FlowProcessType = "Workflow"

	FlowProcessTypeCustomEvent FlowProcessType = "CustomEvent"

	FlowProcessTypeInvocableProcess FlowProcessType = "InvocableProcess"

	FlowProcessTypeLoginFlow FlowProcessType = "LoginFlow"

	FlowProcessTypeActionPlan FlowProcessType = "ActionPlan"

	FlowProcessTypeJourneyBuilderIntegration FlowProcessType = "JourneyBuilderIntegration"

	FlowProcessTypeUserProvisioningFlow FlowProcessType = "UserProvisioningFlow"
)

type FolderAccessTypes string

const (
	FolderAccessTypesShared FolderAccessTypes = "Shared"

	FolderAccessTypesPublic FolderAccessTypes = "Public"

	FolderAccessTypesHidden FolderAccessTypes = "Hidden"

	FolderAccessTypesPublicInternal FolderAccessTypes = "PublicInternal"
)

type FolderShareAccessLevel string

const (
	FolderShareAccessLevelView FolderShareAccessLevel = "View"

	FolderShareAccessLevelEditAllContents FolderShareAccessLevel = "EditAllContents"

	FolderShareAccessLevelManage FolderShareAccessLevel = "Manage"
)

type FolderSharedToType string

const (
	FolderSharedToTypeGroup FolderSharedToType = "Group"

	FolderSharedToTypeRole FolderSharedToType = "Role"

	FolderSharedToTypeRoleAndSubordinates FolderSharedToType = "RoleAndSubordinates"

	FolderSharedToTypeRoleAndSubordinatesInternal FolderSharedToType = "RoleAndSubordinatesInternal"

	FolderSharedToTypeManager FolderSharedToType = "Manager"

	FolderSharedToTypeManagerAndSubordinatesInternal FolderSharedToType = "ManagerAndSubordinatesInternal"

	FolderSharedToTypeOrganization FolderSharedToType = "Organization"

	FolderSharedToTypeTerritory FolderSharedToType = "Territory"

	FolderSharedToTypeTerritoryAndSubordinates FolderSharedToType = "TerritoryAndSubordinates"

	FolderSharedToTypeAllPrmUsers FolderSharedToType = "AllPrmUsers"

	FolderSharedToTypeUser FolderSharedToType = "User"

	FolderSharedToTypePartnerUser FolderSharedToType = "PartnerUser"

	FolderSharedToTypeAllCspUsers FolderSharedToType = "AllCspUsers"

	FolderSharedToTypeCustomerPortalUser FolderSharedToType = "CustomerPortalUser"

	FolderSharedToTypePortalRole FolderSharedToType = "PortalRole"

	FolderSharedToTypePortalRoleAndSubordinates FolderSharedToType = "PortalRoleAndSubordinates"
)

type PublicFolderAccess string

const (
	PublicFolderAccessReadOnly PublicFolderAccess = "ReadOnly"

	PublicFolderAccessReadWrite PublicFolderAccess = "ReadWrite"
)

type DisplayCurrency string

const (
	DisplayCurrencyCORPORATE DisplayCurrency = "CORPORATE"

	DisplayCurrencyPERSONAL DisplayCurrency = "PERSONAL"
)

type PeriodTypes string

const (
	PeriodTypesMonth PeriodTypes = "Month"

	PeriodTypesQuarter PeriodTypes = "Quarter"

	PeriodTypesWeek PeriodTypes = "Week"

	PeriodTypesYear PeriodTypes = "Year"
)

type PageComponentType string

const (
	PageComponentTypeLinks PageComponentType = "links"

	PageComponentTypeHtmlArea PageComponentType = "htmlArea"

	PageComponentTypeImageOrNote PageComponentType = "imageOrNote"

	PageComponentTypeVisualforcePage PageComponentType = "visualforcePage"
)

type PageComponentWidth string

const (
	PageComponentWidthNarrow PageComponentWidth = "narrow"

	PageComponentWidthWide PageComponentWidth = "wide"
)

type KnowledgeCaseEditor string

const (
	KnowledgeCaseEditorSimple KnowledgeCaseEditor = "simple"

	KnowledgeCaseEditorStandard KnowledgeCaseEditor = "standard"
)

type KnowledgeLanguageLookupValueType string

const (
	KnowledgeLanguageLookupValueTypeUser KnowledgeLanguageLookupValueType = "User"

	KnowledgeLanguageLookupValueTypeQueue KnowledgeLanguageLookupValueType = "Queue"
)

type FeedLayoutFilterPosition string

const (
	FeedLayoutFilterPositionCenterDropDown FeedLayoutFilterPosition = "CenterDropDown"

	FeedLayoutFilterPositionLeftFixed FeedLayoutFilterPosition = "LeftFixed"

	FeedLayoutFilterPositionLeftFloat FeedLayoutFilterPosition = "LeftFloat"
)

type FeedLayoutFilterType string

const (
	FeedLayoutFilterTypeAllUpdates FeedLayoutFilterType = "AllUpdates"

	FeedLayoutFilterTypeFeedItemType FeedLayoutFilterType = "FeedItemType"

	FeedLayoutFilterTypeCustom FeedLayoutFilterType = "Custom"
)

type FeedLayoutComponentType string

const (
	FeedLayoutComponentTypeHelpAndToolLinks FeedLayoutComponentType = "HelpAndToolLinks"

	FeedLayoutComponentTypeCustomButtons FeedLayoutComponentType = "CustomButtons"

	FeedLayoutComponentTypeFollowing FeedLayoutComponentType = "Following"

	FeedLayoutComponentTypeFollowers FeedLayoutComponentType = "Followers"

	FeedLayoutComponentTypeCustomLinks FeedLayoutComponentType = "CustomLinks"

	FeedLayoutComponentTypeMilestones FeedLayoutComponentType = "Milestones"

	FeedLayoutComponentTypeTopics FeedLayoutComponentType = "Topics"

	FeedLayoutComponentTypeCaseUnifiedFiles FeedLayoutComponentType = "CaseUnifiedFiles"

	FeedLayoutComponentTypeVisualforce FeedLayoutComponentType = "Visualforce"
)

type LayoutHeader string

const (
	LayoutHeaderPersonalTagging LayoutHeader = "PersonalTagging"

	LayoutHeaderPublicTagging LayoutHeader = "PublicTagging"
)

type UiBehavior string

const (
	UiBehaviorEdit UiBehavior = "Edit"

	UiBehaviorRequired UiBehavior = "Required"

	UiBehaviorReadonly UiBehavior = "Readonly"
)

type ReportChartComponentSize string

const (
	ReportChartComponentSizeSMALL ReportChartComponentSize = "SMALL"

	ReportChartComponentSizeMEDIUM ReportChartComponentSize = "MEDIUM"

	ReportChartComponentSizeLARGE ReportChartComponentSize = "LARGE"
)

type LayoutSectionStyle string

const (
	LayoutSectionStyleTwoColumnsTopToBottom LayoutSectionStyle = "TwoColumnsTopToBottom"

	LayoutSectionStyleTwoColumnsLeftToRight LayoutSectionStyle = "TwoColumnsLeftToRight"

	LayoutSectionStyleOneColumn LayoutSectionStyle = "OneColumn"

	LayoutSectionStyleCustomLinks LayoutSectionStyle = "CustomLinks"
)

type SummaryLayoutStyle string

const (
	SummaryLayoutStyleDefault SummaryLayoutStyle = "Default"

	SummaryLayoutStyleQuoteTemplate SummaryLayoutStyle = "QuoteTemplate"

	SummaryLayoutStyleDefaultQuoteTemplate SummaryLayoutStyle = "DefaultQuoteTemplate"

	SummaryLayoutStyleCaseInteraction SummaryLayoutStyle = "CaseInteraction"

	SummaryLayoutStyleQuickActionLayoutLeftRight SummaryLayoutStyle = "QuickActionLayoutLeftRight"

	SummaryLayoutStyleQuickActionLayoutTopDown SummaryLayoutStyle = "QuickActionLayoutTopDown"

	SummaryLayoutStylePathAssistant SummaryLayoutStyle = "PathAssistant"
)

type VisibleOrRequired string

const (
	VisibleOrRequiredVisibleOptional VisibleOrRequired = "VisibleOptional"

	VisibleOrRequiredVisibleRequired VisibleOrRequired = "VisibleRequired"

	VisibleOrRequiredNotVisible VisibleOrRequired = "NotVisible"
)

type LetterheadHorizontalAlignment string

const (
	LetterheadHorizontalAlignmentNone LetterheadHorizontalAlignment = "None"

	LetterheadHorizontalAlignmentLeft LetterheadHorizontalAlignment = "Left"

	LetterheadHorizontalAlignmentCenter LetterheadHorizontalAlignment = "Center"

	LetterheadHorizontalAlignmentRight LetterheadHorizontalAlignment = "Right"
)

type LetterheadVerticalAlignment string

const (
	LetterheadVerticalAlignmentNone LetterheadVerticalAlignment = "None"

	LetterheadVerticalAlignmentTop LetterheadVerticalAlignment = "Top"

	LetterheadVerticalAlignmentMiddle LetterheadVerticalAlignment = "Middle"

	LetterheadVerticalAlignmentBottom LetterheadVerticalAlignment = "Bottom"
)

type SupervisorAgentStatusFilter string

const (
	SupervisorAgentStatusFilterOnline SupervisorAgentStatusFilter = "Online"

	SupervisorAgentStatusFilterAway SupervisorAgentStatusFilter = "Away"

	SupervisorAgentStatusFilterOffline SupervisorAgentStatusFilter = "Offline"
)

type LiveChatButtonPresentation string

const (
	LiveChatButtonPresentationSlide LiveChatButtonPresentation = "Slide"

	LiveChatButtonPresentationFade LiveChatButtonPresentation = "Fade"

	LiveChatButtonPresentationAppear LiveChatButtonPresentation = "Appear"

	LiveChatButtonPresentationCustom LiveChatButtonPresentation = "Custom"
)

type LiveChatButtonInviteEndPosition string

const (
	LiveChatButtonInviteEndPositionTopLeft LiveChatButtonInviteEndPosition = "TopLeft"

	LiveChatButtonInviteEndPositionTop LiveChatButtonInviteEndPosition = "Top"

	LiveChatButtonInviteEndPositionTopRight LiveChatButtonInviteEndPosition = "TopRight"

	LiveChatButtonInviteEndPositionLeft LiveChatButtonInviteEndPosition = "Left"

	LiveChatButtonInviteEndPositionCenter LiveChatButtonInviteEndPosition = "Center"

	LiveChatButtonInviteEndPositionRight LiveChatButtonInviteEndPosition = "Right"

	LiveChatButtonInviteEndPositionBottomLeft LiveChatButtonInviteEndPosition = "BottomLeft"

	LiveChatButtonInviteEndPositionBottom LiveChatButtonInviteEndPosition = "Bottom"

	LiveChatButtonInviteEndPositionBottomRight LiveChatButtonInviteEndPosition = "BottomRight"
)

type LiveChatButtonInviteStartPosition string

const (
	LiveChatButtonInviteStartPositionTopLeft LiveChatButtonInviteStartPosition = "TopLeft"

	LiveChatButtonInviteStartPositionTopLeftTop LiveChatButtonInviteStartPosition = "TopLeftTop"

	LiveChatButtonInviteStartPositionTop LiveChatButtonInviteStartPosition = "Top"

	LiveChatButtonInviteStartPositionTopRightTop LiveChatButtonInviteStartPosition = "TopRightTop"

	LiveChatButtonInviteStartPositionTopRight LiveChatButtonInviteStartPosition = "TopRight"

	LiveChatButtonInviteStartPositionTopRightRight LiveChatButtonInviteStartPosition = "TopRightRight"

	LiveChatButtonInviteStartPositionRight LiveChatButtonInviteStartPosition = "Right"

	LiveChatButtonInviteStartPositionBottomRightRight LiveChatButtonInviteStartPosition = "BottomRightRight"

	LiveChatButtonInviteStartPositionBottomRight LiveChatButtonInviteStartPosition = "BottomRight"

	LiveChatButtonInviteStartPositionBottomRightBottom LiveChatButtonInviteStartPosition = "BottomRightBottom"

	LiveChatButtonInviteStartPositionBottom LiveChatButtonInviteStartPosition = "Bottom"

	LiveChatButtonInviteStartPositionBottomLeftBottom LiveChatButtonInviteStartPosition = "BottomLeftBottom"

	LiveChatButtonInviteStartPositionBottomLeft LiveChatButtonInviteStartPosition = "BottomLeft"

	LiveChatButtonInviteStartPositionBottomLeftLeft LiveChatButtonInviteStartPosition = "BottomLeftLeft"

	LiveChatButtonInviteStartPositionLeft LiveChatButtonInviteStartPosition = "Left"

	LiveChatButtonInviteStartPositionTopLeftLeft LiveChatButtonInviteStartPosition = "TopLeftLeft"
)

type LiveChatButtonRoutingType string

const (
	LiveChatButtonRoutingTypeChoice LiveChatButtonRoutingType = "Choice"

	LiveChatButtonRoutingTypeLeastActive LiveChatButtonRoutingType = "LeastActive"

	LiveChatButtonRoutingTypeMostAvailable LiveChatButtonRoutingType = "MostAvailable"
)

type LiveChatButtonType string

const (
	LiveChatButtonTypeStandard LiveChatButtonType = "Standard"

	LiveChatButtonTypeInvite LiveChatButtonType = "Invite"
)

type SensitiveDataActionType string

const (
	SensitiveDataActionTypeRemove SensitiveDataActionType = "Remove"

	SensitiveDataActionTypeReplace SensitiveDataActionType = "Replace"
)

type BlankValueBehavior string

const (
	BlankValueBehaviorMatchBlanks BlankValueBehavior = "MatchBlanks"

	BlankValueBehaviorNullNotAllowed BlankValueBehavior = "NullNotAllowed"
)

type MatchingMethod string

const (
	MatchingMethodExact MatchingMethod = "Exact"

	MatchingMethodFirstName MatchingMethod = "FirstName"

	MatchingMethodLastName MatchingMethod = "LastName"

	MatchingMethodCompanyName MatchingMethod = "CompanyName"

	MatchingMethodPhone MatchingMethod = "Phone"

	MatchingMethodCity MatchingMethod = "City"

	MatchingMethodStreet MatchingMethod = "Street"

	MatchingMethodZip MatchingMethod = "Zip"

	MatchingMethodTitle MatchingMethod = "Title"
)

type MatchingRuleStatus string

const (
	MatchingRuleStatusInactive MatchingRuleStatus = "Inactive"

	MatchingRuleStatusDeactivationFailed MatchingRuleStatus = "DeactivationFailed"

	MatchingRuleStatusActivating MatchingRuleStatus = "Activating"

	MatchingRuleStatusDeactivating MatchingRuleStatus = "Deactivating"

	MatchingRuleStatusActive MatchingRuleStatus = "Active"

	MatchingRuleStatusActivationFailed MatchingRuleStatus = "ActivationFailed"
)

type ApexCodeUnitStatus string

const (
	ApexCodeUnitStatusInactive ApexCodeUnitStatus = "Inactive"

	ApexCodeUnitStatusActive ApexCodeUnitStatus = "Active"

	ApexCodeUnitStatusDeleted ApexCodeUnitStatus = "Deleted"
)

type ContentAssetFormat string

const (
	ContentAssetFormatOriginal ContentAssetFormat = "Original"

	ContentAssetFormatZippedVersions ContentAssetFormat = "ZippedVersions"
)

type ContentAssetAccess string

const (
	ContentAssetAccessVIEWER ContentAssetAccess = "VIEWER"

	ContentAssetAccessCOLLABORATOR ContentAssetAccess = "COLLABORATOR"

	ContentAssetAccessINFERRED ContentAssetAccess = "INFERRED"
)

type DataPipelineType string

const (
	DataPipelineTypePig DataPipelineType = "Pig"
)

type EmailTemplateStyle string

const (
	EmailTemplateStyleNone EmailTemplateStyle = "none"

	EmailTemplateStyleFreeForm EmailTemplateStyle = "freeForm"

	EmailTemplateStyleFormalLetter EmailTemplateStyle = "formalLetter"

	EmailTemplateStylePromotionRight EmailTemplateStyle = "promotionRight"

	EmailTemplateStylePromotionLeft EmailTemplateStyle = "promotionLeft"

	EmailTemplateStyleNewsletter EmailTemplateStyle = "newsletter"

	EmailTemplateStyleProducts EmailTemplateStyle = "products"
)

type EmailTemplateType string

const (
	EmailTemplateTypeText EmailTemplateType = "text"

	EmailTemplateTypeHtml EmailTemplateType = "html"

	EmailTemplateTypeCustom EmailTemplateType = "custom"

	EmailTemplateTypeVisualforce EmailTemplateType = "visualforce"
)

type SControlContentSource string

const (
	SControlContentSourceHTML SControlContentSource = "HTML"

	SControlContentSourceURL SControlContentSource = "URL"

	SControlContentSourceSnippet SControlContentSource = "Snippet"
)

type StaticResourceCacheControl string

const (
	StaticResourceCacheControlPrivate StaticResourceCacheControl = "Private"

	StaticResourceCacheControlPublic StaticResourceCacheControl = "Public"
)

type MilestoneTypeRecurrenceType string

const (
	MilestoneTypeRecurrenceTypeNone MilestoneTypeRecurrenceType = "none"

	MilestoneTypeRecurrenceTypeRecursIndependently MilestoneTypeRecurrenceType = "recursIndependently"

	MilestoneTypeRecurrenceTypeRecursChained MilestoneTypeRecurrenceType = "recursChained"
)

type ModerationRuleAction string

const (
	ModerationRuleActionBlock ModerationRuleAction = "Block"

	ModerationRuleActionFreezeAndNotify ModerationRuleAction = "FreezeAndNotify"

	ModerationRuleActionReview ModerationRuleAction = "Review"

	ModerationRuleActionReplace ModerationRuleAction = "Replace"

	ModerationRuleActionFlag ModerationRuleAction = "Flag"
)

type NetworkStatus string

const (
	NetworkStatusUnderConstruction NetworkStatus = "UnderConstruction"

	NetworkStatusLive NetworkStatus = "Live"

	NetworkStatusDownForMaintenance NetworkStatus = "DownForMaintenance"
)

type APIAccessLevel string

const (
	APIAccessLevelUnrestricted APIAccessLevel = "Unrestricted"

	APIAccessLevelRestricted APIAccessLevel = "Restricted"
)

type PermissionSetTabVisibility string

const (
	PermissionSetTabVisibilityNone PermissionSetTabVisibility = "None"

	PermissionSetTabVisibilityAvailable PermissionSetTabVisibility = "Available"

	PermissionSetTabVisibilityVisible PermissionSetTabVisibility = "Visible"
)

type PlatformCacheType string

const (
	PlatformCacheTypeSession PlatformCacheType = "Session"

	PlatformCacheTypeOrganization PlatformCacheType = "Organization"
)

type PortalRoles string

const (
	PortalRolesExecutive PortalRoles = "Executive"

	PortalRolesManager PortalRoles = "Manager"

	PortalRolesWorker PortalRoles = "Worker"

	PortalRolesPersonAccount PortalRoles = "PersonAccount"
)

type PortalType string

const (
	PortalTypeCustomerSuccess PortalType = "CustomerSuccess"

	PortalTypePartner PortalType = "Partner"

	PortalTypeNetwork PortalType = "Network"
)

type TabVisibility string

const (
	TabVisibilityHidden TabVisibility = "Hidden"

	TabVisibilityDefaultOff TabVisibility = "DefaultOff"

	TabVisibilityDefaultOn TabVisibility = "DefaultOn"
)

type QuickActionLabel string

const (
	QuickActionLabelLogACall QuickActionLabel = "LogACall"

	QuickActionLabelLogANote QuickActionLabel = "LogANote"

	QuickActionLabelNew QuickActionLabel = "New"

	QuickActionLabelNewRecordType QuickActionLabel = "NewRecordType"

	QuickActionLabelUpdate QuickActionLabel = "Update"

	QuickActionLabelNewChild QuickActionLabel = "NewChild"

	QuickActionLabelNewChildRecordType QuickActionLabel = "NewChildRecordType"

	QuickActionLabelCreateNew QuickActionLabel = "CreateNew"

	QuickActionLabelCreateNewRecordType QuickActionLabel = "CreateNewRecordType"

	QuickActionLabelSendEmail QuickActionLabel = "SendEmail"

	QuickActionLabelQuickRecordType QuickActionLabel = "QuickRecordType"

	QuickActionLabelQuick QuickActionLabel = "Quick"

	QuickActionLabelEditDescription QuickActionLabel = "EditDescription"

	QuickActionLabelDefer QuickActionLabel = "Defer"

	QuickActionLabelChangeDueDate QuickActionLabel = "ChangeDueDate"

	QuickActionLabelChangePriority QuickActionLabel = "ChangePriority"

	QuickActionLabelChangeStatus QuickActionLabel = "ChangeStatus"

	QuickActionLabelSocialPost QuickActionLabel = "SocialPost"

	QuickActionLabelEscalate QuickActionLabel = "Escalate"

	QuickActionLabelEscalateToRecord QuickActionLabel = "EscalateToRecord"

	QuickActionLabelOfferFeedback QuickActionLabel = "OfferFeedback"

	QuickActionLabelRequestFeedback QuickActionLabel = "RequestFeedback"

	QuickActionLabelAddRecord QuickActionLabel = "AddRecord"

	QuickActionLabelAddMember QuickActionLabel = "AddMember"
)

type QuickActionType string

const (
	QuickActionTypeCreate QuickActionType = "Create"

	QuickActionTypeVisualforcePage QuickActionType = "VisualforcePage"

	QuickActionTypePost QuickActionType = "Post"

	QuickActionTypeSendEmail QuickActionType = "SendEmail"

	QuickActionTypeLogACall QuickActionType = "LogACall"

	QuickActionTypeSocialPost QuickActionType = "SocialPost"

	QuickActionTypeCanvas QuickActionType = "Canvas"

	QuickActionTypeUpdate QuickActionType = "Update"

	QuickActionTypeLightningComponent QuickActionType = "LightningComponent"
)

type ReportAggregateDatatype string

const (
	ReportAggregateDatatypeCurrency ReportAggregateDatatype = "currency"

	ReportAggregateDatatypePercent ReportAggregateDatatype = "percent"

	ReportAggregateDatatypeNumber ReportAggregateDatatype = "number"
)

type ReportBucketFieldType string

const (
	ReportBucketFieldTypeText ReportBucketFieldType = "text"

	ReportBucketFieldTypeNumber ReportBucketFieldType = "number"

	ReportBucketFieldTypePicklist ReportBucketFieldType = "picklist"
)

type ReportFormulaNullTreatment string

const (
	ReportFormulaNullTreatmentN ReportFormulaNullTreatment = "n"

	ReportFormulaNullTreatmentZ ReportFormulaNullTreatment = "z"
)

type ChartType string

const (
	ChartTypeNone ChartType = "None"

	ChartTypeScatter ChartType = "Scatter"

	ChartTypeScatterGrouped ChartType = "ScatterGrouped"

	ChartTypeBubble ChartType = "Bubble"

	ChartTypeBubbleGrouped ChartType = "BubbleGrouped"

	ChartTypeHorizontalBar ChartType = "HorizontalBar"

	ChartTypeHorizontalBarGrouped ChartType = "HorizontalBarGrouped"

	ChartTypeHorizontalBarStacked ChartType = "HorizontalBarStacked"

	ChartTypeHorizontalBarStackedTo100 ChartType = "HorizontalBarStackedTo100"

	ChartTypeVerticalColumn ChartType = "VerticalColumn"

	ChartTypeVerticalColumnGrouped ChartType = "VerticalColumnGrouped"

	ChartTypeVerticalColumnStacked ChartType = "VerticalColumnStacked"

	ChartTypeVerticalColumnStackedTo100 ChartType = "VerticalColumnStackedTo100"

	ChartTypeLine ChartType = "Line"

	ChartTypeLineGrouped ChartType = "LineGrouped"

	ChartTypeLineCumulative ChartType = "LineCumulative"

	ChartTypeLineCumulativeGrouped ChartType = "LineCumulativeGrouped"

	ChartTypePie ChartType = "Pie"

	ChartTypeDonut ChartType = "Donut"

	ChartTypeFunnel ChartType = "Funnel"

	ChartTypeVerticalColumnLine ChartType = "VerticalColumnLine"

	ChartTypeVerticalColumnGroupedLine ChartType = "VerticalColumnGroupedLine"

	ChartTypeVerticalColumnStackedLine ChartType = "VerticalColumnStackedLine"

	ChartTypePlugin ChartType = "Plugin"
)

type ChartPosition string

const (
	ChartPositionCHARTTOP ChartPosition = "CHARTTOP"

	ChartPositionCHARTBOTTOM ChartPosition = "CHARTBOTTOM"
)

type ReportChartSize string

const (
	ReportChartSizeTiny ReportChartSize = "Tiny"

	ReportChartSizeSmall ReportChartSize = "Small"

	ReportChartSizeMedium ReportChartSize = "Medium"

	ReportChartSizeLarge ReportChartSize = "Large"

	ReportChartSizeHuge ReportChartSize = "Huge"
)

type ObjectFilterOperator string

const (
	ObjectFilterOperatorWith ObjectFilterOperator = "with"

	ObjectFilterOperatorWithout ObjectFilterOperator = "without"
)

type CurrencyIsoCode string

const (
	CurrencyIsoCodeADP CurrencyIsoCode = "ADP"

	CurrencyIsoCodeAED CurrencyIsoCode = "AED"

	CurrencyIsoCodeAFA CurrencyIsoCode = "AFA"

	CurrencyIsoCodeAFN CurrencyIsoCode = "AFN"

	CurrencyIsoCodeALL CurrencyIsoCode = "ALL"

	CurrencyIsoCodeAMD CurrencyIsoCode = "AMD"

	CurrencyIsoCodeANG CurrencyIsoCode = "ANG"

	CurrencyIsoCodeAOA CurrencyIsoCode = "AOA"

	CurrencyIsoCodeARS CurrencyIsoCode = "ARS"

	CurrencyIsoCodeATS CurrencyIsoCode = "ATS"

	CurrencyIsoCodeAUD CurrencyIsoCode = "AUD"

	CurrencyIsoCodeAWG CurrencyIsoCode = "AWG"

	CurrencyIsoCodeAZM CurrencyIsoCode = "AZM"

	CurrencyIsoCodeAZN CurrencyIsoCode = "AZN"

	CurrencyIsoCodeBAM CurrencyIsoCode = "BAM"

	CurrencyIsoCodeBBD CurrencyIsoCode = "BBD"

	CurrencyIsoCodeBDT CurrencyIsoCode = "BDT"

	CurrencyIsoCodeBEF CurrencyIsoCode = "BEF"

	CurrencyIsoCodeBGL CurrencyIsoCode = "BGL"

	CurrencyIsoCodeBGN CurrencyIsoCode = "BGN"

	CurrencyIsoCodeBHD CurrencyIsoCode = "BHD"

	CurrencyIsoCodeBIF CurrencyIsoCode = "BIF"

	CurrencyIsoCodeBMD CurrencyIsoCode = "BMD"

	CurrencyIsoCodeBND CurrencyIsoCode = "BND"

	CurrencyIsoCodeBOB CurrencyIsoCode = "BOB"

	CurrencyIsoCodeBOV CurrencyIsoCode = "BOV"

	CurrencyIsoCodeBRB CurrencyIsoCode = "BRB"

	CurrencyIsoCodeBRL CurrencyIsoCode = "BRL"

	CurrencyIsoCodeBSD CurrencyIsoCode = "BSD"

	CurrencyIsoCodeBTN CurrencyIsoCode = "BTN"

	CurrencyIsoCodeBWP CurrencyIsoCode = "BWP"

	CurrencyIsoCodeBYB CurrencyIsoCode = "BYB"

	CurrencyIsoCodeBYR CurrencyIsoCode = "BYR"

	CurrencyIsoCodeBZD CurrencyIsoCode = "BZD"

	CurrencyIsoCodeCAD CurrencyIsoCode = "CAD"

	CurrencyIsoCodeCDF CurrencyIsoCode = "CDF"

	CurrencyIsoCodeCHF CurrencyIsoCode = "CHF"

	CurrencyIsoCodeCLF CurrencyIsoCode = "CLF"

	CurrencyIsoCodeCLP CurrencyIsoCode = "CLP"

	CurrencyIsoCodeCNY CurrencyIsoCode = "CNY"

	CurrencyIsoCodeCOP CurrencyIsoCode = "COP"

	CurrencyIsoCodeCRC CurrencyIsoCode = "CRC"

	CurrencyIsoCodeCSD CurrencyIsoCode = "CSD"

	CurrencyIsoCodeCUC CurrencyIsoCode = "CUC"

	CurrencyIsoCodeCUP CurrencyIsoCode = "CUP"

	CurrencyIsoCodeCVE CurrencyIsoCode = "CVE"

	CurrencyIsoCodeCYP CurrencyIsoCode = "CYP"

	CurrencyIsoCodeCZK CurrencyIsoCode = "CZK"

	CurrencyIsoCodeDEM CurrencyIsoCode = "DEM"

	CurrencyIsoCodeDJF CurrencyIsoCode = "DJF"

	CurrencyIsoCodeDKK CurrencyIsoCode = "DKK"

	CurrencyIsoCodeDOP CurrencyIsoCode = "DOP"

	CurrencyIsoCodeDZD CurrencyIsoCode = "DZD"

	CurrencyIsoCodeECS CurrencyIsoCode = "ECS"

	CurrencyIsoCodeEEK CurrencyIsoCode = "EEK"

	CurrencyIsoCodeEGP CurrencyIsoCode = "EGP"

	CurrencyIsoCodeERN CurrencyIsoCode = "ERN"

	CurrencyIsoCodeESP CurrencyIsoCode = "ESP"

	CurrencyIsoCodeETB CurrencyIsoCode = "ETB"

	CurrencyIsoCodeEUR CurrencyIsoCode = "EUR"

	CurrencyIsoCodeFIM CurrencyIsoCode = "FIM"

	CurrencyIsoCodeFJD CurrencyIsoCode = "FJD"

	CurrencyIsoCodeFKP CurrencyIsoCode = "FKP"

	CurrencyIsoCodeFRF CurrencyIsoCode = "FRF"

	CurrencyIsoCodeGBP CurrencyIsoCode = "GBP"

	CurrencyIsoCodeGEL CurrencyIsoCode = "GEL"

	CurrencyIsoCodeGHC CurrencyIsoCode = "GHC"

	CurrencyIsoCodeGHS CurrencyIsoCode = "GHS"

	CurrencyIsoCodeGIP CurrencyIsoCode = "GIP"

	CurrencyIsoCodeGMD CurrencyIsoCode = "GMD"

	CurrencyIsoCodeGNF CurrencyIsoCode = "GNF"

	CurrencyIsoCodeGRD CurrencyIsoCode = "GRD"

	CurrencyIsoCodeGTQ CurrencyIsoCode = "GTQ"

	CurrencyIsoCodeGWP CurrencyIsoCode = "GWP"

	CurrencyIsoCodeGYD CurrencyIsoCode = "GYD"

	CurrencyIsoCodeHKD CurrencyIsoCode = "HKD"

	CurrencyIsoCodeHNL CurrencyIsoCode = "HNL"

	CurrencyIsoCodeHRD CurrencyIsoCode = "HRD"

	CurrencyIsoCodeHRK CurrencyIsoCode = "HRK"

	CurrencyIsoCodeHTG CurrencyIsoCode = "HTG"

	CurrencyIsoCodeHUF CurrencyIsoCode = "HUF"

	CurrencyIsoCodeIDR CurrencyIsoCode = "IDR"

	CurrencyIsoCodeIEP CurrencyIsoCode = "IEP"

	CurrencyIsoCodeILS CurrencyIsoCode = "ILS"

	CurrencyIsoCodeINR CurrencyIsoCode = "INR"

	CurrencyIsoCodeIQD CurrencyIsoCode = "IQD"

	CurrencyIsoCodeIRR CurrencyIsoCode = "IRR"

	CurrencyIsoCodeISK CurrencyIsoCode = "ISK"

	CurrencyIsoCodeITL CurrencyIsoCode = "ITL"

	CurrencyIsoCodeJMD CurrencyIsoCode = "JMD"

	CurrencyIsoCodeJOD CurrencyIsoCode = "JOD"

	CurrencyIsoCodeJPY CurrencyIsoCode = "JPY"

	CurrencyIsoCodeKES CurrencyIsoCode = "KES"

	CurrencyIsoCodeKGS CurrencyIsoCode = "KGS"

	CurrencyIsoCodeKHR CurrencyIsoCode = "KHR"

	CurrencyIsoCodeKMF CurrencyIsoCode = "KMF"

	CurrencyIsoCodeKPW CurrencyIsoCode = "KPW"

	CurrencyIsoCodeKRW CurrencyIsoCode = "KRW"

	CurrencyIsoCodeKWD CurrencyIsoCode = "KWD"

	CurrencyIsoCodeKYD CurrencyIsoCode = "KYD"

	CurrencyIsoCodeKZT CurrencyIsoCode = "KZT"

	CurrencyIsoCodeLAK CurrencyIsoCode = "LAK"

	CurrencyIsoCodeLBP CurrencyIsoCode = "LBP"

	CurrencyIsoCodeLKR CurrencyIsoCode = "LKR"

	CurrencyIsoCodeLRD CurrencyIsoCode = "LRD"

	CurrencyIsoCodeLSL CurrencyIsoCode = "LSL"

	CurrencyIsoCodeLTL CurrencyIsoCode = "LTL"

	CurrencyIsoCodeLUF CurrencyIsoCode = "LUF"

	CurrencyIsoCodeLVL CurrencyIsoCode = "LVL"

	CurrencyIsoCodeLYD CurrencyIsoCode = "LYD"

	CurrencyIsoCodeMAD CurrencyIsoCode = "MAD"

	CurrencyIsoCodeMDL CurrencyIsoCode = "MDL"

	CurrencyIsoCodeMGA CurrencyIsoCode = "MGA"

	CurrencyIsoCodeMGF CurrencyIsoCode = "MGF"

	CurrencyIsoCodeMKD CurrencyIsoCode = "MKD"

	CurrencyIsoCodeMMK CurrencyIsoCode = "MMK"

	CurrencyIsoCodeMNT CurrencyIsoCode = "MNT"

	CurrencyIsoCodeMOP CurrencyIsoCode = "MOP"

	CurrencyIsoCodeMRO CurrencyIsoCode = "MRO"

	CurrencyIsoCodeMTL CurrencyIsoCode = "MTL"

	CurrencyIsoCodeMUR CurrencyIsoCode = "MUR"

	CurrencyIsoCodeMVR CurrencyIsoCode = "MVR"

	CurrencyIsoCodeMWK CurrencyIsoCode = "MWK"

	CurrencyIsoCodeMXN CurrencyIsoCode = "MXN"

	CurrencyIsoCodeMXV CurrencyIsoCode = "MXV"

	CurrencyIsoCodeMYR CurrencyIsoCode = "MYR"

	CurrencyIsoCodeMZM CurrencyIsoCode = "MZM"

	CurrencyIsoCodeMZN CurrencyIsoCode = "MZN"

	CurrencyIsoCodeNAD CurrencyIsoCode = "NAD"

	CurrencyIsoCodeNGN CurrencyIsoCode = "NGN"

	CurrencyIsoCodeNIO CurrencyIsoCode = "NIO"

	CurrencyIsoCodeNLG CurrencyIsoCode = "NLG"

	CurrencyIsoCodeNOK CurrencyIsoCode = "NOK"

	CurrencyIsoCodeNPR CurrencyIsoCode = "NPR"

	CurrencyIsoCodeNZD CurrencyIsoCode = "NZD"

	CurrencyIsoCodeOMR CurrencyIsoCode = "OMR"

	CurrencyIsoCodePAB CurrencyIsoCode = "PAB"

	CurrencyIsoCodePEN CurrencyIsoCode = "PEN"

	CurrencyIsoCodePGK CurrencyIsoCode = "PGK"

	CurrencyIsoCodePHP CurrencyIsoCode = "PHP"

	CurrencyIsoCodePKR CurrencyIsoCode = "PKR"

	CurrencyIsoCodePLN CurrencyIsoCode = "PLN"

	CurrencyIsoCodePTE CurrencyIsoCode = "PTE"

	CurrencyIsoCodePYG CurrencyIsoCode = "PYG"

	CurrencyIsoCodeQAR CurrencyIsoCode = "QAR"

	CurrencyIsoCodeRMB CurrencyIsoCode = "RMB"

	CurrencyIsoCodeROL CurrencyIsoCode = "ROL"

	CurrencyIsoCodeRON CurrencyIsoCode = "RON"

	CurrencyIsoCodeRSD CurrencyIsoCode = "RSD"

	CurrencyIsoCodeRUB CurrencyIsoCode = "RUB"

	CurrencyIsoCodeRUR CurrencyIsoCode = "RUR"

	CurrencyIsoCodeRWF CurrencyIsoCode = "RWF"

	CurrencyIsoCodeSAR CurrencyIsoCode = "SAR"

	CurrencyIsoCodeSBD CurrencyIsoCode = "SBD"

	CurrencyIsoCodeSCR CurrencyIsoCode = "SCR"

	CurrencyIsoCodeSDD CurrencyIsoCode = "SDD"

	CurrencyIsoCodeSDG CurrencyIsoCode = "SDG"

	CurrencyIsoCodeSEK CurrencyIsoCode = "SEK"

	CurrencyIsoCodeSGD CurrencyIsoCode = "SGD"

	CurrencyIsoCodeSHP CurrencyIsoCode = "SHP"

	CurrencyIsoCodeSIT CurrencyIsoCode = "SIT"

	CurrencyIsoCodeSKK CurrencyIsoCode = "SKK"

	CurrencyIsoCodeSLL CurrencyIsoCode = "SLL"

	CurrencyIsoCodeSOS CurrencyIsoCode = "SOS"

	CurrencyIsoCodeSRD CurrencyIsoCode = "SRD"

	CurrencyIsoCodeSRG CurrencyIsoCode = "SRG"

	CurrencyIsoCodeSSP CurrencyIsoCode = "SSP"

	CurrencyIsoCodeSTD CurrencyIsoCode = "STD"

	CurrencyIsoCodeSUR CurrencyIsoCode = "SUR"

	CurrencyIsoCodeSVC CurrencyIsoCode = "SVC"

	CurrencyIsoCodeSYP CurrencyIsoCode = "SYP"

	CurrencyIsoCodeSZL CurrencyIsoCode = "SZL"

	CurrencyIsoCodeTHB CurrencyIsoCode = "THB"

	CurrencyIsoCodeTJR CurrencyIsoCode = "TJR"

	CurrencyIsoCodeTJS CurrencyIsoCode = "TJS"

	CurrencyIsoCodeTMM CurrencyIsoCode = "TMM"

	CurrencyIsoCodeTMT CurrencyIsoCode = "TMT"

	CurrencyIsoCodeTND CurrencyIsoCode = "TND"

	CurrencyIsoCodeTOP CurrencyIsoCode = "TOP"

	CurrencyIsoCodeTPE CurrencyIsoCode = "TPE"

	CurrencyIsoCodeTRL CurrencyIsoCode = "TRL"

	CurrencyIsoCodeTRY CurrencyIsoCode = "TRY"

	CurrencyIsoCodeTTD CurrencyIsoCode = "TTD"

	CurrencyIsoCodeTWD CurrencyIsoCode = "TWD"

	CurrencyIsoCodeTZS CurrencyIsoCode = "TZS"

	CurrencyIsoCodeUAH CurrencyIsoCode = "UAH"

	CurrencyIsoCodeUGX CurrencyIsoCode = "UGX"

	CurrencyIsoCodeUSD CurrencyIsoCode = "USD"

	CurrencyIsoCodeUYU CurrencyIsoCode = "UYU"

	CurrencyIsoCodeUZS CurrencyIsoCode = "UZS"

	CurrencyIsoCodeVEB CurrencyIsoCode = "VEB"

	CurrencyIsoCodeVEF CurrencyIsoCode = "VEF"

	CurrencyIsoCodeVND CurrencyIsoCode = "VND"

	CurrencyIsoCodeVUV CurrencyIsoCode = "VUV"

	CurrencyIsoCodeWST CurrencyIsoCode = "WST"

	CurrencyIsoCodeXAF CurrencyIsoCode = "XAF"

	CurrencyIsoCodeXCD CurrencyIsoCode = "XCD"

	CurrencyIsoCodeXOF CurrencyIsoCode = "XOF"

	CurrencyIsoCodeXPF CurrencyIsoCode = "XPF"

	CurrencyIsoCodeYER CurrencyIsoCode = "YER"

	CurrencyIsoCodeYUM CurrencyIsoCode = "YUM"

	CurrencyIsoCodeZAR CurrencyIsoCode = "ZAR"

	CurrencyIsoCodeZMK CurrencyIsoCode = "ZMK"

	CurrencyIsoCodeZMW CurrencyIsoCode = "ZMW"

	CurrencyIsoCodeZWD CurrencyIsoCode = "ZWD"

	CurrencyIsoCodeZWL CurrencyIsoCode = "ZWL"
)

type DataCategoryFilterOperation string

const (
	DataCategoryFilterOperationAbove DataCategoryFilterOperation = "above"

	DataCategoryFilterOperationBelow DataCategoryFilterOperation = "below"

	DataCategoryFilterOperationAt DataCategoryFilterOperation = "at"

	DataCategoryFilterOperationAboveOrBelow DataCategoryFilterOperation = "aboveOrBelow"
)

type ReportFormat string

const (
	ReportFormatMultiBlock ReportFormat = "MultiBlock"

	ReportFormatMatrix ReportFormat = "Matrix"

	ReportFormatSummary ReportFormat = "Summary"

	ReportFormatTabular ReportFormat = "Tabular"
)

type ReportAggrType string

const (
	ReportAggrTypeSum ReportAggrType = "Sum"

	ReportAggrTypeAverage ReportAggrType = "Average"

	ReportAggrTypeMaximum ReportAggrType = "Maximum"

	ReportAggrTypeMinimum ReportAggrType = "Minimum"

	ReportAggrTypeRowCount ReportAggrType = "RowCount"
)

type UserDateGranularity string

const (
	UserDateGranularityNone UserDateGranularity = "None"

	UserDateGranularityDay UserDateGranularity = "Day"

	UserDateGranularityWeek UserDateGranularity = "Week"

	UserDateGranularityMonth UserDateGranularity = "Month"

	UserDateGranularityQuarter UserDateGranularity = "Quarter"

	UserDateGranularityYear UserDateGranularity = "Year"

	UserDateGranularityFiscalQuarter UserDateGranularity = "FiscalQuarter"

	UserDateGranularityFiscalYear UserDateGranularity = "FiscalYear"

	UserDateGranularityMonthInYear UserDateGranularity = "MonthInYear"

	UserDateGranularityDayInMonth UserDateGranularity = "DayInMonth"

	UserDateGranularityFiscalPeriod UserDateGranularity = "FiscalPeriod"

	UserDateGranularityFiscalWeek UserDateGranularity = "FiscalWeek"
)

type ReportSortType string

const (
	ReportSortTypeColumn ReportSortType = "Column"

	ReportSortTypeAggregate ReportSortType = "Aggregate"

	ReportSortTypeCustomSummaryFormula ReportSortType = "CustomSummaryFormula"
)

type UserDateInterval string

const (
	UserDateIntervalINTERVALCURRENT UserDateInterval = "INTERVALCURRENT"

	UserDateIntervalINTERVALCURNEXT1 UserDateInterval = "INTERVALCURNEXT1"

	UserDateIntervalINTERVALCURPREV1 UserDateInterval = "INTERVALCURPREV1"

	UserDateIntervalINTERVALNEXT1 UserDateInterval = "INTERVALNEXT1"

	UserDateIntervalINTERVALPREV1 UserDateInterval = "INTERVALPREV1"

	UserDateIntervalINTERVALCURNEXT3 UserDateInterval = "INTERVALCURNEXT3"

	UserDateIntervalINTERVALCURFY UserDateInterval = "INTERVALCURFY"

	UserDateIntervalINTERVALPREVFY UserDateInterval = "INTERVALPREVFY"

	UserDateIntervalINTERVALPREV2FY UserDateInterval = "INTERVALPREV2FY"

	UserDateIntervalINTERVALAGO2FY UserDateInterval = "INTERVALAGO2FY"

	UserDateIntervalINTERVALNEXTFY UserDateInterval = "INTERVALNEXTFY"

	UserDateIntervalINTERVALPREVCURFY UserDateInterval = "INTERVALPREVCURFY"

	UserDateIntervalINTERVALPREVCUR2FY UserDateInterval = "INTERVALPREVCUR2FY"

	UserDateIntervalINTERVALCURNEXTFY UserDateInterval = "INTERVALCURNEXTFY"

	UserDateIntervalINTERVALCUSTOM UserDateInterval = "INTERVALCUSTOM"

	UserDateIntervalINTERVALYESTERDAY UserDateInterval = "INTERVALYESTERDAY"

	UserDateIntervalINTERVALTODAY UserDateInterval = "INTERVALTODAY"

	UserDateIntervalINTERVALTOMORROW UserDateInterval = "INTERVALTOMORROW"

	UserDateIntervalINTERVALLASTWEEK UserDateInterval = "INTERVALLASTWEEK"

	UserDateIntervalINTERVALTHISWEEK UserDateInterval = "INTERVALTHISWEEK"

	UserDateIntervalINTERVALNEXTWEEK UserDateInterval = "INTERVALNEXTWEEK"

	UserDateIntervalINTERVALLASTMONTH UserDateInterval = "INTERVALLASTMONTH"

	UserDateIntervalINTERVALTHISMONTH UserDateInterval = "INTERVALTHISMONTH"

	UserDateIntervalINTERVALNEXTMONTH UserDateInterval = "INTERVALNEXTMONTH"

	UserDateIntervalINTERVALLASTTHISMONTH UserDateInterval = "INTERVALLASTTHISMONTH"

	UserDateIntervalINTERVALTHISNEXTMONTH UserDateInterval = "INTERVALTHISNEXTMONTH"

	UserDateIntervalINTERVALCURRENTQ UserDateInterval = "INTERVALCURRENTQ"

	UserDateIntervalINTERVALCURNEXTQ UserDateInterval = "INTERVALCURNEXTQ"

	UserDateIntervalINTERVALCURPREVQ UserDateInterval = "INTERVALCURPREVQ"

	UserDateIntervalINTERVALNEXTQ UserDateInterval = "INTERVALNEXTQ"

	UserDateIntervalINTERVALPREVQ UserDateInterval = "INTERVALPREVQ"

	UserDateIntervalINTERVALCURNEXT3Q UserDateInterval = "INTERVALCURNEXT3Q"

	UserDateIntervalINTERVALCURY UserDateInterval = "INTERVALCURY"

	UserDateIntervalINTERVALPREVY UserDateInterval = "INTERVALPREVY"

	UserDateIntervalINTERVALPREV2Y UserDateInterval = "INTERVALPREV2Y"

	UserDateIntervalINTERVALAGO2Y UserDateInterval = "INTERVALAGO2Y"

	UserDateIntervalINTERVALNEXTY UserDateInterval = "INTERVALNEXTY"

	UserDateIntervalINTERVALPREVCURY UserDateInterval = "INTERVALPREVCURY"

	UserDateIntervalINTERVALPREVCUR2Y UserDateInterval = "INTERVALPREVCUR2Y"

	UserDateIntervalINTERVALCURNEXTY UserDateInterval = "INTERVALCURNEXTY"

	UserDateIntervalINTERVALLAST7 UserDateInterval = "INTERVALLAST7"

	UserDateIntervalINTERVALLAST30 UserDateInterval = "INTERVALLAST30"

	UserDateIntervalINTERVALLAST60 UserDateInterval = "INTERVALLAST60"

	UserDateIntervalINTERVALLAST90 UserDateInterval = "INTERVALLAST90"

	UserDateIntervalINTERVALLAST120 UserDateInterval = "INTERVALLAST120"

	UserDateIntervalINTERVALNEXT7 UserDateInterval = "INTERVALNEXT7"

	UserDateIntervalINTERVALNEXT30 UserDateInterval = "INTERVALNEXT30"

	UserDateIntervalINTERVALNEXT60 UserDateInterval = "INTERVALNEXT60"

	UserDateIntervalINTERVALNEXT90 UserDateInterval = "INTERVALNEXT90"

	UserDateIntervalINTERVALNEXT120 UserDateInterval = "INTERVALNEXT120"

	UserDateIntervalLASTFISCALWEEK UserDateInterval = "LASTFISCALWEEK"

	UserDateIntervalTHISFISCALWEEK UserDateInterval = "THISFISCALWEEK"

	UserDateIntervalNEXTFISCALWEEK UserDateInterval = "NEXTFISCALWEEK"

	UserDateIntervalLASTFISCALPERIOD UserDateInterval = "LASTFISCALPERIOD"

	UserDateIntervalTHISFISCALPERIOD UserDateInterval = "THISFISCALPERIOD"

	UserDateIntervalNEXTFISCALPERIOD UserDateInterval = "NEXTFISCALPERIOD"

	UserDateIntervalLASTTHISFISCALPERIOD UserDateInterval = "LASTTHISFISCALPERIOD"

	UserDateIntervalTHISNEXTFISCALPERIOD UserDateInterval = "THISNEXTFISCALPERIOD"

	UserDateIntervalCURRENTENTITLEMENTPERIOD UserDateInterval = "CURRENTENTITLEMENTPERIOD"

	UserDateIntervalPREVIOUSENTITLEMENTPERIOD UserDateInterval = "PREVIOUSENTITLEMENTPERIOD"

	UserDateIntervalPREVIOUSTWOENTITLEMENTPERIODS UserDateInterval = "PREVIOUSTWOENTITLEMENTPERIODS"

	UserDateIntervalTWOENTITLEMENTPERIODSAGO UserDateInterval = "TWOENTITLEMENTPERIODSAGO"

	UserDateIntervalCURRENTANDPREVIOUSENTITLEMENTPERIOD UserDateInterval = "CURRENTANDPREVIOUSENTITLEMENTPERIOD"

	UserDateIntervalCURRENTANDPREVIOUSTWOENTITLEMENTPERIODS UserDateInterval = "CURRENTANDPREVIOUSTWOENTITLEMENTPERIODS"
)

type ReportTypeCategory string

const (
	ReportTypeCategoryAccounts ReportTypeCategory = "accounts"

	ReportTypeCategoryOpportunities ReportTypeCategory = "opportunities"

	ReportTypeCategoryForecasts ReportTypeCategory = "forecasts"

	ReportTypeCategoryCases ReportTypeCategory = "cases"

	ReportTypeCategoryLeads ReportTypeCategory = "leads"

	ReportTypeCategoryCampaigns ReportTypeCategory = "campaigns"

	ReportTypeCategoryActivities ReportTypeCategory = "activities"

	ReportTypeCategoryBusop ReportTypeCategory = "busop"

	ReportTypeCategoryProducts ReportTypeCategory = "products"

	ReportTypeCategoryAdmin ReportTypeCategory = "admin"

	ReportTypeCategoryTerritory ReportTypeCategory = "territory"

	ReportTypeCategoryOther ReportTypeCategory = "other"

	ReportTypeCategoryContent ReportTypeCategory = "content"

	ReportTypeCategoryUsageentitlement ReportTypeCategory = "usageentitlement"

	ReportTypeCategoryWdc ReportTypeCategory = "wdc"

	ReportTypeCategoryCalibration ReportTypeCategory = "calibration"

	ReportTypeCategoryTerritory2 ReportTypeCategory = "territory2"
)

type SamlIdentityLocationType string

const (
	SamlIdentityLocationTypeSubjectNameId SamlIdentityLocationType = "SubjectNameId"

	SamlIdentityLocationTypeAttribute SamlIdentityLocationType = "Attribute"
)

type SamlIdentityType string

const (
	SamlIdentityTypeUsername SamlIdentityType = "Username"

	SamlIdentityTypeFederationId SamlIdentityType = "FederationId"

	SamlIdentityTypeUserId SamlIdentityType = "UserId"
)

type SamlType string

const (
	SamlTypeSAML11 SamlType = "SAML11"

	SamlTypeSAML20 SamlType = "SAML20"
)

type Complexity string

const (
	ComplexityNoRestriction Complexity = "NoRestriction"

	ComplexityAlphaNumeric Complexity = "AlphaNumeric"

	ComplexitySpecialCharacters Complexity = "SpecialCharacters"

	ComplexityUpperLowerCaseNumeric Complexity = "UpperLowerCaseNumeric"

	ComplexityUpperLowerCaseNumericSpecialCharacters Complexity = "UpperLowerCaseNumericSpecialCharacters"
)

type Expiration string

const (
	ExpirationThirtyDays Expiration = "ThirtyDays"

	ExpirationSixtyDays Expiration = "SixtyDays"

	ExpirationNinetyDays Expiration = "NinetyDays"

	ExpirationSixMonths Expiration = "SixMonths"

	ExpirationOneYear Expiration = "OneYear"

	ExpirationNever Expiration = "Never"
)

type LockoutInterval string

const (
	LockoutIntervalFifteenMinutes LockoutInterval = "FifteenMinutes"

	LockoutIntervalThirtyMinutes LockoutInterval = "ThirtyMinutes"

	LockoutIntervalSixtyMinutes LockoutInterval = "SixtyMinutes"

	LockoutIntervalForever LockoutInterval = "Forever"
)

type MaxLoginAttempts string

const (
	MaxLoginAttemptsThreeAttempts MaxLoginAttempts = "ThreeAttempts"

	MaxLoginAttemptsFiveAttempts MaxLoginAttempts = "FiveAttempts"

	MaxLoginAttemptsTenAttempts MaxLoginAttempts = "TenAttempts"

	MaxLoginAttemptsNoLimit MaxLoginAttempts = "NoLimit"
)

type QuestionRestriction string

const (
	QuestionRestrictionNone QuestionRestriction = "None"

	QuestionRestrictionDoesNotContainPassword QuestionRestriction = "DoesNotContainPassword"
)

type SessionTimeout string

const (
	SessionTimeoutTwentyFourHours SessionTimeout = "TwentyFourHours"

	SessionTimeoutTwelveHours SessionTimeout = "TwelveHours"

	SessionTimeoutEightHours SessionTimeout = "EightHours"

	SessionTimeoutFourHours SessionTimeout = "FourHours"

	SessionTimeoutTwoHours SessionTimeout = "TwoHours"

	SessionTimeoutSixtyMinutes SessionTimeout = "SixtyMinutes"

	SessionTimeoutThirtyMinutes SessionTimeout = "ThirtyMinutes"

	SessionTimeoutFifteenMinutes SessionTimeout = "FifteenMinutes"
)

type MonitoredEvents string

const (
	MonitoredEventsAuditTrail MonitoredEvents = "AuditTrail"

	MonitoredEventsLogin MonitoredEvents = "Login"

	MonitoredEventsEntity MonitoredEvents = "Entity"

	MonitoredEventsDataExport MonitoredEvents = "DataExport"

	MonitoredEventsAccessResource MonitoredEvents = "AccessResource"
)

type VisualizationResourceType string

const (
	VisualizationResourceTypeJs VisualizationResourceType = "js"

	VisualizationResourceTypeCss VisualizationResourceType = "css"
)

type LookupValueType string

const (
	LookupValueTypeUser LookupValueType = "User"

	LookupValueTypeQueue LookupValueType = "Queue"

	LookupValueTypeRecordType LookupValueType = "RecordType"
)

type FieldUpdateOperation string

const (
	FieldUpdateOperationFormula FieldUpdateOperation = "Formula"

	FieldUpdateOperationLiteral FieldUpdateOperation = "Literal"

	FieldUpdateOperationNull FieldUpdateOperation = "Null"

	FieldUpdateOperationNextValue FieldUpdateOperation = "NextValue"

	FieldUpdateOperationPreviousValue FieldUpdateOperation = "PreviousValue"

	FieldUpdateOperationLookupValue FieldUpdateOperation = "LookupValue"
)

type KnowledgeWorkflowAction string

const (
	KnowledgeWorkflowActionPublishAsNew KnowledgeWorkflowAction = "PublishAsNew"

	KnowledgeWorkflowActionPublish KnowledgeWorkflowAction = "Publish"
)

type SendAction string

const (
	SendActionSend SendAction = "Send"
)

type ActionTaskAssignedToTypes string

const (
	ActionTaskAssignedToTypesUser ActionTaskAssignedToTypes = "user"

	ActionTaskAssignedToTypesRole ActionTaskAssignedToTypes = "role"

	ActionTaskAssignedToTypesOpportunityTeam ActionTaskAssignedToTypes = "opportunityTeam"

	ActionTaskAssignedToTypesAccountTeam ActionTaskAssignedToTypes = "accountTeam"

	ActionTaskAssignedToTypesOwner ActionTaskAssignedToTypes = "owner"

	ActionTaskAssignedToTypesAccountOwner ActionTaskAssignedToTypes = "accountOwner"

	ActionTaskAssignedToTypesCreator ActionTaskAssignedToTypes = "creator"

	ActionTaskAssignedToTypesAccountCreator ActionTaskAssignedToTypes = "accountCreator"

	ActionTaskAssignedToTypesPartnerUser ActionTaskAssignedToTypes = "partnerUser"

	ActionTaskAssignedToTypesPortalRole ActionTaskAssignedToTypes = "portalRole"
)

type ActionEmailRecipientTypes string

const (
	ActionEmailRecipientTypesGroup ActionEmailRecipientTypes = "group"

	ActionEmailRecipientTypesRole ActionEmailRecipientTypes = "role"

	ActionEmailRecipientTypesUser ActionEmailRecipientTypes = "user"

	ActionEmailRecipientTypesOpportunityTeam ActionEmailRecipientTypes = "opportunityTeam"

	ActionEmailRecipientTypesAccountTeam ActionEmailRecipientTypes = "accountTeam"

	ActionEmailRecipientTypesRoleSubordinates ActionEmailRecipientTypes = "roleSubordinates"

	ActionEmailRecipientTypesOwner ActionEmailRecipientTypes = "owner"

	ActionEmailRecipientTypesCreator ActionEmailRecipientTypes = "creator"

	ActionEmailRecipientTypesPartnerUser ActionEmailRecipientTypes = "partnerUser"

	ActionEmailRecipientTypesAccountOwner ActionEmailRecipientTypes = "accountOwner"

	ActionEmailRecipientTypesCustomerPortalUser ActionEmailRecipientTypes = "customerPortalUser"

	ActionEmailRecipientTypesPortalRole ActionEmailRecipientTypes = "portalRole"

	ActionEmailRecipientTypesPortalRoleSubordinates ActionEmailRecipientTypes = "portalRoleSubordinates"

	ActionEmailRecipientTypesContactLookup ActionEmailRecipientTypes = "contactLookup"

	ActionEmailRecipientTypesUserLookup ActionEmailRecipientTypes = "userLookup"

	ActionEmailRecipientTypesRoleSubordinatesInternal ActionEmailRecipientTypes = "roleSubordinatesInternal"

	ActionEmailRecipientTypesEmail ActionEmailRecipientTypes = "email"

	ActionEmailRecipientTypesCaseTeam ActionEmailRecipientTypes = "caseTeam"

	ActionEmailRecipientTypesCampaignMemberDerivedOwner ActionEmailRecipientTypes = "campaignMemberDerivedOwner"
)

type ActionEmailSenderType string

const (
	ActionEmailSenderTypeCurrentUser ActionEmailSenderType = "CurrentUser"

	ActionEmailSenderTypeOrgWideEmailAddress ActionEmailSenderType = "OrgWideEmailAddress"

	ActionEmailSenderTypeDefaultWorkflowUser ActionEmailSenderType = "DefaultWorkflowUser"
)

type WorkflowTriggerTypes string

const (
	WorkflowTriggerTypesOnCreateOnly WorkflowTriggerTypes = "onCreateOnly"

	WorkflowTriggerTypesOnCreateOrTriggeringUpdate WorkflowTriggerTypes = "onCreateOrTriggeringUpdate"

	WorkflowTriggerTypesOnAllChanges WorkflowTriggerTypes = "onAllChanges"

	WorkflowTriggerTypesOnRecursiveUpdate WorkflowTriggerTypes = "OnRecursiveUpdate"
)

type WorkflowTimeUnits string

const (
	WorkflowTimeUnitsHours WorkflowTimeUnits = "Hours"

	WorkflowTimeUnitsDays WorkflowTimeUnits = "Days"
)

type ExtendedErrorCode string

const ()

type TestLevel string

const (
	TestLevelNoTestRun TestLevel = "NoTestRun"

	TestLevelRunSpecifiedTests TestLevel = "RunSpecifiedTests"

	TestLevelRunLocalTests TestLevel = "RunLocalTests"

	TestLevelRunAllTestsInOrg TestLevel = "RunAllTestsInOrg"
)

type AsyncRequestState string

const (
	AsyncRequestStateQueued AsyncRequestState = "Queued"

	AsyncRequestStateInProgress AsyncRequestState = "InProgress"

	AsyncRequestStateCompleted AsyncRequestState = "Completed"

	AsyncRequestStateError AsyncRequestState = "Error"
)

type LogCategory string

const (
	LogCategoryDb LogCategory = "Db"

	LogCategoryWorkflow LogCategory = "Workflow"

	LogCategoryValidation LogCategory = "Validation"

	LogCategoryCallout LogCategory = "Callout"

	LogCategoryApexcode LogCategory = "Apexcode"

	LogCategoryApexprofiling LogCategory = "Apexprofiling"

	LogCategoryVisualforce LogCategory = "Visualforce"

	LogCategorySystem LogCategory = "System"

	LogCategoryAll LogCategory = "All"
)

type LogCategoryLevel string

const (
	LogCategoryLevelNone LogCategoryLevel = "None"

	LogCategoryLevelFinest LogCategoryLevel = "Finest"

	LogCategoryLevelFiner LogCategoryLevel = "Finer"

	LogCategoryLevelFine LogCategoryLevel = "Fine"

	LogCategoryLevelDebug LogCategoryLevel = "Debug"

	LogCategoryLevelInfo LogCategoryLevel = "Info"

	LogCategoryLevelWarn LogCategoryLevel = "Warn"

	LogCategoryLevelError LogCategoryLevel = "Error"
)

type LogType string

const (
	LogTypeNone LogType = "None"

	LogTypeDebugonly LogType = "Debugonly"

	LogTypeDb LogType = "Db"

	LogTypeProfiling LogType = "Profiling"

	LogTypeCallout LogType = "Callout"

	LogTypeDetail LogType = "Detail"
)

type ID string

const ()

type StatusCode string

const (
	StatusCodeALLORNONEOPERATIONROLLEDBACK StatusCode = "ALLORNONEOPERATIONROLLEDBACK"

	StatusCodeALREADYINPROCESS StatusCode = "ALREADYINPROCESS"

	StatusCodeAPEXDATAACCESSRESTRICTION StatusCode = "APEXDATAACCESSRESTRICTION"

	StatusCodeASSIGNEETYPEREQUIRED StatusCode = "ASSIGNEETYPEREQUIRED"

	StatusCodeAURACOMPILEERROR StatusCode = "AURACOMPILEERROR"

	StatusCodeBADCUSTOMENTITYPARENTDOMAIN StatusCode = "BADCUSTOMENTITYPARENTDOMAIN"

	StatusCodeBCCNOTALLOWEDIFBCCCOMPLIANCEENABLED StatusCode = "BCCNOTALLOWEDIFBCCCOMPLIANCEENABLED"

	StatusCodeCANNOTCASCADEPRODUCTACTIVE StatusCode = "CANNOTCASCADEPRODUCTACTIVE"

	StatusCodeCANNOTCHANGEFIELDTYPEOFAPEXREFERENCEDFIELD StatusCode = "CANNOTCHANGEFIELDTYPEOFAPEXREFERENCEDFIELD"

	StatusCodeCANNOTCHANGEFIELDTYPEOFREFERENCEDFIELD StatusCode = "CANNOTCHANGEFIELDTYPEOFREFERENCEDFIELD"

	StatusCodeCANNOTCREATEANOTHERMANAGEDPACKAGE StatusCode = "CANNOTCREATEANOTHERMANAGEDPACKAGE"

	StatusCodeCANNOTDEACTIVATEDIVISION StatusCode = "CANNOTDEACTIVATEDIVISION"

	StatusCodeCANNOTDELETEGLOBALACTIONLIST StatusCode = "CANNOTDELETEGLOBALACTIONLIST"

	StatusCodeCANNOTDELETELASTDATEDCONVERSIONRATE StatusCode = "CANNOTDELETELASTDATEDCONVERSIONRATE"

	StatusCodeCANNOTDELETEMANAGEDOBJECT StatusCode = "CANNOTDELETEMANAGEDOBJECT"

	StatusCodeCANNOTDISABLELASTADMIN StatusCode = "CANNOTDISABLELASTADMIN"

	StatusCodeCANNOTENABLEIPRESTRICTREQUESTS StatusCode = "CANNOTENABLEIPRESTRICTREQUESTS"

	StatusCodeCANNOTEXECUTEFLOWTRIGGER StatusCode = "CANNOTEXECUTEFLOWTRIGGER"

	StatusCodeCANNOTFREEZESELF StatusCode = "CANNOTFREEZESELF"

	StatusCodeCANNOTINSERTUPDATEACTIVATEENTITY StatusCode = "CANNOTINSERTUPDATEACTIVATEENTITY"

	StatusCodeCANNOTMODIFYMANAGEDOBJECT StatusCode = "CANNOTMODIFYMANAGEDOBJECT"

	StatusCodeCANNOTPASSWORDLOCKOUT StatusCode = "CANNOTPASSWORDLOCKOUT"

	StatusCodeCANNOTPOSTTOARCHIVEDGROUP StatusCode = "CANNOTPOSTTOARCHIVEDGROUP"

	StatusCodeCANNOTRENAMEAPEXREFERENCEDFIELD StatusCode = "CANNOTRENAMEAPEXREFERENCEDFIELD"

	StatusCodeCANNOTRENAMEAPEXREFERENCEDOBJECT StatusCode = "CANNOTRENAMEAPEXREFERENCEDOBJECT"

	StatusCodeCANNOTRENAMEREFERENCEDFIELD StatusCode = "CANNOTRENAMEREFERENCEDFIELD"

	StatusCodeCANNOTRENAMEREFERENCEDOBJECT StatusCode = "CANNOTRENAMEREFERENCEDOBJECT"

	StatusCodeCANNOTREPARENTRECORD StatusCode = "CANNOTREPARENTRECORD"

	StatusCodeCANNOTUPDATECONVERTEDLEAD StatusCode = "CANNOTUPDATECONVERTEDLEAD"

	StatusCodeCANTDISABLECORPCURRENCY StatusCode = "CANTDISABLECORPCURRENCY"

	StatusCodeCANTUNSETCORPCURRENCY StatusCode = "CANTUNSETCORPCURRENCY"

	StatusCodeCHILDSHAREFAILSPARENT StatusCode = "CHILDSHAREFAILSPARENT"

	StatusCodeCIRCULARDEPENDENCY StatusCode = "CIRCULARDEPENDENCY"

	StatusCodeCLEANSERVICEERROR StatusCode = "CLEANSERVICEERROR"

	StatusCodeCOLLISIONDETECTED StatusCode = "COLLISIONDETECTED"

	StatusCodeCOMMUNITYNOTACCESSIBLE StatusCode = "COMMUNITYNOTACCESSIBLE"

	StatusCodeCONFLICTINGENVIRONMENTHUBMEMBER StatusCode = "CONFLICTINGENVIRONMENTHUBMEMBER"

	StatusCodeCONFLICTINGSSOUSERMAPPING StatusCode = "CONFLICTINGSSOUSERMAPPING"

	StatusCodeCUSTOMAPEXERROR StatusCode = "CUSTOMAPEXERROR"

	StatusCodeCUSTOMCLOBFIELDLIMITEXCEEDED StatusCode = "CUSTOMCLOBFIELDLIMITEXCEEDED"

	StatusCodeCUSTOMENTITYORFIELDLIMIT StatusCode = "CUSTOMENTITYORFIELDLIMIT"

	StatusCodeCUSTOMFIELDINDEXLIMITEXCEEDED StatusCode = "CUSTOMFIELDINDEXLIMITEXCEEDED"

	StatusCodeCUSTOMINDEXEXISTS StatusCode = "CUSTOMINDEXEXISTS"

	StatusCodeCUSTOMLINKLIMITEXCEEDED StatusCode = "CUSTOMLINKLIMITEXCEEDED"

	StatusCodeCUSTOMMETADATALIMITEXCEEDED StatusCode = "CUSTOMMETADATALIMITEXCEEDED"

	StatusCodeCUSTOMSETTINGSLIMITEXCEEDED StatusCode = "CUSTOMSETTINGSLIMITEXCEEDED"

	StatusCodeCUSTOMTABLIMITEXCEEDED StatusCode = "CUSTOMTABLIMITEXCEEDED"

	StatusCodeDATACLOUDADDRESSNORECORDSFOUND StatusCode = "DATACLOUDADDRESSNORECORDSFOUND"

	StatusCodeDATACLOUDADDRESSPROCESSINGERROR StatusCode = "DATACLOUDADDRESSPROCESSINGERROR"

	StatusCodeDATACLOUDADDRESSSERVERERROR StatusCode = "DATACLOUDADDRESSSERVERERROR"

	StatusCodeDELETEFAILED StatusCode = "DELETEFAILED"

	StatusCodeDELETEOPERATIONTOOLARGE StatusCode = "DELETEOPERATIONTOOLARGE"

	StatusCodeDELETEREQUIREDONCASCADE StatusCode = "DELETEREQUIREDONCASCADE"

	StatusCodeDEPENDENCYEXISTS StatusCode = "DEPENDENCYEXISTS"

	StatusCodeDUPLICATESDETECTED StatusCode = "DUPLICATESDETECTED"

	StatusCodeDUPLICATECASESOLUTION StatusCode = "DUPLICATECASESOLUTION"

	StatusCodeDUPLICATECOMMNICKNAME StatusCode = "DUPLICATECOMMNICKNAME"

	StatusCodeDUPLICATECUSTOMENTITYDEFINITION StatusCode = "DUPLICATECUSTOMENTITYDEFINITION"

	StatusCodeDUPLICATECUSTOMTABMOTIF StatusCode = "DUPLICATECUSTOMTABMOTIF"

	StatusCodeDUPLICATEDEVELOPERNAME StatusCode = "DUPLICATEDEVELOPERNAME"

	StatusCodeDUPLICATEEXTERNALID StatusCode = "DUPLICATEEXTERNALID"

	StatusCodeDUPLICATEMASTERLABEL StatusCode = "DUPLICATEMASTERLABEL"

	StatusCodeDUPLICATESENDERDISPLAYNAME StatusCode = "DUPLICATESENDERDISPLAYNAME"

	StatusCodeDUPLICATEUSERNAME StatusCode = "DUPLICATEUSERNAME"

	StatusCodeDUPLICATEVALUE StatusCode = "DUPLICATEVALUE"

	StatusCodeEMAILADDRESSBOUNCED StatusCode = "EMAILADDRESSBOUNCED"

	StatusCodeEMAILEXTERNALTRANSPORTCONNECTIONERROR StatusCode = "EMAILEXTERNALTRANSPORTCONNECTIONERROR"

	StatusCodeEMAILEXTERNALTRANSPORTTOKENERROR StatusCode = "EMAILEXTERNALTRANSPORTTOKENERROR"

	StatusCodeEMAILEXTERNALTRANSPORTTOOMANYREQUESTSERROR StatusCode = "EMAILEXTERNALTRANSPORTTOOMANYREQUESTSERROR"

	StatusCodeEMAILEXTERNALTRANSPORTUNKNOWNERROR StatusCode = "EMAILEXTERNALTRANSPORTUNKNOWNERROR"

	StatusCodeEMAILNOTPROCESSEDDUETOPRIORERROR StatusCode = "EMAILNOTPROCESSEDDUETOPRIORERROR"

	StatusCodeEMAILOPTEDOUT StatusCode = "EMAILOPTEDOUT"

	StatusCodeEMAILTEMPLATEFORMULAERROR StatusCode = "EMAILTEMPLATEFORMULAERROR"

	StatusCodeEMAILTEMPLATEMERGEFIELDACCESSERROR StatusCode = "EMAILTEMPLATEMERGEFIELDACCESSERROR"

	StatusCodeEMAILTEMPLATEMERGEFIELDERROR StatusCode = "EMAILTEMPLATEMERGEFIELDERROR"

	StatusCodeEMAILTEMPLATEMERGEFIELDVALUEERROR StatusCode = "EMAILTEMPLATEMERGEFIELDVALUEERROR"

	StatusCodeEMAILTEMPLATEPROCESSINGERROR StatusCode = "EMAILTEMPLATEPROCESSINGERROR"

	StatusCodeEMPTYSCONTROLFILENAME StatusCode = "EMPTYSCONTROLFILENAME"

	StatusCodeENTITYFAILEDIFLASTMODIFIEDONUPDATE StatusCode = "ENTITYFAILEDIFLASTMODIFIEDONUPDATE"

	StatusCodeENTITYISARCHIVED StatusCode = "ENTITYISARCHIVED"

	StatusCodeENTITYISDELETED StatusCode = "ENTITYISDELETED"

	StatusCodeENTITYISLOCKED StatusCode = "ENTITYISLOCKED"

	StatusCodeENTITYSAVEERROR StatusCode = "ENTITYSAVEERROR"

	StatusCodeENTITYSAVEVALIDATIONERROR StatusCode = "ENTITYSAVEVALIDATIONERROR"

	StatusCodeENVIRONMENTHUBMEMBERSHIPCONFLICT StatusCode = "ENVIRONMENTHUBMEMBERSHIPCONFLICT"

	StatusCodeENVIRONMENTHUBMEMBERSHIPERRORJOININGHUB StatusCode = "ENVIRONMENTHUBMEMBERSHIPERRORJOININGHUB"

	StatusCodeENVIRONMENTHUBMEMBERSHIPUSERALREADYINHUB StatusCode = "ENVIRONMENTHUBMEMBERSHIPUSERALREADYINHUB"

	StatusCodeENVIRONMENTHUBMEMBERSHIPUSERNOTORGADMIN StatusCode = "ENVIRONMENTHUBMEMBERSHIPUSERNOTORGADMIN"

	StatusCodeERRORINMAILER StatusCode = "ERRORINMAILER"

	StatusCodeEXCHANGEWEBSERVICESURLINVALID StatusCode = "EXCHANGEWEBSERVICESURLINVALID"

	StatusCodeFAILEDACTIVATION StatusCode = "FAILEDACTIVATION"

	StatusCodeFIELDCUSTOMVALIDATIONEXCEPTION StatusCode = "FIELDCUSTOMVALIDATIONEXCEPTION"

	StatusCodeFIELDFILTERVALIDATIONEXCEPTION StatusCode = "FIELDFILTERVALIDATIONEXCEPTION"

	StatusCodeFIELDINTEGRITYEXCEPTION StatusCode = "FIELDINTEGRITYEXCEPTION"

	StatusCodeFIELDKEYWORDLISTMATCHLIMIT StatusCode = "FIELDKEYWORDLISTMATCHLIMIT"

	StatusCodeFIELDMAPPINGERROR StatusCode = "FIELDMAPPINGERROR"

	StatusCodeFIELDMODERATIONRULEBLOCK StatusCode = "FIELDMODERATIONRULEBLOCK"

	StatusCodeFIELDNOTUPDATABLE StatusCode = "FIELDNOTUPDATABLE"

	StatusCodeFILEEXTENSIONNOTALLOWED StatusCode = "FILEEXTENSIONNOTALLOWED"

	StatusCodeFILESIZELIMITEXCEEDED StatusCode = "FILESIZELIMITEXCEEDED"

	StatusCodeFILTEREDLOOKUPLIMITEXCEEDED StatusCode = "FILTEREDLOOKUPLIMITEXCEEDED"

	StatusCodeFINDDUPLICATESERROR StatusCode = "FINDDUPLICATESERROR"

	StatusCodeFUNCTIONALITYNOTENABLED StatusCode = "FUNCTIONALITYNOTENABLED"

	StatusCodeHASPUBLICREFERENCES StatusCode = "HASPUBLICREFERENCES"

	StatusCodeHTMLFILEUPLOADNOTALLOWED StatusCode = "HTMLFILEUPLOADNOTALLOWED"

	StatusCodeIMAGETOOLARGE StatusCode = "IMAGETOOLARGE"

	StatusCodeINACTIVEOWNERORUSER StatusCode = "INACTIVEOWNERORUSER"

	StatusCodeINACTIVERULEERROR StatusCode = "INACTIVERULEERROR"

	StatusCodeINSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE StatusCode = "INSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE"

	StatusCodeINSUFFICIENTACCESSONCROSSREFERENCEENTITY StatusCode = "INSUFFICIENTACCESSONCROSSREFERENCEENTITY"

	StatusCodeINSUFFICIENTACCESSORREADONLY StatusCode = "INSUFFICIENTACCESSORREADONLY"

	StatusCodeINSUFFICIENTACCESSTOINSIGHTSEXTERNALDATA StatusCode = "INSUFFICIENTACCESSTOINSIGHTSEXTERNALDATA"

	StatusCodeINSUFFICIENTCREDITS StatusCode = "INSUFFICIENTCREDITS"

	StatusCodeINVALIDACCESSLEVEL StatusCode = "INVALIDACCESSLEVEL"

	StatusCodeINVALIDARGUMENTTYPE StatusCode = "INVALIDARGUMENTTYPE"

	StatusCodeINVALIDASSIGNEETYPE StatusCode = "INVALIDASSIGNEETYPE"

	StatusCodeINVALIDASSIGNMENTRULE StatusCode = "INVALIDASSIGNMENTRULE"

	StatusCodeINVALIDBATCHOPERATION StatusCode = "INVALIDBATCHOPERATION"

	StatusCodeINVALIDCONTENTTYPE StatusCode = "INVALIDCONTENTTYPE"

	StatusCodeINVALIDCREDITCARDINFO StatusCode = "INVALIDCREDITCARDINFO"

	StatusCodeINVALIDCROSSREFERENCEKEY StatusCode = "INVALIDCROSSREFERENCEKEY"

	StatusCodeINVALIDCROSSREFERENCETYPEFORFIELD StatusCode = "INVALIDCROSSREFERENCETYPEFORFIELD"

	StatusCodeINVALIDCURRENCYCONVRATE StatusCode = "INVALIDCURRENCYCONVRATE"

	StatusCodeINVALIDCURRENCYCORPRATE StatusCode = "INVALIDCURRENCYCORPRATE"

	StatusCodeINVALIDCURRENCYISO StatusCode = "INVALIDCURRENCYISO"

	StatusCodeINVALIDDATACATEGORYGROUPREFERENCE StatusCode = "INVALIDDATACATEGORYGROUPREFERENCE"

	StatusCodeINVALIDDATAURI StatusCode = "INVALIDDATAURI"

	StatusCodeINVALIDEMAILADDRESS StatusCode = "INVALIDEMAILADDRESS"

	StatusCodeINVALIDEMPTYKEYOWNER StatusCode = "INVALIDEMPTYKEYOWNER"

	StatusCodeINVALIDENTITYFORMATCHENGINEERROR StatusCode = "INVALIDENTITYFORMATCHENGINEERROR"

	StatusCodeINVALIDENTITYFORMATCHOPERATIONERROR StatusCode = "INVALIDENTITYFORMATCHOPERATIONERROR"

	StatusCodeINVALIDENTITYFORUPSERT StatusCode = "INVALIDENTITYFORUPSERT"

	StatusCodeINVALIDENVIRONMENTHUBMEMBER StatusCode = "INVALIDENVIRONMENTHUBMEMBER"

	StatusCodeINVALIDEVENTDELIVERY StatusCode = "INVALIDEVENTDELIVERY"

	StatusCodeINVALIDEVENTSUBSCRIPTION StatusCode = "INVALIDEVENTSUBSCRIPTION"

	StatusCodeINVALIDFIELD StatusCode = "INVALIDFIELD"

	StatusCodeINVALIDFIELDFORINSERTUPDATE StatusCode = "INVALIDFIELDFORINSERTUPDATE"

	StatusCodeINVALIDFIELDWHENUSINGTEMPLATE StatusCode = "INVALIDFIELDWHENUSINGTEMPLATE"

	StatusCodeINVALIDFILTERACTION StatusCode = "INVALIDFILTERACTION"

	StatusCodeINVALIDGOOGLEDOCSURL StatusCode = "INVALIDGOOGLEDOCSURL"

	StatusCodeINVALIDIDFIELD StatusCode = "INVALIDIDFIELD"

	StatusCodeINVALIDINETADDRESS StatusCode = "INVALIDINETADDRESS"

	StatusCodeINVALIDINPUT StatusCode = "INVALIDINPUT"

	StatusCodeINVALIDLINEITEMCLONESTATE StatusCode = "INVALIDLINEITEMCLONESTATE"

	StatusCodeINVALIDMARKUP StatusCode = "INVALIDMARKUP"

	StatusCodeINVALIDMASTERORTRANSLATEDSOLUTION StatusCode = "INVALIDMASTERORTRANSLATEDSOLUTION"

	StatusCodeINVALIDMESSAGEIDREFERENCE StatusCode = "INVALIDMESSAGEIDREFERENCE"

	StatusCodeINVALIDNAMESPACEPREFIX StatusCode = "INVALIDNAMESPACEPREFIX"

	StatusCodeINVALIDOAUTHURL StatusCode = "INVALIDOAUTHURL"

	StatusCodeINVALIDOPERATION StatusCode = "INVALIDOPERATION"

	StatusCodeINVALIDOPERATOR StatusCode = "INVALIDOPERATOR"

	StatusCodeINVALIDORNULLFORRESTRICTEDPICKLIST StatusCode = "INVALIDORNULLFORRESTRICTEDPICKLIST"

	StatusCodeINVALIDOWNER StatusCode = "INVALIDOWNER"

	StatusCodeINVALIDPACKAGELICENSE StatusCode = "INVALIDPACKAGELICENSE"

	StatusCodeINVALIDPACKAGEVERSION StatusCode = "INVALIDPACKAGEVERSION"

	StatusCodeINVALIDPARTNERNETWORKSTATUS StatusCode = "INVALIDPARTNERNETWORKSTATUS"

	StatusCodeINVALIDPERSONACCOUNTOPERATION StatusCode = "INVALIDPERSONACCOUNTOPERATION"

	StatusCodeINVALIDQUERYLOCATOR StatusCode = "INVALIDQUERYLOCATOR"

	StatusCodeINVALIDREADONLYUSERDML StatusCode = "INVALIDREADONLYUSERDML"

	StatusCodeINVALIDRUNTIMEVALUE StatusCode = "INVALIDRUNTIMEVALUE"

	StatusCodeINVALIDSAVEASACTIVITYFLAG StatusCode = "INVALIDSAVEASACTIVITYFLAG"

	StatusCodeINVALIDSESSIONID StatusCode = "INVALIDSESSIONID"

	StatusCodeINVALIDSETUPOWNER StatusCode = "INVALIDSETUPOWNER"

	StatusCodeINVALIDSIGNUPCOUNTRY StatusCode = "INVALIDSIGNUPCOUNTRY"

	StatusCodeINVALIDSIGNUPOPTION StatusCode = "INVALIDSIGNUPOPTION"

	StatusCodeINVALIDSITEDELETEEXCEPTION StatusCode = "INVALIDSITEDELETEEXCEPTION"

	StatusCodeINVALIDSITEFILEIMPORTEDEXCEPTION StatusCode = "INVALIDSITEFILEIMPORTEDEXCEPTION"

	StatusCodeINVALIDSITEFILETYPEEXCEPTION StatusCode = "INVALIDSITEFILETYPEEXCEPTION"

	StatusCodeINVALIDSTATUS StatusCode = "INVALIDSTATUS"

	StatusCodeINVALIDSUBDOMAIN StatusCode = "INVALIDSUBDOMAIN"

	StatusCodeINVALIDTYPE StatusCode = "INVALIDTYPE"

	StatusCodeINVALIDTYPEFOROPERATION StatusCode = "INVALIDTYPEFOROPERATION"

	StatusCodeINVALIDTYPEONFIELDINRECORD StatusCode = "INVALIDTYPEONFIELDINRECORD"

	StatusCodeINVALIDUSERID StatusCode = "INVALIDUSERID"

	StatusCodeIPRANGELIMITEXCEEDED StatusCode = "IPRANGELIMITEXCEEDED"

	StatusCodeJIGSAWIMPORTLIMITEXCEEDED StatusCode = "JIGSAWIMPORTLIMITEXCEEDED"

	StatusCodeLICENSELIMITEXCEEDED StatusCode = "LICENSELIMITEXCEEDED"

	StatusCodeLIGHTPORTALUSEREXCEPTION StatusCode = "LIGHTPORTALUSEREXCEPTION"

	StatusCodeLIMITEXCEEDED StatusCode = "LIMITEXCEEDED"

	StatusCodeMALFORMEDID StatusCode = "MALFORMEDID"

	StatusCodeMANAGERNOTDEFINED StatusCode = "MANAGERNOTDEFINED"

	StatusCodeMASSMAILRETRYLIMITEXCEEDED StatusCode = "MASSMAILRETRYLIMITEXCEEDED"

	StatusCodeMASSMAILLIMITEXCEEDED StatusCode = "MASSMAILLIMITEXCEEDED"

	StatusCodeMATCHDEFINITIONERROR StatusCode = "MATCHDEFINITIONERROR"

	StatusCodeMATCHOPERATIONERROR StatusCode = "MATCHOPERATIONERROR"

	StatusCodeMATCHOPERATIONINVALIDENGINEERROR StatusCode = "MATCHOPERATIONINVALIDENGINEERROR"

	StatusCodeMATCHOPERATIONINVALIDRULEERROR StatusCode = "MATCHOPERATIONINVALIDRULEERROR"

	StatusCodeMATCHOPERATIONMISSINGENGINEERROR StatusCode = "MATCHOPERATIONMISSINGENGINEERROR"

	StatusCodeMATCHOPERATIONMISSINGOBJECTTYPEERROR StatusCode = "MATCHOPERATIONMISSINGOBJECTTYPEERROR"

	StatusCodeMATCHOPERATIONMISSINGOPTIONSERROR StatusCode = "MATCHOPERATIONMISSINGOPTIONSERROR"

	StatusCodeMATCHOPERATIONMISSINGRULEERROR StatusCode = "MATCHOPERATIONMISSINGRULEERROR"

	StatusCodeMATCHOPERATIONUNKNOWNRULEERROR StatusCode = "MATCHOPERATIONUNKNOWNRULEERROR"

	StatusCodeMATCHOPERATIONUNSUPPORTEDVERSIONERROR StatusCode = "MATCHOPERATIONUNSUPPORTEDVERSIONERROR"

	StatusCodeMATCHPRECONDITIONFAILED StatusCode = "MATCHPRECONDITIONFAILED"

	StatusCodeMATCHRUNTIMEERROR StatusCode = "MATCHRUNTIMEERROR"

	StatusCodeMATCHSERVICEERROR StatusCode = "MATCHSERVICEERROR"

	StatusCodeMATCHSERVICETIMEDOUT StatusCode = "MATCHSERVICETIMEDOUT"

	StatusCodeMATCHSERVICEUNAVAILABLEERROR StatusCode = "MATCHSERVICEUNAVAILABLEERROR"

	StatusCodeMAXIMUMCCEMAILSEXCEEDED StatusCode = "MAXIMUMCCEMAILSEXCEEDED"

	StatusCodeMAXIMUMDASHBOARDCOMPONENTSEXCEEDED StatusCode = "MAXIMUMDASHBOARDCOMPONENTSEXCEEDED"

	StatusCodeMAXIMUMHIERARCHYCHILDRENREACHED StatusCode = "MAXIMUMHIERARCHYCHILDRENREACHED"

	StatusCodeMAXIMUMHIERARCHYLEVELSREACHED StatusCode = "MAXIMUMHIERARCHYLEVELSREACHED"

	StatusCodeMAXIMUMHIERARCHYTREESIZEREACHED StatusCode = "MAXIMUMHIERARCHYTREESIZEREACHED"

	StatusCodeMAXIMUMSIZEOFATTACHMENT StatusCode = "MAXIMUMSIZEOFATTACHMENT"

	StatusCodeMAXIMUMSIZEOFDOCUMENT StatusCode = "MAXIMUMSIZEOFDOCUMENT"

	StatusCodeMAXACTIONSPERRULEEXCEEDED StatusCode = "MAXACTIONSPERRULEEXCEEDED"

	StatusCodeMAXACTIVERULESEXCEEDED StatusCode = "MAXACTIVERULESEXCEEDED"

	StatusCodeMAXAPPROVALSTEPSEXCEEDED StatusCode = "MAXAPPROVALSTEPSEXCEEDED"

	StatusCodeMAXDEPTHINFLOWEXECUTION StatusCode = "MAXDEPTHINFLOWEXECUTION"

	StatusCodeMAXFORMULASPERRULEEXCEEDED StatusCode = "MAXFORMULASPERRULEEXCEEDED"

	StatusCodeMAXRULESEXCEEDED StatusCode = "MAXRULESEXCEEDED"

	StatusCodeMAXRULEENTRIESEXCEEDED StatusCode = "MAXRULEENTRIESEXCEEDED"

	StatusCodeMAXTASKDESCRIPTIONEXCEEEDED StatusCode = "MAXTASKDESCRIPTIONEXCEEEDED"

	StatusCodeMAXTMRULESEXCEEDED StatusCode = "MAXTMRULESEXCEEDED"

	StatusCodeMAXTMRULEITEMSEXCEEDED StatusCode = "MAXTMRULEITEMSEXCEEDED"

	StatusCodeMERGEFAILED StatusCode = "MERGEFAILED"

	StatusCodeMETADATAFIELDUPDATEERROR StatusCode = "METADATAFIELDUPDATEERROR"

	StatusCodeMISSINGARGUMENT StatusCode = "MISSINGARGUMENT"

	StatusCodeMISSINGRECORD StatusCode = "MISSINGRECORD"

	StatusCodeMIXEDDMLOPERATION StatusCode = "MIXEDDMLOPERATION"

	StatusCodeNONUNIQUESHIPPINGADDRESS StatusCode = "NONUNIQUESHIPPINGADDRESS"

	StatusCodeNOAPPLICABLEPROCESS StatusCode = "NOAPPLICABLEPROCESS"

	StatusCodeNOATTACHMENTPERMISSION StatusCode = "NOATTACHMENTPERMISSION"

	StatusCodeNOINACTIVEDIVISIONMEMBERS StatusCode = "NOINACTIVEDIVISIONMEMBERS"

	StatusCodeNOMASSMAILPERMISSION StatusCode = "NOMASSMAILPERMISSION"

	StatusCodeNOPARTNERPERMISSION StatusCode = "NOPARTNERPERMISSION"

	StatusCodeNOSUCHUSEREXISTS StatusCode = "NOSUCHUSEREXISTS"

	StatusCodeNUMBEROUTSIDEVALIDRANGE StatusCode = "NUMBEROUTSIDEVALIDRANGE"

	StatusCodeNUMHISTORYFIELDSBYSOBJECTEXCEEDED StatusCode = "NUMHISTORYFIELDSBYSOBJECTEXCEEDED"

	StatusCodeOPTEDOUTOFMASSMAIL StatusCode = "OPTEDOUTOFMASSMAIL"

	StatusCodeOPWITHINVALIDUSERTYPEEXCEPTION StatusCode = "OPWITHINVALIDUSERTYPEEXCEPTION"

	StatusCodePACKAGELICENSEREQUIRED StatusCode = "PACKAGELICENSEREQUIRED"

	StatusCodePACKAGINGAPIINSTALLFAILED StatusCode = "PACKAGINGAPIINSTALLFAILED"

	StatusCodePACKAGINGAPIUNINSTALLFAILED StatusCode = "PACKAGINGAPIUNINSTALLFAILED"

	StatusCodePALIINVALIDACTIONID StatusCode = "PALIINVALIDACTIONID"

	StatusCodePALIINVALIDACTIONNAME StatusCode = "PALIINVALIDACTIONNAME"

	StatusCodePALIINVALIDACTIONTYPE StatusCode = "PALIINVALIDACTIONTYPE"

	StatusCodePALINVALIDASSISTANTRECOMMENDATIONTYPEID StatusCode = "PALINVALIDASSISTANTRECOMMENDATIONTYPEID"

	StatusCodePALINVALIDENTITYID StatusCode = "PALINVALIDENTITYID"

	StatusCodePALINVALIDFLEXIPAGEID StatusCode = "PALINVALIDFLEXIPAGEID"

	StatusCodePALINVALIDLAYOUTID StatusCode = "PALINVALIDLAYOUTID"

	StatusCodePALINVALIDPARAMETERS StatusCode = "PALINVALIDPARAMETERS"

	StatusCodePAAPIEXCEPTION StatusCode = "PAAPIEXCEPTION"

	StatusCodePAAXISFAULT StatusCode = "PAAXISFAULT"

	StatusCodePAINVALIDIDEXCEPTION StatusCode = "PAINVALIDIDEXCEPTION"

	StatusCodePANOACCESSEXCEPTION StatusCode = "PANOACCESSEXCEPTION"

	StatusCodePANODATAFOUNDEXCEPTION StatusCode = "PANODATAFOUNDEXCEPTION"

	StatusCodePAURISYNTAXEXCEPTION StatusCode = "PAURISYNTAXEXCEPTION"

	StatusCodePAVISIBLEACTIONSFILTERORDERINGEXCEPTION StatusCode = "PAVISIBLEACTIONSFILTERORDERINGEXCEPTION"

	StatusCodePORTALNOACCESS StatusCode = "PORTALNOACCESS"

	StatusCodePORTALUSERALREADYEXISTSFORCONTACT StatusCode = "PORTALUSERALREADYEXISTSFORCONTACT"

	StatusCodePORTALUSERCREATIONRESTRICTEDWITHENCRYPTION StatusCode = "PORTALUSERCREATIONRESTRICTEDWITHENCRYPTION"

	StatusCodePRIVATECONTACTONASSET StatusCode = "PRIVATECONTACTONASSET"

	StatusCodePROCESSINGHALTED StatusCode = "PROCESSINGHALTED"

	StatusCodeQAINVALIDCREATEFEEDITEM StatusCode = "QAINVALIDCREATEFEEDITEM"

	StatusCodeQAINVALIDSUCCESSMESSAGE StatusCode = "QAINVALIDSUCCESSMESSAGE"

	StatusCodeQUERYTIMEOUT StatusCode = "QUERYTIMEOUT"

	StatusCodeQUICKACTIONLISTITEMNOTALLOWED StatusCode = "QUICKACTIONLISTITEMNOTALLOWED"

	StatusCodeQUICKACTIONLISTNOTALLOWED StatusCode = "QUICKACTIONLISTNOTALLOWED"

	StatusCodeRECORDINUSEBYWORKFLOW StatusCode = "RECORDINUSEBYWORKFLOW"

	StatusCodeRELFIELDBADACCESSIBILITY StatusCode = "RELFIELDBADACCESSIBILITY"

	StatusCodeREPUTATIONMINIMUMNUMBERNOTREACHED StatusCode = "REPUTATIONMINIMUMNUMBERNOTREACHED"

	StatusCodeREQUESTRUNNINGTOOLONG StatusCode = "REQUESTRUNNINGTOOLONG"

	StatusCodeREQUIREDFEATUREMISSING StatusCode = "REQUIREDFEATUREMISSING"

	StatusCodeREQUIREDFIELDMISSING StatusCode = "REQUIREDFIELDMISSING"

	StatusCodeRETRIEVEEXCHANGEATTACHMENTFAILED StatusCode = "RETRIEVEEXCHANGEATTACHMENTFAILED"

	StatusCodeRETRIEVEEXCHANGEEMAILFAILED StatusCode = "RETRIEVEEXCHANGEEMAILFAILED"

	StatusCodeRETRIEVEEXCHANGEEVENTFAILED StatusCode = "RETRIEVEEXCHANGEEVENTFAILED"

	StatusCodeSALESFORCEINBOXTRANSPORTCONNECTIONERROR StatusCode = "SALESFORCEINBOXTRANSPORTCONNECTIONERROR"

	StatusCodeSALESFORCEINBOXTRANSPORTTOKENERROR StatusCode = "SALESFORCEINBOXTRANSPORTTOKENERROR"

	StatusCodeSALESFORCEINBOXTRANSPORTUNKNOWNERROR StatusCode = "SALESFORCEINBOXTRANSPORTUNKNOWNERROR"

	StatusCodeSELFREFERENCEFROMFLOW StatusCode = "SELFREFERENCEFROMFLOW"

	StatusCodeSELFREFERENCEFROMTRIGGER StatusCode = "SELFREFERENCEFROMTRIGGER"

	StatusCodeSHARENEEDEDFORCHILDOWNER StatusCode = "SHARENEEDEDFORCHILDOWNER"

	StatusCodeSINGLEEMAILLIMITEXCEEDED StatusCode = "SINGLEEMAILLIMITEXCEEDED"

	StatusCodeSOCIALACCOUNTNOTFOUND StatusCode = "SOCIALACCOUNTNOTFOUND"

	StatusCodeSOCIALACTIONINVALID StatusCode = "SOCIALACTIONINVALID"

	StatusCodeSOCIALPOSTINVALID StatusCode = "SOCIALPOSTINVALID"

	StatusCodeSOCIALPOSTNOTFOUND StatusCode = "SOCIALPOSTNOTFOUND"

	StatusCodeSTANDARDPRICENOTDEFINED StatusCode = "STANDARDPRICENOTDEFINED"

	StatusCodeSTORAGELIMITEXCEEDED StatusCode = "STORAGELIMITEXCEEDED"

	StatusCodeSTRINGTOOLONG StatusCode = "STRINGTOOLONG"

	StatusCodeSUBDOMAININUSE StatusCode = "SUBDOMAININUSE"

	StatusCodeTABSETLIMITEXCEEDED StatusCode = "TABSETLIMITEXCEEDED"

	StatusCodeTEMPLATENOTACTIVE StatusCode = "TEMPLATENOTACTIVE"

	StatusCodeTEMPLATENOTFOUND StatusCode = "TEMPLATENOTFOUND"

	StatusCodeTERRITORYREALIGNINPROGRESS StatusCode = "TERRITORYREALIGNINPROGRESS"

	StatusCodeTEXTDATAOUTSIDESUPPORTEDCHARSET StatusCode = "TEXTDATAOUTSIDESUPPORTEDCHARSET"

	StatusCodeTOOMANYAPEXREQUESTS StatusCode = "TOOMANYAPEXREQUESTS"

	StatusCodeTOOMANYENUMVALUE StatusCode = "TOOMANYENUMVALUE"

	StatusCodeTOOMANYPOSSIBLEUSERSEXIST StatusCode = "TOOMANYPOSSIBLEUSERSEXIST"

	StatusCodeTRANSFERREQUIRESREAD StatusCode = "TRANSFERREQUIRESREAD"

	StatusCodeUNABLETOLOCKROW StatusCode = "UNABLETOLOCKROW"

	StatusCodeUNAVAILABLERECORDTYPEEXCEPTION StatusCode = "UNAVAILABLERECORDTYPEEXCEPTION"

	StatusCodeUNAVAILABLEREF StatusCode = "UNAVAILABLEREF"

	StatusCodeUNDELETEFAILED StatusCode = "UNDELETEFAILED"

	StatusCodeUNKNOWNEXCEPTION StatusCode = "UNKNOWNEXCEPTION"

	StatusCodeUNSAFEHTMLCONTENT StatusCode = "UNSAFEHTMLCONTENT"

	StatusCodeUNSPECIFIEDEMAILADDRESS StatusCode = "UNSPECIFIEDEMAILADDRESS"

	StatusCodeUNSUPPORTEDAPEXTRIGGEROPERATON StatusCode = "UNSUPPORTEDAPEXTRIGGEROPERATON"

	StatusCodeUNVERIFIEDSENDERADDRESS StatusCode = "UNVERIFIEDSENDERADDRESS"

	StatusCodeUSEROWNSPORTALACCOUNTEXCEPTION StatusCode = "USEROWNSPORTALACCOUNTEXCEPTION"

	StatusCodeUSERWITHAPEXSHARESEXCEPTION StatusCode = "USERWITHAPEXSHARESEXCEPTION"

	StatusCodeVFCOMPILEERROR StatusCode = "VFCOMPILEERROR"

	StatusCodeWEBLINKSIZELIMITEXCEEDED StatusCode = "WEBLINKSIZELIMITEXCEEDED"

	StatusCodeWEBLINKURLINVALID StatusCode = "WEBLINKURLINVALID"

	StatusCodeWRONGCONTROLLERTYPE StatusCode = "WRONGCONTROLLERTYPE"

	StatusCodeXCLEANUNEXPECTEDERROR StatusCode = "XCLEANUNEXPECTEDERROR"
)

type AllOrNoneHeader struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AllOrNoneHeader"`

	AllOrNone bool `xml:"allOrNone,omitempty"`
}

type CallOptions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CallOptions"`

	Client string `xml:"client,omitempty"`
}

type DebuggingHeader struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DebuggingHeader"`

	Categories []*LogInfo `xml:"categories,omitempty"`

	DebugLevel *LogType `xml:"debugLevel,omitempty"`
}

type DebuggingInfo struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DebuggingInfo"`

	DebugLog string `xml:"debugLog,omitempty"`
}

type SessionHeader struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SessionHeader"`

	SessionId string `xml:"sessionId,omitempty"`
}

type CancelDeploy struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata cancelDeploy"`

	String *ID `xml:"String,omitempty"`
}

type CancelDeployResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata cancelDeployResponse"`

	Result *CancelDeployResult `xml:"result,omitempty"`
}

type CheckDeployStatus struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata checkDeployStatus"`

	AsyncProcessId *ID `xml:"asyncProcessId,omitempty"`

	IncludeDetails bool `xml:"includeDetails,omitempty"`
}

type CheckDeployStatusResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata checkDeployStatusResponse"`

	Result *DeployResult `xml:"result,omitempty"`
}

type CheckRetrieveStatus struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata checkRetrieveStatus"`

	AsyncProcessId *ID `xml:"asyncProcessId,omitempty"`

	IncludeZip bool `xml:"includeZip,omitempty"`
}

type CheckRetrieveStatusResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata checkRetrieveStatusResponse"`

	Result *RetrieveResult `xml:"result,omitempty"`
}

type CreateMetadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata createMetadata"`

	Metadata []*Metadata `xml:"metadata,omitempty"`
}

type CreateMetadataResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata createMetadataResponse"`

	Result []*SaveResult `xml:"result,omitempty"`
}

type DeleteMetadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata deleteMetadata"`

	Type_ string `xml:"type,omitempty"`

	FullNames []string `xml:"fullNames,omitempty"`
}

type DeleteMetadataResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata deleteMetadataResponse"`

	Result []*DeleteResult `xml:"result,omitempty"`
}

type Deploy struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata deploy"`

	ZipFile string `xml:"ZipFile,omitempty"`

	DeployOptions *DeployOptions `xml:"DeployOptions,omitempty"`
}

type DeployResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata deployResponse"`

	Result *AsyncResult `xml:"result,omitempty"`
}

type DeployRecentValidation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata deployRecentValidation"`

	ValidationId *ID `xml:"validationId,omitempty"`
}

type DeployRecentValidationResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata deployRecentValidationResponse"`

	Result string `xml:"result,omitempty"`
}

type DescribeMetadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata describeMetadata"`

	AsOfVersion float64 `xml:"asOfVersion,omitempty"`
}

type DescribeMetadataResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata describeMetadataResponse"`

	Result *DescribeMetadataResult `xml:"result,omitempty"`
}

type DescribeValueType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata describeValueType"`

	Type_ string `xml:"type,omitempty"`
}

type DescribeValueTypeResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata describeValueTypeResponse"`

	Result *DescribeValueTypeResult `xml:"result,omitempty"`
}

type ListMetadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata listMetadata"`

	Queries []*ListMetadataQuery `xml:"queries,omitempty"`

	AsOfVersion float64 `xml:"asOfVersion,omitempty"`
}

type ListMetadataResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata listMetadataResponse"`

	Result []*FileProperties `xml:"result,omitempty"`
}

type ReadMetadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata readMetadata"`

	Type_ string `xml:"type,omitempty"`

	FullNames []string `xml:"fullNames,omitempty"`
}

type ReadMetadataResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata readMetadataResponse"`

	Result *ReadResult `xml:"result,omitempty"`
}

type RenameMetadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata renameMetadata"`

	Type_ string `xml:"type,omitempty"`

	OldFullName string `xml:"oldFullName,omitempty"`

	NewFullName string `xml:"newFullName,omitempty"`
}

type RenameMetadataResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata renameMetadataResponse"`

	Result *SaveResult `xml:"result,omitempty"`
}

type Retrieve struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata retrieve"`

	RetrieveRequest *RetrieveRequest `xml:"retrieveRequest,omitempty"`
}

type RetrieveResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata retrieveResponse"`

	Result *AsyncResult `xml:"result,omitempty"`
}

type UpdateMetadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata updateMetadata"`

	Metadata []*Metadata `xml:"metadata,omitempty"`
}

type UpdateMetadataResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata updateMetadataResponse"`

	Result []*SaveResult `xml:"result,omitempty"`
}

type UpsertMetadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata upsertMetadata"`

	Metadata []*Metadata `xml:"metadata,omitempty"`
}

type UpsertMetadataResponse struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata upsertMetadataResponse"`

	Result []*UpsertResult `xml:"result,omitempty"`
}

type CancelDeployResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CancelDeployResult"`

	Done bool `xml:"done,omitempty"`

	Id *ID `xml:"id,omitempty"`
}

type DeployResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DeployResult"`

	CanceledBy string `xml:"canceledBy,omitempty"`

	CanceledByName string `xml:"canceledByName,omitempty"`

	CheckOnly bool `xml:"checkOnly,omitempty"`

	CompletedDate time.Time `xml:"completedDate,omitempty"`

	CreatedBy string `xml:"createdBy,omitempty"`

	CreatedByName string `xml:"createdByName,omitempty"`

	CreatedDate time.Time `xml:"createdDate,omitempty"`

	Details *DeployDetails `xml:"details,omitempty"`

	Done bool `xml:"done,omitempty"`

	ErrorMessage string `xml:"errorMessage,omitempty"`

	ErrorStatusCode *StatusCode `xml:"errorStatusCode,omitempty"`

	Id *ID `xml:"id,omitempty"`

	IgnoreWarnings bool `xml:"ignoreWarnings,omitempty"`

	LastModifiedDate time.Time `xml:"lastModifiedDate,omitempty"`

	NumberComponentErrors int32 `xml:"numberComponentErrors,omitempty"`

	NumberComponentsDeployed int32 `xml:"numberComponentsDeployed,omitempty"`

	NumberComponentsTotal int32 `xml:"numberComponentsTotal,omitempty"`

	NumberTestErrors int32 `xml:"numberTestErrors,omitempty"`

	NumberTestsCompleted int32 `xml:"numberTestsCompleted,omitempty"`

	NumberTestsTotal int32 `xml:"numberTestsTotal,omitempty"`

	RollbackOnError bool `xml:"rollbackOnError,omitempty"`

	RunTestsEnabled bool `xml:"runTestsEnabled,omitempty"`

	StartDate time.Time `xml:"startDate,omitempty"`

	StateDetail string `xml:"stateDetail,omitempty"`

	Status *DeployStatus `xml:"status,omitempty"`

	Success bool `xml:"success,omitempty"`
}

type DeployDetails struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DeployDetails"`

	ComponentFailures []*DeployMessage `xml:"componentFailures,omitempty"`

	ComponentSuccesses []*DeployMessage `xml:"componentSuccesses,omitempty"`

	RetrieveResult *RetrieveResult `xml:"retrieveResult,omitempty"`

	RunTestResult *RunTestsResult `xml:"runTestResult,omitempty"`
}

type DeployMessage struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DeployMessage"`

	Changed bool `xml:"changed,omitempty"`

	ColumnNumber int32 `xml:"columnNumber,omitempty"`

	ComponentType string `xml:"componentType,omitempty"`

	Created bool `xml:"created,omitempty"`

	CreatedDate time.Time `xml:"createdDate,omitempty"`

	Deleted bool `xml:"deleted,omitempty"`

	FileName string `xml:"fileName,omitempty"`

	FullName string `xml:"fullName,omitempty"`

	Id string `xml:"id,omitempty"`

	LineNumber int32 `xml:"lineNumber,omitempty"`

	Problem string `xml:"problem,omitempty"`

	ProblemType *DeployProblemType `xml:"problemType,omitempty"`

	Success bool `xml:"success,omitempty"`
}

type RetrieveResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RetrieveResult"`

	Done bool `xml:"done,omitempty"`

	ErrorMessage string `xml:"errorMessage,omitempty"`

	ErrorStatusCode *StatusCode `xml:"errorStatusCode,omitempty"`

	FileProperties []*FileProperties `xml:"fileProperties,omitempty"`

	Id string `xml:"id,omitempty"`

	Messages []*RetrieveMessage `xml:"messages,omitempty"`

	Status *RetrieveStatus `xml:"status,omitempty"`

	Success bool `xml:"success,omitempty"`

	ZipFile []byte `xml:"zipFile,omitempty"`
}

type FileProperties struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FileProperties"`

	CreatedById string `xml:"createdById,omitempty"`

	CreatedByName string `xml:"createdByName,omitempty"`

	CreatedDate time.Time `xml:"createdDate,omitempty"`

	FileName string `xml:"fileName,omitempty"`

	FullName string `xml:"fullName,omitempty"`

	Id string `xml:"id,omitempty"`

	LastModifiedById string `xml:"lastModifiedById,omitempty"`

	LastModifiedByName string `xml:"lastModifiedByName,omitempty"`

	LastModifiedDate time.Time `xml:"lastModifiedDate,omitempty"`

	ManageableState *ManageableState `xml:"manageableState,omitempty"`

	NamespacePrefix string `xml:"namespacePrefix,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type RetrieveMessage struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RetrieveMessage"`

	FileName string `xml:"fileName,omitempty"`

	Problem string `xml:"problem,omitempty"`
}

type RunTestsResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RunTestsResult"`

	ApexLogId string `xml:"apexLogId,omitempty"`

	CodeCoverage []*CodeCoverageResult `xml:"codeCoverage,omitempty"`

	CodeCoverageWarnings []*CodeCoverageWarning `xml:"codeCoverageWarnings,omitempty"`

	Failures []*RunTestFailure `xml:"failures,omitempty"`

	NumFailures int32 `xml:"numFailures,omitempty"`

	NumTestsRun int32 `xml:"numTestsRun,omitempty"`

	Successes []*RunTestSuccess `xml:"successes,omitempty"`

	TotalTime float64 `xml:"totalTime,omitempty"`
}

type CodeCoverageResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CodeCoverageResult"`

	DmlInfo []*CodeLocation `xml:"dmlInfo,omitempty"`

	Id *ID `xml:"id,omitempty"`

	LocationsNotCovered []*CodeLocation `xml:"locationsNotCovered,omitempty"`

	MethodInfo []*CodeLocation `xml:"methodInfo,omitempty"`

	Name string `xml:"name,omitempty"`

	Namespace string `xml:"namespace,omitempty"`

	NumLocations int32 `xml:"numLocations,omitempty"`

	NumLocationsNotCovered int32 `xml:"numLocationsNotCovered,omitempty"`

	SoqlInfo []*CodeLocation `xml:"soqlInfo,omitempty"`

	SoslInfo []*CodeLocation `xml:"soslInfo,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type CodeLocation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CodeLocation"`

	Column int32 `xml:"column,omitempty"`

	Line int32 `xml:"line,omitempty"`

	NumExecutions int32 `xml:"numExecutions,omitempty"`

	Time float64 `xml:"time,omitempty"`
}

type CodeCoverageWarning struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CodeCoverageWarning"`

	Id *ID `xml:"id,omitempty"`

	Message string `xml:"message,omitempty"`

	Name string `xml:"name,omitempty"`

	Namespace string `xml:"namespace,omitempty"`
}

type RunTestFailure struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RunTestFailure"`

	Id *ID `xml:"id,omitempty"`

	Message string `xml:"message,omitempty"`

	MethodName string `xml:"methodName,omitempty"`

	Name string `xml:"name,omitempty"`

	Namespace string `xml:"namespace,omitempty"`

	PackageName string `xml:"packageName,omitempty"`

	SeeAllData bool `xml:"seeAllData,omitempty"`

	StackTrace string `xml:"stackTrace,omitempty"`

	Time float64 `xml:"time,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type RunTestSuccess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RunTestSuccess"`

	Id *ID `xml:"id,omitempty"`

	MethodName string `xml:"methodName,omitempty"`

	Name string `xml:"name,omitempty"`

	Namespace string `xml:"namespace,omitempty"`

	SeeAllData bool `xml:"seeAllData,omitempty"`

	Time float64 `xml:"time,omitempty"`
}

type Metadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Metadata"`

	FullName string `xml:"fullName,omitempty"`
}

type AccountSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AccountSettings"`

	*Metadata

	EnableAccountOwnerReport bool `xml:"enableAccountOwnerReport,omitempty"`

	EnableAccountTeams bool `xml:"enableAccountTeams,omitempty"`

	ShowViewHierarchyLink bool `xml:"showViewHierarchyLink,omitempty"`
}

type ActionLinkGroupTemplate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ActionLinkGroupTemplate"`

	*Metadata

	ActionLinkTemplates []*ActionLinkTemplate `xml:"actionLinkTemplates,omitempty"`

	Category *PlatformActionGroupCategory `xml:"category,omitempty"`

	ExecutionsAllowed *ActionLinkExecutionsAllowed `xml:"executionsAllowed,omitempty"`

	HoursUntilExpiration int32 `xml:"hoursUntilExpiration,omitempty"`

	IsPublished bool `xml:"isPublished,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ActionLinkTemplate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ActionLinkTemplate"`

	ActionUrl string `xml:"actionUrl,omitempty"`

	Headers string `xml:"headers,omitempty"`

	IsConfirmationRequired bool `xml:"isConfirmationRequired,omitempty"`

	IsGroupDefault bool `xml:"isGroupDefault,omitempty"`

	Label string `xml:"label,omitempty"`

	LabelKey string `xml:"labelKey,omitempty"`

	LinkType *ActionLinkType `xml:"linkType,omitempty"`

	Method *ActionLinkHttpMethod `xml:"method,omitempty"`

	Position int32 `xml:"position,omitempty"`

	RequestBody string `xml:"requestBody,omitempty"`

	UserAlias string `xml:"userAlias,omitempty"`

	UserVisibility *ActionLinkUserVisibility `xml:"userVisibility,omitempty"`
}

type ActivitiesSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ActivitiesSettings"`

	*Metadata

	AllowUsersToRelateMultipleContactsToTasksAndEvents bool `xml:"allowUsersToRelateMultipleContactsToTasksAndEvents,omitempty"`

	EnableActivityReminders bool `xml:"enableActivityReminders,omitempty"`

	EnableClickCreateEvents bool `xml:"enableClickCreateEvents,omitempty"`

	EnableDragAndDropScheduling bool `xml:"enableDragAndDropScheduling,omitempty"`

	EnableEmailTracking bool `xml:"enableEmailTracking,omitempty"`

	EnableGroupTasks bool `xml:"enableGroupTasks,omitempty"`

	EnableListViewScheduling bool `xml:"enableListViewScheduling,omitempty"`

	EnableLogNote bool `xml:"enableLogNote,omitempty"`

	EnableMultidayEvents bool `xml:"enableMultidayEvents,omitempty"`

	EnableRecurringEvents bool `xml:"enableRecurringEvents,omitempty"`

	EnableRecurringTasks bool `xml:"enableRecurringTasks,omitempty"`

	EnableSidebarCalendarShortcut bool `xml:"enableSidebarCalendarShortcut,omitempty"`

	EnableSimpleTaskCreateUI bool `xml:"enableSimpleTaskCreateUI,omitempty"`

	EnableUNSTaskDelegatedToNotifications bool `xml:"enableUNSTaskDelegatedToNotifications,omitempty"`

	MeetingRequestsLogo string `xml:"meetingRequestsLogo,omitempty"`

	ShowCustomLogoMeetingRequests bool `xml:"showCustomLogoMeetingRequests,omitempty"`

	ShowEventDetailsMultiUserCalendar bool `xml:"showEventDetailsMultiUserCalendar,omitempty"`

	ShowHomePageHoverLinksForEvents bool `xml:"showHomePageHoverLinksForEvents,omitempty"`

	ShowMyTasksHoverLinks bool `xml:"showMyTasksHoverLinks,omitempty"`

	ShowRequestedMeetingsOnHomePage bool `xml:"showRequestedMeetingsOnHomePage,omitempty"`
}

type AddressSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AddressSettings"`

	*Metadata

	CountriesAndStates *CountriesAndStates `xml:"countriesAndStates,omitempty"`
}

type CountriesAndStates struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CountriesAndStates"`

	Countries []*Country `xml:"countries,omitempty"`
}

type Country struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Country"`

	Active bool `xml:"active,omitempty"`

	IntegrationValue string `xml:"integrationValue,omitempty"`

	IsoCode string `xml:"isoCode,omitempty"`

	Label string `xml:"label,omitempty"`

	OrgDefault bool `xml:"orgDefault,omitempty"`

	Standard bool `xml:"standard,omitempty"`

	States []*State `xml:"states,omitempty"`

	Visible bool `xml:"visible,omitempty"`
}

type State struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata State"`

	Active bool `xml:"active,omitempty"`

	IntegrationValue string `xml:"integrationValue,omitempty"`

	IsoCode string `xml:"isoCode,omitempty"`

	Label string `xml:"label,omitempty"`

	Standard bool `xml:"standard,omitempty"`

	Visible bool `xml:"visible,omitempty"`
}

type AnalyticSnapshot struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AnalyticSnapshot"`

	*Metadata

	Description string `xml:"description,omitempty"`

	GroupColumn string `xml:"groupColumn,omitempty"`

	Mappings []*AnalyticSnapshotMapping `xml:"mappings,omitempty"`

	Name string `xml:"name,omitempty"`

	RunningUser string `xml:"runningUser,omitempty"`

	SourceReport string `xml:"sourceReport,omitempty"`

	TargetObject string `xml:"targetObject,omitempty"`
}

type AnalyticSnapshotMapping struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AnalyticSnapshotMapping"`

	AggregateType *ReportSummaryType `xml:"aggregateType,omitempty"`

	SourceField string `xml:"sourceField,omitempty"`

	SourceType *ReportJobSourceTypes `xml:"sourceType,omitempty"`

	TargetField string `xml:"targetField,omitempty"`
}

type ApexTestSuite struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApexTestSuite"`

	*Metadata

	TestClassName []string `xml:"testClassName,omitempty"`
}

type AppMenu struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AppMenu"`

	*Metadata

	AppMenuItems []*AppMenuItem `xml:"appMenuItems,omitempty"`
}

type AppMenuItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AppMenuItem"`

	Name string `xml:"name,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type ApprovalProcess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApprovalProcess"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	AllowRecall bool `xml:"allowRecall,omitempty"`

	AllowedSubmitters []*ApprovalSubmitter `xml:"allowedSubmitters,omitempty"`

	ApprovalPageFields *ApprovalPageField `xml:"approvalPageFields,omitempty"`

	ApprovalStep []*ApprovalStep `xml:"approvalStep,omitempty"`

	Description string `xml:"description,omitempty"`

	EmailTemplate string `xml:"emailTemplate,omitempty"`

	EnableMobileDeviceAccess bool `xml:"enableMobileDeviceAccess,omitempty"`

	EntryCriteria *ApprovalEntryCriteria `xml:"entryCriteria,omitempty"`

	FinalApprovalActions *ApprovalAction `xml:"finalApprovalActions,omitempty"`

	FinalApprovalRecordLock bool `xml:"finalApprovalRecordLock,omitempty"`

	FinalRejectionActions *ApprovalAction `xml:"finalRejectionActions,omitempty"`

	FinalRejectionRecordLock bool `xml:"finalRejectionRecordLock,omitempty"`

	InitialSubmissionActions *ApprovalAction `xml:"initialSubmissionActions,omitempty"`

	Label string `xml:"label,omitempty"`

	NextAutomatedApprover *NextAutomatedApprover `xml:"nextAutomatedApprover,omitempty"`

	PostTemplate string `xml:"postTemplate,omitempty"`

	RecallActions *ApprovalAction `xml:"recallActions,omitempty"`

	RecordEditability *RecordEditabilityType `xml:"recordEditability,omitempty"`

	ShowApprovalHistory bool `xml:"showApprovalHistory,omitempty"`
}

type ApprovalSubmitter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApprovalSubmitter"`

	Submitter string `xml:"submitter,omitempty"`

	Type_ *ProcessSubmitterType `xml:"type,omitempty"`
}

type ApprovalPageField struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApprovalPageField"`

	Field []string `xml:"field,omitempty"`
}

type ApprovalStep struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApprovalStep"`

	AllowDelegate bool `xml:"allowDelegate,omitempty"`

	ApprovalActions *ApprovalAction `xml:"approvalActions,omitempty"`

	AssignedApprover *ApprovalStepApprover `xml:"assignedApprover,omitempty"`

	Description string `xml:"description,omitempty"`

	EntryCriteria *ApprovalEntryCriteria `xml:"entryCriteria,omitempty"`

	IfCriteriaNotMet *StepCriteriaNotMetType `xml:"ifCriteriaNotMet,omitempty"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`

	RejectBehavior *ApprovalStepRejectBehavior `xml:"rejectBehavior,omitempty"`

	RejectionActions *ApprovalAction `xml:"rejectionActions,omitempty"`
}

type ApprovalAction struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApprovalAction"`

	Action []*WorkflowActionReference `xml:"action,omitempty"`
}

type WorkflowActionReference struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowActionReference"`

	Name string `xml:"name,omitempty"`

	Type_ *WorkflowActionType `xml:"type,omitempty"`
}

type ApprovalStepApprover struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApprovalStepApprover"`

	Approver []*Approver `xml:"approver,omitempty"`

	WhenMultipleApprovers *RoutingType `xml:"whenMultipleApprovers,omitempty"`
}

type Approver struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Approver"`

	Name string `xml:"name,omitempty"`

	Type_ *NextOwnerType `xml:"type,omitempty"`
}

type ApprovalEntryCriteria struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApprovalEntryCriteria"`

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	CriteriaItems []*FilterItem `xml:"criteriaItems,omitempty"`

	Formula string `xml:"formula,omitempty"`
}

type FilterItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FilterItem"`

	Field string `xml:"field,omitempty"`

	Operation *FilterOperation `xml:"operation,omitempty"`

	Value string `xml:"value,omitempty"`

	ValueField string `xml:"valueField,omitempty"`
}

type DuplicateRuleFilterItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DuplicateRuleFilterItem"`

	*FilterItem

	SortOrder int32 `xml:"sortOrder,omitempty"`

	Table string `xml:"table,omitempty"`
}

type ApprovalStepRejectBehavior struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApprovalStepRejectBehavior"`

	Type_ *StepRejectBehaviorType `xml:"type,omitempty"`
}

type NextAutomatedApprover struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata NextAutomatedApprover"`

	UseApproverFieldOfRecordOwner bool `xml:"useApproverFieldOfRecordOwner,omitempty"`

	UserHierarchyField string `xml:"userHierarchyField,omitempty"`
}

type AssignmentRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AssignmentRule"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	RuleEntry []*RuleEntry `xml:"ruleEntry,omitempty"`
}

type RuleEntry struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RuleEntry"`

	AssignedTo string `xml:"assignedTo,omitempty"`

	AssignedToType *AssignToLookupValueType `xml:"assignedToType,omitempty"`

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	BusinessHours string `xml:"businessHours,omitempty"`

	BusinessHoursSource *BusinessHoursSourceType `xml:"businessHoursSource,omitempty"`

	CriteriaItems []*FilterItem `xml:"criteriaItems,omitempty"`

	DisableEscalationWhenModified bool `xml:"disableEscalationWhenModified,omitempty"`

	EscalationAction []*EscalationAction `xml:"escalationAction,omitempty"`

	EscalationStartTime *EscalationStartTimeType `xml:"escalationStartTime,omitempty"`

	Formula string `xml:"formula,omitempty"`

	NotifyCcRecipients bool `xml:"notifyCcRecipients,omitempty"`

	OverrideExistingTeams bool `xml:"overrideExistingTeams,omitempty"`

	ReplyToEmail string `xml:"replyToEmail,omitempty"`

	SenderEmail string `xml:"senderEmail,omitempty"`

	SenderName string `xml:"senderName,omitempty"`

	Team []string `xml:"team,omitempty"`

	Template string `xml:"template,omitempty"`
}

type EscalationAction struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EscalationAction"`

	AssignedTo string `xml:"assignedTo,omitempty"`

	AssignedToTemplate string `xml:"assignedToTemplate,omitempty"`

	AssignedToType *AssignToLookupValueType `xml:"assignedToType,omitempty"`

	MinutesToEscalation int32 `xml:"minutesToEscalation,omitempty"`

	NotifyCaseOwner bool `xml:"notifyCaseOwner,omitempty"`

	NotifyEmail []string `xml:"notifyEmail,omitempty"`

	NotifyTo string `xml:"notifyTo,omitempty"`

	NotifyToTemplate string `xml:"notifyToTemplate,omitempty"`
}

type AssignmentRules struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AssignmentRules"`

	*Metadata

	AssignmentRule []*AssignmentRule `xml:"assignmentRule,omitempty"`
}

type AssistantRecommendationType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AssistantRecommendationType"`

	*Metadata

	Description string `xml:"description,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	PlatformActionlist *PlatformActionList `xml:"platformActionlist,omitempty"`

	SobjectType string `xml:"sobjectType,omitempty"`

	Title string `xml:"title,omitempty"`
}

type PlatformActionList struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PlatformActionList"`

	*Metadata

	ActionListContext *PlatformActionListContext `xml:"actionListContext,omitempty"`

	PlatformActionListItems []*PlatformActionListItem `xml:"platformActionListItems,omitempty"`

	RelatedSourceEntity string `xml:"relatedSourceEntity,omitempty"`
}

type PlatformActionListItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PlatformActionListItem"`

	ActionName string `xml:"actionName,omitempty"`

	ActionType *PlatformActionType `xml:"actionType,omitempty"`

	SortOrder int32 `xml:"sortOrder,omitempty"`

	Subtype string `xml:"subtype,omitempty"`
}

type AuraDefinitionBundle struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AuraDefinitionBundle"`

	*Metadata

	SVGContent []byte `xml:"SVGContent,omitempty"`

	ApiVersion float64 `xml:"apiVersion,omitempty"`

	ControllerContent []byte `xml:"controllerContent,omitempty"`

	Description string `xml:"description,omitempty"`

	DesignContent []byte `xml:"designContent,omitempty"`

	DocumentationContent []byte `xml:"documentationContent,omitempty"`

	HelperContent []byte `xml:"helperContent,omitempty"`

	Markup []byte `xml:"markup,omitempty"`

	ModelContent []byte `xml:"modelContent,omitempty"`

	PackageVersions []*PackageVersion `xml:"packageVersions,omitempty"`

	RendererContent []byte `xml:"rendererContent,omitempty"`

	StyleContent []byte `xml:"styleContent,omitempty"`

	TestsuiteContent []byte `xml:"testsuiteContent,omitempty"`

	Type_ *AuraBundleType `xml:"type,omitempty"`
}

type PackageVersion struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PackageVersion"`

	MajorNumber int32 `xml:"majorNumber,omitempty"`

	MinorNumber int32 `xml:"minorNumber,omitempty"`

	Namespace string `xml:"namespace,omitempty"`
}

type AuthProvider struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AuthProvider"`

	*Metadata

	AuthorizeUrl string `xml:"authorizeUrl,omitempty"`

	ConsumerKey string `xml:"consumerKey,omitempty"`

	ConsumerSecret string `xml:"consumerSecret,omitempty"`

	CustomMetadataTypeRecord string `xml:"customMetadataTypeRecord,omitempty"`

	DefaultScopes string `xml:"defaultScopes,omitempty"`

	ErrorUrl string `xml:"errorUrl,omitempty"`

	ExecutionUser string `xml:"executionUser,omitempty"`

	FriendlyName string `xml:"friendlyName,omitempty"`

	IconUrl string `xml:"iconUrl,omitempty"`

	IdTokenIssuer string `xml:"idTokenIssuer,omitempty"`

	IncludeOrgIdInIdentifier bool `xml:"includeOrgIdInIdentifier,omitempty"`

	LogoutUrl string `xml:"logoutUrl,omitempty"`

	Plugin string `xml:"plugin,omitempty"`

	Portal string `xml:"portal,omitempty"`

	ProviderType *AuthProviderType `xml:"providerType,omitempty"`

	RegistrationHandler string `xml:"registrationHandler,omitempty"`

	SendAccessTokenInHeader bool `xml:"sendAccessTokenInHeader,omitempty"`

	SendClientCredentialsInHeader bool `xml:"sendClientCredentialsInHeader,omitempty"`

	TokenUrl string `xml:"tokenUrl,omitempty"`

	UserInfoUrl string `xml:"userInfoUrl,omitempty"`
}

type AutoResponseRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AutoResponseRule"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	RuleEntry []*RuleEntry `xml:"ruleEntry,omitempty"`
}

type AutoResponseRules struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AutoResponseRules"`

	*Metadata

	AutoResponseRule []*AutoResponseRule `xml:"autoResponseRule,omitempty"`
}

type BusinessHoursEntry struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata BusinessHoursEntry"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	Default_ bool `xml:"default,omitempty"`

	FridayEndTime time.Time `xml:"fridayEndTime,omitempty"`

	FridayStartTime time.Time `xml:"fridayStartTime,omitempty"`

	MondayEndTime time.Time `xml:"mondayEndTime,omitempty"`

	MondayStartTime time.Time `xml:"mondayStartTime,omitempty"`

	Name string `xml:"name,omitempty"`

	SaturdayEndTime time.Time `xml:"saturdayEndTime,omitempty"`

	SaturdayStartTime time.Time `xml:"saturdayStartTime,omitempty"`

	SundayEndTime time.Time `xml:"sundayEndTime,omitempty"`

	SundayStartTime time.Time `xml:"sundayStartTime,omitempty"`

	ThursdayEndTime time.Time `xml:"thursdayEndTime,omitempty"`

	ThursdayStartTime time.Time `xml:"thursdayStartTime,omitempty"`

	TimeZoneId string `xml:"timeZoneId,omitempty"`

	TuesdayEndTime time.Time `xml:"tuesdayEndTime,omitempty"`

	TuesdayStartTime time.Time `xml:"tuesdayStartTime,omitempty"`

	WednesdayEndTime time.Time `xml:"wednesdayEndTime,omitempty"`

	WednesdayStartTime time.Time `xml:"wednesdayStartTime,omitempty"`
}

type BusinessHoursSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata BusinessHoursSettings"`

	*Metadata

	BusinessHours []*BusinessHoursEntry `xml:"businessHours,omitempty"`

	Holidays []*Holiday `xml:"holidays,omitempty"`
}

type Holiday struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Holiday"`

	ActivityDate time.Time `xml:"activityDate,omitempty"`

	BusinessHours []string `xml:"businessHours,omitempty"`

	Description string `xml:"description,omitempty"`

	EndTime time.Time `xml:"endTime,omitempty"`

	IsRecurring bool `xml:"isRecurring,omitempty"`

	Name string `xml:"name,omitempty"`

	RecurrenceDayOfMonth int32 `xml:"recurrenceDayOfMonth,omitempty"`

	RecurrenceDayOfWeek []string `xml:"recurrenceDayOfWeek,omitempty"`

	RecurrenceDayOfWeekMask int32 `xml:"recurrenceDayOfWeekMask,omitempty"`

	RecurrenceEndDate time.Time `xml:"recurrenceEndDate,omitempty"`

	RecurrenceInstance string `xml:"recurrenceInstance,omitempty"`

	RecurrenceInterval int32 `xml:"recurrenceInterval,omitempty"`

	RecurrenceMonthOfYear string `xml:"recurrenceMonthOfYear,omitempty"`

	RecurrenceStartDate time.Time `xml:"recurrenceStartDate,omitempty"`

	RecurrenceType string `xml:"recurrenceType,omitempty"`

	StartTime time.Time `xml:"startTime,omitempty"`
}

type BusinessProcess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata BusinessProcess"`

	*Metadata

	Description string `xml:"description,omitempty"`

	IsActive bool `xml:"isActive,omitempty"`

	Values []*PicklistValue `xml:"values,omitempty"`
}

type PicklistValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PicklistValue"`

	*GlobalPicklistValue

	AllowEmail bool `xml:"allowEmail,omitempty"`

	Closed bool `xml:"closed,omitempty"`

	ControllingFieldValues []string `xml:"controllingFieldValues,omitempty"`

	Converted bool `xml:"converted,omitempty"`

	CssExposed bool `xml:"cssExposed,omitempty"`

	ForecastCategory *ForecastCategories `xml:"forecastCategory,omitempty"`

	HighPriority bool `xml:"highPriority,omitempty"`

	Probability int32 `xml:"probability,omitempty"`

	ReverseRole string `xml:"reverseRole,omitempty"`

	Reviewed bool `xml:"reviewed,omitempty"`

	Won bool `xml:"won,omitempty"`
}

type GlobalPicklistValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata GlobalPicklistValue"`

	*Metadata

	Color string `xml:"color,omitempty"`

	Default_ bool `xml:"default,omitempty"`

	Description string `xml:"description,omitempty"`

	IsActive bool `xml:"isActive,omitempty"`
}

type CallCenter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CallCenter"`

	*Metadata

	AdapterUrl string `xml:"adapterUrl,omitempty"`

	CustomSettings string `xml:"customSettings,omitempty"`

	DisplayName string `xml:"displayName,omitempty"`

	DisplayNameLabel string `xml:"displayNameLabel,omitempty"`

	InternalNameLabel string `xml:"internalNameLabel,omitempty"`

	Sections []*CallCenterSection `xml:"sections,omitempty"`

	Version string `xml:"version,omitempty"`
}

type CallCenterSection struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CallCenterSection"`

	Items []*CallCenterItem `xml:"items,omitempty"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type CallCenterItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CallCenterItem"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`

	Value string `xml:"value,omitempty"`
}

type CampaignInfluenceModel struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CampaignInfluenceModel"`

	*Metadata

	IsDefaultModel bool `xml:"isDefaultModel,omitempty"`

	IsModelLocked bool `xml:"isModelLocked,omitempty"`

	ModelDescription string `xml:"modelDescription,omitempty"`

	Name string `xml:"name,omitempty"`
}

type CaseSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CaseSettings"`

	*Metadata

	CaseAssignNotificationTemplate string `xml:"caseAssignNotificationTemplate,omitempty"`

	CaseCloseNotificationTemplate string `xml:"caseCloseNotificationTemplate,omitempty"`

	CaseCommentNotificationTemplate string `xml:"caseCommentNotificationTemplate,omitempty"`

	CaseCreateNotificationTemplate string `xml:"caseCreateNotificationTemplate,omitempty"`

	CaseFeedItemSettings []*FeedItemSettings `xml:"caseFeedItemSettings,omitempty"`

	CloseCaseThroughStatusChange bool `xml:"closeCaseThroughStatusChange,omitempty"`

	DefaultCaseOwner string `xml:"defaultCaseOwner,omitempty"`

	DefaultCaseOwnerType string `xml:"defaultCaseOwnerType,omitempty"`

	DefaultCaseUser string `xml:"defaultCaseUser,omitempty"`

	EmailActionDefaultsHandlerClass string `xml:"emailActionDefaultsHandlerClass,omitempty"`

	EmailToCase *EmailToCaseSettings `xml:"emailToCase,omitempty"`

	EnableCaseFeed bool `xml:"enableCaseFeed,omitempty"`

	EnableDraftEmails bool `xml:"enableDraftEmails,omitempty"`

	EnableEarlyEscalationRuleTriggers bool `xml:"enableEarlyEscalationRuleTriggers,omitempty"`

	EnableEmailActionDefaultsHandler bool `xml:"enableEmailActionDefaultsHandler,omitempty"`

	EnableSuggestedArticlesApplication bool `xml:"enableSuggestedArticlesApplication,omitempty"`

	EnableSuggestedArticlesCustomerPortal bool `xml:"enableSuggestedArticlesCustomerPortal,omitempty"`

	EnableSuggestedArticlesPartnerPortal bool `xml:"enableSuggestedArticlesPartnerPortal,omitempty"`

	EnableSuggestedSolutions bool `xml:"enableSuggestedSolutions,omitempty"`

	KeepRecordTypeOnAssignmentRule bool `xml:"keepRecordTypeOnAssignmentRule,omitempty"`

	NotifyContactOnCaseComment bool `xml:"notifyContactOnCaseComment,omitempty"`

	NotifyDefaultCaseOwner bool `xml:"notifyDefaultCaseOwner,omitempty"`

	NotifyOwnerOnCaseComment bool `xml:"notifyOwnerOnCaseComment,omitempty"`

	NotifyOwnerOnCaseOwnerChange bool `xml:"notifyOwnerOnCaseOwnerChange,omitempty"`

	ShowFewerCloseActions bool `xml:"showFewerCloseActions,omitempty"`

	UseSystemEmailAddress bool `xml:"useSystemEmailAddress,omitempty"`

	WebToCase *WebToCaseSettings `xml:"webToCase,omitempty"`
}

type FeedItemSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FeedItemSettings"`

	CharacterLimit int32 `xml:"characterLimit,omitempty"`

	CollapseThread bool `xml:"collapseThread,omitempty"`

	DisplayFormat *FeedItemDisplayFormat `xml:"displayFormat,omitempty"`

	FeedItemType *FeedItemType `xml:"feedItemType,omitempty"`
}

type EmailToCaseSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EmailToCaseSettings"`

	EnableEmailToCase bool `xml:"enableEmailToCase,omitempty"`

	EnableHtmlEmail bool `xml:"enableHtmlEmail,omitempty"`

	EnableOnDemandEmailToCase bool `xml:"enableOnDemandEmailToCase,omitempty"`

	EnableThreadIDInBody bool `xml:"enableThreadIDInBody,omitempty"`

	EnableThreadIDInSubject bool `xml:"enableThreadIDInSubject,omitempty"`

	NotifyOwnerOnNewCaseEmail bool `xml:"notifyOwnerOnNewCaseEmail,omitempty"`

	OverEmailLimitAction *EmailToCaseOnFailureActionType `xml:"overEmailLimitAction,omitempty"`

	PreQuoteSignature bool `xml:"preQuoteSignature,omitempty"`

	RoutingAddresses []*EmailToCaseRoutingAddress `xml:"routingAddresses,omitempty"`

	UnauthorizedSenderAction *EmailToCaseOnFailureActionType `xml:"unauthorizedSenderAction,omitempty"`
}

type EmailToCaseRoutingAddress struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EmailToCaseRoutingAddress"`

	AddressType *EmailToCaseRoutingAddressType `xml:"addressType,omitempty"`

	AuthorizedSenders string `xml:"authorizedSenders,omitempty"`

	CaseOrigin string `xml:"caseOrigin,omitempty"`

	CaseOwner string `xml:"caseOwner,omitempty"`

	CaseOwnerType string `xml:"caseOwnerType,omitempty"`

	CasePriority string `xml:"casePriority,omitempty"`

	CreateTask bool `xml:"createTask,omitempty"`

	EmailAddress string `xml:"emailAddress,omitempty"`

	EmailServicesAddress string `xml:"emailServicesAddress,omitempty"`

	IsVerified bool `xml:"isVerified,omitempty"`

	RoutingName string `xml:"routingName,omitempty"`

	SaveEmailHeaders bool `xml:"saveEmailHeaders,omitempty"`

	TaskStatus string `xml:"taskStatus,omitempty"`
}

type WebToCaseSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WebToCaseSettings"`

	CaseOrigin string `xml:"caseOrigin,omitempty"`

	DefaultResponseTemplate string `xml:"defaultResponseTemplate,omitempty"`

	EnableWebToCase bool `xml:"enableWebToCase,omitempty"`
}

type ChannelLayout struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ChannelLayout"`

	*Metadata

	EnabledChannels []string `xml:"enabledChannels,omitempty"`

	Label string `xml:"label,omitempty"`

	LayoutItems []*ChannelLayoutItem `xml:"layoutItems,omitempty"`
}

type ChannelLayoutItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ChannelLayoutItem"`

	Field string `xml:"field,omitempty"`
}

type ChatterAnswersSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ChatterAnswersSettings"`

	*Metadata

	EmailFollowersOnBestAnswer bool `xml:"emailFollowersOnBestAnswer,omitempty"`

	EmailFollowersOnReply bool `xml:"emailFollowersOnReply,omitempty"`

	EmailOwnerOnPrivateReply bool `xml:"emailOwnerOnPrivateReply,omitempty"`

	EmailOwnerOnReply bool `xml:"emailOwnerOnReply,omitempty"`

	EnableAnswerViaEmail bool `xml:"enableAnswerViaEmail,omitempty"`

	EnableChatterAnswers bool `xml:"enableChatterAnswers,omitempty"`

	EnableFacebookSSO bool `xml:"enableFacebookSSO,omitempty"`

	EnableInlinePublisher bool `xml:"enableInlinePublisher,omitempty"`

	EnableReputation bool `xml:"enableReputation,omitempty"`

	EnableRichTextEditor bool `xml:"enableRichTextEditor,omitempty"`

	FacebookAuthProvider string `xml:"facebookAuthProvider,omitempty"`

	ShowInPortals bool `xml:"showInPortals,omitempty"`
}

type CleanDataService struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CleanDataService"`

	*Metadata

	CleanRules []*CleanRule `xml:"cleanRules,omitempty"`

	Description string `xml:"description,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	MatchEngine string `xml:"matchEngine,omitempty"`
}

type CleanRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CleanRule"`

	BulkEnabled bool `xml:"bulkEnabled,omitempty"`

	BypassTriggers bool `xml:"bypassTriggers,omitempty"`

	BypassWorkflow bool `xml:"bypassWorkflow,omitempty"`

	Description string `xml:"description,omitempty"`

	DeveloperName string `xml:"developerName,omitempty"`

	FieldMappings []*FieldMapping `xml:"fieldMappings,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	MatchRule string `xml:"matchRule,omitempty"`

	SourceSobjectType string `xml:"sourceSobjectType,omitempty"`

	Status *CleanRuleStatus `xml:"status,omitempty"`

	TargetSobjectType string `xml:"targetSobjectType,omitempty"`
}

type FieldMapping struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FieldMapping"`

	SObjectType string `xml:"SObjectType,omitempty"`

	DeveloperName string `xml:"developerName,omitempty"`

	FieldMappingRows []*FieldMappingRow `xml:"fieldMappingRows,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`
}

type FieldMappingRow struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FieldMappingRow"`

	SObjectType string `xml:"SObjectType,omitempty"`

	FieldMappingFields []*FieldMappingField `xml:"fieldMappingFields,omitempty"`

	FieldName string `xml:"fieldName,omitempty"`

	MappingOperation *MappingOperation `xml:"mappingOperation,omitempty"`
}

type FieldMappingField struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FieldMappingField"`

	DataServiceField string `xml:"dataServiceField,omitempty"`

	DataServiceObjectName string `xml:"dataServiceObjectName,omitempty"`

	Priority int32 `xml:"priority,omitempty"`
}

type Community struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Community"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	ChatterAnswersFacebookSsoUrl string `xml:"chatterAnswersFacebookSsoUrl,omitempty"`

	CommunityFeedPage string `xml:"communityFeedPage,omitempty"`

	DataCategoryName string `xml:"dataCategoryName,omitempty"`

	Description string `xml:"description,omitempty"`

	EmailFooterDocument string `xml:"emailFooterDocument,omitempty"`

	EmailHeaderDocument string `xml:"emailHeaderDocument,omitempty"`

	EmailNotificationUrl string `xml:"emailNotificationUrl,omitempty"`

	EnableChatterAnswers bool `xml:"enableChatterAnswers,omitempty"`

	EnablePrivateQuestions bool `xml:"enablePrivateQuestions,omitempty"`

	ExpertsGroup string `xml:"expertsGroup,omitempty"`

	Portal string `xml:"portal,omitempty"`

	ReputationLevels *ReputationLevels `xml:"reputationLevels,omitempty"`

	ShowInPortal bool `xml:"showInPortal,omitempty"`

	Site string `xml:"site,omitempty"`
}

type ReputationLevels struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReputationLevels"`

	ChatterAnswersReputationLevels []*ChatterAnswersReputationLevel `xml:"chatterAnswersReputationLevels,omitempty"`

	IdeaReputationLevels []*IdeaReputationLevel `xml:"ideaReputationLevels,omitempty"`
}

type ChatterAnswersReputationLevel struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ChatterAnswersReputationLevel"`

	Name string `xml:"name,omitempty"`

	Value int32 `xml:"value,omitempty"`
}

type IdeaReputationLevel struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata IdeaReputationLevel"`

	Name string `xml:"name,omitempty"`

	Value int32 `xml:"value,omitempty"`
}

type CommunityTemplateDefinition struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CommunityTemplateDefinition"`

	*Metadata

	BundlesInfo []*CommunityTemplateBundleInfo `xml:"bundlesInfo,omitempty"`

	Category *CommunityTemplateCategory `xml:"category,omitempty"`

	DefaultThemeDefinition string `xml:"defaultThemeDefinition,omitempty"`

	Description string `xml:"description,omitempty"`

	EnableExtendedCleanUpOnDelete bool `xml:"enableExtendedCleanUpOnDelete,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	PageSetting []*CommunityTemplatePageSetting `xml:"pageSetting,omitempty"`
}

type CommunityTemplateBundleInfo struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CommunityTemplateBundleInfo"`

	Description string `xml:"description,omitempty"`

	Image string `xml:"image,omitempty"`

	Order int32 `xml:"order,omitempty"`

	Title string `xml:"title,omitempty"`

	Type_ *CommunityTemplateBundleInfoType `xml:"type,omitempty"`
}

type CommunityTemplatePageSetting struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CommunityTemplatePageSetting"`

	Page string `xml:"page,omitempty"`
}

type CommunityThemeDefinition struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CommunityThemeDefinition"`

	*Metadata

	Description string `xml:"description,omitempty"`

	EnableExtendedCleanUpOnDelete bool `xml:"enableExtendedCleanUpOnDelete,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	ThemeSetting []*CommunityThemeSetting `xml:"themeSetting,omitempty"`
}

type CommunityThemeSetting struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CommunityThemeSetting"`

	ThemeLayout string `xml:"themeLayout,omitempty"`

	ThemeLayoutType *CommunityThemeLayoutType `xml:"themeLayoutType,omitempty"`
}

type CompactLayout struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CompactLayout"`

	*Metadata

	Fields []string `xml:"fields,omitempty"`

	Label string `xml:"label,omitempty"`
}

type CompanySettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CompanySettings"`

	*Metadata

	FiscalYear *FiscalYearSettings `xml:"fiscalYear,omitempty"`
}

type FiscalYearSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FiscalYearSettings"`

	FiscalYearNameBasedOn string `xml:"fiscalYearNameBasedOn,omitempty"`

	StartMonth string `xml:"startMonth,omitempty"`
}

type ConnectedApp struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ConnectedApp"`

	*Metadata

	Attributes []*ConnectedAppAttribute `xml:"attributes,omitempty"`

	CanvasConfig *ConnectedAppCanvasConfig `xml:"canvasConfig,omitempty"`

	ContactEmail string `xml:"contactEmail,omitempty"`

	ContactPhone string `xml:"contactPhone,omitempty"`

	Description string `xml:"description,omitempty"`

	IconUrl string `xml:"iconUrl,omitempty"`

	InfoUrl string `xml:"infoUrl,omitempty"`

	IpRanges []*ConnectedAppIpRange `xml:"ipRanges,omitempty"`

	Label string `xml:"label,omitempty"`

	LogoUrl string `xml:"logoUrl,omitempty"`

	MobileAppConfig *ConnectedAppMobileDetailConfig `xml:"mobileAppConfig,omitempty"`

	MobileStartUrl string `xml:"mobileStartUrl,omitempty"`

	OauthConfig *ConnectedAppOauthConfig `xml:"oauthConfig,omitempty"`

	Plugin string `xml:"plugin,omitempty"`

	SamlConfig *ConnectedAppSamlConfig `xml:"samlConfig,omitempty"`

	StartUrl string `xml:"startUrl,omitempty"`
}

type ConnectedAppAttribute struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ConnectedAppAttribute"`

	Formula string `xml:"formula,omitempty"`

	Key string `xml:"key,omitempty"`
}

type ConnectedAppCanvasConfig struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ConnectedAppCanvasConfig"`

	AccessMethod *AccessMethod `xml:"accessMethod,omitempty"`

	CanvasUrl string `xml:"canvasUrl,omitempty"`

	LifecycleClass string `xml:"lifecycleClass,omitempty"`

	Locations []*CanvasLocationOptions `xml:"locations,omitempty"`

	Options []*CanvasOptions `xml:"options,omitempty"`

	SamlInitiationMethod *SamlInitiationMethod `xml:"samlInitiationMethod,omitempty"`
}

type ConnectedAppIpRange struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ConnectedAppIpRange"`

	Description string `xml:"description,omitempty"`

	End string `xml:"end,omitempty"`

	Start string `xml:"start,omitempty"`
}

type ConnectedAppMobileDetailConfig struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ConnectedAppMobileDetailConfig"`

	ApplicationBinaryFile []byte `xml:"applicationBinaryFile,omitempty"`

	ApplicationBinaryFileName string `xml:"applicationBinaryFileName,omitempty"`

	ApplicationBundleIdentifier string `xml:"applicationBundleIdentifier,omitempty"`

	ApplicationFileLength int32 `xml:"applicationFileLength,omitempty"`

	ApplicationIconFile string `xml:"applicationIconFile,omitempty"`

	ApplicationIconFileName string `xml:"applicationIconFileName,omitempty"`

	ApplicationInstallUrl string `xml:"applicationInstallUrl,omitempty"`

	DevicePlatform *DevicePlatformType `xml:"devicePlatform,omitempty"`

	DeviceType *DeviceType `xml:"deviceType,omitempty"`

	MinimumOsVersion string `xml:"minimumOsVersion,omitempty"`

	PrivateApp bool `xml:"privateApp,omitempty"`

	Version string `xml:"version,omitempty"`
}

type ConnectedAppOauthConfig struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ConnectedAppOauthConfig"`

	CallbackUrl string `xml:"callbackUrl,omitempty"`

	Certificate string `xml:"certificate,omitempty"`

	ConsumerKey string `xml:"consumerKey,omitempty"`

	ConsumerSecret string `xml:"consumerSecret,omitempty"`

	Scopes []*ConnectedAppOauthAccessScope `xml:"scopes,omitempty"`
}

type ConnectedAppSamlConfig struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ConnectedAppSamlConfig"`

	AcsUrl string `xml:"acsUrl,omitempty"`

	Certificate string `xml:"certificate,omitempty"`

	EncryptionCertificate string `xml:"encryptionCertificate,omitempty"`

	EncryptionType *SamlEncryptionType `xml:"encryptionType,omitempty"`

	EntityUrl string `xml:"entityUrl,omitempty"`

	Issuer string `xml:"issuer,omitempty"`

	SamlNameIdFormat *SamlNameIdFormatType `xml:"samlNameIdFormat,omitempty"`

	SamlSubjectCustomAttr string `xml:"samlSubjectCustomAttr,omitempty"`

	SamlSubjectType *SamlSubjectType `xml:"samlSubjectType,omitempty"`
}

type ContractSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ContractSettings"`

	*Metadata

	AutoCalculateEndDate bool `xml:"autoCalculateEndDate,omitempty"`

	AutoExpirationDelay string `xml:"autoExpirationDelay,omitempty"`

	AutoExpirationRecipient string `xml:"autoExpirationRecipient,omitempty"`

	AutoExpireContracts bool `xml:"autoExpireContracts,omitempty"`

	EnableContractHistoryTracking bool `xml:"enableContractHistoryTracking,omitempty"`

	NotifyOwnersOnContractExpiration bool `xml:"notifyOwnersOnContractExpiration,omitempty"`
}

type CorsWhitelistOrigin struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CorsWhitelistOrigin"`

	*Metadata

	UrlPattern string `xml:"urlPattern,omitempty"`
}

type CustomApplication struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomApplication"`

	*Metadata

	ActionOverrides []*AppActionOverride `xml:"actionOverrides,omitempty"`

	Brand *AppBrand `xml:"brand,omitempty"`

	CustomApplicationComponents *CustomApplicationComponents `xml:"customApplicationComponents,omitempty"`

	DefaultLandingTab string `xml:"defaultLandingTab,omitempty"`

	Description string `xml:"description,omitempty"`

	DetailPageRefreshMethod string `xml:"detailPageRefreshMethod,omitempty"`

	DomainWhitelist *DomainWhitelist `xml:"domainWhitelist,omitempty"`

	EnableCustomizeMyTabs bool `xml:"enableCustomizeMyTabs,omitempty"`

	EnableKeyboardShortcuts bool `xml:"enableKeyboardShortcuts,omitempty"`

	EnableListViewHover bool `xml:"enableListViewHover,omitempty"`

	EnableListViewReskin bool `xml:"enableListViewReskin,omitempty"`

	EnableMultiMonitorComponents bool `xml:"enableMultiMonitorComponents,omitempty"`

	EnablePinTabs bool `xml:"enablePinTabs,omitempty"`

	EnableTabHover bool `xml:"enableTabHover,omitempty"`

	EnableTabLimits bool `xml:"enableTabLimits,omitempty"`

	FooterColor string `xml:"footerColor,omitempty"`

	FormFactors []*FormFactor `xml:"formFactors,omitempty"`

	HeaderColor string `xml:"headerColor,omitempty"`

	IsServiceCloudConsole bool `xml:"isServiceCloudConsole,omitempty"`

	KeyboardShortcuts *KeyboardShortcuts `xml:"keyboardShortcuts,omitempty"`

	Label string `xml:"label,omitempty"`

	ListPlacement *ListPlacement `xml:"listPlacement,omitempty"`

	ListRefreshMethod string `xml:"listRefreshMethod,omitempty"`

	LiveAgentConfig *LiveAgentConfig `xml:"liveAgentConfig,omitempty"`

	Logo string `xml:"logo,omitempty"`

	NavType *NavType `xml:"navType,omitempty"`

	PrimaryTabColor string `xml:"primaryTabColor,omitempty"`

	PushNotifications *PushNotifications `xml:"pushNotifications,omitempty"`

	SaveUserSessions bool `xml:"saveUserSessions,omitempty"`

	Tab []string `xml:"tab,omitempty"`

	TabLimitConfig *TabLimitConfig `xml:"tabLimitConfig,omitempty"`

	UiType *UiType `xml:"uiType,omitempty"`

	UtilityBar string `xml:"utilityBar,omitempty"`

	WorkspaceMappings *WorkspaceMappings `xml:"workspaceMappings,omitempty"`
}

type AppActionOverride struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AppActionOverride"`

	*ActionOverride

	PageOrSobjectType string `xml:"pageOrSobjectType,omitempty"`
}

type ActionOverride struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ActionOverride"`

	ActionName string `xml:"actionName,omitempty"`

	Comment string `xml:"comment,omitempty"`

	Content string `xml:"content,omitempty"`

	FormFactor *FormFactor `xml:"formFactor,omitempty"`

	SkipRecordTypeSelect bool `xml:"skipRecordTypeSelect,omitempty"`

	Type_ *ActionOverrideType `xml:"type,omitempty"`
}

type AppBrand struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AppBrand"`

	FooterColor string `xml:"footerColor,omitempty"`

	HeaderColor string `xml:"headerColor,omitempty"`

	Logo string `xml:"logo,omitempty"`

	LogoVersion int32 `xml:"logoVersion,omitempty"`
}

type CustomApplicationComponents struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomApplicationComponents"`

	Alignment string `xml:"alignment,omitempty"`

	CustomApplicationComponent []string `xml:"customApplicationComponent,omitempty"`
}

type DomainWhitelist struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DomainWhitelist"`

	Domain []string `xml:"domain,omitempty"`
}

type KeyboardShortcuts struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KeyboardShortcuts"`

	CustomShortcut []*CustomShortcut `xml:"customShortcut,omitempty"`

	DefaultShortcut []*DefaultShortcut `xml:"defaultShortcut,omitempty"`
}

type CustomShortcut struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomShortcut"`

	*DefaultShortcut

	Description string `xml:"description,omitempty"`

	EventName string `xml:"eventName,omitempty"`
}

type DefaultShortcut struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DefaultShortcut"`

	Action string `xml:"action,omitempty"`

	Active bool `xml:"active,omitempty"`

	KeyCommand string `xml:"keyCommand,omitempty"`
}

type ListPlacement struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ListPlacement"`

	Height int32 `xml:"height,omitempty"`

	Location string `xml:"location,omitempty"`

	Units string `xml:"units,omitempty"`

	Width int32 `xml:"width,omitempty"`
}

type LiveAgentConfig struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LiveAgentConfig"`

	EnableLiveChat bool `xml:"enableLiveChat,omitempty"`

	OpenNewAccountSubtab bool `xml:"openNewAccountSubtab,omitempty"`

	OpenNewCaseSubtab bool `xml:"openNewCaseSubtab,omitempty"`

	OpenNewContactSubtab bool `xml:"openNewContactSubtab,omitempty"`

	OpenNewLeadSubtab bool `xml:"openNewLeadSubtab,omitempty"`

	OpenNewVFPageSubtab bool `xml:"openNewVFPageSubtab,omitempty"`

	PagesToOpen *PagesToOpen `xml:"pagesToOpen,omitempty"`

	ShowKnowledgeArticles bool `xml:"showKnowledgeArticles,omitempty"`
}

type PagesToOpen struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PagesToOpen"`

	PageToOpen []string `xml:"pageToOpen,omitempty"`
}

type PushNotifications struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PushNotifications"`

	PushNotification []*PushNotification `xml:"pushNotification,omitempty"`
}

type PushNotification struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PushNotification"`

	FieldNames []string `xml:"fieldNames,omitempty"`

	ObjectName string `xml:"objectName,omitempty"`
}

type TabLimitConfig struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata TabLimitConfig"`

	MaxNumberOfPrimaryTabs string `xml:"maxNumberOfPrimaryTabs,omitempty"`

	MaxNumberOfSubTabs string `xml:"maxNumberOfSubTabs,omitempty"`
}

type WorkspaceMappings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkspaceMappings"`

	Mapping []*WorkspaceMapping `xml:"mapping,omitempty"`
}

type WorkspaceMapping struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkspaceMapping"`

	FieldName string `xml:"fieldName,omitempty"`

	Tab string `xml:"tab,omitempty"`
}

type CustomApplicationComponent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomApplicationComponent"`

	*Metadata

	ButtonIconUrl string `xml:"buttonIconUrl,omitempty"`

	ButtonStyle string `xml:"buttonStyle,omitempty"`

	ButtonText string `xml:"buttonText,omitempty"`

	ButtonWidth int32 `xml:"buttonWidth,omitempty"`

	Height int32 `xml:"height,omitempty"`

	IsHeightFixed bool `xml:"isHeightFixed,omitempty"`

	IsHidden bool `xml:"isHidden,omitempty"`

	IsWidthFixed bool `xml:"isWidthFixed,omitempty"`

	VisualforcePage string `xml:"visualforcePage,omitempty"`

	Width int32 `xml:"width,omitempty"`
}

type CustomDataType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomDataType"`

	*Metadata

	CustomDataTypeComponents []*CustomDataTypeComponent `xml:"customDataTypeComponents,omitempty"`

	Description string `xml:"description,omitempty"`

	DisplayFormula string `xml:"displayFormula,omitempty"`

	EditComponentsOnSeparateLines bool `xml:"editComponentsOnSeparateLines,omitempty"`

	Label string `xml:"label,omitempty"`

	RightAligned bool `xml:"rightAligned,omitempty"`

	SupportComponentsInReports bool `xml:"supportComponentsInReports,omitempty"`
}

type CustomDataTypeComponent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomDataTypeComponent"`

	DeveloperSuffix string `xml:"developerSuffix,omitempty"`

	EnforceFieldRequiredness bool `xml:"enforceFieldRequiredness,omitempty"`

	Label string `xml:"label,omitempty"`

	Length int32 `xml:"length,omitempty"`

	Precision int32 `xml:"precision,omitempty"`

	Scale int32 `xml:"scale,omitempty"`

	SortOrder *SortOrder `xml:"sortOrder,omitempty"`

	SortPriority int32 `xml:"sortPriority,omitempty"`

	Type_ *FieldType `xml:"type,omitempty"`
}

type CustomExperience struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomExperience"`

	*Metadata

	AllowInternalUserLogin bool `xml:"allowInternalUserLogin,omitempty"`

	Branding *CustomExperienceBranding `xml:"branding,omitempty"`

	ChangePasswordEmailTemplate string `xml:"changePasswordEmailTemplate,omitempty"`

	EmailFooterLogo string `xml:"emailFooterLogo,omitempty"`

	EmailFooterText string `xml:"emailFooterText,omitempty"`

	EmailSenderAddress string `xml:"emailSenderAddress,omitempty"`

	EmailSenderName string `xml:"emailSenderName,omitempty"`

	EnableErrorPageOverridesForVisualforce bool `xml:"enableErrorPageOverridesForVisualforce,omitempty"`

	ForgotPasswordEmailTemplate string `xml:"forgotPasswordEmailTemplate,omitempty"`

	PicassoSite string `xml:"picassoSite,omitempty"`

	SObjectType string `xml:"sObjectType,omitempty"`

	SendWelcomeEmail bool `xml:"sendWelcomeEmail,omitempty"`

	Site string `xml:"site,omitempty"`

	SiteAsContainerEnabled bool `xml:"siteAsContainerEnabled,omitempty"`

	Tabs *CustomExperienceTabSet `xml:"tabs,omitempty"`

	UrlPathPrefix string `xml:"urlPathPrefix,omitempty"`

	WelcomeEmailTemplate string `xml:"welcomeEmailTemplate,omitempty"`
}

type CustomExperienceBranding struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomExperienceBranding"`

	LoginFooterText string `xml:"loginFooterText,omitempty"`

	LoginLogo string `xml:"loginLogo,omitempty"`

	PageFooter string `xml:"pageFooter,omitempty"`

	PageHeader string `xml:"pageHeader,omitempty"`

	PrimaryColor string `xml:"primaryColor,omitempty"`

	PrimaryComplementColor string `xml:"primaryComplementColor,omitempty"`

	QuaternaryColor string `xml:"quaternaryColor,omitempty"`

	QuaternaryComplementColor string `xml:"quaternaryComplementColor,omitempty"`

	SecondaryColor string `xml:"secondaryColor,omitempty"`

	TertiaryColor string `xml:"tertiaryColor,omitempty"`

	TertiaryComplementColor string `xml:"tertiaryComplementColor,omitempty"`

	ZeronaryColor string `xml:"zeronaryColor,omitempty"`

	ZeronaryComplementColor string `xml:"zeronaryComplementColor,omitempty"`
}

type CustomExperienceTabSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomExperienceTabSet"`

	CustomTab []string `xml:"customTab,omitempty"`

	DefaultTab string `xml:"defaultTab,omitempty"`

	StandardTab []string `xml:"standardTab,omitempty"`
}

type CustomFeedFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomFeedFilter"`

	*Metadata

	Criteria []*FeedFilterCriterion `xml:"criteria,omitempty"`

	Description string `xml:"description,omitempty"`

	IsProtected bool `xml:"isProtected,omitempty"`

	Label string `xml:"label,omitempty"`
}

type FeedFilterCriterion struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FeedFilterCriterion"`

	FeedItemType *FeedItemType `xml:"feedItemType,omitempty"`

	FeedItemVisibility *FeedItemVisibility `xml:"feedItemVisibility,omitempty"`

	RelatedSObjectType string `xml:"relatedSObjectType,omitempty"`
}

type CustomField struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomField"`

	*Metadata

	CaseSensitive bool `xml:"caseSensitive,omitempty"`

	CustomDataType string `xml:"customDataType,omitempty"`

	DefaultValue string `xml:"defaultValue,omitempty"`

	DeleteConstraint *DeleteConstraint `xml:"deleteConstraint,omitempty"`

	Deprecated bool `xml:"deprecated,omitempty"`

	Description string `xml:"description,omitempty"`

	DisplayFormat string `xml:"displayFormat,omitempty"`

	Encrypted bool `xml:"encrypted,omitempty"`

	EscapeMarkup bool `xml:"escapeMarkup,omitempty"`

	ExternalDeveloperName string `xml:"externalDeveloperName,omitempty"`

	ExternalId bool `xml:"externalId,omitempty"`

	FieldManageability *FieldManageability `xml:"fieldManageability,omitempty"`

	Formula string `xml:"formula,omitempty"`

	FormulaTreatBlanksAs *TreatBlanksAs `xml:"formulaTreatBlanksAs,omitempty"`

	InlineHelpText string `xml:"inlineHelpText,omitempty"`

	IsConvertLeadDisabled bool `xml:"isConvertLeadDisabled,omitempty"`

	IsFilteringDisabled bool `xml:"isFilteringDisabled,omitempty"`

	IsNameField bool `xml:"isNameField,omitempty"`

	IsSortingDisabled bool `xml:"isSortingDisabled,omitempty"`

	Label string `xml:"label,omitempty"`

	Length int32 `xml:"length,omitempty"`

	LookupFilter *LookupFilter `xml:"lookupFilter,omitempty"`

	MaskChar *EncryptedFieldMaskChar `xml:"maskChar,omitempty"`

	MaskType *EncryptedFieldMaskType `xml:"maskType,omitempty"`

	Picklist *Picklist `xml:"picklist,omitempty"`

	PopulateExistingRows bool `xml:"populateExistingRows,omitempty"`

	Precision int32 `xml:"precision,omitempty"`

	ReferenceTargetField string `xml:"referenceTargetField,omitempty"`

	ReferenceTo string `xml:"referenceTo,omitempty"`

	RelationshipLabel string `xml:"relationshipLabel,omitempty"`

	RelationshipName string `xml:"relationshipName,omitempty"`

	RelationshipOrder int32 `xml:"relationshipOrder,omitempty"`

	ReparentableMasterDetail bool `xml:"reparentableMasterDetail,omitempty"`

	Required bool `xml:"required,omitempty"`

	RestrictedAdminField bool `xml:"restrictedAdminField,omitempty"`

	Scale int32 `xml:"scale,omitempty"`

	StartingNumber int32 `xml:"startingNumber,omitempty"`

	StripMarkup bool `xml:"stripMarkup,omitempty"`

	SummarizedField string `xml:"summarizedField,omitempty"`

	SummaryFilterItems []*FilterItem `xml:"summaryFilterItems,omitempty"`

	SummaryForeignKey string `xml:"summaryForeignKey,omitempty"`

	SummaryOperation *SummaryOperations `xml:"summaryOperation,omitempty"`

	TrackFeedHistory bool `xml:"trackFeedHistory,omitempty"`

	TrackHistory bool `xml:"trackHistory,omitempty"`

	TrackTrending bool `xml:"trackTrending,omitempty"`

	Type_ *FieldType `xml:"type,omitempty"`

	Unique bool `xml:"unique,omitempty"`

	ValueSet *ValueSet `xml:"valueSet,omitempty"`

	VisibleLines int32 `xml:"visibleLines,omitempty"`

	WriteRequiresMasterRead bool `xml:"writeRequiresMasterRead,omitempty"`
}

type LookupFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LookupFilter"`

	Active bool `xml:"active,omitempty"`

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	Description string `xml:"description,omitempty"`

	ErrorMessage string `xml:"errorMessage,omitempty"`

	FilterItems []*FilterItem `xml:"filterItems,omitempty"`

	InfoMessage string `xml:"infoMessage,omitempty"`

	IsOptional bool `xml:"isOptional,omitempty"`
}

type Picklist struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Picklist"`

	ControllingField string `xml:"controllingField,omitempty"`

	PicklistValues []*PicklistValue `xml:"picklistValues,omitempty"`

	RestrictedPicklist bool `xml:"restrictedPicklist,omitempty"`

	Sorted bool `xml:"sorted,omitempty"`
}

type ValueSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ValueSet"`

	ControllingField string `xml:"controllingField,omitempty"`

	Restricted bool `xml:"restricted,omitempty"`

	ValueSetDefinition *ValueSetValuesDefinition `xml:"valueSetDefinition,omitempty"`

	ValueSetName string `xml:"valueSetName,omitempty"`

	ValueSettings []*ValueSettings `xml:"valueSettings,omitempty"`
}

type ValueSetValuesDefinition struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ValueSetValuesDefinition"`

	Sorted bool `xml:"sorted,omitempty"`

	Value []*CustomValue `xml:"value,omitempty"`
}

type CustomValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomValue"`

	*Metadata

	Color string `xml:"color,omitempty"`

	Default_ bool `xml:"default,omitempty"`

	Description string `xml:"description,omitempty"`

	IsActive bool `xml:"isActive,omitempty"`
}

type StandardValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata StandardValue"`

	*CustomValue

	AllowEmail bool `xml:"allowEmail,omitempty"`

	Closed bool `xml:"closed,omitempty"`

	Converted bool `xml:"converted,omitempty"`

	CssExposed bool `xml:"cssExposed,omitempty"`

	ForecastCategory *ForecastCategories `xml:"forecastCategory,omitempty"`

	HighPriority bool `xml:"highPriority,omitempty"`

	Probability int32 `xml:"probability,omitempty"`

	ReverseRole string `xml:"reverseRole,omitempty"`

	Reviewed bool `xml:"reviewed,omitempty"`

	Won bool `xml:"won,omitempty"`
}

type ValueSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ValueSettings"`

	ControllingFieldValue []string `xml:"controllingFieldValue,omitempty"`

	ValueName string `xml:"valueName,omitempty"`
}

type CustomLabel struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomLabel"`

	*Metadata

	Categories string `xml:"categories,omitempty"`

	Language string `xml:"language,omitempty"`

	Protected bool `xml:"protected,omitempty"`

	ShortDescription string `xml:"shortDescription,omitempty"`

	Value string `xml:"value,omitempty"`
}

type CustomLabels struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomLabels"`

	*Metadata

	Labels []*CustomLabel `xml:"labels,omitempty"`
}

type CustomMetadata struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomMetadata"`

	*Metadata

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	Protected bool `xml:"protected,omitempty"`

	Values []*CustomMetadataValue `xml:"values,omitempty"`
}

type CustomMetadataValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomMetadataValue"`

	Field string `xml:"field,omitempty"`

	Value interface{} `xml:"value,omitempty"`
}

type CustomObject struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomObject"`

	*Metadata

	ActionOverrides []*ActionOverride `xml:"actionOverrides,omitempty"`

	AllowInChatterGroups bool `xml:"allowInChatterGroups,omitempty"`

	ArticleTypeChannelDisplay *ArticleTypeChannelDisplay `xml:"articleTypeChannelDisplay,omitempty"`

	BusinessProcesses []*BusinessProcess `xml:"businessProcesses,omitempty"`

	CompactLayoutAssignment string `xml:"compactLayoutAssignment,omitempty"`

	CompactLayouts []*CompactLayout `xml:"compactLayouts,omitempty"`

	CustomHelp string `xml:"customHelp,omitempty"`

	CustomHelpPage string `xml:"customHelpPage,omitempty"`

	CustomSettingsType *CustomSettingsType `xml:"customSettingsType,omitempty"`

	DeploymentStatus *DeploymentStatus `xml:"deploymentStatus,omitempty"`

	Deprecated bool `xml:"deprecated,omitempty"`

	Description string `xml:"description,omitempty"`

	EnableActivities bool `xml:"enableActivities,omitempty"`

	EnableBulkApi bool `xml:"enableBulkApi,omitempty"`

	EnableDivisions bool `xml:"enableDivisions,omitempty"`

	EnableEnhancedLookup bool `xml:"enableEnhancedLookup,omitempty"`

	EnableFeeds bool `xml:"enableFeeds,omitempty"`

	EnableHistory bool `xml:"enableHistory,omitempty"`

	EnableReports bool `xml:"enableReports,omitempty"`

	EnableSearch bool `xml:"enableSearch,omitempty"`

	EnableSharing bool `xml:"enableSharing,omitempty"`

	EnableStreamingApi bool `xml:"enableStreamingApi,omitempty"`

	ExternalDataSource string `xml:"externalDataSource,omitempty"`

	ExternalName string `xml:"externalName,omitempty"`

	ExternalRepository string `xml:"externalRepository,omitempty"`

	ExternalSharingModel *SharingModel `xml:"externalSharingModel,omitempty"`

	FieldSets []*FieldSet `xml:"fieldSets,omitempty"`

	Fields []*CustomField `xml:"fields,omitempty"`

	Gender *Gender `xml:"gender,omitempty"`

	HistoryRetentionPolicy *HistoryRetentionPolicy `xml:"historyRetentionPolicy,omitempty"`

	Household bool `xml:"household,omitempty"`

	Label string `xml:"label,omitempty"`

	ListViews []*ListView `xml:"listViews,omitempty"`

	NameField *CustomField `xml:"nameField,omitempty"`

	PluralLabel string `xml:"pluralLabel,omitempty"`

	RecordTypeTrackFeedHistory bool `xml:"recordTypeTrackFeedHistory,omitempty"`

	RecordTypeTrackHistory bool `xml:"recordTypeTrackHistory,omitempty"`

	RecordTypes []*RecordType `xml:"recordTypes,omitempty"`

	SearchLayouts *SearchLayouts `xml:"searchLayouts,omitempty"`

	SharingModel *SharingModel `xml:"sharingModel,omitempty"`

	SharingReasons []*SharingReason `xml:"sharingReasons,omitempty"`

	SharingRecalculations []*SharingRecalculation `xml:"sharingRecalculations,omitempty"`

	StartsWith *StartsWith `xml:"startsWith,omitempty"`

	ValidationRules []*ValidationRule `xml:"validationRules,omitempty"`

	Visibility *SetupObjectVisibility `xml:"visibility,omitempty"`

	WebLinks []*WebLink `xml:"webLinks,omitempty"`
}

type ArticleTypeChannelDisplay struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ArticleTypeChannelDisplay"`

	ArticleTypeTemplates []*ArticleTypeTemplate `xml:"articleTypeTemplates,omitempty"`
}

type ArticleTypeTemplate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ArticleTypeTemplate"`

	Channel *Channel `xml:"channel,omitempty"`

	Page string `xml:"page,omitempty"`

	Template *Template `xml:"template,omitempty"`
}

type FieldSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FieldSet"`

	*Metadata

	AvailableFields []*FieldSetItem `xml:"availableFields,omitempty"`

	Description string `xml:"description,omitempty"`

	DisplayedFields []*FieldSetItem `xml:"displayedFields,omitempty"`

	Label string `xml:"label,omitempty"`
}

type FieldSetItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FieldSetItem"`

	Field string `xml:"field,omitempty"`

	IsFieldManaged bool `xml:"isFieldManaged,omitempty"`

	IsRequired bool `xml:"isRequired,omitempty"`
}

type HistoryRetentionPolicy struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata HistoryRetentionPolicy"`

	ArchiveAfterMonths int32 `xml:"archiveAfterMonths,omitempty"`

	ArchiveRetentionYears int32 `xml:"archiveRetentionYears,omitempty"`

	Description string `xml:"description,omitempty"`
}

type ListView struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ListView"`

	*Metadata

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	Columns []string `xml:"columns,omitempty"`

	Division string `xml:"division,omitempty"`

	FilterScope *FilterScope `xml:"filterScope,omitempty"`

	Filters []*ListViewFilter `xml:"filters,omitempty"`

	Label string `xml:"label,omitempty"`

	Language *Language `xml:"language,omitempty"`

	Queue string `xml:"queue,omitempty"`

	SharedTo *SharedTo `xml:"sharedTo,omitempty"`
}

type ListViewFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ListViewFilter"`

	Field string `xml:"field,omitempty"`

	Operation *FilterOperation `xml:"operation,omitempty"`

	Value string `xml:"value,omitempty"`
}

type SharedTo struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharedTo"`

	AllCustomerPortalUsers string `xml:"allCustomerPortalUsers,omitempty"`

	AllInternalUsers string `xml:"allInternalUsers,omitempty"`

	AllPartnerUsers string `xml:"allPartnerUsers,omitempty"`

	Group []string `xml:"group,omitempty"`

	Groups []string `xml:"groups,omitempty"`

	ManagerSubordinates []string `xml:"managerSubordinates,omitempty"`

	Managers []string `xml:"managers,omitempty"`

	PortalRole []string `xml:"portalRole,omitempty"`

	PortalRoleAndSubordinates []string `xml:"portalRoleAndSubordinates,omitempty"`

	Queue []string `xml:"queue,omitempty"`

	Role []string `xml:"role,omitempty"`

	RoleAndSubordinates []string `xml:"roleAndSubordinates,omitempty"`

	RoleAndSubordinatesInternal []string `xml:"roleAndSubordinatesInternal,omitempty"`

	Roles []string `xml:"roles,omitempty"`

	RolesAndSubordinates []string `xml:"rolesAndSubordinates,omitempty"`

	Territories []string `xml:"territories,omitempty"`

	TerritoriesAndSubordinates []string `xml:"territoriesAndSubordinates,omitempty"`

	Territory []string `xml:"territory,omitempty"`

	TerritoryAndSubordinates []string `xml:"territoryAndSubordinates,omitempty"`
}

type RecordType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RecordType"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	BusinessProcess string `xml:"businessProcess,omitempty"`

	CompactLayoutAssignment string `xml:"compactLayoutAssignment,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	PicklistValues []*RecordTypePicklistValue `xml:"picklistValues,omitempty"`
}

type RecordTypePicklistValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RecordTypePicklistValue"`

	Picklist string `xml:"picklist,omitempty"`

	Values []*PicklistValue `xml:"values,omitempty"`
}

type SearchLayouts struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SearchLayouts"`

	CustomTabListAdditionalFields []string `xml:"customTabListAdditionalFields,omitempty"`

	ExcludedStandardButtons []string `xml:"excludedStandardButtons,omitempty"`

	ListViewButtons []string `xml:"listViewButtons,omitempty"`

	LookupDialogsAdditionalFields []string `xml:"lookupDialogsAdditionalFields,omitempty"`

	LookupFilterFields []string `xml:"lookupFilterFields,omitempty"`

	LookupPhoneDialogsAdditionalFields []string `xml:"lookupPhoneDialogsAdditionalFields,omitempty"`

	SearchFilterFields []string `xml:"searchFilterFields,omitempty"`

	SearchResultsAdditionalFields []string `xml:"searchResultsAdditionalFields,omitempty"`

	SearchResultsCustomButtons []string `xml:"searchResultsCustomButtons,omitempty"`
}

type SharingReason struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharingReason"`

	*Metadata

	Label string `xml:"label,omitempty"`
}

type SharingRecalculation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharingRecalculation"`

	ClassName string `xml:"className,omitempty"`
}

type ValidationRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ValidationRule"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	Description string `xml:"description,omitempty"`

	ErrorConditionFormula string `xml:"errorConditionFormula,omitempty"`

	ErrorDisplayField string `xml:"errorDisplayField,omitempty"`

	ErrorMessage string `xml:"errorMessage,omitempty"`
}

type WebLink struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WebLink"`

	*Metadata

	Availability *WebLinkAvailability `xml:"availability,omitempty"`

	Description string `xml:"description,omitempty"`

	DisplayType *WebLinkDisplayType `xml:"displayType,omitempty"`

	EncodingKey *Encoding `xml:"encodingKey,omitempty"`

	HasMenubar bool `xml:"hasMenubar,omitempty"`

	HasScrollbars bool `xml:"hasScrollbars,omitempty"`

	HasToolbar bool `xml:"hasToolbar,omitempty"`

	Height int32 `xml:"height,omitempty"`

	IsResizable bool `xml:"isResizable,omitempty"`

	LinkType *WebLinkType `xml:"linkType,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	OpenType *WebLinkWindowType `xml:"openType,omitempty"`

	Page string `xml:"page,omitempty"`

	Position *WebLinkPosition `xml:"position,omitempty"`

	Protected bool `xml:"protected,omitempty"`

	RequireRowSelection bool `xml:"requireRowSelection,omitempty"`

	Scontrol string `xml:"scontrol,omitempty"`

	ShowsLocation bool `xml:"showsLocation,omitempty"`

	ShowsStatus bool `xml:"showsStatus,omitempty"`

	Url string `xml:"url,omitempty"`

	Width int32 `xml:"width,omitempty"`
}

type CustomObjectTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomObjectTranslation"`

	*Metadata

	CaseValues []*ObjectNameCaseValue `xml:"caseValues,omitempty"`

	Fields []*CustomFieldTranslation `xml:"fields,omitempty"`

	Gender *Gender `xml:"gender,omitempty"`

	Layouts []*LayoutTranslation `xml:"layouts,omitempty"`

	NameFieldLabel string `xml:"nameFieldLabel,omitempty"`

	QuickActions []*QuickActionTranslation `xml:"quickActions,omitempty"`

	RecordTypes []*RecordTypeTranslation `xml:"recordTypes,omitempty"`

	SharingReasons []*SharingReasonTranslation `xml:"sharingReasons,omitempty"`

	StandardFields []*StandardFieldTranslation `xml:"standardFields,omitempty"`

	StartsWith *StartsWith `xml:"startsWith,omitempty"`

	ValidationRules []*ValidationRuleTranslation `xml:"validationRules,omitempty"`

	WebLinks []*WebLinkTranslation `xml:"webLinks,omitempty"`

	WorkflowTasks []*WorkflowTaskTranslation `xml:"workflowTasks,omitempty"`
}

type ObjectNameCaseValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ObjectNameCaseValue"`

	Article *Article `xml:"article,omitempty"`

	CaseType *CaseType `xml:"caseType,omitempty"`

	Plural bool `xml:"plural,omitempty"`

	Possessive *Possessive `xml:"possessive,omitempty"`

	Value string `xml:"value,omitempty"`
}

type CustomFieldTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomFieldTranslation"`

	CaseValues []*ObjectNameCaseValue `xml:"caseValues,omitempty"`

	Gender *Gender `xml:"gender,omitempty"`

	Help string `xml:"help,omitempty"`

	Label string `xml:"label,omitempty"`

	LookupFilter *LookupFilterTranslation `xml:"lookupFilter,omitempty"`

	Name string `xml:"name,omitempty"`

	PicklistValues []*PicklistValueTranslation `xml:"picklistValues,omitempty"`

	RelationshipLabel string `xml:"relationshipLabel,omitempty"`

	StartsWith *StartsWith `xml:"startsWith,omitempty"`
}

type LookupFilterTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LookupFilterTranslation"`

	ErrorMessage string `xml:"errorMessage,omitempty"`

	InformationalMessage string `xml:"informationalMessage,omitempty"`
}

type PicklistValueTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PicklistValueTranslation"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	Translation string `xml:"translation,omitempty"`
}

type LayoutTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LayoutTranslation"`

	Layout string `xml:"layout,omitempty"`

	LayoutType string `xml:"layoutType,omitempty"`

	Sections []*LayoutSectionTranslation `xml:"sections,omitempty"`
}

type LayoutSectionTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LayoutSectionTranslation"`

	Label string `xml:"label,omitempty"`

	Section string `xml:"section,omitempty"`
}

type QuickActionTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuickActionTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type RecordTypeTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RecordTypeTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type SharingReasonTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharingReasonTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type StandardFieldTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata StandardFieldTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ValidationRuleTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ValidationRuleTranslation"`

	ErrorMessage string `xml:"errorMessage,omitempty"`

	Name string `xml:"name,omitempty"`
}

type WebLinkTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WebLinkTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type WorkflowTaskTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowTaskTranslation"`

	Description string `xml:"description,omitempty"`

	Name string `xml:"name,omitempty"`

	Subject string `xml:"subject,omitempty"`
}

type CustomPageWebLink struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomPageWebLink"`

	*Metadata

	Availability *WebLinkAvailability `xml:"availability,omitempty"`

	Description string `xml:"description,omitempty"`

	DisplayType *WebLinkDisplayType `xml:"displayType,omitempty"`

	EncodingKey *Encoding `xml:"encodingKey,omitempty"`

	HasMenubar bool `xml:"hasMenubar,omitempty"`

	HasScrollbars bool `xml:"hasScrollbars,omitempty"`

	HasToolbar bool `xml:"hasToolbar,omitempty"`

	Height int32 `xml:"height,omitempty"`

	IsResizable bool `xml:"isResizable,omitempty"`

	LinkType *WebLinkType `xml:"linkType,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	OpenType *WebLinkWindowType `xml:"openType,omitempty"`

	Page string `xml:"page,omitempty"`

	Position *WebLinkPosition `xml:"position,omitempty"`

	Protected bool `xml:"protected,omitempty"`

	RequireRowSelection bool `xml:"requireRowSelection,omitempty"`

	Scontrol string `xml:"scontrol,omitempty"`

	ShowsLocation bool `xml:"showsLocation,omitempty"`

	ShowsStatus bool `xml:"showsStatus,omitempty"`

	Url string `xml:"url,omitempty"`

	Width int32 `xml:"width,omitempty"`
}

type CustomPermission struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomPermission"`

	*Metadata

	ConnectedApp string `xml:"connectedApp,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	RequiredPermission []*CustomPermissionDependencyRequired `xml:"requiredPermission,omitempty"`
}

type CustomPermissionDependencyRequired struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomPermissionDependencyRequired"`

	CustomPermission string `xml:"customPermission,omitempty"`

	Dependency bool `xml:"dependency,omitempty"`
}

type CustomSite struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomSite"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	AllowHomePage bool `xml:"allowHomePage,omitempty"`

	AllowStandardAnswersPages bool `xml:"allowStandardAnswersPages,omitempty"`

	AllowStandardIdeasPages bool `xml:"allowStandardIdeasPages,omitempty"`

	AllowStandardLookups bool `xml:"allowStandardLookups,omitempty"`

	AllowStandardSearch bool `xml:"allowStandardSearch,omitempty"`

	AnalyticsTrackingCode string `xml:"analyticsTrackingCode,omitempty"`

	AuthorizationRequiredPage string `xml:"authorizationRequiredPage,omitempty"`

	BandwidthExceededPage string `xml:"bandwidthExceededPage,omitempty"`

	ChangePasswordPage string `xml:"changePasswordPage,omitempty"`

	ChatterAnswersForgotPasswordConfirmPage string `xml:"chatterAnswersForgotPasswordConfirmPage,omitempty"`

	ChatterAnswersForgotPasswordPage string `xml:"chatterAnswersForgotPasswordPage,omitempty"`

	ChatterAnswersHelpPage string `xml:"chatterAnswersHelpPage,omitempty"`

	ChatterAnswersLoginPage string `xml:"chatterAnswersLoginPage,omitempty"`

	ChatterAnswersRegistrationPage string `xml:"chatterAnswersRegistrationPage,omitempty"`

	ClickjackProtectionLevel *SiteClickjackProtectionLevel `xml:"clickjackProtectionLevel,omitempty"`

	CustomWebAddresses []*SiteWebAddress `xml:"customWebAddresses,omitempty"`

	Description string `xml:"description,omitempty"`

	FavoriteIcon string `xml:"favoriteIcon,omitempty"`

	FileNotFoundPage string `xml:"fileNotFoundPage,omitempty"`

	ForgotPasswordPage string `xml:"forgotPasswordPage,omitempty"`

	GenericErrorPage string `xml:"genericErrorPage,omitempty"`

	GuestProfile string `xml:"guestProfile,omitempty"`

	InMaintenancePage string `xml:"inMaintenancePage,omitempty"`

	InactiveIndexPage string `xml:"inactiveIndexPage,omitempty"`

	IndexPage string `xml:"indexPage,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	MyProfilePage string `xml:"myProfilePage,omitempty"`

	Portal string `xml:"portal,omitempty"`

	RequireHttps bool `xml:"requireHttps,omitempty"`

	RequireInsecurePortalAccess bool `xml:"requireInsecurePortalAccess,omitempty"`

	RobotsTxtPage string `xml:"robotsTxtPage,omitempty"`

	RootComponent string `xml:"rootComponent,omitempty"`

	SelfRegPage string `xml:"selfRegPage,omitempty"`

	ServerIsDown string `xml:"serverIsDown,omitempty"`

	SiteAdmin string `xml:"siteAdmin,omitempty"`

	SiteRedirectMappings []*SiteRedirectMapping `xml:"siteRedirectMappings,omitempty"`

	SiteTemplate string `xml:"siteTemplate,omitempty"`

	SiteType *SiteType `xml:"siteType,omitempty"`

	Subdomain string `xml:"subdomain,omitempty"`

	UrlPathPrefix string `xml:"urlPathPrefix,omitempty"`
}

type SiteWebAddress struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SiteWebAddress"`

	Certificate string `xml:"certificate,omitempty"`

	DomainName string `xml:"domainName,omitempty"`

	Primary bool `xml:"primary,omitempty"`
}

type SiteRedirectMapping struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SiteRedirectMapping"`

	Action *SiteRedirect `xml:"action,omitempty"`

	IsActive bool `xml:"isActive,omitempty"`

	Source string `xml:"source,omitempty"`

	Target string `xml:"target,omitempty"`
}

type CustomTab struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomTab"`

	*Metadata

	ActionOverrides []*ActionOverride `xml:"actionOverrides,omitempty"`

	AuraComponent string `xml:"auraComponent,omitempty"`

	CustomObject bool `xml:"customObject,omitempty"`

	Description string `xml:"description,omitempty"`

	FlexiPage string `xml:"flexiPage,omitempty"`

	FrameHeight int32 `xml:"frameHeight,omitempty"`

	HasSidebar bool `xml:"hasSidebar,omitempty"`

	Icon string `xml:"icon,omitempty"`

	Label string `xml:"label,omitempty"`

	MobileReady bool `xml:"mobileReady,omitempty"`

	Motif string `xml:"motif,omitempty"`

	Page string `xml:"page,omitempty"`

	Scontrol string `xml:"scontrol,omitempty"`

	SplashPageLink string `xml:"splashPageLink,omitempty"`

	Url string `xml:"url,omitempty"`

	UrlEncodingKey *Encoding `xml:"urlEncodingKey,omitempty"`
}

type Dashboard struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Dashboard"`

	*Metadata

	BackgroundEndColor string `xml:"backgroundEndColor,omitempty"`

	BackgroundFadeDirection *ChartBackgroundDirection `xml:"backgroundFadeDirection,omitempty"`

	BackgroundStartColor string `xml:"backgroundStartColor,omitempty"`

	DashboardFilters []*DashboardFilter `xml:"dashboardFilters,omitempty"`

	DashboardGridLayout *DashboardGridLayout `xml:"dashboardGridLayout,omitempty"`

	DashboardResultRefreshedDate string `xml:"dashboardResultRefreshedDate,omitempty"`

	DashboardResultRunningUser string `xml:"dashboardResultRunningUser,omitempty"`

	DashboardType *DashboardType `xml:"dashboardType,omitempty"`

	Description string `xml:"description,omitempty"`

	FolderName string `xml:"folderName,omitempty"`

	IsGridLayout bool `xml:"isGridLayout,omitempty"`

	LeftSection *DashboardComponentSection `xml:"leftSection,omitempty"`

	MiddleSection *DashboardComponentSection `xml:"middleSection,omitempty"`

	RightSection *DashboardComponentSection `xml:"rightSection,omitempty"`

	RunningUser string `xml:"runningUser,omitempty"`

	TextColor string `xml:"textColor,omitempty"`

	Title string `xml:"title,omitempty"`

	TitleColor string `xml:"titleColor,omitempty"`

	TitleSize int32 `xml:"titleSize,omitempty"`
}

type DashboardFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardFilter"`

	DashboardFilterOptions []*DashboardFilterOption `xml:"dashboardFilterOptions,omitempty"`

	Name string `xml:"name,omitempty"`
}

type DashboardFilterOption struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardFilterOption"`

	Operator *DashboardFilterOperation `xml:"operator,omitempty"`

	Values []string `xml:"values,omitempty"`
}

type DashboardGridLayout struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardGridLayout"`

	DashboardGridComponents []*DashboardGridComponent `xml:"dashboardGridComponents,omitempty"`

	NumberOfColumns int32 `xml:"numberOfColumns,omitempty"`

	RowHeight int32 `xml:"rowHeight,omitempty"`
}

type DashboardGridComponent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardGridComponent"`

	ColSpan int32 `xml:"colSpan,omitempty"`

	ColumnIndex int32 `xml:"columnIndex,omitempty"`

	DashboardComponent *DashboardComponent `xml:"dashboardComponent,omitempty"`

	RowIndex int32 `xml:"rowIndex,omitempty"`

	RowSpan int32 `xml:"rowSpan,omitempty"`
}

type DashboardComponent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardComponent"`

	AutoselectColumnsFromReport bool `xml:"autoselectColumnsFromReport,omitempty"`

	ChartAxisRange *ChartRangeType `xml:"chartAxisRange,omitempty"`

	ChartAxisRangeMax float64 `xml:"chartAxisRangeMax,omitempty"`

	ChartAxisRangeMin float64 `xml:"chartAxisRangeMin,omitempty"`

	ChartSummary []*ChartSummary `xml:"chartSummary,omitempty"`

	ComponentType *DashboardComponentType `xml:"componentType,omitempty"`

	DashboardFilterColumns []*DashboardFilterColumn `xml:"dashboardFilterColumns,omitempty"`

	DashboardTableColumn []*DashboardTableColumn `xml:"dashboardTableColumn,omitempty"`

	DisplayUnits *ChartUnits `xml:"displayUnits,omitempty"`

	DrillDownUrl string `xml:"drillDownUrl,omitempty"`

	DrillEnabled bool `xml:"drillEnabled,omitempty"`

	DrillToDetailEnabled bool `xml:"drillToDetailEnabled,omitempty"`

	EnableHover bool `xml:"enableHover,omitempty"`

	ExpandOthers bool `xml:"expandOthers,omitempty"`

	Footer string `xml:"footer,omitempty"`

	GaugeMax float64 `xml:"gaugeMax,omitempty"`

	GaugeMin float64 `xml:"gaugeMin,omitempty"`

	GroupingColumn []string `xml:"groupingColumn,omitempty"`

	Header string `xml:"header,omitempty"`

	IndicatorBreakpoint1 float64 `xml:"indicatorBreakpoint1,omitempty"`

	IndicatorBreakpoint2 float64 `xml:"indicatorBreakpoint2,omitempty"`

	IndicatorHighColor string `xml:"indicatorHighColor,omitempty"`

	IndicatorLowColor string `xml:"indicatorLowColor,omitempty"`

	IndicatorMiddleColor string `xml:"indicatorMiddleColor,omitempty"`

	LegendPosition *ChartLegendPosition `xml:"legendPosition,omitempty"`

	MaxValuesDisplayed int32 `xml:"maxValuesDisplayed,omitempty"`

	MetricLabel string `xml:"metricLabel,omitempty"`

	Page string `xml:"page,omitempty"`

	PageHeightInPixels int32 `xml:"pageHeightInPixels,omitempty"`

	Report string `xml:"report,omitempty"`

	Scontrol string `xml:"scontrol,omitempty"`

	ScontrolHeightInPixels int32 `xml:"scontrolHeightInPixels,omitempty"`

	ShowPercentage bool `xml:"showPercentage,omitempty"`

	ShowPicturesOnCharts bool `xml:"showPicturesOnCharts,omitempty"`

	ShowPicturesOnTables bool `xml:"showPicturesOnTables,omitempty"`

	ShowTotal bool `xml:"showTotal,omitempty"`

	ShowValues bool `xml:"showValues,omitempty"`

	SortBy *DashboardComponentFilter `xml:"sortBy,omitempty"`

	Title string `xml:"title,omitempty"`

	UseReportChart bool `xml:"useReportChart,omitempty"`
}

type ChartSummary struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ChartSummary"`

	Aggregate *ReportSummaryType `xml:"aggregate,omitempty"`

	AxisBinding *ChartAxis `xml:"axisBinding,omitempty"`

	Column string `xml:"column,omitempty"`
}

type DashboardFilterColumn struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardFilterColumn"`

	Column string `xml:"column,omitempty"`
}

type DashboardTableColumn struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardTableColumn"`

	AggregateType *ReportSummaryType `xml:"aggregateType,omitempty"`

	CalculatePercent bool `xml:"calculatePercent,omitempty"`

	Column string `xml:"column,omitempty"`

	DecimalPlaces int32 `xml:"decimalPlaces,omitempty"`

	ShowTotal bool `xml:"showTotal,omitempty"`

	SortBy *DashboardComponentFilter `xml:"sortBy,omitempty"`
}

type DashboardComponentSection struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardComponentSection"`

	ColumnSize *DashboardComponentSize `xml:"columnSize,omitempty"`

	Components []*DashboardComponent `xml:"components,omitempty"`
}

type DataCategoryGroup struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DataCategoryGroup"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	DataCategory *DataCategory `xml:"dataCategory,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	ObjectUsage *ObjectUsage `xml:"objectUsage,omitempty"`
}

type DataCategory struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DataCategory"`

	DataCategory []*DataCategory `xml:"dataCategory,omitempty"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ObjectUsage struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ObjectUsage"`

	Object []string `xml:"object,omitempty"`
}

type DelegateGroup struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DelegateGroup"`

	*Metadata

	CustomObjects []string `xml:"customObjects,omitempty"`

	Groups []string `xml:"groups,omitempty"`

	Label string `xml:"label,omitempty"`

	LoginAccess bool `xml:"loginAccess,omitempty"`

	PermissionSets []string `xml:"permissionSets,omitempty"`

	Profiles []string `xml:"profiles,omitempty"`

	Roles []string `xml:"roles,omitempty"`
}

type DuplicateRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DuplicateRule"`

	*Metadata

	ActionOnInsert *DupeActionType `xml:"actionOnInsert,omitempty"`

	ActionOnUpdate *DupeActionType `xml:"actionOnUpdate,omitempty"`

	AlertText string `xml:"alertText,omitempty"`

	Description string `xml:"description,omitempty"`

	DuplicateRuleFilter *DuplicateRuleFilter `xml:"duplicateRuleFilter,omitempty"`

	DuplicateRuleMatchRules []*DuplicateRuleMatchRule `xml:"duplicateRuleMatchRules,omitempty"`

	IsActive bool `xml:"isActive,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	OperationsOnInsert []string `xml:"operationsOnInsert,omitempty"`

	OperationsOnUpdate []string `xml:"operationsOnUpdate,omitempty"`

	SecurityOption *DupeSecurityOptionType `xml:"securityOption,omitempty"`

	SortOrder int32 `xml:"sortOrder,omitempty"`
}

type DuplicateRuleFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DuplicateRuleFilter"`

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	DuplicateRuleFilterItems []*DuplicateRuleFilterItem `xml:"duplicateRuleFilterItems,omitempty"`
}

type DuplicateRuleMatchRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DuplicateRuleMatchRule"`

	MatchRuleSObjectType string `xml:"matchRuleSObjectType,omitempty"`

	MatchingRule string `xml:"matchingRule,omitempty"`

	ObjectMapping *ObjectMapping `xml:"objectMapping,omitempty"`
}

type ObjectMapping struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ObjectMapping"`

	InputObject string `xml:"inputObject,omitempty"`

	MappingFields []*ObjectMappingField `xml:"mappingFields,omitempty"`

	OutputObject string `xml:"outputObject,omitempty"`
}

type ObjectMappingField struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ObjectMappingField"`

	InputField string `xml:"inputField,omitempty"`

	OutputField string `xml:"outputField,omitempty"`
}

type EmbeddedServiceConfig struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EmbeddedServiceConfig"`

	*Metadata

	MasterLabel string `xml:"masterLabel,omitempty"`

	Site string `xml:"site,omitempty"`
}

type EmbeddedServiceLiveAgent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EmbeddedServiceLiveAgent"`

	*Metadata

	EmbeddedServiceConfig string `xml:"embeddedServiceConfig,omitempty"`

	LiveAgentChatUrl string `xml:"liveAgentChatUrl,omitempty"`

	LiveAgentContentUrl string `xml:"liveAgentContentUrl,omitempty"`

	LiveChatButton string `xml:"liveChatButton,omitempty"`

	LiveChatDeployment string `xml:"liveChatDeployment,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`
}

type EntitlementProcess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EntitlementProcess"`

	*Metadata

	SObjectType string `xml:"SObjectType,omitempty"`

	Active bool `xml:"active,omitempty"`

	BusinessHours string `xml:"businessHours,omitempty"`

	Description string `xml:"description,omitempty"`

	EntryStartDateField string `xml:"entryStartDateField,omitempty"`

	ExitCriteriaBooleanFilter string `xml:"exitCriteriaBooleanFilter,omitempty"`

	ExitCriteriaFilterItems []*FilterItem `xml:"exitCriteriaFilterItems,omitempty"`

	ExitCriteriaFormula string `xml:"exitCriteriaFormula,omitempty"`

	IsVersionDefault bool `xml:"isVersionDefault,omitempty"`

	Milestones []*EntitlementProcessMilestoneItem `xml:"milestones,omitempty"`

	Name string `xml:"name,omitempty"`

	VersionMaster string `xml:"versionMaster,omitempty"`

	VersionNotes string `xml:"versionNotes,omitempty"`

	VersionNumber int32 `xml:"versionNumber,omitempty"`
}

type EntitlementProcessMilestoneItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EntitlementProcessMilestoneItem"`

	BusinessHours string `xml:"businessHours,omitempty"`

	CriteriaBooleanFilter string `xml:"criteriaBooleanFilter,omitempty"`

	MilestoneCriteriaFilterItems []*FilterItem `xml:"milestoneCriteriaFilterItems,omitempty"`

	MilestoneCriteriaFormula string `xml:"milestoneCriteriaFormula,omitempty"`

	MilestoneName string `xml:"milestoneName,omitempty"`

	MinutesCustomClass string `xml:"minutesCustomClass,omitempty"`

	MinutesToComplete int32 `xml:"minutesToComplete,omitempty"`

	SuccessActions []*WorkflowActionReference `xml:"successActions,omitempty"`

	TimeTriggers []*EntitlementProcessMilestoneTimeTrigger `xml:"timeTriggers,omitempty"`

	UseCriteriaStartTime bool `xml:"useCriteriaStartTime,omitempty"`
}

type EntitlementProcessMilestoneTimeTrigger struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EntitlementProcessMilestoneTimeTrigger"`

	Actions []*WorkflowActionReference `xml:"actions,omitempty"`

	TimeLength int32 `xml:"timeLength,omitempty"`

	WorkflowTimeTriggerUnit *MilestoneTimeUnits `xml:"workflowTimeTriggerUnit,omitempty"`
}

type EntitlementSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EntitlementSettings"`

	*Metadata

	AssetLookupLimitedToActiveEntitlementsOnAccount bool `xml:"assetLookupLimitedToActiveEntitlementsOnAccount,omitempty"`

	AssetLookupLimitedToActiveEntitlementsOnContact bool `xml:"assetLookupLimitedToActiveEntitlementsOnContact,omitempty"`

	AssetLookupLimitedToSameAccount bool `xml:"assetLookupLimitedToSameAccount,omitempty"`

	AssetLookupLimitedToSameContact bool `xml:"assetLookupLimitedToSameContact,omitempty"`

	EnableEntitlementVersioning bool `xml:"enableEntitlementVersioning,omitempty"`

	EnableEntitlements bool `xml:"enableEntitlements,omitempty"`

	EntitlementLookupLimitedToActiveStatus bool `xml:"entitlementLookupLimitedToActiveStatus,omitempty"`

	EntitlementLookupLimitedToSameAccount bool `xml:"entitlementLookupLimitedToSameAccount,omitempty"`

	EntitlementLookupLimitedToSameAsset bool `xml:"entitlementLookupLimitedToSameAsset,omitempty"`

	EntitlementLookupLimitedToSameContact bool `xml:"entitlementLookupLimitedToSameContact,omitempty"`
}

type EntitlementTemplate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EntitlementTemplate"`

	*Metadata

	BusinessHours string `xml:"businessHours,omitempty"`

	CasesPerEntitlement int32 `xml:"casesPerEntitlement,omitempty"`

	EntitlementProcess string `xml:"entitlementProcess,omitempty"`

	IsPerIncident bool `xml:"isPerIncident,omitempty"`

	Term int32 `xml:"term,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type EscalationRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EscalationRule"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	RuleEntry []*RuleEntry `xml:"ruleEntry,omitempty"`
}

type EscalationRules struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EscalationRules"`

	*Metadata

	EscalationRule []*EscalationRule `xml:"escalationRule,omitempty"`
}

type EventDelivery struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EventDelivery"`

	*Metadata

	EventParameters []*EventParameterMap `xml:"eventParameters,omitempty"`

	EventSubscription string `xml:"eventSubscription,omitempty"`

	ReferenceData string `xml:"referenceData,omitempty"`

	Type_ *EventDeliveryType `xml:"type,omitempty"`
}

type EventParameterMap struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EventParameterMap"`

	ParameterName string `xml:"parameterName,omitempty"`

	ParameterValue string `xml:"parameterValue,omitempty"`
}

type EventSubscription struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EventSubscription"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	EventParameters []*EventParameterMap `xml:"eventParameters,omitempty"`

	EventType string `xml:"eventType,omitempty"`

	ReferenceData string `xml:"referenceData,omitempty"`
}

type EventType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EventType"`

	*Metadata

	Description string `xml:"description,omitempty"`

	Fields []*EventTypeParameter `xml:"fields,omitempty"`

	Label string `xml:"label,omitempty"`
}

type EventTypeParameter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EventTypeParameter"`

	DefaultValue string `xml:"defaultValue,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	MaxOccurs int32 `xml:"maxOccurs,omitempty"`

	MinOccurs int32 `xml:"minOccurs,omitempty"`

	Name string `xml:"name,omitempty"`

	SObjectType string `xml:"sObjectType,omitempty"`

	Type_ *FieldType `xml:"type,omitempty"`
}

type ExternalDataSource struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ExternalDataSource"`

	*Metadata

	AuthProvider string `xml:"authProvider,omitempty"`

	Certificate string `xml:"certificate,omitempty"`

	CustomConfiguration string `xml:"customConfiguration,omitempty"`

	Endpoint string `xml:"endpoint,omitempty"`

	IsWritable bool `xml:"isWritable,omitempty"`

	Label string `xml:"label,omitempty"`

	OauthRefreshToken string `xml:"oauthRefreshToken,omitempty"`

	OauthScope string `xml:"oauthScope,omitempty"`

	OauthToken string `xml:"oauthToken,omitempty"`

	Password string `xml:"password,omitempty"`

	PrincipalType *ExternalPrincipalType `xml:"principalType,omitempty"`

	Protocol *AuthenticationProtocol `xml:"protocol,omitempty"`

	Repository string `xml:"repository,omitempty"`

	Type_ *ExternalDataSourceType `xml:"type,omitempty"`

	Username string `xml:"username,omitempty"`

	Version string `xml:"version,omitempty"`
}

type ExternalServiceRegistration struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ExternalServiceRegistration"`

	*Metadata

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	NamedCredential string `xml:"namedCredential,omitempty"`

	Schema string `xml:"schema,omitempty"`

	SchemaType string `xml:"schemaType,omitempty"`

	SchemaUrl string `xml:"schemaUrl,omitempty"`

	Status string `xml:"status,omitempty"`
}

type FlexiPage struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlexiPage"`

	*Metadata

	Description string `xml:"description,omitempty"`

	FlexiPageRegions []*FlexiPageRegion `xml:"flexiPageRegions,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	PageTemplate string `xml:"pageTemplate,omitempty"`

	ParentFlexiPage string `xml:"parentFlexiPage,omitempty"`

	PlatformActionlist *PlatformActionList `xml:"platformActionlist,omitempty"`

	QuickActionList *QuickActionList `xml:"quickActionList,omitempty"`

	SobjectType string `xml:"sobjectType,omitempty"`

	Type_ *FlexiPageType `xml:"type,omitempty"`
}

type FlexiPageRegion struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlexiPageRegion"`

	Appendable *RegionFlagStatus `xml:"appendable,omitempty"`

	ComponentInstances []*ComponentInstance `xml:"componentInstances,omitempty"`

	Mode *FlexiPageRegionMode `xml:"mode,omitempty"`

	Name string `xml:"name,omitempty"`

	Prependable *RegionFlagStatus `xml:"prependable,omitempty"`

	Replaceable *RegionFlagStatus `xml:"replaceable,omitempty"`

	Type_ *FlexiPageRegionType `xml:"type,omitempty"`
}

type ComponentInstance struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ComponentInstance"`

	ComponentInstanceProperties []*ComponentInstanceProperty `xml:"componentInstanceProperties,omitempty"`

	ComponentName string `xml:"componentName,omitempty"`
}

type ComponentInstanceProperty struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ComponentInstanceProperty"`

	Name string `xml:"name,omitempty"`

	Type_ *ComponentInstancePropertyTypeEnum `xml:"type,omitempty"`

	Value string `xml:"value,omitempty"`
}

type QuickActionList struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuickActionList"`

	QuickActionListItems []*QuickActionListItem `xml:"quickActionListItems,omitempty"`
}

type QuickActionListItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuickActionListItem"`

	QuickActionName string `xml:"quickActionName,omitempty"`
}

type Flow struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Flow"`

	*Metadata

	ActionCalls []*FlowActionCall `xml:"actionCalls,omitempty"`

	ApexPluginCalls []*FlowApexPluginCall `xml:"apexPluginCalls,omitempty"`

	Assignments []*FlowAssignment `xml:"assignments,omitempty"`

	Choices []*FlowChoice `xml:"choices,omitempty"`

	Constants []*FlowConstant `xml:"constants,omitempty"`

	Decisions []*FlowDecision `xml:"decisions,omitempty"`

	Description string `xml:"description,omitempty"`

	DynamicChoiceSets []*FlowDynamicChoiceSet `xml:"dynamicChoiceSets,omitempty"`

	Formulas []*FlowFormula `xml:"formulas,omitempty"`

	InterviewLabel string `xml:"interviewLabel,omitempty"`

	Label string `xml:"label,omitempty"`

	Loops []*FlowLoop `xml:"loops,omitempty"`

	ProcessMetadataValues []*FlowMetadataValue `xml:"processMetadataValues,omitempty"`

	ProcessType *FlowProcessType `xml:"processType,omitempty"`

	RecordCreates []*FlowRecordCreate `xml:"recordCreates,omitempty"`

	RecordDeletes []*FlowRecordDelete `xml:"recordDeletes,omitempty"`

	RecordLookups []*FlowRecordLookup `xml:"recordLookups,omitempty"`

	RecordUpdates []*FlowRecordUpdate `xml:"recordUpdates,omitempty"`

	Screens []*FlowScreen `xml:"screens,omitempty"`

	StartElementReference string `xml:"startElementReference,omitempty"`

	Steps []*FlowStep `xml:"steps,omitempty"`

	Subflows []*FlowSubflow `xml:"subflows,omitempty"`

	TextTemplates []*FlowTextTemplate `xml:"textTemplates,omitempty"`

	Variables []*FlowVariable `xml:"variables,omitempty"`

	Waits []*FlowWait `xml:"waits,omitempty"`
}

type FlowActionCall struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowActionCall"`

	*FlowNode

	ActionName string `xml:"actionName,omitempty"`

	ActionType *InvocableActionType `xml:"actionType,omitempty"`

	Connector *FlowConnector `xml:"connector,omitempty"`

	FaultConnector *FlowConnector `xml:"faultConnector,omitempty"`

	InputParameters []*FlowActionCallInputParameter `xml:"inputParameters,omitempty"`

	OutputParameters []*FlowActionCallOutputParameter `xml:"outputParameters,omitempty"`
}

type FlowNode struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowNode"`

	*FlowElement

	Label string `xml:"label,omitempty"`

	LocationX int32 `xml:"locationX,omitempty"`

	LocationY int32 `xml:"locationY,omitempty"`
}

type FlowElement struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowElement"`

	*FlowBaseElement

	Description string `xml:"description,omitempty"`

	Name string `xml:"name,omitempty"`
}

type FlowBaseElement struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowBaseElement"`

	ProcessMetadataValues []*FlowMetadataValue `xml:"processMetadataValues,omitempty"`
}

type FlowMetadataValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowMetadataValue"`

	Name string `xml:"name,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowElementReferenceOrValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowElementReferenceOrValue"`

	BooleanValue bool `xml:"booleanValue,omitempty"`

	DateTimeValue time.Time `xml:"dateTimeValue,omitempty"`

	DateValue time.Time `xml:"dateValue,omitempty"`

	ElementReference string `xml:"elementReference,omitempty"`

	NumberValue float64 `xml:"numberValue,omitempty"`

	StringValue string `xml:"stringValue,omitempty"`
}

type FlowActionCallInputParameter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowActionCallInputParameter"`

	*FlowBaseElement

	Name string `xml:"name,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowActionCallOutputParameter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowActionCallOutputParameter"`

	*FlowBaseElement

	AssignToReference string `xml:"assignToReference,omitempty"`

	Name string `xml:"name,omitempty"`
}

type FlowApexPluginCallInputParameter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowApexPluginCallInputParameter"`

	*FlowBaseElement

	Name string `xml:"name,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowApexPluginCallOutputParameter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowApexPluginCallOutputParameter"`

	*FlowBaseElement

	AssignToReference string `xml:"assignToReference,omitempty"`

	Name string `xml:"name,omitempty"`
}

type FlowAssignmentItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowAssignmentItem"`

	*FlowBaseElement

	AssignToReference string `xml:"assignToReference,omitempty"`

	Operator *FlowAssignmentOperator `xml:"operator,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowChoiceUserInput struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowChoiceUserInput"`

	*FlowBaseElement

	IsRequired bool `xml:"isRequired,omitempty"`

	PromptText string `xml:"promptText,omitempty"`

	ValidationRule *FlowInputValidationRule `xml:"validationRule,omitempty"`
}

type FlowInputValidationRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowInputValidationRule"`

	ErrorMessage string `xml:"errorMessage,omitempty"`

	FormulaExpression string `xml:"formulaExpression,omitempty"`
}

type FlowCondition struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowCondition"`

	*FlowBaseElement

	LeftValueReference string `xml:"leftValueReference,omitempty"`

	Operator *FlowComparisonOperator `xml:"operator,omitempty"`

	RightValue *FlowElementReferenceOrValue `xml:"rightValue,omitempty"`
}

type FlowConnector struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowConnector"`

	*FlowBaseElement

	TargetReference string `xml:"targetReference,omitempty"`
}

type FlowInputFieldAssignment struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowInputFieldAssignment"`

	*FlowBaseElement

	Field string `xml:"field,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowOutputFieldAssignment struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowOutputFieldAssignment"`

	*FlowBaseElement

	AssignToReference string `xml:"assignToReference,omitempty"`

	Field string `xml:"field,omitempty"`
}

type FlowRecordFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowRecordFilter"`

	*FlowBaseElement

	Field string `xml:"field,omitempty"`

	Operator *FlowRecordFilterOperator `xml:"operator,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowScreenRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowScreenRule"`

	*FlowBaseElement

	ConditionLogic string `xml:"conditionLogic,omitempty"`

	Conditions []*FlowCondition `xml:"conditions,omitempty"`

	Label string `xml:"label,omitempty"`

	RuleActions []*FlowScreenRuleAction `xml:"ruleActions,omitempty"`
}

type FlowScreenRuleAction struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowScreenRuleAction"`

	*FlowBaseElement

	Attribute string `xml:"attribute,omitempty"`

	FieldReference string `xml:"fieldReference,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowSubflowInputAssignment struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowSubflowInputAssignment"`

	*FlowBaseElement

	Name string `xml:"name,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowSubflowOutputAssignment struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowSubflowOutputAssignment"`

	*FlowBaseElement

	AssignToReference string `xml:"assignToReference,omitempty"`

	Name string `xml:"name,omitempty"`
}

type FlowWaitEventInputParameter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowWaitEventInputParameter"`

	*FlowBaseElement

	Name string `xml:"name,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowWaitEventOutputParameter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowWaitEventOutputParameter"`

	*FlowBaseElement

	AssignToReference string `xml:"assignToReference,omitempty"`

	Name string `xml:"name,omitempty"`
}

type FlowChoice struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowChoice"`

	*FlowElement

	ChoiceText string `xml:"choiceText,omitempty"`

	DataType *FlowDataType `xml:"dataType,omitempty"`

	UserInput *FlowChoiceUserInput `xml:"userInput,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowConstant struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowConstant"`

	*FlowElement

	DataType *FlowDataType `xml:"dataType,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowDynamicChoiceSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowDynamicChoiceSet"`

	*FlowElement

	DataType *FlowDataType `xml:"dataType,omitempty"`

	DisplayField string `xml:"displayField,omitempty"`

	Filters []*FlowRecordFilter `xml:"filters,omitempty"`

	Limit int32 `xml:"limit,omitempty"`

	Object string `xml:"object,omitempty"`

	OutputAssignments []*FlowOutputFieldAssignment `xml:"outputAssignments,omitempty"`

	PicklistField string `xml:"picklistField,omitempty"`

	PicklistObject string `xml:"picklistObject,omitempty"`

	SortField string `xml:"sortField,omitempty"`

	SortOrder *SortOrder `xml:"sortOrder,omitempty"`

	ValueField string `xml:"valueField,omitempty"`
}

type FlowFormula struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowFormula"`

	*FlowElement

	DataType *FlowDataType `xml:"dataType,omitempty"`

	Expression string `xml:"expression,omitempty"`

	Scale int32 `xml:"scale,omitempty"`
}

type FlowRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowRule"`

	*FlowElement

	ConditionLogic string `xml:"conditionLogic,omitempty"`

	Conditions []*FlowCondition `xml:"conditions,omitempty"`

	Connector *FlowConnector `xml:"connector,omitempty"`

	Label string `xml:"label,omitempty"`
}

type FlowScreenField struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowScreenField"`

	*FlowElement

	ChoiceReferences []string `xml:"choiceReferences,omitempty"`

	DataType *FlowDataType `xml:"dataType,omitempty"`

	DefaultSelectedChoiceReference string `xml:"defaultSelectedChoiceReference,omitempty"`

	DefaultValue *FlowElementReferenceOrValue `xml:"defaultValue,omitempty"`

	FieldText string `xml:"fieldText,omitempty"`

	FieldType *FlowScreenFieldType `xml:"fieldType,omitempty"`

	HelpText string `xml:"helpText,omitempty"`

	IsRequired bool `xml:"isRequired,omitempty"`

	Scale int32 `xml:"scale,omitempty"`

	ValidationRule *FlowInputValidationRule `xml:"validationRule,omitempty"`
}

type FlowTextTemplate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowTextTemplate"`

	*FlowElement

	Text string `xml:"text,omitempty"`
}

type FlowVariable struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowVariable"`

	*FlowElement

	DataType *FlowDataType `xml:"dataType,omitempty"`

	IsCollection bool `xml:"isCollection,omitempty"`

	IsInput bool `xml:"isInput,omitempty"`

	IsOutput bool `xml:"isOutput,omitempty"`

	ObjectType string `xml:"objectType,omitempty"`

	Scale int32 `xml:"scale,omitempty"`

	Value *FlowElementReferenceOrValue `xml:"value,omitempty"`
}

type FlowWaitEvent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowWaitEvent"`

	*FlowElement

	ConditionLogic string `xml:"conditionLogic,omitempty"`

	Conditions []*FlowCondition `xml:"conditions,omitempty"`

	Connector *FlowConnector `xml:"connector,omitempty"`

	EventType string `xml:"eventType,omitempty"`

	InputParameters []*FlowWaitEventInputParameter `xml:"inputParameters,omitempty"`

	Label string `xml:"label,omitempty"`

	OutputParameters []*FlowWaitEventOutputParameter `xml:"outputParameters,omitempty"`
}

type FlowApexPluginCall struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowApexPluginCall"`

	*FlowNode

	ApexClass string `xml:"apexClass,omitempty"`

	Connector *FlowConnector `xml:"connector,omitempty"`

	FaultConnector *FlowConnector `xml:"faultConnector,omitempty"`

	InputParameters []*FlowApexPluginCallInputParameter `xml:"inputParameters,omitempty"`

	OutputParameters []*FlowApexPluginCallOutputParameter `xml:"outputParameters,omitempty"`
}

type FlowAssignment struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowAssignment"`

	*FlowNode

	AssignmentItems []*FlowAssignmentItem `xml:"assignmentItems,omitempty"`

	Connector *FlowConnector `xml:"connector,omitempty"`
}

type FlowDecision struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowDecision"`

	*FlowNode

	DefaultConnector *FlowConnector `xml:"defaultConnector,omitempty"`

	DefaultConnectorLabel string `xml:"defaultConnectorLabel,omitempty"`

	Rules []*FlowRule `xml:"rules,omitempty"`
}

type FlowLoop struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowLoop"`

	*FlowNode

	AssignNextValueToReference string `xml:"assignNextValueToReference,omitempty"`

	CollectionReference string `xml:"collectionReference,omitempty"`

	IterationOrder *IterationOrder `xml:"iterationOrder,omitempty"`

	NextValueConnector *FlowConnector `xml:"nextValueConnector,omitempty"`

	NoMoreValuesConnector *FlowConnector `xml:"noMoreValuesConnector,omitempty"`
}

type FlowRecordCreate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowRecordCreate"`

	*FlowNode

	AssignRecordIdToReference string `xml:"assignRecordIdToReference,omitempty"`

	Connector *FlowConnector `xml:"connector,omitempty"`

	FaultConnector *FlowConnector `xml:"faultConnector,omitempty"`

	InputAssignments []*FlowInputFieldAssignment `xml:"inputAssignments,omitempty"`

	InputReference string `xml:"inputReference,omitempty"`

	Object string `xml:"object,omitempty"`
}

type FlowRecordDelete struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowRecordDelete"`

	*FlowNode

	Connector *FlowConnector `xml:"connector,omitempty"`

	FaultConnector *FlowConnector `xml:"faultConnector,omitempty"`

	Filters []*FlowRecordFilter `xml:"filters,omitempty"`

	InputReference string `xml:"inputReference,omitempty"`

	Object string `xml:"object,omitempty"`
}

type FlowRecordLookup struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowRecordLookup"`

	*FlowNode

	AssignNullValuesIfNoRecordsFound bool `xml:"assignNullValuesIfNoRecordsFound,omitempty"`

	Connector *FlowConnector `xml:"connector,omitempty"`

	FaultConnector *FlowConnector `xml:"faultConnector,omitempty"`

	Filters []*FlowRecordFilter `xml:"filters,omitempty"`

	Object string `xml:"object,omitempty"`

	OutputAssignments []*FlowOutputFieldAssignment `xml:"outputAssignments,omitempty"`

	OutputReference string `xml:"outputReference,omitempty"`

	QueriedFields []string `xml:"queriedFields,omitempty"`

	SortField string `xml:"sortField,omitempty"`

	SortOrder *SortOrder `xml:"sortOrder,omitempty"`
}

type FlowRecordUpdate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowRecordUpdate"`

	*FlowNode

	Connector *FlowConnector `xml:"connector,omitempty"`

	FaultConnector *FlowConnector `xml:"faultConnector,omitempty"`

	Filters []*FlowRecordFilter `xml:"filters,omitempty"`

	InputAssignments []*FlowInputFieldAssignment `xml:"inputAssignments,omitempty"`

	InputReference string `xml:"inputReference,omitempty"`

	Object string `xml:"object,omitempty"`
}

type FlowScreen struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowScreen"`

	*FlowNode

	AllowBack bool `xml:"allowBack,omitempty"`

	AllowFinish bool `xml:"allowFinish,omitempty"`

	AllowPause bool `xml:"allowPause,omitempty"`

	Connector *FlowConnector `xml:"connector,omitempty"`

	Fields []*FlowScreenField `xml:"fields,omitempty"`

	HelpText string `xml:"helpText,omitempty"`

	PausedText string `xml:"pausedText,omitempty"`

	Rules []*FlowScreenRule `xml:"rules,omitempty"`
}

type FlowStep struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowStep"`

	*FlowNode

	Connectors []*FlowConnector `xml:"connectors,omitempty"`
}

type FlowSubflow struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowSubflow"`

	*FlowNode

	Connector *FlowConnector `xml:"connector,omitempty"`

	FlowName string `xml:"flowName,omitempty"`

	InputAssignments []*FlowSubflowInputAssignment `xml:"inputAssignments,omitempty"`

	OutputAssignments []*FlowSubflowOutputAssignment `xml:"outputAssignments,omitempty"`
}

type FlowWait struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowWait"`

	*FlowNode

	DefaultConnector *FlowConnector `xml:"defaultConnector,omitempty"`

	DefaultConnectorLabel string `xml:"defaultConnectorLabel,omitempty"`

	FaultConnector *FlowConnector `xml:"faultConnector,omitempty"`

	WaitEvents []*FlowWaitEvent `xml:"waitEvents,omitempty"`
}

type FlowDefinition struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FlowDefinition"`

	*Metadata

	ActiveVersionNumber int32 `xml:"activeVersionNumber,omitempty"`

	Description string `xml:"description,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`
}

type Folder struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Folder"`

	*Metadata

	AccessType *FolderAccessTypes `xml:"accessType,omitempty"`

	FolderShares []*FolderShare `xml:"folderShares,omitempty"`

	Name string `xml:"name,omitempty"`

	PublicFolderAccess *PublicFolderAccess `xml:"publicFolderAccess,omitempty"`

	SharedTo *SharedTo `xml:"sharedTo,omitempty"`
}

type FolderShare struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FolderShare"`

	AccessLevel *FolderShareAccessLevel `xml:"accessLevel,omitempty"`

	SharedTo string `xml:"sharedTo,omitempty"`

	SharedToType *FolderSharedToType `xml:"sharedToType,omitempty"`
}

type DashboardFolder struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardFolder"`

	*Folder
}

type DocumentFolder struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DocumentFolder"`

	*Folder
}

type EmailFolder struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EmailFolder"`

	*Folder
}

type ReportFolder struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportFolder"`

	*Folder
}

type ForecastingSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ForecastingSettings"`

	*Metadata

	DisplayCurrency *DisplayCurrency `xml:"displayCurrency,omitempty"`

	EnableForecasts bool `xml:"enableForecasts,omitempty"`

	ForecastingCategoryMappings []*ForecastingCategoryMapping `xml:"forecastingCategoryMappings,omitempty"`

	ForecastingTypeSettings []*ForecastingTypeSettings `xml:"forecastingTypeSettings,omitempty"`
}

type ForecastingCategoryMapping struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ForecastingCategoryMapping"`

	ForecastingItemCategoryApiName string `xml:"forecastingItemCategoryApiName,omitempty"`

	WeightedSourceCategories []*WeightedSourceCategory `xml:"weightedSourceCategories,omitempty"`
}

type WeightedSourceCategory struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WeightedSourceCategory"`

	SourceCategoryApiName string `xml:"sourceCategoryApiName,omitempty"`

	Weight float64 `xml:"weight,omitempty"`
}

type ForecastingTypeSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ForecastingTypeSettings"`

	Active bool `xml:"active,omitempty"`

	AdjustmentsSettings *AdjustmentsSettings `xml:"adjustmentsSettings,omitempty"`

	DisplayedCategoryApiNames []string `xml:"displayedCategoryApiNames,omitempty"`

	ForecastRangeSettings *ForecastRangeSettings `xml:"forecastRangeSettings,omitempty"`

	ForecastedCategoryApiNames []string `xml:"forecastedCategoryApiNames,omitempty"`

	IsAmount bool `xml:"isAmount,omitempty"`

	IsAvailable bool `xml:"isAvailable,omitempty"`

	IsQuantity bool `xml:"isQuantity,omitempty"`

	ManagerAdjustableCategoryApiNames []string `xml:"managerAdjustableCategoryApiNames,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	Name string `xml:"name,omitempty"`

	OpportunityListFieldsLabelMappings []*OpportunityListFieldsLabelMapping `xml:"opportunityListFieldsLabelMappings,omitempty"`

	OpportunityListFieldsSelectedSettings *OpportunityListFieldsSelectedSettings `xml:"opportunityListFieldsSelectedSettings,omitempty"`

	OpportunityListFieldsUnselectedSettings *OpportunityListFieldsUnselectedSettings `xml:"opportunityListFieldsUnselectedSettings,omitempty"`

	OwnerAdjustableCategoryApiNames []string `xml:"ownerAdjustableCategoryApiNames,omitempty"`

	QuotasSettings *QuotasSettings `xml:"quotasSettings,omitempty"`
}

type AdjustmentsSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AdjustmentsSettings"`

	EnableAdjustments bool `xml:"enableAdjustments,omitempty"`

	EnableOwnerAdjustments bool `xml:"enableOwnerAdjustments,omitempty"`
}

type ForecastRangeSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ForecastRangeSettings"`

	Beginning int32 `xml:"beginning,omitempty"`

	Displaying int32 `xml:"displaying,omitempty"`

	PeriodType *PeriodTypes `xml:"periodType,omitempty"`
}

type OpportunityListFieldsLabelMapping struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata OpportunityListFieldsLabelMapping"`

	Field string `xml:"field,omitempty"`

	Label string `xml:"label,omitempty"`
}

type OpportunityListFieldsSelectedSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata OpportunityListFieldsSelectedSettings"`

	Field []string `xml:"field,omitempty"`
}

type OpportunityListFieldsUnselectedSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata OpportunityListFieldsUnselectedSettings"`

	Field []string `xml:"field,omitempty"`
}

type QuotasSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuotasSettings"`

	ShowQuotas bool `xml:"showQuotas,omitempty"`
}

type GlobalValueSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata GlobalValueSet"`

	*Metadata

	CustomValue []*CustomValue `xml:"customValue,omitempty"`

	Description string `xml:"description,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	Sorted bool `xml:"sorted,omitempty"`
}

type GlobalValueSetTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata GlobalValueSetTranslation"`

	*Metadata

	ValueTranslation []*ValueTranslation `xml:"valueTranslation,omitempty"`
}

type ValueTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ValueTranslation"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	Translation string `xml:"translation,omitempty"`
}

type Group struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Group"`

	*Metadata

	DoesIncludeBosses bool `xml:"doesIncludeBosses,omitempty"`

	Name string `xml:"name,omitempty"`
}

type HomePageComponent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata HomePageComponent"`

	*Metadata

	Body string `xml:"body,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Links []string `xml:"links,omitempty"`

	Page string `xml:"page,omitempty"`

	PageComponentType *PageComponentType `xml:"pageComponentType,omitempty"`

	ShowLabel bool `xml:"showLabel,omitempty"`

	ShowScrollbars bool `xml:"showScrollbars,omitempty"`

	Width *PageComponentWidth `xml:"width,omitempty"`
}

type HomePageLayout struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata HomePageLayout"`

	*Metadata

	NarrowComponents []string `xml:"narrowComponents,omitempty"`

	WideComponents []string `xml:"wideComponents,omitempty"`
}

type IdeasSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata IdeasSettings"`

	*Metadata

	EnableChatterProfile bool `xml:"enableChatterProfile,omitempty"`

	EnableIdeaThemes bool `xml:"enableIdeaThemes,omitempty"`

	EnableIdeas bool `xml:"enableIdeas,omitempty"`

	EnableIdeasReputation bool `xml:"enableIdeasReputation,omitempty"`

	HalfLife float64 `xml:"halfLife,omitempty"`

	IdeasProfilePage string `xml:"ideasProfilePage,omitempty"`
}

type InstalledPackage struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata InstalledPackage"`

	*Metadata

	Password string `xml:"password,omitempty"`

	VersionNumber string `xml:"versionNumber,omitempty"`
}

type KeywordList struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KeywordList"`

	*Metadata

	Description string `xml:"description,omitempty"`

	Keywords []*Keyword `xml:"keywords,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`
}

type Keyword struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Keyword"`

	Keyword string `xml:"keyword,omitempty"`
}

type KnowledgeSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeSettings"`

	*Metadata

	Answers *KnowledgeAnswerSettings `xml:"answers,omitempty"`

	Cases *KnowledgeCaseSettings `xml:"cases,omitempty"`

	DefaultLanguage string `xml:"defaultLanguage,omitempty"`

	EnableChatterQuestionKBDeflection bool `xml:"enableChatterQuestionKBDeflection,omitempty"`

	EnableCreateEditOnArticlesTab bool `xml:"enableCreateEditOnArticlesTab,omitempty"`

	EnableExternalMediaContent bool `xml:"enableExternalMediaContent,omitempty"`

	EnableKnowledge bool `xml:"enableKnowledge,omitempty"`

	Languages *KnowledgeLanguageSettings `xml:"languages,omitempty"`

	ShowArticleSummariesCustomerPortal bool `xml:"showArticleSummariesCustomerPortal,omitempty"`

	ShowArticleSummariesInternalApp bool `xml:"showArticleSummariesInternalApp,omitempty"`

	ShowArticleSummariesPartnerPortal bool `xml:"showArticleSummariesPartnerPortal,omitempty"`

	ShowValidationStatusField bool `xml:"showValidationStatusField,omitempty"`

	SuggestedArticles *KnowledgeSuggestedArticlesSettings `xml:"suggestedArticles,omitempty"`
}

type KnowledgeAnswerSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeAnswerSettings"`

	AssignTo string `xml:"assignTo,omitempty"`

	DefaultArticleType string `xml:"defaultArticleType,omitempty"`

	EnableArticleCreation bool `xml:"enableArticleCreation,omitempty"`
}

type KnowledgeCaseSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeCaseSettings"`

	ArticlePDFCreationProfile string `xml:"articlePDFCreationProfile,omitempty"`

	ArticlePublicSharingCommunities *KnowledgeCommunitiesSettings `xml:"articlePublicSharingCommunities,omitempty"`

	ArticlePublicSharingSites *KnowledgeSitesSettings `xml:"articlePublicSharingSites,omitempty"`

	ArticlePublicSharingSitesChatterAnswers *KnowledgeSitesSettings `xml:"articlePublicSharingSitesChatterAnswers,omitempty"`

	AssignTo string `xml:"assignTo,omitempty"`

	CustomizationClass string `xml:"customizationClass,omitempty"`

	DefaultContributionArticleType string `xml:"defaultContributionArticleType,omitempty"`

	Editor *KnowledgeCaseEditor `xml:"editor,omitempty"`

	EnableArticleCreation bool `xml:"enableArticleCreation,omitempty"`

	EnableArticlePublicSharingSites bool `xml:"enableArticlePublicSharingSites,omitempty"`

	UseProfileForPDFCreation bool `xml:"useProfileForPDFCreation,omitempty"`
}

type KnowledgeCommunitiesSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeCommunitiesSettings"`

	Community []string `xml:"community,omitempty"`
}

type KnowledgeSitesSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeSitesSettings"`

	Site []string `xml:"site,omitempty"`
}

type KnowledgeLanguageSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeLanguageSettings"`

	Language []*KnowledgeLanguage `xml:"language,omitempty"`
}

type KnowledgeLanguage struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeLanguage"`

	Active bool `xml:"active,omitempty"`

	DefaultAssignee string `xml:"defaultAssignee,omitempty"`

	DefaultAssigneeType *KnowledgeLanguageLookupValueType `xml:"defaultAssigneeType,omitempty"`

	DefaultReviewer string `xml:"defaultReviewer,omitempty"`

	DefaultReviewerType *KnowledgeLanguageLookupValueType `xml:"defaultReviewerType,omitempty"`

	Name string `xml:"name,omitempty"`
}

type KnowledgeSuggestedArticlesSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeSuggestedArticlesSettings"`

	CaseFields *KnowledgeCaseFieldsSettings `xml:"caseFields,omitempty"`

	UseSuggestedArticlesForCase bool `xml:"useSuggestedArticlesForCase,omitempty"`
}

type KnowledgeCaseFieldsSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeCaseFieldsSettings"`

	Field []*KnowledgeCaseField `xml:"field,omitempty"`
}

type KnowledgeCaseField struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata KnowledgeCaseField"`

	Name string `xml:"name,omitempty"`
}

type Layout struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Layout"`

	*Metadata

	CustomButtons []string `xml:"customButtons,omitempty"`

	CustomConsoleComponents *CustomConsoleComponents `xml:"customConsoleComponents,omitempty"`

	EmailDefault bool `xml:"emailDefault,omitempty"`

	ExcludeButtons []string `xml:"excludeButtons,omitempty"`

	FeedLayout *FeedLayout `xml:"feedLayout,omitempty"`

	Headers []*LayoutHeader `xml:"headers,omitempty"`

	LayoutSections []*LayoutSection `xml:"layoutSections,omitempty"`

	MiniLayout *MiniLayout `xml:"miniLayout,omitempty"`

	MultilineLayoutFields []string `xml:"multilineLayoutFields,omitempty"`

	PlatformActionList *PlatformActionList `xml:"platformActionList,omitempty"`

	QuickActionList *QuickActionList `xml:"quickActionList,omitempty"`

	RelatedContent *RelatedContent `xml:"relatedContent,omitempty"`

	RelatedLists []*RelatedListItem `xml:"relatedLists,omitempty"`

	RelatedObjects []string `xml:"relatedObjects,omitempty"`

	RunAssignmentRulesDefault bool `xml:"runAssignmentRulesDefault,omitempty"`

	ShowEmailCheckbox bool `xml:"showEmailCheckbox,omitempty"`

	ShowHighlightsPanel bool `xml:"showHighlightsPanel,omitempty"`

	ShowInteractionLogPanel bool `xml:"showInteractionLogPanel,omitempty"`

	ShowKnowledgeComponent bool `xml:"showKnowledgeComponent,omitempty"`

	ShowRunAssignmentRulesCheckbox bool `xml:"showRunAssignmentRulesCheckbox,omitempty"`

	ShowSolutionSection bool `xml:"showSolutionSection,omitempty"`

	ShowSubmitAndAttachButton bool `xml:"showSubmitAndAttachButton,omitempty"`

	SummaryLayout *SummaryLayout `xml:"summaryLayout,omitempty"`
}

type CustomConsoleComponents struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomConsoleComponents"`

	PrimaryTabComponents *PrimaryTabComponents `xml:"primaryTabComponents,omitempty"`

	SubtabComponents *SubtabComponents `xml:"subtabComponents,omitempty"`
}

type PrimaryTabComponents struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PrimaryTabComponents"`

	Containers []*Container `xml:"containers,omitempty"`
}

type Container struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Container"`

	Height int32 `xml:"height,omitempty"`

	IsContainerAutoSizeEnabled bool `xml:"isContainerAutoSizeEnabled,omitempty"`

	Region string `xml:"region,omitempty"`

	SidebarComponents []*SidebarComponent `xml:"sidebarComponents,omitempty"`

	Style string `xml:"style,omitempty"`

	Unit string `xml:"unit,omitempty"`

	Width int32 `xml:"width,omitempty"`
}

type SidebarComponent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SidebarComponent"`

	ComponentType string `xml:"componentType,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Label string `xml:"label,omitempty"`

	Lookup string `xml:"lookup,omitempty"`

	Page string `xml:"page,omitempty"`

	RelatedLists []*RelatedList `xml:"relatedLists,omitempty"`

	Unit string `xml:"unit,omitempty"`

	Width int32 `xml:"width,omitempty"`
}

type RelatedList struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RelatedList"`

	HideOnDetail bool `xml:"hideOnDetail,omitempty"`

	Name string `xml:"name,omitempty"`
}

type SubtabComponents struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SubtabComponents"`

	Containers []*Container `xml:"containers,omitempty"`
}

type FeedLayout struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FeedLayout"`

	AutocollapsePublisher bool `xml:"autocollapsePublisher,omitempty"`

	CompactFeed bool `xml:"compactFeed,omitempty"`

	FeedFilterPosition *FeedLayoutFilterPosition `xml:"feedFilterPosition,omitempty"`

	FeedFilters []*FeedLayoutFilter `xml:"feedFilters,omitempty"`

	FullWidthFeed bool `xml:"fullWidthFeed,omitempty"`

	HideSidebar bool `xml:"hideSidebar,omitempty"`

	HighlightExternalFeedItems bool `xml:"highlightExternalFeedItems,omitempty"`

	LeftComponents []*FeedLayoutComponent `xml:"leftComponents,omitempty"`

	RightComponents []*FeedLayoutComponent `xml:"rightComponents,omitempty"`

	UseInlineFiltersInConsole bool `xml:"useInlineFiltersInConsole,omitempty"`
}

type FeedLayoutFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FeedLayoutFilter"`

	FeedFilterName string `xml:"feedFilterName,omitempty"`

	FeedFilterType *FeedLayoutFilterType `xml:"feedFilterType,omitempty"`

	FeedItemType *FeedItemType `xml:"feedItemType,omitempty"`
}

type FeedLayoutComponent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FeedLayoutComponent"`

	ComponentType *FeedLayoutComponentType `xml:"componentType,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Page string `xml:"page,omitempty"`
}

type LayoutSection struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LayoutSection"`

	CustomLabel bool `xml:"customLabel,omitempty"`

	DetailHeading bool `xml:"detailHeading,omitempty"`

	EditHeading bool `xml:"editHeading,omitempty"`

	Label string `xml:"label,omitempty"`

	LayoutColumns []*LayoutColumn `xml:"layoutColumns,omitempty"`

	Style *LayoutSectionStyle `xml:"style,omitempty"`
}

type LayoutColumn struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LayoutColumn"`

	LayoutItems []*LayoutItem `xml:"layoutItems,omitempty"`

	Reserved string `xml:"reserved,omitempty"`
}

type LayoutItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LayoutItem"`

	AnalyticsCloudComponent *AnalyticsCloudComponentLayoutItem `xml:"analyticsCloudComponent,omitempty"`

	Behavior *UiBehavior `xml:"behavior,omitempty"`

	Canvas string `xml:"canvas,omitempty"`

	Component string `xml:"component,omitempty"`

	CustomLink string `xml:"customLink,omitempty"`

	EmptySpace bool `xml:"emptySpace,omitempty"`

	Field string `xml:"field,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Page string `xml:"page,omitempty"`

	ReportChartComponent *ReportChartComponentLayoutItem `xml:"reportChartComponent,omitempty"`

	Scontrol string `xml:"scontrol,omitempty"`

	ShowLabel bool `xml:"showLabel,omitempty"`

	ShowScrollbars bool `xml:"showScrollbars,omitempty"`

	Width string `xml:"width,omitempty"`
}

type AnalyticsCloudComponentLayoutItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AnalyticsCloudComponentLayoutItem"`

	AssetType string `xml:"assetType,omitempty"`

	DevName string `xml:"devName,omitempty"`

	Error string `xml:"error,omitempty"`

	Filter string `xml:"filter,omitempty"`

	Height int32 `xml:"height,omitempty"`

	HideOnError bool `xml:"hideOnError,omitempty"`

	ShowSharing bool `xml:"showSharing,omitempty"`

	ShowTitle bool `xml:"showTitle,omitempty"`

	Width string `xml:"width,omitempty"`
}

type ReportChartComponentLayoutItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportChartComponentLayoutItem"`

	CacheData bool `xml:"cacheData,omitempty"`

	ContextFilterableField string `xml:"contextFilterableField,omitempty"`

	Error string `xml:"error,omitempty"`

	HideOnError bool `xml:"hideOnError,omitempty"`

	IncludeContext bool `xml:"includeContext,omitempty"`

	ReportName string `xml:"reportName,omitempty"`

	ShowTitle bool `xml:"showTitle,omitempty"`

	Size *ReportChartComponentSize `xml:"size,omitempty"`
}

type MiniLayout struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata MiniLayout"`

	Fields []string `xml:"fields,omitempty"`

	RelatedLists []*RelatedListItem `xml:"relatedLists,omitempty"`
}

type RelatedListItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RelatedListItem"`

	CustomButtons []string `xml:"customButtons,omitempty"`

	ExcludeButtons []string `xml:"excludeButtons,omitempty"`

	Fields []string `xml:"fields,omitempty"`

	RelatedList string `xml:"relatedList,omitempty"`

	SortField string `xml:"sortField,omitempty"`

	SortOrder *SortOrder `xml:"sortOrder,omitempty"`
}

type RelatedContent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RelatedContent"`

	RelatedContentItems []*RelatedContentItem `xml:"relatedContentItems,omitempty"`
}

type RelatedContentItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RelatedContentItem"`

	LayoutItem *LayoutItem `xml:"layoutItem,omitempty"`
}

type SummaryLayout struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SummaryLayout"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	SizeX int32 `xml:"sizeX,omitempty"`

	SizeY int32 `xml:"sizeY,omitempty"`

	SizeZ int32 `xml:"sizeZ,omitempty"`

	SummaryLayoutItems []*SummaryLayoutItem `xml:"summaryLayoutItems,omitempty"`

	SummaryLayoutStyle *SummaryLayoutStyle `xml:"summaryLayoutStyle,omitempty"`
}

type SummaryLayoutItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SummaryLayoutItem"`

	CustomLink string `xml:"customLink,omitempty"`

	Field string `xml:"field,omitempty"`

	PosX int32 `xml:"posX,omitempty"`

	PosY int32 `xml:"posY,omitempty"`

	PosZ int32 `xml:"posZ,omitempty"`
}

type LeadConvertSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LeadConvertSettings"`

	*Metadata

	AllowOwnerChange bool `xml:"allowOwnerChange,omitempty"`

	OpportunityCreationOptions *VisibleOrRequired `xml:"opportunityCreationOptions,omitempty"`
}

type Letterhead struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Letterhead"`

	*Metadata

	Available bool `xml:"available,omitempty"`

	BackgroundColor string `xml:"backgroundColor,omitempty"`

	BodyColor string `xml:"bodyColor,omitempty"`

	BottomLine *LetterheadLine `xml:"bottomLine,omitempty"`

	Description string `xml:"description,omitempty"`

	Footer *LetterheadHeaderFooter `xml:"footer,omitempty"`

	Header *LetterheadHeaderFooter `xml:"header,omitempty"`

	MiddleLine *LetterheadLine `xml:"middleLine,omitempty"`

	Name string `xml:"name,omitempty"`

	TopLine *LetterheadLine `xml:"topLine,omitempty"`
}

type LetterheadLine struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LetterheadLine"`

	Color string `xml:"color,omitempty"`

	Height int32 `xml:"height,omitempty"`
}

type LetterheadHeaderFooter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LetterheadHeaderFooter"`

	BackgroundColor string `xml:"backgroundColor,omitempty"`

	Height int32 `xml:"height,omitempty"`

	HorizontalAlignment *LetterheadHorizontalAlignment `xml:"horizontalAlignment,omitempty"`

	Logo string `xml:"logo,omitempty"`

	VerticalAlignment *LetterheadVerticalAlignment `xml:"verticalAlignment,omitempty"`
}

type LicenseDefinition struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LicenseDefinition"`

	*Metadata

	AggregationGroup string `xml:"aggregationGroup,omitempty"`

	Description string `xml:"description,omitempty"`

	IsPublished bool `xml:"isPublished,omitempty"`

	Label string `xml:"label,omitempty"`

	LicensedCustomPermissions []*LicensedCustomPermissions `xml:"licensedCustomPermissions,omitempty"`

	LicensingAuthority string `xml:"licensingAuthority,omitempty"`

	LicensingAuthorityProvider string `xml:"licensingAuthorityProvider,omitempty"`

	MinPlatformVersion int32 `xml:"minPlatformVersion,omitempty"`

	Origin string `xml:"origin,omitempty"`

	Revision int32 `xml:"revision,omitempty"`

	TrialLicenseDuration int32 `xml:"trialLicenseDuration,omitempty"`

	TrialLicenseQuantity int32 `xml:"trialLicenseQuantity,omitempty"`
}

type LicensedCustomPermissions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LicensedCustomPermissions"`

	CustomPermission string `xml:"customPermission,omitempty"`

	LicenseDefinition string `xml:"licenseDefinition,omitempty"`
}

type LiveAgentSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LiveAgentSettings"`

	*Metadata

	EnableLiveAgent bool `xml:"enableLiveAgent,omitempty"`
}

type LiveChatAgentConfig struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LiveChatAgentConfig"`

	*Metadata

	Assignments *AgentConfigAssignments `xml:"assignments,omitempty"`

	AutoGreeting string `xml:"autoGreeting,omitempty"`

	Capacity int32 `xml:"capacity,omitempty"`

	CriticalWaitTime int32 `xml:"criticalWaitTime,omitempty"`

	CustomAgentName string `xml:"customAgentName,omitempty"`

	EnableAgentFileTransfer bool `xml:"enableAgentFileTransfer,omitempty"`

	EnableAgentSneakPeek bool `xml:"enableAgentSneakPeek,omitempty"`

	EnableAssistanceFlag bool `xml:"enableAssistanceFlag,omitempty"`

	EnableAutoAwayOnDecline bool `xml:"enableAutoAwayOnDecline,omitempty"`

	EnableAutoAwayOnPushTimeout bool `xml:"enableAutoAwayOnPushTimeout,omitempty"`

	EnableChatConferencing bool `xml:"enableChatConferencing,omitempty"`

	EnableChatMonitoring bool `xml:"enableChatMonitoring,omitempty"`

	EnableChatTransferToAgent bool `xml:"enableChatTransferToAgent,omitempty"`

	EnableChatTransferToButton bool `xml:"enableChatTransferToButton,omitempty"`

	EnableChatTransferToSkill bool `xml:"enableChatTransferToSkill,omitempty"`

	EnableLogoutSound bool `xml:"enableLogoutSound,omitempty"`

	EnableNotifications bool `xml:"enableNotifications,omitempty"`

	EnableRequestSound bool `xml:"enableRequestSound,omitempty"`

	EnableSneakPeek bool `xml:"enableSneakPeek,omitempty"`

	EnableVisitorBlocking bool `xml:"enableVisitorBlocking,omitempty"`

	EnableWhisperMessage bool `xml:"enableWhisperMessage,omitempty"`

	Label string `xml:"label,omitempty"`

	SupervisorDefaultAgentStatusFilter *SupervisorAgentStatusFilter `xml:"supervisorDefaultAgentStatusFilter,omitempty"`

	SupervisorDefaultButtonFilter string `xml:"supervisorDefaultButtonFilter,omitempty"`

	SupervisorDefaultSkillFilter string `xml:"supervisorDefaultSkillFilter,omitempty"`

	SupervisorSkills *SupervisorAgentConfigSkills `xml:"supervisorSkills,omitempty"`

	TransferableButtons *AgentConfigButtons `xml:"transferableButtons,omitempty"`

	TransferableSkills *AgentConfigSkills `xml:"transferableSkills,omitempty"`
}

type AgentConfigAssignments struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AgentConfigAssignments"`

	Profiles *AgentConfigProfileAssignments `xml:"profiles,omitempty"`

	Users *AgentConfigUserAssignments `xml:"users,omitempty"`
}

type AgentConfigProfileAssignments struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AgentConfigProfileAssignments"`

	Profile []string `xml:"profile,omitempty"`
}

type AgentConfigUserAssignments struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AgentConfigUserAssignments"`

	User []string `xml:"user,omitempty"`
}

type SupervisorAgentConfigSkills struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SupervisorAgentConfigSkills"`

	Skill []string `xml:"skill,omitempty"`
}

type AgentConfigButtons struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AgentConfigButtons"`

	Button []string `xml:"button,omitempty"`
}

type AgentConfigSkills struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AgentConfigSkills"`

	Skill []string `xml:"skill,omitempty"`
}

type LiveChatButton struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LiveChatButton"`

	*Metadata

	Animation *LiveChatButtonPresentation `xml:"animation,omitempty"`

	AutoGreeting string `xml:"autoGreeting,omitempty"`

	ChasitorIdleTimeout int32 `xml:"chasitorIdleTimeout,omitempty"`

	ChasitorIdleTimeoutWarning int32 `xml:"chasitorIdleTimeoutWarning,omitempty"`

	ChatPage string `xml:"chatPage,omitempty"`

	CustomAgentName string `xml:"customAgentName,omitempty"`

	Deployments *LiveChatButtonDeployments `xml:"deployments,omitempty"`

	EnableQueue bool `xml:"enableQueue,omitempty"`

	InviteEndPosition *LiveChatButtonInviteEndPosition `xml:"inviteEndPosition,omitempty"`

	InviteImage string `xml:"inviteImage,omitempty"`

	InviteStartPosition *LiveChatButtonInviteStartPosition `xml:"inviteStartPosition,omitempty"`

	IsActive bool `xml:"isActive,omitempty"`

	Label string `xml:"label,omitempty"`

	NumberOfReroutingAttempts int32 `xml:"numberOfReroutingAttempts,omitempty"`

	OfflineImage string `xml:"offlineImage,omitempty"`

	OnlineImage string `xml:"onlineImage,omitempty"`

	OptionsCustomRoutingIsEnabled bool `xml:"optionsCustomRoutingIsEnabled,omitempty"`

	OptionsHasChasitorIdleTimeout bool `xml:"optionsHasChasitorIdleTimeout,omitempty"`

	OptionsHasInviteAfterAccept bool `xml:"optionsHasInviteAfterAccept,omitempty"`

	OptionsHasInviteAfterReject bool `xml:"optionsHasInviteAfterReject,omitempty"`

	OptionsHasRerouteDeclinedRequest bool `xml:"optionsHasRerouteDeclinedRequest,omitempty"`

	OptionsIsAutoAccept bool `xml:"optionsIsAutoAccept,omitempty"`

	OptionsIsInviteAutoRemove bool `xml:"optionsIsInviteAutoRemove,omitempty"`

	OverallQueueLength int32 `xml:"overallQueueLength,omitempty"`

	PerAgentQueueLength int32 `xml:"perAgentQueueLength,omitempty"`

	PostChatPage string `xml:"postChatPage,omitempty"`

	PostChatUrl string `xml:"postChatUrl,omitempty"`

	PreChatFormPage string `xml:"preChatFormPage,omitempty"`

	PreChatFormUrl string `xml:"preChatFormUrl,omitempty"`

	PushTimeOut int32 `xml:"pushTimeOut,omitempty"`

	RoutingType *LiveChatButtonRoutingType `xml:"routingType,omitempty"`

	Site string `xml:"site,omitempty"`

	Skills *LiveChatButtonSkills `xml:"skills,omitempty"`

	TimeToRemoveInvite int32 `xml:"timeToRemoveInvite,omitempty"`

	Type_ *LiveChatButtonType `xml:"type,omitempty"`

	WindowLanguage *Language `xml:"windowLanguage,omitempty"`
}

type LiveChatButtonDeployments struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LiveChatButtonDeployments"`

	Deployment []string `xml:"deployment,omitempty"`
}

type LiveChatButtonSkills struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LiveChatButtonSkills"`

	Skill []string `xml:"skill,omitempty"`
}

type LiveChatDeployment struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LiveChatDeployment"`

	*Metadata

	BrandingImage string `xml:"brandingImage,omitempty"`

	ConnectionTimeoutDuration int32 `xml:"connectionTimeoutDuration,omitempty"`

	ConnectionWarningDuration int32 `xml:"connectionWarningDuration,omitempty"`

	DisplayQueuePosition bool `xml:"displayQueuePosition,omitempty"`

	DomainWhiteList *LiveChatDeploymentDomainWhitelist `xml:"domainWhiteList,omitempty"`

	EnablePrechatApi bool `xml:"enablePrechatApi,omitempty"`

	EnableTranscriptSave bool `xml:"enableTranscriptSave,omitempty"`

	Label string `xml:"label,omitempty"`

	MobileBrandingImage string `xml:"mobileBrandingImage,omitempty"`

	Site string `xml:"site,omitempty"`

	WindowTitle string `xml:"windowTitle,omitempty"`
}

type LiveChatDeploymentDomainWhitelist struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LiveChatDeploymentDomainWhitelist"`

	Domain []string `xml:"domain,omitempty"`
}

type LiveChatSensitiveDataRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LiveChatSensitiveDataRule"`

	*Metadata

	ActionType *SensitiveDataActionType `xml:"actionType,omitempty"`

	Description string `xml:"description,omitempty"`

	EnforceOn int32 `xml:"enforceOn,omitempty"`

	IsEnabled bool `xml:"isEnabled,omitempty"`

	Pattern string `xml:"pattern,omitempty"`

	Replacement string `xml:"replacement,omitempty"`
}

type ManagedTopic struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ManagedTopic"`

	*Metadata

	ManagedTopicType string `xml:"managedTopicType,omitempty"`

	Name string `xml:"name,omitempty"`

	ParentName string `xml:"parentName,omitempty"`

	Position int32 `xml:"position,omitempty"`

	TopicDescription string `xml:"topicDescription,omitempty"`
}

type ManagedTopics struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ManagedTopics"`

	*Metadata

	ManagedTopic []*ManagedTopic `xml:"managedTopic,omitempty"`
}

type MarketingActionSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata MarketingActionSettings"`

	*Metadata

	EnableMarketingAction bool `xml:"enableMarketingAction,omitempty"`
}

type MarketingResourceType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata MarketingResourceType"`

	*Metadata

	Description string `xml:"description,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	Object string `xml:"object,omitempty"`

	Provider string `xml:"provider,omitempty"`
}

type MatchingRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata MatchingRule"`

	*Metadata

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	MatchingRuleItems []*MatchingRuleItem `xml:"matchingRuleItems,omitempty"`

	RuleStatus *MatchingRuleStatus `xml:"ruleStatus,omitempty"`
}

type MatchingRuleItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata MatchingRuleItem"`

	BlankValueBehavior *BlankValueBehavior `xml:"blankValueBehavior,omitempty"`

	FieldName string `xml:"fieldName,omitempty"`

	MatchingMethod *MatchingMethod `xml:"matchingMethod,omitempty"`
}

type MatchingRules struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata MatchingRules"`

	*Metadata

	MatchingRules []*MatchingRule `xml:"matchingRules,omitempty"`
}

type MetadataWithContent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata MetadataWithContent"`

	*Metadata

	Content []byte `xml:"content,omitempty"`
}

type ApexClass struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApexClass"`

	*MetadataWithContent

	ApiVersion float64 `xml:"apiVersion,omitempty"`

	PackageVersions []*PackageVersion `xml:"packageVersions,omitempty"`

	Status *ApexCodeUnitStatus `xml:"status,omitempty"`
}

type ApexComponent struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApexComponent"`

	*MetadataWithContent

	ApiVersion float64 `xml:"apiVersion,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	PackageVersions []*PackageVersion `xml:"packageVersions,omitempty"`
}

type ApexPage struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApexPage"`

	*MetadataWithContent

	ApiVersion float64 `xml:"apiVersion,omitempty"`

	AvailableInTouch bool `xml:"availableInTouch,omitempty"`

	ConfirmationTokenRequired bool `xml:"confirmationTokenRequired,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	PackageVersions []*PackageVersion `xml:"packageVersions,omitempty"`
}

type ApexTrigger struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ApexTrigger"`

	*MetadataWithContent

	ApiVersion float64 `xml:"apiVersion,omitempty"`

	PackageVersions []*PackageVersion `xml:"packageVersions,omitempty"`

	Status *ApexCodeUnitStatus `xml:"status,omitempty"`
}

type Certificate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Certificate"`

	*MetadataWithContent

	CaSigned bool `xml:"caSigned,omitempty"`

	EncryptedWithPlatformEncryption bool `xml:"encryptedWithPlatformEncryption,omitempty"`

	ExpirationDate time.Time `xml:"expirationDate,omitempty"`

	KeySize int32 `xml:"keySize,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	PrivateKeyExportable bool `xml:"privateKeyExportable,omitempty"`
}

type ContentAsset struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ContentAsset"`

	*MetadataWithContent

	Format *ContentAssetFormat `xml:"format,omitempty"`

	Language string `xml:"language,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	OriginNetwork string `xml:"originNetwork,omitempty"`

	Relationships *ContentAssetRelationships `xml:"relationships,omitempty"`

	Versions *ContentAssetVersions `xml:"versions,omitempty"`
}

type ContentAssetRelationships struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ContentAssetRelationships"`

	Organization *ContentAssetLink `xml:"organization,omitempty"`
}

type ContentAssetLink struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ContentAssetLink"`

	Access *ContentAssetAccess `xml:"access,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ContentAssetVersions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ContentAssetVersions"`

	Version []*ContentAssetVersion `xml:"version,omitempty"`
}

type ContentAssetVersion struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ContentAssetVersion"`

	Number string `xml:"number,omitempty"`

	PathOnClient string `xml:"pathOnClient,omitempty"`

	ZipEntry string `xml:"zipEntry,omitempty"`
}

type DataPipeline struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DataPipeline"`

	*MetadataWithContent

	ApiVersion float64 `xml:"apiVersion,omitempty"`

	Label string `xml:"label,omitempty"`

	ScriptType *DataPipelineType `xml:"scriptType,omitempty"`
}

type Document struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Document"`

	*MetadataWithContent

	Description string `xml:"description,omitempty"`

	InternalUseOnly bool `xml:"internalUseOnly,omitempty"`

	Keywords string `xml:"keywords,omitempty"`

	Name string `xml:"name,omitempty"`

	Public bool `xml:"public,omitempty"`
}

type EmailTemplate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata EmailTemplate"`

	*MetadataWithContent

	ApiVersion float64 `xml:"apiVersion,omitempty"`

	AttachedDocuments []string `xml:"attachedDocuments,omitempty"`

	Attachments []*Attachment `xml:"attachments,omitempty"`

	Available bool `xml:"available,omitempty"`

	Description string `xml:"description,omitempty"`

	EncodingKey *Encoding `xml:"encodingKey,omitempty"`

	Letterhead string `xml:"letterhead,omitempty"`

	Name string `xml:"name,omitempty"`

	PackageVersions []*PackageVersion `xml:"packageVersions,omitempty"`

	Style *EmailTemplateStyle `xml:"style,omitempty"`

	Subject string `xml:"subject,omitempty"`

	TextOnly string `xml:"textOnly,omitempty"`

	Type_ *EmailTemplateType `xml:"type,omitempty"`
}

type Attachment struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Attachment"`

	Content []byte `xml:"content,omitempty"`

	Name string `xml:"name,omitempty"`
}

type Scontrol struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Scontrol"`

	*MetadataWithContent

	ContentSource *SControlContentSource `xml:"contentSource,omitempty"`

	Description string `xml:"description,omitempty"`

	EncodingKey *Encoding `xml:"encodingKey,omitempty"`

	FileContent []byte `xml:"fileContent,omitempty"`

	FileName string `xml:"fileName,omitempty"`

	Name string `xml:"name,omitempty"`

	SupportsCaching bool `xml:"supportsCaching,omitempty"`
}

type SiteDotCom struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SiteDotCom"`

	*MetadataWithContent

	Label string `xml:"label,omitempty"`

	SiteType *SiteType `xml:"siteType,omitempty"`
}

type StaticResource struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata StaticResource"`

	*MetadataWithContent

	CacheControl *StaticResourceCacheControl `xml:"cacheControl,omitempty"`

	ContentType string `xml:"contentType,omitempty"`

	Description string `xml:"description,omitempty"`
}

type UiPlugin struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata UiPlugin"`

	*MetadataWithContent

	Description string `xml:"description,omitempty"`

	ExtensionPointIdentifier string `xml:"extensionPointIdentifier,omitempty"`

	IsEnabled bool `xml:"isEnabled,omitempty"`

	Language string `xml:"language,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`
}

type WaveDashboard struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WaveDashboard"`

	*MetadataWithContent

	Application string `xml:"application,omitempty"`

	Description string `xml:"description,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	TemplateAssetSourceName string `xml:"templateAssetSourceName,omitempty"`
}

type WaveDataflow struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WaveDataflow"`

	*MetadataWithContent

	Description string `xml:"description,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`
}

type WaveLens struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WaveLens"`

	*MetadataWithContent

	Application string `xml:"application,omitempty"`

	Datasets []string `xml:"datasets,omitempty"`

	Description string `xml:"description,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	TemplateAssetSourceName string `xml:"templateAssetSourceName,omitempty"`

	VisualizationType string `xml:"visualizationType,omitempty"`
}

type MilestoneType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata MilestoneType"`

	*Metadata

	Description string `xml:"description,omitempty"`

	RecurrenceType *MilestoneTypeRecurrenceType `xml:"recurrenceType,omitempty"`
}

type MobileSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata MobileSettings"`

	*Metadata

	ChatterMobile *ChatterMobileSettings `xml:"chatterMobile,omitempty"`

	DashboardMobile *DashboardMobileSettings `xml:"dashboardMobile,omitempty"`

	SalesforceMobile *SFDCMobileSettings `xml:"salesforceMobile,omitempty"`

	TouchMobile *TouchMobileSettings `xml:"touchMobile,omitempty"`
}

type ChatterMobileSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ChatterMobileSettings"`

	EnablePushNotifications bool `xml:"enablePushNotifications,omitempty"`
}

type DashboardMobileSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DashboardMobileSettings"`

	EnableDashboardIPadApp bool `xml:"enableDashboardIPadApp,omitempty"`
}

type SFDCMobileSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SFDCMobileSettings"`

	EnableMobileLite bool `xml:"enableMobileLite,omitempty"`

	EnableUserToDeviceLinking bool `xml:"enableUserToDeviceLinking,omitempty"`
}

type TouchMobileSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata TouchMobileSettings"`

	EnableTouchAppIPad bool `xml:"enableTouchAppIPad,omitempty"`

	EnableTouchAppIPhone bool `xml:"enableTouchAppIPhone,omitempty"`

	EnableTouchBrowserIPad bool `xml:"enableTouchBrowserIPad,omitempty"`

	EnableTouchIosPhone bool `xml:"enableTouchIosPhone,omitempty"`

	EnableVisualforceInTouch bool `xml:"enableVisualforceInTouch,omitempty"`
}

type ModerationRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ModerationRule"`

	*Metadata

	Action *ModerationRuleAction `xml:"action,omitempty"`

	Active bool `xml:"active,omitempty"`

	Description string `xml:"description,omitempty"`

	EntitiesAndFields []*ModeratedEntityField `xml:"entitiesAndFields,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	UserMessage string `xml:"userMessage,omitempty"`
}

type ModeratedEntityField struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ModeratedEntityField"`

	EntityName string `xml:"entityName,omitempty"`

	FieldName string `xml:"fieldName,omitempty"`

	KeywordList string `xml:"keywordList,omitempty"`
}

type NameSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata NameSettings"`

	*Metadata

	EnableMiddleName bool `xml:"enableMiddleName,omitempty"`

	EnableNameSuffix bool `xml:"enableNameSuffix,omitempty"`
}

type NamedCredential struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata NamedCredential"`

	*Metadata

	AuthProvider string `xml:"authProvider,omitempty"`

	Certificate string `xml:"certificate,omitempty"`

	Endpoint string `xml:"endpoint,omitempty"`

	Label string `xml:"label,omitempty"`

	OauthRefreshToken string `xml:"oauthRefreshToken,omitempty"`

	OauthScope string `xml:"oauthScope,omitempty"`

	OauthToken string `xml:"oauthToken,omitempty"`

	Password string `xml:"password,omitempty"`

	PrincipalType *ExternalPrincipalType `xml:"principalType,omitempty"`

	Protocol *AuthenticationProtocol `xml:"protocol,omitempty"`

	Username string `xml:"username,omitempty"`
}

type Network struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Network"`

	*Metadata

	AllowMembersToFlag bool `xml:"allowMembersToFlag,omitempty"`

	AllowedExtensions string `xml:"allowedExtensions,omitempty"`

	Branding *Branding `xml:"branding,omitempty"`

	CaseCommentEmailTemplate string `xml:"caseCommentEmailTemplate,omitempty"`

	ChangePasswordTemplate string `xml:"changePasswordTemplate,omitempty"`

	Description string `xml:"description,omitempty"`

	EmailSenderAddress string `xml:"emailSenderAddress,omitempty"`

	EmailSenderName string `xml:"emailSenderName,omitempty"`

	EnableGuestChatter bool `xml:"enableGuestChatter,omitempty"`

	EnableInvitation bool `xml:"enableInvitation,omitempty"`

	EnableKnowledgeable bool `xml:"enableKnowledgeable,omitempty"`

	EnableNicknameDisplay bool `xml:"enableNicknameDisplay,omitempty"`

	EnablePrivateMessages bool `xml:"enablePrivateMessages,omitempty"`

	EnableReputation bool `xml:"enableReputation,omitempty"`

	EnableSiteAsContainer bool `xml:"enableSiteAsContainer,omitempty"`

	FeedChannel string `xml:"feedChannel,omitempty"`

	ForgotPasswordTemplate string `xml:"forgotPasswordTemplate,omitempty"`

	LogoutUrl string `xml:"logoutUrl,omitempty"`

	MaxFileSizeKb int32 `xml:"maxFileSizeKb,omitempty"`

	NavigationLinkSet *NavigationLinkSet `xml:"navigationLinkSet,omitempty"`

	NetworkMemberGroups *NetworkMemberGroup `xml:"networkMemberGroups,omitempty"`

	NewSenderAddress string `xml:"newSenderAddress,omitempty"`

	PicassoSite string `xml:"picassoSite,omitempty"`

	ReputationLevels *ReputationLevelDefinitions `xml:"reputationLevels,omitempty"`

	ReputationPointsRules *ReputationPointsRules `xml:"reputationPointsRules,omitempty"`

	SelfRegProfile string `xml:"selfRegProfile,omitempty"`

	SelfRegistration bool `xml:"selfRegistration,omitempty"`

	SendWelcomeEmail bool `xml:"sendWelcomeEmail,omitempty"`

	Site string `xml:"site,omitempty"`

	Status *NetworkStatus `xml:"status,omitempty"`

	Tabs *NetworkTabSet `xml:"tabs,omitempty"`

	UrlPathPrefix string `xml:"urlPathPrefix,omitempty"`

	WelcomeTemplate string `xml:"welcomeTemplate,omitempty"`
}

type Branding struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Branding"`

	LoginFooterText string `xml:"loginFooterText,omitempty"`

	LoginLogo string `xml:"loginLogo,omitempty"`

	PageFooter string `xml:"pageFooter,omitempty"`

	PageHeader string `xml:"pageHeader,omitempty"`

	PrimaryColor string `xml:"primaryColor,omitempty"`

	PrimaryComplementColor string `xml:"primaryComplementColor,omitempty"`

	QuaternaryColor string `xml:"quaternaryColor,omitempty"`

	QuaternaryComplementColor string `xml:"quaternaryComplementColor,omitempty"`

	SecondaryColor string `xml:"secondaryColor,omitempty"`

	TertiaryColor string `xml:"tertiaryColor,omitempty"`

	TertiaryComplementColor string `xml:"tertiaryComplementColor,omitempty"`

	ZeronaryColor string `xml:"zeronaryColor,omitempty"`

	ZeronaryComplementColor string `xml:"zeronaryComplementColor,omitempty"`
}

type NavigationLinkSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata NavigationLinkSet"`

	NavigationMenuItem []*NavigationMenuItem `xml:"navigationMenuItem,omitempty"`
}

type NavigationMenuItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata NavigationMenuItem"`

	DefaultListViewId string `xml:"defaultListViewId,omitempty"`

	Label string `xml:"label,omitempty"`

	Position int32 `xml:"position,omitempty"`

	PubliclyAvailable bool `xml:"publiclyAvailable,omitempty"`

	Target string `xml:"target,omitempty"`

	TargetPreference string `xml:"targetPreference,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type NetworkMemberGroup struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata NetworkMemberGroup"`

	PermissionSet []string `xml:"permissionSet,omitempty"`

	Profile []string `xml:"profile,omitempty"`
}

type ReputationLevelDefinitions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReputationLevelDefinitions"`

	Level []*ReputationLevel `xml:"level,omitempty"`
}

type ReputationLevel struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReputationLevel"`

	Branding *ReputationBranding `xml:"branding,omitempty"`

	Label string `xml:"label,omitempty"`

	LowerThreshold float64 `xml:"lowerThreshold,omitempty"`
}

type ReputationBranding struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReputationBranding"`

	SmallImage string `xml:"smallImage,omitempty"`
}

type ReputationPointsRules struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReputationPointsRules"`

	PointsRule []*ReputationPointsRule `xml:"pointsRule,omitempty"`
}

type ReputationPointsRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReputationPointsRule"`

	EventType string `xml:"eventType,omitempty"`

	Points int32 `xml:"points,omitempty"`
}

type NetworkTabSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata NetworkTabSet"`

	CustomTab []string `xml:"customTab,omitempty"`

	DefaultTab string `xml:"defaultTab,omitempty"`

	StandardTab []string `xml:"standardTab,omitempty"`
}

type OpportunitySettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata OpportunitySettings"`

	*Metadata

	AutoActivateNewReminders bool `xml:"autoActivateNewReminders,omitempty"`

	EnableFindSimilarOpportunities bool `xml:"enableFindSimilarOpportunities,omitempty"`

	EnableOpportunityTeam bool `xml:"enableOpportunityTeam,omitempty"`

	EnableUpdateReminders bool `xml:"enableUpdateReminders,omitempty"`

	FindSimilarOppFilter *FindSimilarOppFilter `xml:"findSimilarOppFilter,omitempty"`

	PromptToAddProducts bool `xml:"promptToAddProducts,omitempty"`
}

type FindSimilarOppFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FindSimilarOppFilter"`

	SimilarOpportunitiesDisplayColumns []string `xml:"similarOpportunitiesDisplayColumns,omitempty"`

	SimilarOpportunitiesMatchFields []string `xml:"similarOpportunitiesMatchFields,omitempty"`
}

type OrderSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata OrderSettings"`

	*Metadata

	EnableNegativeQuantity bool `xml:"enableNegativeQuantity,omitempty"`

	EnableOrders bool `xml:"enableOrders,omitempty"`

	EnableReductionOrders bool `xml:"enableReductionOrders,omitempty"`
}

type OrgPreferenceSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata OrgPreferenceSettings"`

	*Metadata

	Preferences []*OrganizationSettingsDetail `xml:"preferences,omitempty"`
}

type OrganizationSettingsDetail struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata OrganizationSettingsDetail"`

	SettingName string `xml:"settingName,omitempty"`

	SettingValue bool `xml:"settingValue,omitempty"`
}

type Package struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Package"`

	*Metadata

	ApiAccessLevel *APIAccessLevel `xml:"apiAccessLevel,omitempty"`

	Description string `xml:"description,omitempty"`

	NamespacePrefix string `xml:"namespacePrefix,omitempty"`

	ObjectPermissions []*ProfileObjectPermissions `xml:"objectPermissions,omitempty"`

	PackageType string `xml:"packageType,omitempty"`

	PostInstallClass string `xml:"postInstallClass,omitempty"`

	SetupWeblink string `xml:"setupWeblink,omitempty"`

	Types []*PackageTypeMembers `xml:"types,omitempty"`

	UninstallClass string `xml:"uninstallClass,omitempty"`

	Version string `xml:"version,omitempty"`
}

type ProfileObjectPermissions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileObjectPermissions"`

	AllowCreate bool `xml:"allowCreate,omitempty"`

	AllowDelete bool `xml:"allowDelete,omitempty"`

	AllowEdit bool `xml:"allowEdit,omitempty"`

	AllowRead bool `xml:"allowRead,omitempty"`

	ModifyAllRecords bool `xml:"modifyAllRecords,omitempty"`

	Object string `xml:"object,omitempty"`

	ViewAllRecords bool `xml:"viewAllRecords,omitempty"`
}

type PackageTypeMembers struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PackageTypeMembers"`

	Members []string `xml:"members,omitempty"`

	Name string `xml:"name,omitempty"`
}

type PathAssistant struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PathAssistant"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	EntityName string `xml:"entityName,omitempty"`

	FieldName string `xml:"fieldName,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	PathAssistantSteps []*PathAssistantStep `xml:"pathAssistantSteps,omitempty"`

	RecordTypeName string `xml:"recordTypeName,omitempty"`
}

type PathAssistantStep struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PathAssistantStep"`

	FieldNames []string `xml:"fieldNames,omitempty"`

	Info string `xml:"info,omitempty"`

	PicklistValueName string `xml:"picklistValueName,omitempty"`
}

type PathAssistantSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PathAssistantSettings"`

	*Metadata

	PathAssistantEnabled bool `xml:"pathAssistantEnabled,omitempty"`
}

type PermissionSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSet"`

	*Metadata

	ApplicationVisibilities []*PermissionSetApplicationVisibility `xml:"applicationVisibilities,omitempty"`

	ClassAccesses []*PermissionSetApexClassAccess `xml:"classAccesses,omitempty"`

	CustomPermissions []*PermissionSetCustomPermissions `xml:"customPermissions,omitempty"`

	Description string `xml:"description,omitempty"`

	ExternalDataSourceAccesses []*PermissionSetExternalDataSourceAccess `xml:"externalDataSourceAccesses,omitempty"`

	FieldPermissions []*PermissionSetFieldPermissions `xml:"fieldPermissions,omitempty"`

	HasActivationRequired bool `xml:"hasActivationRequired,omitempty"`

	Label string `xml:"label,omitempty"`

	License string `xml:"license,omitempty"`

	ObjectPermissions []*PermissionSetObjectPermissions `xml:"objectPermissions,omitempty"`

	PageAccesses []*PermissionSetApexPageAccess `xml:"pageAccesses,omitempty"`

	RecordTypeVisibilities []*PermissionSetRecordTypeVisibility `xml:"recordTypeVisibilities,omitempty"`

	TabSettings []*PermissionSetTabSetting `xml:"tabSettings,omitempty"`

	UserPermissions []*PermissionSetUserPermission `xml:"userPermissions,omitempty"`
}

type PermissionSetApplicationVisibility struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetApplicationVisibility"`

	Application string `xml:"application,omitempty"`

	Visible bool `xml:"visible,omitempty"`
}

type PermissionSetApexClassAccess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetApexClassAccess"`

	ApexClass string `xml:"apexClass,omitempty"`

	Enabled bool `xml:"enabled,omitempty"`
}

type PermissionSetCustomPermissions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetCustomPermissions"`

	Enabled bool `xml:"enabled,omitempty"`

	Name string `xml:"name,omitempty"`
}

type PermissionSetExternalDataSourceAccess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetExternalDataSourceAccess"`

	Enabled bool `xml:"enabled,omitempty"`

	ExternalDataSource string `xml:"externalDataSource,omitempty"`
}

type PermissionSetFieldPermissions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetFieldPermissions"`

	Editable bool `xml:"editable,omitempty"`

	Field string `xml:"field,omitempty"`

	Readable bool `xml:"readable,omitempty"`
}

type PermissionSetObjectPermissions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetObjectPermissions"`

	AllowCreate bool `xml:"allowCreate,omitempty"`

	AllowDelete bool `xml:"allowDelete,omitempty"`

	AllowEdit bool `xml:"allowEdit,omitempty"`

	AllowRead bool `xml:"allowRead,omitempty"`

	ModifyAllRecords bool `xml:"modifyAllRecords,omitempty"`

	Object string `xml:"object,omitempty"`

	ViewAllRecords bool `xml:"viewAllRecords,omitempty"`
}

type PermissionSetApexPageAccess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetApexPageAccess"`

	ApexPage string `xml:"apexPage,omitempty"`

	Enabled bool `xml:"enabled,omitempty"`
}

type PermissionSetRecordTypeVisibility struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetRecordTypeVisibility"`

	RecordType string `xml:"recordType,omitempty"`

	Visible bool `xml:"visible,omitempty"`
}

type PermissionSetTabSetting struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetTabSetting"`

	Tab string `xml:"tab,omitempty"`

	Visibility *PermissionSetTabVisibility `xml:"visibility,omitempty"`
}

type PermissionSetUserPermission struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PermissionSetUserPermission"`

	Enabled bool `xml:"enabled,omitempty"`

	Name string `xml:"name,omitempty"`
}

type PersonListSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PersonListSettings"`

	*Metadata

	EnablePersonList bool `xml:"enablePersonList,omitempty"`
}

type PersonalJourneySettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PersonalJourneySettings"`

	*Metadata

	EnableExactTargetForSalesforceApps bool `xml:"enableExactTargetForSalesforceApps,omitempty"`
}

type PlatformCachePartition struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PlatformCachePartition"`

	*Metadata

	Description string `xml:"description,omitempty"`

	IsDefaultPartition bool `xml:"isDefaultPartition,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	PlatformCachePartitionTypes []*PlatformCachePartitionType `xml:"platformCachePartitionTypes,omitempty"`
}

type PlatformCachePartitionType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PlatformCachePartitionType"`

	AllocatedCapacity int32 `xml:"allocatedCapacity,omitempty"`

	AllocatedPurchasedCapacity int32 `xml:"allocatedPurchasedCapacity,omitempty"`

	AllocatedTrialCapacity int32 `xml:"allocatedTrialCapacity,omitempty"`

	CacheType *PlatformCacheType `xml:"cacheType,omitempty"`
}

type Portal struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Portal"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	Admin string `xml:"admin,omitempty"`

	DefaultLanguage string `xml:"defaultLanguage,omitempty"`

	Description string `xml:"description,omitempty"`

	EmailSenderAddress string `xml:"emailSenderAddress,omitempty"`

	EmailSenderName string `xml:"emailSenderName,omitempty"`

	EnableSelfCloseCase bool `xml:"enableSelfCloseCase,omitempty"`

	FooterDocument string `xml:"footerDocument,omitempty"`

	ForgotPassTemplate string `xml:"forgotPassTemplate,omitempty"`

	HeaderDocument string `xml:"headerDocument,omitempty"`

	IsSelfRegistrationActivated bool `xml:"isSelfRegistrationActivated,omitempty"`

	LoginHeaderDocument string `xml:"loginHeaderDocument,omitempty"`

	LogoDocument string `xml:"logoDocument,omitempty"`

	LogoutUrl string `xml:"logoutUrl,omitempty"`

	NewCommentTemplate string `xml:"newCommentTemplate,omitempty"`

	NewPassTemplate string `xml:"newPassTemplate,omitempty"`

	NewUserTemplate string `xml:"newUserTemplate,omitempty"`

	OwnerNotifyTemplate string `xml:"ownerNotifyTemplate,omitempty"`

	SelfRegNewUserUrl string `xml:"selfRegNewUserUrl,omitempty"`

	SelfRegUserDefaultProfile string `xml:"selfRegUserDefaultProfile,omitempty"`

	SelfRegUserDefaultRole *PortalRoles `xml:"selfRegUserDefaultRole,omitempty"`

	SelfRegUserTemplate string `xml:"selfRegUserTemplate,omitempty"`

	ShowActionConfirmation bool `xml:"showActionConfirmation,omitempty"`

	StylesheetDocument string `xml:"stylesheetDocument,omitempty"`

	Type_ *PortalType `xml:"type,omitempty"`
}

type PostTemplate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PostTemplate"`

	*Metadata

	Default_ bool `xml:"default,omitempty"`

	Description string `xml:"description,omitempty"`

	Fields []string `xml:"fields,omitempty"`

	Label string `xml:"label,omitempty"`
}

type ProductSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProductSettings"`

	*Metadata

	EnableCascadeActivateToRelatedPrices bool `xml:"enableCascadeActivateToRelatedPrices,omitempty"`

	EnableQuantitySchedule bool `xml:"enableQuantitySchedule,omitempty"`

	EnableRevenueSchedule bool `xml:"enableRevenueSchedule,omitempty"`
}

type Profile struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Profile"`

	*Metadata

	ApplicationVisibilities []*ProfileApplicationVisibility `xml:"applicationVisibilities,omitempty"`

	ClassAccesses []*ProfileApexClassAccess `xml:"classAccesses,omitempty"`

	Custom bool `xml:"custom,omitempty"`

	CustomPermissions []*ProfileCustomPermissions `xml:"customPermissions,omitempty"`

	Description string `xml:"description,omitempty"`

	ExternalDataSourceAccesses []*ProfileExternalDataSourceAccess `xml:"externalDataSourceAccesses,omitempty"`

	FieldPermissions []*ProfileFieldLevelSecurity `xml:"fieldPermissions,omitempty"`

	LayoutAssignments []*ProfileLayoutAssignment `xml:"layoutAssignments,omitempty"`

	LoginHours *ProfileLoginHours `xml:"loginHours,omitempty"`

	LoginIpRanges []*ProfileLoginIpRange `xml:"loginIpRanges,omitempty"`

	ObjectPermissions []*ProfileObjectPermissions `xml:"objectPermissions,omitempty"`

	PageAccesses []*ProfileApexPageAccess `xml:"pageAccesses,omitempty"`

	ProfileActionOverrides []*ProfileActionOverride `xml:"profileActionOverrides,omitempty"`

	RecordTypeVisibilities []*ProfileRecordTypeVisibility `xml:"recordTypeVisibilities,omitempty"`

	TabVisibilities []*ProfileTabVisibility `xml:"tabVisibilities,omitempty"`

	UserLicense string `xml:"userLicense,omitempty"`

	UserPermissions []*ProfileUserPermission `xml:"userPermissions,omitempty"`
}

type ProfileApplicationVisibility struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileApplicationVisibility"`

	Application string `xml:"application,omitempty"`

	Default_ bool `xml:"default,omitempty"`

	Visible bool `xml:"visible,omitempty"`
}

type ProfileApexClassAccess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileApexClassAccess"`

	ApexClass string `xml:"apexClass,omitempty"`

	Enabled bool `xml:"enabled,omitempty"`
}

type ProfileCustomPermissions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileCustomPermissions"`

	Enabled bool `xml:"enabled,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ProfileExternalDataSourceAccess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileExternalDataSourceAccess"`

	Enabled bool `xml:"enabled,omitempty"`

	ExternalDataSource string `xml:"externalDataSource,omitempty"`
}

type ProfileFieldLevelSecurity struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileFieldLevelSecurity"`

	Editable bool `xml:"editable,omitempty"`

	Field string `xml:"field,omitempty"`

	Readable bool `xml:"readable,omitempty"`
}

type ProfileLayoutAssignment struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileLayoutAssignment"`

	Layout string `xml:"layout,omitempty"`

	RecordType string `xml:"recordType,omitempty"`
}

type ProfileLoginHours struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileLoginHours"`

	FridayEnd string `xml:"fridayEnd,omitempty"`

	FridayStart string `xml:"fridayStart,omitempty"`

	MondayEnd string `xml:"mondayEnd,omitempty"`

	MondayStart string `xml:"mondayStart,omitempty"`

	SaturdayEnd string `xml:"saturdayEnd,omitempty"`

	SaturdayStart string `xml:"saturdayStart,omitempty"`

	SundayEnd string `xml:"sundayEnd,omitempty"`

	SundayStart string `xml:"sundayStart,omitempty"`

	ThursdayEnd string `xml:"thursdayEnd,omitempty"`

	ThursdayStart string `xml:"thursdayStart,omitempty"`

	TuesdayEnd string `xml:"tuesdayEnd,omitempty"`

	TuesdayStart string `xml:"tuesdayStart,omitempty"`

	WednesdayEnd string `xml:"wednesdayEnd,omitempty"`

	WednesdayStart string `xml:"wednesdayStart,omitempty"`
}

type ProfileLoginIpRange struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileLoginIpRange"`

	Description string `xml:"description,omitempty"`

	EndAddress string `xml:"endAddress,omitempty"`

	StartAddress string `xml:"startAddress,omitempty"`
}

type ProfileApexPageAccess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileApexPageAccess"`

	ApexPage string `xml:"apexPage,omitempty"`

	Enabled bool `xml:"enabled,omitempty"`
}

type ProfileActionOverride struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileActionOverride"`

	ActionName string `xml:"actionName,omitempty"`

	Content string `xml:"content,omitempty"`

	FormFactor *FormFactor `xml:"formFactor,omitempty"`

	PageOrSobjectType string `xml:"pageOrSobjectType,omitempty"`

	RecordType string `xml:"recordType,omitempty"`

	Type_ *ActionOverrideType `xml:"type,omitempty"`
}

type ProfileRecordTypeVisibility struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileRecordTypeVisibility"`

	Default_ bool `xml:"default,omitempty"`

	PersonAccountDefault bool `xml:"personAccountDefault,omitempty"`

	RecordType string `xml:"recordType,omitempty"`

	Visible bool `xml:"visible,omitempty"`
}

type ProfileTabVisibility struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileTabVisibility"`

	Tab string `xml:"tab,omitempty"`

	Visibility *TabVisibility `xml:"visibility,omitempty"`
}

type ProfileUserPermission struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ProfileUserPermission"`

	Enabled bool `xml:"enabled,omitempty"`

	Name string `xml:"name,omitempty"`
}

type Queue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Queue"`

	*Metadata

	DoesSendEmailToMembers bool `xml:"doesSendEmailToMembers,omitempty"`

	Email string `xml:"email,omitempty"`

	Name string `xml:"name,omitempty"`

	QueueSobject []*QueueSobject `xml:"queueSobject,omitempty"`
}

type QueueSobject struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QueueSobject"`

	SobjectType string `xml:"sobjectType,omitempty"`
}

type QuickAction struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuickAction"`

	*Metadata

	Canvas string `xml:"canvas,omitempty"`

	Description string `xml:"description,omitempty"`

	FieldOverrides []*FieldOverride `xml:"fieldOverrides,omitempty"`

	Height int32 `xml:"height,omitempty"`

	Icon string `xml:"icon,omitempty"`

	IsProtected bool `xml:"isProtected,omitempty"`

	Label string `xml:"label,omitempty"`

	LightningComponent string `xml:"lightningComponent,omitempty"`

	OptionsCreateFeedItem bool `xml:"optionsCreateFeedItem,omitempty"`

	Page string `xml:"page,omitempty"`

	QuickActionLayout *QuickActionLayout `xml:"quickActionLayout,omitempty"`

	QuickActionSendEmailOptions *QuickActionSendEmailOptions `xml:"quickActionSendEmailOptions,omitempty"`

	StandardLabel *QuickActionLabel `xml:"standardLabel,omitempty"`

	SuccessMessage string `xml:"successMessage,omitempty"`

	TargetObject string `xml:"targetObject,omitempty"`

	TargetParentField string `xml:"targetParentField,omitempty"`

	TargetRecordType string `xml:"targetRecordType,omitempty"`

	Type_ *QuickActionType `xml:"type,omitempty"`

	Width int32 `xml:"width,omitempty"`
}

type FieldOverride struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FieldOverride"`

	Field string `xml:"field,omitempty"`

	Formula string `xml:"formula,omitempty"`

	LiteralValue string `xml:"literalValue,omitempty"`
}

type QuickActionLayout struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuickActionLayout"`

	LayoutSectionStyle *LayoutSectionStyle `xml:"layoutSectionStyle,omitempty"`

	QuickActionLayoutColumns []*QuickActionLayoutColumn `xml:"quickActionLayoutColumns,omitempty"`
}

type QuickActionLayoutColumn struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuickActionLayoutColumn"`

	QuickActionLayoutItems []*QuickActionLayoutItem `xml:"quickActionLayoutItems,omitempty"`
}

type QuickActionLayoutItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuickActionLayoutItem"`

	EmptySpace bool `xml:"emptySpace,omitempty"`

	Field string `xml:"field,omitempty"`

	UiBehavior *UiBehavior `xml:"uiBehavior,omitempty"`
}

type QuickActionSendEmailOptions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuickActionSendEmailOptions"`

	DefaultEmailTemplateName string `xml:"defaultEmailTemplateName,omitempty"`

	IgnoreDefaultEmailTemplateSubject bool `xml:"ignoreDefaultEmailTemplateSubject,omitempty"`
}

type QuoteSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata QuoteSettings"`

	*Metadata

	EnableQuote bool `xml:"enableQuote,omitempty"`
}

type RemoteSiteSetting struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RemoteSiteSetting"`

	*Metadata

	Description string `xml:"description,omitempty"`

	DisableProtocolSecurity bool `xml:"disableProtocolSecurity,omitempty"`

	IsActive bool `xml:"isActive,omitempty"`

	Url string `xml:"url,omitempty"`
}

type Report struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Report"`

	*Metadata

	Aggregates []*ReportAggregate `xml:"aggregates,omitempty"`

	Block []*Report `xml:"block,omitempty"`

	BlockInfo *ReportBlockInfo `xml:"blockInfo,omitempty"`

	Buckets []*ReportBucketField `xml:"buckets,omitempty"`

	Chart *ReportChart `xml:"chart,omitempty"`

	ColorRanges []*ReportColorRange `xml:"colorRanges,omitempty"`

	Columns []*ReportColumn `xml:"columns,omitempty"`

	CrossFilters []*ReportCrossFilter `xml:"crossFilters,omitempty"`

	Currency *CurrencyIsoCode `xml:"currency,omitempty"`

	DataCategoryFilters []*ReportDataCategoryFilter `xml:"dataCategoryFilters,omitempty"`

	Description string `xml:"description,omitempty"`

	Division string `xml:"division,omitempty"`

	Filter *ReportFilter `xml:"filter,omitempty"`

	FolderName string `xml:"folderName,omitempty"`

	Format *ReportFormat `xml:"format,omitempty"`

	GroupingsAcross []*ReportGrouping `xml:"groupingsAcross,omitempty"`

	GroupingsDown []*ReportGrouping `xml:"groupingsDown,omitempty"`

	HistoricalSelector *ReportHistoricalSelector `xml:"historicalSelector,omitempty"`

	Name string `xml:"name,omitempty"`

	NumSubscriptions int32 `xml:"numSubscriptions,omitempty"`

	Params []*ReportParam `xml:"params,omitempty"`

	ReportType string `xml:"reportType,omitempty"`

	RoleHierarchyFilter string `xml:"roleHierarchyFilter,omitempty"`

	RowLimit int32 `xml:"rowLimit,omitempty"`

	Scope string `xml:"scope,omitempty"`

	ShowCurrentDate bool `xml:"showCurrentDate,omitempty"`

	ShowDetails bool `xml:"showDetails,omitempty"`

	SortColumn string `xml:"sortColumn,omitempty"`

	SortOrder *SortOrder `xml:"sortOrder,omitempty"`

	TerritoryHierarchyFilter string `xml:"territoryHierarchyFilter,omitempty"`

	TimeFrameFilter *ReportTimeFrameFilter `xml:"timeFrameFilter,omitempty"`

	UserFilter string `xml:"userFilter,omitempty"`
}

type ReportAggregate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportAggregate"`

	AcrossGroupingContext string `xml:"acrossGroupingContext,omitempty"`

	CalculatedFormula string `xml:"calculatedFormula,omitempty"`

	Datatype *ReportAggregateDatatype `xml:"datatype,omitempty"`

	Description string `xml:"description,omitempty"`

	DeveloperName string `xml:"developerName,omitempty"`

	DownGroupingContext string `xml:"downGroupingContext,omitempty"`

	IsActive bool `xml:"isActive,omitempty"`

	IsCrossBlock bool `xml:"isCrossBlock,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	ReportType string `xml:"reportType,omitempty"`

	Scale int32 `xml:"scale,omitempty"`
}

type ReportBlockInfo struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportBlockInfo"`

	AggregateReferences []*ReportAggregateReference `xml:"aggregateReferences,omitempty"`

	BlockId string `xml:"blockId,omitempty"`

	JoinTable string `xml:"joinTable,omitempty"`
}

type ReportAggregateReference struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportAggregateReference"`

	Aggregate string `xml:"aggregate,omitempty"`
}

type ReportBucketField struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportBucketField"`

	BucketType *ReportBucketFieldType `xml:"bucketType,omitempty"`

	DeveloperName string `xml:"developerName,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	NullTreatment *ReportFormulaNullTreatment `xml:"nullTreatment,omitempty"`

	OtherBucketLabel string `xml:"otherBucketLabel,omitempty"`

	SourceColumnName string `xml:"sourceColumnName,omitempty"`

	UseOther bool `xml:"useOther,omitempty"`

	Values []*ReportBucketFieldValue `xml:"values,omitempty"`
}

type ReportBucketFieldValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportBucketFieldValue"`

	SourceValues []*ReportBucketFieldSourceValue `xml:"sourceValues,omitempty"`

	Value string `xml:"value,omitempty"`
}

type ReportBucketFieldSourceValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportBucketFieldSourceValue"`

	From string `xml:"from,omitempty"`

	SourceValue string `xml:"sourceValue,omitempty"`

	To string `xml:"to,omitempty"`
}

type ReportChart struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportChart"`

	BackgroundColor1 string `xml:"backgroundColor1,omitempty"`

	BackgroundColor2 string `xml:"backgroundColor2,omitempty"`

	BackgroundFadeDir *ChartBackgroundDirection `xml:"backgroundFadeDir,omitempty"`

	ChartSummaries []*ChartSummary `xml:"chartSummaries,omitempty"`

	ChartType *ChartType `xml:"chartType,omitempty"`

	EnableHoverLabels bool `xml:"enableHoverLabels,omitempty"`

	ExpandOthers bool `xml:"expandOthers,omitempty"`

	GroupingColumn string `xml:"groupingColumn,omitempty"`

	LegendPosition *ChartLegendPosition `xml:"legendPosition,omitempty"`

	Location *ChartPosition `xml:"location,omitempty"`

	SecondaryGroupingColumn string `xml:"secondaryGroupingColumn,omitempty"`

	ShowAxisLabels bool `xml:"showAxisLabels,omitempty"`

	ShowPercentage bool `xml:"showPercentage,omitempty"`

	ShowTotal bool `xml:"showTotal,omitempty"`

	ShowValues bool `xml:"showValues,omitempty"`

	Size *ReportChartSize `xml:"size,omitempty"`

	SummaryAxisManualRangeEnd float64 `xml:"summaryAxisManualRangeEnd,omitempty"`

	SummaryAxisManualRangeStart float64 `xml:"summaryAxisManualRangeStart,omitempty"`

	SummaryAxisRange *ChartRangeType `xml:"summaryAxisRange,omitempty"`

	TextColor string `xml:"textColor,omitempty"`

	TextSize int32 `xml:"textSize,omitempty"`

	Title string `xml:"title,omitempty"`

	TitleColor string `xml:"titleColor,omitempty"`

	TitleSize int32 `xml:"titleSize,omitempty"`
}

type ReportColorRange struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportColorRange"`

	Aggregate *ReportSummaryType `xml:"aggregate,omitempty"`

	ColumnName string `xml:"columnName,omitempty"`

	HighBreakpoint float64 `xml:"highBreakpoint,omitempty"`

	HighColor string `xml:"highColor,omitempty"`

	LowBreakpoint float64 `xml:"lowBreakpoint,omitempty"`

	LowColor string `xml:"lowColor,omitempty"`

	MidColor string `xml:"midColor,omitempty"`
}

type ReportColumn struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportColumn"`

	AggregateTypes []*ReportSummaryType `xml:"aggregateTypes,omitempty"`

	Field string `xml:"field,omitempty"`

	ReverseColors bool `xml:"reverseColors,omitempty"`

	ShowChanges bool `xml:"showChanges,omitempty"`
}

type ReportCrossFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportCrossFilter"`

	CriteriaItems []*ReportFilterItem `xml:"criteriaItems,omitempty"`

	Operation *ObjectFilterOperator `xml:"operation,omitempty"`

	PrimaryTableColumn string `xml:"primaryTableColumn,omitempty"`

	RelatedTable string `xml:"relatedTable,omitempty"`

	RelatedTableJoinColumn string `xml:"relatedTableJoinColumn,omitempty"`
}

type ReportFilterItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportFilterItem"`

	Column string `xml:"column,omitempty"`

	ColumnToColumn bool `xml:"columnToColumn,omitempty"`

	IsUnlocked bool `xml:"isUnlocked,omitempty"`

	Operator *FilterOperation `xml:"operator,omitempty"`

	Snapshot string `xml:"snapshot,omitempty"`

	Value string `xml:"value,omitempty"`
}

type ReportDataCategoryFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportDataCategoryFilter"`

	DataCategory string `xml:"dataCategory,omitempty"`

	DataCategoryGroup string `xml:"dataCategoryGroup,omitempty"`

	Operator *DataCategoryFilterOperation `xml:"operator,omitempty"`
}

type ReportFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportFilter"`

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	CriteriaItems []*ReportFilterItem `xml:"criteriaItems,omitempty"`

	Language *Language `xml:"language,omitempty"`
}

type ReportGrouping struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportGrouping"`

	AggregateType *ReportAggrType `xml:"aggregateType,omitempty"`

	DateGranularity *UserDateGranularity `xml:"dateGranularity,omitempty"`

	Field string `xml:"field,omitempty"`

	SortByName string `xml:"sortByName,omitempty"`

	SortOrder *SortOrder `xml:"sortOrder,omitempty"`

	SortType *ReportSortType `xml:"sortType,omitempty"`
}

type ReportHistoricalSelector struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportHistoricalSelector"`

	Snapshot []string `xml:"snapshot,omitempty"`
}

type ReportParam struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportParam"`

	Name string `xml:"name,omitempty"`

	Value string `xml:"value,omitempty"`
}

type ReportTimeFrameFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportTimeFrameFilter"`

	DateColumn string `xml:"dateColumn,omitempty"`

	EndDate time.Time `xml:"endDate,omitempty"`

	Interval *UserDateInterval `xml:"interval,omitempty"`

	StartDate time.Time `xml:"startDate,omitempty"`
}

type ReportType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportType"`

	*Metadata

	Autogenerated bool `xml:"autogenerated,omitempty"`

	BaseObject string `xml:"baseObject,omitempty"`

	Category *ReportTypeCategory `xml:"category,omitempty"`

	Deployed bool `xml:"deployed,omitempty"`

	Description string `xml:"description,omitempty"`

	Join *ObjectRelationship `xml:"join,omitempty"`

	Label string `xml:"label,omitempty"`

	Sections []*ReportLayoutSection `xml:"sections,omitempty"`
}

type ObjectRelationship struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ObjectRelationship"`

	Join *ObjectRelationship `xml:"join,omitempty"`

	OuterJoin bool `xml:"outerJoin,omitempty"`

	Relationship string `xml:"relationship,omitempty"`
}

type ReportLayoutSection struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportLayoutSection"`

	Columns []*ReportTypeColumn `xml:"columns,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`
}

type ReportTypeColumn struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportTypeColumn"`

	CheckedByDefault bool `xml:"checkedByDefault,omitempty"`

	DisplayNameOverride string `xml:"displayNameOverride,omitempty"`

	Field string `xml:"field,omitempty"`

	Table string `xml:"table,omitempty"`
}

type RoleOrTerritory struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RoleOrTerritory"`

	*Metadata

	CaseAccessLevel string `xml:"caseAccessLevel,omitempty"`

	ContactAccessLevel string `xml:"contactAccessLevel,omitempty"`

	Description string `xml:"description,omitempty"`

	MayForecastManagerShare bool `xml:"mayForecastManagerShare,omitempty"`

	Name string `xml:"name,omitempty"`

	OpportunityAccessLevel string `xml:"opportunityAccessLevel,omitempty"`
}

type Role struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Role"`

	*RoleOrTerritory

	ParentRole string `xml:"parentRole,omitempty"`
}

type Territory struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Territory"`

	*RoleOrTerritory

	AccountAccessLevel string `xml:"accountAccessLevel,omitempty"`

	ParentTerritory string `xml:"parentTerritory,omitempty"`
}

type SamlSsoConfig struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SamlSsoConfig"`

	*Metadata

	AttributeName string `xml:"attributeName,omitempty"`

	AttributeNameIdFormat string `xml:"attributeNameIdFormat,omitempty"`

	DecryptionCertificate string `xml:"decryptionCertificate,omitempty"`

	ErrorUrl string `xml:"errorUrl,omitempty"`

	ExecutionUserId string `xml:"executionUserId,omitempty"`

	IdentityLocation *SamlIdentityLocationType `xml:"identityLocation,omitempty"`

	IdentityMapping *SamlIdentityType `xml:"identityMapping,omitempty"`

	Issuer string `xml:"issuer,omitempty"`

	LoginUrl string `xml:"loginUrl,omitempty"`

	LogoutUrl string `xml:"logoutUrl,omitempty"`

	Name string `xml:"name,omitempty"`

	OauthTokenEndpoint string `xml:"oauthTokenEndpoint,omitempty"`

	RedirectBinding bool `xml:"redirectBinding,omitempty"`

	RequestSignatureMethod string `xml:"requestSignatureMethod,omitempty"`

	SalesforceLoginUrl string `xml:"salesforceLoginUrl,omitempty"`

	SamlEntityId string `xml:"samlEntityId,omitempty"`

	SamlJitHandlerId string `xml:"samlJitHandlerId,omitempty"`

	SamlVersion *SamlType `xml:"samlVersion,omitempty"`

	UserProvisioning bool `xml:"userProvisioning,omitempty"`

	ValidationCert string `xml:"validationCert,omitempty"`
}

type SearchSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SearchSettings"`

	*Metadata

	DocumentContentSearchEnabled bool `xml:"documentContentSearchEnabled,omitempty"`

	OptimizeSearchForCJKEnabled bool `xml:"optimizeSearchForCJKEnabled,omitempty"`

	RecentlyViewedUsersForBlankLookupEnabled bool `xml:"recentlyViewedUsersForBlankLookupEnabled,omitempty"`

	SearchSettingsByObject *SearchSettingsByObject `xml:"searchSettingsByObject,omitempty"`

	SidebarAutoCompleteEnabled bool `xml:"sidebarAutoCompleteEnabled,omitempty"`

	SidebarDropDownListEnabled bool `xml:"sidebarDropDownListEnabled,omitempty"`

	SidebarLimitToItemsIOwnCheckboxEnabled bool `xml:"sidebarLimitToItemsIOwnCheckboxEnabled,omitempty"`

	SingleSearchResultShortcutEnabled bool `xml:"singleSearchResultShortcutEnabled,omitempty"`

	SpellCorrectKnowledgeSearchEnabled bool `xml:"spellCorrectKnowledgeSearchEnabled,omitempty"`
}

type SearchSettingsByObject struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SearchSettingsByObject"`

	SearchSettingsByObject []*ObjectSearchSetting `xml:"searchSettingsByObject,omitempty"`
}

type ObjectSearchSetting struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ObjectSearchSetting"`

	EnhancedLookupEnabled bool `xml:"enhancedLookupEnabled,omitempty"`

	LookupAutoCompleteEnabled bool `xml:"lookupAutoCompleteEnabled,omitempty"`

	Name string `xml:"name,omitempty"`

	ResultsPerPageCount int32 `xml:"resultsPerPageCount,omitempty"`
}

type SecuritySettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SecuritySettings"`

	*Metadata

	NetworkAccess *NetworkAccess `xml:"networkAccess,omitempty"`

	PasswordPolicies *PasswordPolicies `xml:"passwordPolicies,omitempty"`

	SessionSettings *SessionSettings `xml:"sessionSettings,omitempty"`
}

type NetworkAccess struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata NetworkAccess"`

	IpRanges []*IpRange `xml:"ipRanges,omitempty"`
}

type IpRange struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata IpRange"`

	Description string `xml:"description,omitempty"`

	End string `xml:"end,omitempty"`

	Start string `xml:"start,omitempty"`
}

type PasswordPolicies struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PasswordPolicies"`

	ApiOnlyUserHomePageURL string `xml:"apiOnlyUserHomePageURL,omitempty"`

	Complexity *Complexity `xml:"complexity,omitempty"`

	Expiration *Expiration `xml:"expiration,omitempty"`

	HistoryRestriction string `xml:"historyRestriction,omitempty"`

	LockoutInterval *LockoutInterval `xml:"lockoutInterval,omitempty"`

	MaxLoginAttempts *MaxLoginAttempts `xml:"maxLoginAttempts,omitempty"`

	MinimumPasswordLength string `xml:"minimumPasswordLength,omitempty"`

	MinimumPasswordLifetime bool `xml:"minimumPasswordLifetime,omitempty"`

	ObscureSecretAnswer bool `xml:"obscureSecretAnswer,omitempty"`

	PasswordAssistanceMessage string `xml:"passwordAssistanceMessage,omitempty"`

	PasswordAssistanceURL string `xml:"passwordAssistanceURL,omitempty"`

	QuestionRestriction *QuestionRestriction `xml:"questionRestriction,omitempty"`
}

type SessionSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SessionSettings"`

	DisableTimeoutWarning bool `xml:"disableTimeoutWarning,omitempty"`

	EnableCSPOnEmail bool `xml:"enableCSPOnEmail,omitempty"`

	EnableCSRFOnGet bool `xml:"enableCSRFOnGet,omitempty"`

	EnableCSRFOnPost bool `xml:"enableCSRFOnPost,omitempty"`

	EnableCacheAndAutocomplete bool `xml:"enableCacheAndAutocomplete,omitempty"`

	EnableClickjackNonsetupSFDC bool `xml:"enableClickjackNonsetupSFDC,omitempty"`

	EnableClickjackNonsetupUser bool `xml:"enableClickjackNonsetupUser,omitempty"`

	EnableClickjackNonsetupUserHeaderless bool `xml:"enableClickjackNonsetupUserHeaderless,omitempty"`

	EnableClickjackSetup bool `xml:"enableClickjackSetup,omitempty"`

	EnablePostForSessions bool `xml:"enablePostForSessions,omitempty"`

	EnableSMSIdentity bool `xml:"enableSMSIdentity,omitempty"`

	EnforceIpRangesEveryRequest bool `xml:"enforceIpRangesEveryRequest,omitempty"`

	ForceLogoutOnSessionTimeout bool `xml:"forceLogoutOnSessionTimeout,omitempty"`

	ForceRelogin bool `xml:"forceRelogin,omitempty"`

	LockSessionsToDomain bool `xml:"lockSessionsToDomain,omitempty"`

	LockSessionsToIp bool `xml:"lockSessionsToIp,omitempty"`

	LogoutURL string `xml:"logoutURL,omitempty"`

	SecurityCentralKillSession bool `xml:"securityCentralKillSession,omitempty"`

	SessionTimeout *SessionTimeout `xml:"sessionTimeout,omitempty"`
}

type SharingBaseRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharingBaseRule"`

	*Metadata

	AccessLevel string `xml:"accessLevel,omitempty"`

	AccountSettings *AccountSharingRuleSettings `xml:"accountSettings,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	SharedTo *SharedTo `xml:"sharedTo,omitempty"`
}

type AccountSharingRuleSettings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AccountSharingRuleSettings"`

	CaseAccessLevel string `xml:"caseAccessLevel,omitempty"`

	ContactAccessLevel string `xml:"contactAccessLevel,omitempty"`

	OpportunityAccessLevel string `xml:"opportunityAccessLevel,omitempty"`
}

type SharingCriteriaRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharingCriteriaRule"`

	*SharingBaseRule

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	CriteriaItems []*FilterItem `xml:"criteriaItems,omitempty"`
}

type SharingOwnerRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharingOwnerRule"`

	*SharingBaseRule

	SharedFrom *SharedTo `xml:"sharedFrom,omitempty"`
}

type SharingTerritoryRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharingTerritoryRule"`

	*SharingOwnerRule
}

type SharingRules struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharingRules"`

	*Metadata

	SharingCriteriaRules []*SharingCriteriaRule `xml:"sharingCriteriaRules,omitempty"`

	SharingOwnerRules []*SharingOwnerRule `xml:"sharingOwnerRules,omitempty"`

	SharingTerritoryRules []*SharingTerritoryRule `xml:"sharingTerritoryRules,omitempty"`
}

type SharingSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SharingSet"`

	*Metadata

	AccessMappings []*AccessMapping `xml:"accessMappings,omitempty"`

	Description string `xml:"description,omitempty"`

	Name string `xml:"name,omitempty"`

	Profiles []string `xml:"profiles,omitempty"`
}

type AccessMapping struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AccessMapping"`

	AccessLevel string `xml:"accessLevel,omitempty"`

	Object string `xml:"object,omitempty"`

	ObjectField string `xml:"objectField,omitempty"`

	UserField string `xml:"userField,omitempty"`
}

type Skill struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Skill"`

	*Metadata

	Assignments *SkillAssignments `xml:"assignments,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`
}

type SkillAssignments struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SkillAssignments"`

	Profiles *SkillProfileAssignments `xml:"profiles,omitempty"`

	Users *SkillUserAssignments `xml:"users,omitempty"`
}

type SkillProfileAssignments struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SkillProfileAssignments"`

	Profile []string `xml:"profile,omitempty"`
}

type SkillUserAssignments struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SkillUserAssignments"`

	User []string `xml:"user,omitempty"`
}

type StandardValueSet struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata StandardValueSet"`

	*Metadata

	Sorted bool `xml:"sorted,omitempty"`

	StandardValue []*StandardValue `xml:"standardValue,omitempty"`
}

type StandardValueSetTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata StandardValueSetTranslation"`

	*Metadata

	ValueTranslation []*ValueTranslation `xml:"valueTranslation,omitempty"`
}

type SynonymDictionary struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SynonymDictionary"`

	*Metadata

	Groups []*SynonymGroup `xml:"groups,omitempty"`

	IsProtected bool `xml:"isProtected,omitempty"`

	Label string `xml:"label,omitempty"`
}

type SynonymGroup struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SynonymGroup"`

	Languages []*Language `xml:"languages,omitempty"`

	Terms []string `xml:"terms,omitempty"`
}

type Territory2 struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Territory2"`

	*Metadata

	AccountAccessLevel string `xml:"accountAccessLevel,omitempty"`

	CaseAccessLevel string `xml:"caseAccessLevel,omitempty"`

	ContactAccessLevel string `xml:"contactAccessLevel,omitempty"`

	CustomFields []*FieldValue `xml:"customFields,omitempty"`

	Description string `xml:"description,omitempty"`

	Name string `xml:"name,omitempty"`

	OpportunityAccessLevel string `xml:"opportunityAccessLevel,omitempty"`

	ParentTerritory string `xml:"parentTerritory,omitempty"`

	RuleAssociations []*Territory2RuleAssociation `xml:"ruleAssociations,omitempty"`

	Territory2Type string `xml:"territory2Type,omitempty"`
}

type FieldValue struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata FieldValue"`

	Name string `xml:"name,omitempty"`

	Value interface{} `xml:"value,omitempty"`
}

type Territory2RuleAssociation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Territory2RuleAssociation"`

	Inherited bool `xml:"inherited,omitempty"`

	RuleName string `xml:"ruleName,omitempty"`
}

type Territory2Model struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Territory2Model"`

	*Metadata

	CustomFields []*FieldValue `xml:"customFields,omitempty"`

	Description string `xml:"description,omitempty"`

	Name string `xml:"name,omitempty"`
}

type Territory2Rule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Territory2Rule"`

	*Metadata

	Active bool `xml:"active,omitempty"`

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	Name string `xml:"name,omitempty"`

	ObjectType string `xml:"objectType,omitempty"`

	RuleItems []*Territory2RuleItem `xml:"ruleItems,omitempty"`
}

type Territory2RuleItem struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Territory2RuleItem"`

	Field string `xml:"field,omitempty"`

	Operation *FilterOperation `xml:"operation,omitempty"`

	Value string `xml:"value,omitempty"`
}

type Territory2Settings struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Territory2Settings"`

	*Metadata

	DefaultAccountAccessLevel string `xml:"defaultAccountAccessLevel,omitempty"`

	DefaultCaseAccessLevel string `xml:"defaultCaseAccessLevel,omitempty"`

	DefaultContactAccessLevel string `xml:"defaultContactAccessLevel,omitempty"`

	DefaultOpportunityAccessLevel string `xml:"defaultOpportunityAccessLevel,omitempty"`

	OpportunityFilterSettings *Territory2SettingsOpportunityFilter `xml:"opportunityFilterSettings,omitempty"`
}

type Territory2SettingsOpportunityFilter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Territory2SettingsOpportunityFilter"`

	ApexClassName string `xml:"apexClassName,omitempty"`

	EnableFilter bool `xml:"enableFilter,omitempty"`

	RunOnCreate bool `xml:"runOnCreate,omitempty"`
}

type Territory2Type struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Territory2Type"`

	*Metadata

	Description string `xml:"description,omitempty"`

	Name string `xml:"name,omitempty"`

	Priority int32 `xml:"priority,omitempty"`
}

type TransactionSecurityPolicy struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata TransactionSecurityPolicy"`

	*Metadata

	Action *TransactionSecurityAction `xml:"action,omitempty"`

	Active bool `xml:"active,omitempty"`

	ApexClass string `xml:"apexClass,omitempty"`

	EventType *MonitoredEvents `xml:"eventType,omitempty"`

	ExecutionUser string `xml:"executionUser,omitempty"`

	ResourceName string `xml:"resourceName,omitempty"`
}

type TransactionSecurityAction struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata TransactionSecurityAction"`

	Block bool `xml:"block,omitempty"`

	EndSession bool `xml:"endSession,omitempty"`

	Notifications []*TransactionSecurityNotification `xml:"notifications,omitempty"`

	TwoFactorAuthentication bool `xml:"twoFactorAuthentication,omitempty"`
}

type TransactionSecurityNotification struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata TransactionSecurityNotification"`

	InApp bool `xml:"inApp,omitempty"`

	SendEmail bool `xml:"sendEmail,omitempty"`

	User string `xml:"user,omitempty"`
}

type Translations struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Translations"`

	*Metadata

	CustomApplications []*CustomApplicationTranslation `xml:"customApplications,omitempty"`

	CustomDataTypeTranslations []*CustomDataTypeTranslation `xml:"customDataTypeTranslations,omitempty"`

	CustomLabels []*CustomLabelTranslation `xml:"customLabels,omitempty"`

	CustomPageWebLinks []*CustomPageWebLinkTranslation `xml:"customPageWebLinks,omitempty"`

	CustomTabs []*CustomTabTranslation `xml:"customTabs,omitempty"`

	QuickActions []*GlobalQuickActionTranslation `xml:"quickActions,omitempty"`

	ReportTypes []*ReportTypeTranslation `xml:"reportTypes,omitempty"`

	Scontrols []*ScontrolTranslation `xml:"scontrols,omitempty"`
}

type CustomApplicationTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomApplicationTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type CustomDataTypeTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomDataTypeTranslation"`

	Components []*CustomDataTypeComponentTranslation `xml:"components,omitempty"`

	CustomDataTypeName string `xml:"customDataTypeName,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`
}

type CustomDataTypeComponentTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomDataTypeComponentTranslation"`

	DeveloperSuffix string `xml:"developerSuffix,omitempty"`

	Label string `xml:"label,omitempty"`
}

type CustomLabelTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomLabelTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type CustomPageWebLinkTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomPageWebLinkTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type CustomTabTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata CustomTabTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type GlobalQuickActionTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata GlobalQuickActionTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ReportTypeTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportTypeTranslation"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`

	Sections []*ReportTypeSectionTranslation `xml:"sections,omitempty"`
}

type ReportTypeSectionTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportTypeSectionTranslation"`

	Columns []*ReportTypeColumnTranslation `xml:"columns,omitempty"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ReportTypeColumnTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReportTypeColumnTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type ScontrolTranslation struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ScontrolTranslation"`

	Label string `xml:"label,omitempty"`

	Name string `xml:"name,omitempty"`
}

type VisualizationPlugin struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata VisualizationPlugin"`

	*Metadata

	Description string `xml:"description,omitempty"`

	DeveloperName string `xml:"developerName,omitempty"`

	Icon string `xml:"icon,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	VisualizationResources []*VisualizationResource `xml:"visualizationResources,omitempty"`

	VisualizationTypes []*VisualizationType `xml:"visualizationTypes,omitempty"`
}

type VisualizationResource struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata VisualizationResource"`

	Description string `xml:"description,omitempty"`

	File string `xml:"file,omitempty"`

	Rank int32 `xml:"rank,omitempty"`

	Type_ *VisualizationResourceType `xml:"type,omitempty"`
}

type VisualizationType struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata VisualizationType"`

	Description string `xml:"description,omitempty"`

	DeveloperName string `xml:"developerName,omitempty"`

	Icon string `xml:"icon,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	ScriptBootstrapMethod string `xml:"scriptBootstrapMethod,omitempty"`
}

type WaveApplication struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WaveApplication"`

	*Metadata

	AssetIcon string `xml:"assetIcon,omitempty"`

	Description string `xml:"description,omitempty"`

	Folder string `xml:"folder,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	Shares []*FolderShare `xml:"shares,omitempty"`

	TemplateOrigin string `xml:"templateOrigin,omitempty"`

	TemplateVersion string `xml:"templateVersion,omitempty"`
}

type WaveDataset struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WaveDataset"`

	*Metadata

	Application string `xml:"application,omitempty"`

	Description string `xml:"description,omitempty"`

	MasterLabel string `xml:"masterLabel,omitempty"`

	TemplateAssetSourceName string `xml:"templateAssetSourceName,omitempty"`
}

type WaveTemplateBundle struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WaveTemplateBundle"`

	*Metadata

	AssetIcon string `xml:"assetIcon,omitempty"`

	AssetVersion float64 `xml:"assetVersion,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	TemplateType string `xml:"templateType,omitempty"`
}

type Workflow struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Workflow"`

	*Metadata

	Alerts []*WorkflowAlert `xml:"alerts,omitempty"`

	FieldUpdates []*WorkflowFieldUpdate `xml:"fieldUpdates,omitempty"`

	FlowActions []*WorkflowFlowAction `xml:"flowActions,omitempty"`

	KnowledgePublishes []*WorkflowKnowledgePublish `xml:"knowledgePublishes,omitempty"`

	OutboundMessages []*WorkflowOutboundMessage `xml:"outboundMessages,omitempty"`

	Rules []*WorkflowRule `xml:"rules,omitempty"`

	Send []*WorkflowSend `xml:"send,omitempty"`

	Tasks []*WorkflowTask `xml:"tasks,omitempty"`
}

type WorkflowAlert struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowAlert"`

	*WorkflowAction

	CcEmails []string `xml:"ccEmails,omitempty"`

	Description string `xml:"description,omitempty"`

	Protected bool `xml:"protected,omitempty"`

	Recipients []*WorkflowEmailRecipient `xml:"recipients,omitempty"`

	SenderAddress string `xml:"senderAddress,omitempty"`

	SenderType *ActionEmailSenderType `xml:"senderType,omitempty"`

	Template string `xml:"template,omitempty"`
}

type WorkflowAction struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowAction"`

	*Metadata
}

type WorkflowFieldUpdate struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowFieldUpdate"`

	*WorkflowAction

	Description string `xml:"description,omitempty"`

	Field string `xml:"field,omitempty"`

	Formula string `xml:"formula,omitempty"`

	LiteralValue string `xml:"literalValue,omitempty"`

	LookupValue string `xml:"lookupValue,omitempty"`

	LookupValueType *LookupValueType `xml:"lookupValueType,omitempty"`

	Name string `xml:"name,omitempty"`

	NotifyAssignee bool `xml:"notifyAssignee,omitempty"`

	Operation *FieldUpdateOperation `xml:"operation,omitempty"`

	Protected bool `xml:"protected,omitempty"`

	ReevaluateOnChange bool `xml:"reevaluateOnChange,omitempty"`

	TargetObject string `xml:"targetObject,omitempty"`
}

type WorkflowFlowAction struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowFlowAction"`

	*WorkflowAction

	Description string `xml:"description,omitempty"`

	Flow string `xml:"flow,omitempty"`

	FlowInputs []*WorkflowFlowActionParameter `xml:"flowInputs,omitempty"`

	Label string `xml:"label,omitempty"`

	Language string `xml:"language,omitempty"`

	Protected bool `xml:"protected,omitempty"`
}

type WorkflowFlowActionParameter struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowFlowActionParameter"`

	Name string `xml:"name,omitempty"`

	Value string `xml:"value,omitempty"`
}

type WorkflowKnowledgePublish struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowKnowledgePublish"`

	*WorkflowAction

	Action *KnowledgeWorkflowAction `xml:"action,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	Language string `xml:"language,omitempty"`

	Protected bool `xml:"protected,omitempty"`
}

type WorkflowOutboundMessage struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowOutboundMessage"`

	*WorkflowAction

	ApiVersion float64 `xml:"apiVersion,omitempty"`

	Description string `xml:"description,omitempty"`

	EndpointUrl string `xml:"endpointUrl,omitempty"`

	Fields []string `xml:"fields,omitempty"`

	IncludeSessionId bool `xml:"includeSessionId,omitempty"`

	IntegrationUser string `xml:"integrationUser,omitempty"`

	Name string `xml:"name,omitempty"`

	Protected bool `xml:"protected,omitempty"`

	UseDeadLetterQueue bool `xml:"useDeadLetterQueue,omitempty"`
}

type WorkflowSend struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowSend"`

	*WorkflowAction

	Action *SendAction `xml:"action,omitempty"`

	Description string `xml:"description,omitempty"`

	Label string `xml:"label,omitempty"`

	Language string `xml:"language,omitempty"`

	Protected bool `xml:"protected,omitempty"`
}

type WorkflowTask struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowTask"`

	*WorkflowAction

	AssignedTo string `xml:"assignedTo,omitempty"`

	AssignedToType *ActionTaskAssignedToTypes `xml:"assignedToType,omitempty"`

	Description string `xml:"description,omitempty"`

	DueDateOffset int32 `xml:"dueDateOffset,omitempty"`

	NotifyAssignee bool `xml:"notifyAssignee,omitempty"`

	OffsetFromField string `xml:"offsetFromField,omitempty"`

	Priority string `xml:"priority,omitempty"`

	Protected bool `xml:"protected,omitempty"`

	Status string `xml:"status,omitempty"`

	Subject string `xml:"subject,omitempty"`
}

type WorkflowEmailRecipient struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowEmailRecipient"`

	Field string `xml:"field,omitempty"`

	Recipient string `xml:"recipient,omitempty"`

	Type_ *ActionEmailRecipientTypes `xml:"type,omitempty"`
}

type WorkflowRule struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowRule"`

	*Metadata

	Actions []*WorkflowActionReference `xml:"actions,omitempty"`

	Active bool `xml:"active,omitempty"`

	BooleanFilter string `xml:"booleanFilter,omitempty"`

	CriteriaItems []*FilterItem `xml:"criteriaItems,omitempty"`

	Description string `xml:"description,omitempty"`

	Formula string `xml:"formula,omitempty"`

	TriggerType *WorkflowTriggerTypes `xml:"triggerType,omitempty"`

	WorkflowTimeTriggers []*WorkflowTimeTrigger `xml:"workflowTimeTriggers,omitempty"`
}

type WorkflowTimeTrigger struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata WorkflowTimeTrigger"`

	Actions []*WorkflowActionReference `xml:"actions,omitempty"`

	OffsetFromField string `xml:"offsetFromField,omitempty"`

	TimeLength string `xml:"timeLength,omitempty"`

	WorkflowTimeTriggerUnit *WorkflowTimeUnits `xml:"workflowTimeTriggerUnit,omitempty"`
}

type XOrgHub struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata XOrgHub"`

	*Metadata

	Label string `xml:"label,omitempty"`

	LockSharedObjects bool `xml:"lockSharedObjects,omitempty"`

	PermissionSets []string `xml:"permissionSets,omitempty"`

	SharedObjects []*XOrgHubSharedObject `xml:"sharedObjects,omitempty"`
}

type XOrgHubSharedObject struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata XOrgHubSharedObject"`

	Fields []string `xml:"fields,omitempty"`

	Name string `xml:"name,omitempty"`
}

type SaveResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata SaveResult"`

	Errors []*Error `xml:"errors,omitempty"`

	FullName string `xml:"fullName,omitempty"`

	Success bool `xml:"success,omitempty"`
}

type Error struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata Error"`

	ExtendedErrorDetails []*ExtendedErrorDetails `xml:"extendedErrorDetails,omitempty"`

	Fields []string `xml:"fields,omitempty"`

	Message string `xml:"message,omitempty"`

	StatusCode *StatusCode `xml:"statusCode,omitempty"`
}

type ExtendedErrorDetails struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ExtendedErrorDetails"`

	ExtendedErrorCode *ExtendedErrorCode `xml:"extendedErrorCode,omitempty"`
}

type DeleteResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DeleteResult"`

	Errors []*Error `xml:"errors,omitempty"`

	FullName string `xml:"fullName,omitempty"`

	Success bool `xml:"success,omitempty"`
}

type DeployOptions struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DeployOptions"`

	AllowMissingFiles bool `xml:"allowMissingFiles,omitempty"`

	AutoUpdatePackage bool `xml:"autoUpdatePackage,omitempty"`

	CheckOnly bool `xml:"checkOnly,omitempty"`

	IgnoreWarnings bool `xml:"ignoreWarnings,omitempty"`

	PerformRetrieve bool `xml:"performRetrieve,omitempty"`

	PurgeOnDelete bool `xml:"purgeOnDelete,omitempty"`

	RollbackOnError bool `xml:"rollbackOnError,omitempty"`

	RunTests []string `xml:"runTests,omitempty"`

	SinglePackage bool `xml:"singlePackage,omitempty"`

	TestLevel *TestLevel `xml:"testLevel,omitempty"`
}

type AsyncResult struct {
// Modify
//	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata AsyncResult"`

	Done bool `xml:"done,omitempty"`

	Id *ID `xml:"id,omitempty"`

	Message string `xml:"message,omitempty"`

	State *AsyncRequestState `xml:"state,omitempty"`

	StatusCode *StatusCode `xml:"statusCode,omitempty"`
}

type DescribeMetadataResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DescribeMetadataResult"`

	MetadataObjects []*DescribeMetadataObject `xml:"metadataObjects,omitempty"`

	OrganizationNamespace string `xml:"organizationNamespace,omitempty"`

	PartialSaveAllowed bool `xml:"partialSaveAllowed,omitempty"`

	TestRequired bool `xml:"testRequired,omitempty"`
}

type DescribeMetadataObject struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DescribeMetadataObject"`

	ChildXmlNames []string `xml:"childXmlNames,omitempty"`

	DirectoryName string `xml:"directoryName,omitempty"`

	InFolder bool `xml:"inFolder,omitempty"`

	MetaFile bool `xml:"metaFile,omitempty"`

	Suffix string `xml:"suffix,omitempty"`

	XmlName string `xml:"xmlName,omitempty"`
}

type DescribeValueTypeResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata DescribeValueTypeResult"`

	ApiCreatable bool `xml:"apiCreatable,omitempty"`

	ApiDeletable bool `xml:"apiDeletable,omitempty"`

	ApiReadable bool `xml:"apiReadable,omitempty"`

	ApiUpdatable bool `xml:"apiUpdatable,omitempty"`

	ParentField *ValueTypeField `xml:"parentField,omitempty"`

	ValueTypeFields []*ValueTypeField `xml:"valueTypeFields,omitempty"`
}

type ValueTypeField struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ValueTypeField"`

	Fields []*ValueTypeField `xml:"fields,omitempty"`

	ForeignKeyDomain []string `xml:"foreignKeyDomain,omitempty"`

	IsForeignKey bool `xml:"isForeignKey,omitempty"`

	IsNameField bool `xml:"isNameField,omitempty"`

	MinOccurs int32 `xml:"minOccurs,omitempty"`

	Name string `xml:"name,omitempty"`

	PicklistValues []*PicklistEntry `xml:"picklistValues,omitempty"`

	SoapType string `xml:"soapType,omitempty"`

	ValueRequired bool `xml:"valueRequired,omitempty"`
}

type PicklistEntry struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata PicklistEntry"`

	Active bool `xml:"active,omitempty"`

	DefaultValue bool `xml:"defaultValue,omitempty"`

	Label string `xml:"label,omitempty"`

	ValidFor string `xml:"validFor,omitempty"`

	Value string `xml:"value,omitempty"`
}

type ListMetadataQuery struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ListMetadataQuery"`

	Folder string `xml:"folder,omitempty"`

	Type_ string `xml:"type,omitempty"`
}

type ReadResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata ReadResult"`

	Records []*Metadata `xml:"records,omitempty"`
}

type RetrieveRequest struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata RetrieveRequest"`

	ApiVersion float64 `xml:"apiVersion,omitempty"`

	PackageNames []string `xml:"packageNames,omitempty"`

	SinglePackage bool `xml:"singlePackage,omitempty"`

	SpecificFiles []string `xml:"specificFiles,omitempty"`

	Unpackaged *Package `xml:"unpackaged,omitempty"`
}

type UpsertResult struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata UpsertResult"`

	Created bool `xml:"created,omitempty"`

	Errors []*Error `xml:"errors,omitempty"`

	FullName string `xml:"fullName,omitempty"`

	Success bool `xml:"success,omitempty"`
}

type LogInfo struct {
	XMLName xml.Name `xml:"http://soap.sforce.com/2006/04/metadata LogInfo"`

	Category *LogCategory `xml:"category,omitempty"`

	Level *LogCategoryLevel `xml:"level,omitempty"`
}

type LoginRequest struct {
	XMLName xml.Name `xml:"urn:partner.soap.sforce.com login"`
	Username string `xml:"username"`
	Password string `xml:"password"`
}

type LoginResponse struct {
	XMLName xml.Name `xml:"urn:partner.soap.sforce.com loginResponse"`
	LoginResult LoginResult `xml:"result"`
}

type LoginResult struct {
	MetadataServerUrl string `xml:"metadataServerUrl"`
	PasswordExpired bool `xml:"passwordExpired"`
	Sandbox bool `xml:"sandbox`
	ServerUrl string `xml:"serverUrl"`
	SessionId string `xml:"sessionId"`
	UserId *ID `xml:"userId"`
//	UserInfo *UserInfo `xml:"userInfo"`
}

type MetadataPortType struct {
	client *SOAPClient
}

func (service *MetadataPortType) SetServerUrl(url string) {
	service.client.SetServerUrl(url)
}


func NewMetadataPortType(url string, tls bool, auth *BasicAuth) *MetadataPortType {
	if url == "" {
		url = ""
	}
	client := NewSOAPClient(url, tls, auth)

	return &MetadataPortType{
		client: client,
	}
}

func (service *MetadataPortType) SetHeader(header interface{}) {
	service.client.SetHeader(header)
}

/* Cancels a metadata deploy. */
func (service *MetadataPortType) CancelDeploy(request *CancelDeploy) (*CancelDeployResponse, error) {
	response := new(CancelDeployResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Check the current status of an asyncronous deploy call. */
func (service *MetadataPortType) CheckDeployStatus(request *CheckDeployStatus) (*CheckDeployStatusResponse, error) {
	response := new(CheckDeployStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Check the current status of an asyncronous deploy call. */
func (service *MetadataPortType) CheckRetrieveStatus(request *CheckRetrieveStatus) (*CheckRetrieveStatusResponse, error) {
	response := new(CheckRetrieveStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates metadata entries synchronously. */
func (service *MetadataPortType) CreateMetadata(request *CreateMetadata) (*CreateMetadataResponse, error) {
	response := new(CreateMetadataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Deletes metadata entries synchronously. */
func (service *MetadataPortType) DeleteMetadata(request *DeleteMetadata) (*DeleteMetadataResponse, error) {
	response := new(DeleteMetadataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Deploys a zipfile full of metadata entries asynchronously. */
func (service *MetadataPortType) Deploy(request *Deploy) (*DeployResponse, error) {
	response := new(DeployResponse)
	// modify
	err := service.client.Call("''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Deploys a previously validated payload without running tests. */
func (service *MetadataPortType) DeployRecentValidation(request *DeployRecentValidation) (*DeployRecentValidationResponse, error) {
	response := new(DeployRecentValidationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describes features of the metadata API. */
func (service *MetadataPortType) DescribeMetadata(request *DescribeMetadata) (*DescribeMetadataResponse, error) {
	response := new(DescribeMetadataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe a complex value type */
func (service *MetadataPortType) DescribeValueType(request *DescribeValueType) (*DescribeValueTypeResponse, error) {
	response := new(DescribeValueTypeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Lists the available metadata components. */
func (service *MetadataPortType) ListMetadata(request *ListMetadata) (*ListMetadataResponse, error) {
	response := new(ListMetadataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Reads metadata entries synchronously. */
func (service *MetadataPortType) ReadMetadata(request *ReadMetadata) (*ReadMetadataResponse, error) {
	response := new(ReadMetadataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Renames a metadata entry synchronously. */
func (service *MetadataPortType) RenameMetadata(request *RenameMetadata) (*RenameMetadataResponse, error) {
	response := new(RenameMetadataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves a set of individually specified metadata entries. */
func (service *MetadataPortType) Retrieve(request *Retrieve) (*RetrieveResponse, error) {
	response := new(RetrieveResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Updates metadata entries synchronously. */
func (service *MetadataPortType) UpdateMetadata(request *UpdateMetadata) (*UpdateMetadataResponse, error) {
	response := new(UpdateMetadataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Upserts metadata entries synchronously. */
func (service *MetadataPortType) UpsertMetadata(request *UpsertMetadata) (*UpsertMetadataResponse, error) {
	response := new(UpsertMetadataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Upserts metadata entries synchronously. */
func (service *MetadataPortType) Login(request *LoginRequest) (*LoginResponse, error) {
	response := new(LoginResponse)
	err := service.client.Call("''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeader
	Body    SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`

	Header interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	Fault   *SOAPFault  `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type BasicAuth struct {
	Login    string
	Password string
}

type SOAPClient struct {
	url    string
	tls    bool
	auth   *BasicAuth
	header interface{}
}

func (b *SOAPBody) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var (
		token    xml.Token
		err      error
		consumed bool
	)

	Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Fault = &SOAPFault{}
				b.Content = nil

				err = d.DecodeElement(b.Fault, &se)
				if err != nil {
					return err
				}

				consumed = true
			} else {
				if err = d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (f *SOAPFault) Error() string {
	return f.String
}

func NewSOAPClient(url string, tls bool, auth *BasicAuth) *SOAPClient {
	return &SOAPClient{
		url:  url,
		tls:  tls,
		auth: auth,
	}
}

func (s *SOAPClient) SetHeader(header interface{}) {
	s.header = header
}

func (s *SOAPClient) SetServerUrl(url string) {
	s.url = url
}

func (s *SOAPClient) Call(soapAction string, request, response interface{}) error {
	envelope := SOAPEnvelope{}

	if s.header != nil {
		envelope.Header = &SOAPHeader{Header: s.header}
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)


	encoder := xml.NewEncoder(buffer)

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	if err != nil {
		return err
	}
	if s.auth != nil {
		req.SetBasicAuth(s.auth.Login, s.auth.Password)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	req.Header.Set("User-Agent", "gowsdl/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if len(rawbody) == 0 {
		log.Println("empty response")
		return nil
	}

	respEnvelope := new(SOAPEnvelope)
	respEnvelope.Body = SOAPBody{Content: response}
	err = xml.Unmarshal(rawbody, respEnvelope)
	if err != nil {
		return err
	}

	fault := respEnvelope.Body.Fault
	if fault != nil {
		return fault
	}


	return nil
}
