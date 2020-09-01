// Code generated by "stringer -type=OperationType"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[opUnused-0]
	_ = x[opPing-1]
	_ = x[opJoin-2]
	_ = x[opCreateAccount-3]
	_ = x[opLogin-4]
	_ = x[opSendCrashLog-5]
	_ = x[opSendTraceRoute-6]
	_ = x[opSendVfxStats-7]
	_ = x[opSendGamePingInfo-8]
	_ = x[opCreateCharacter-9]
	_ = x[opDeleteCharacter-10]
	_ = x[opSelectCharacter-11]
	_ = x[opRedeemKeycode-12]
	_ = x[opGetGameServerByCluster-13]
	_ = x[opGetActiveSubscription-14]
	_ = x[opGetShopPurchaseUrl-15]
	_ = x[opGetBuyTrialDetails-16]
	_ = x[opGetReferralSeasonDetails-17]
	_ = x[opGetReferralLink-18]
	_ = x[opGetAvailableTrialKeys-19]
	_ = x[opGetShopTilesForCategory-20]
	_ = x[opMove-21]
	_ = x[opAttackStart-22]
	_ = x[opCastStart-23]
	_ = x[opCastCancel-24]
	_ = x[opTerminateToggleSpell-25]
	_ = x[opChannelingCancel-26]
	_ = x[opAttackBuildingStart-27]
	_ = x[opInventoryDestroyItem-28]
	_ = x[opInventoryMoveItem-29]
	_ = x[opInventoryRecoverItem-30]
	_ = x[opInventoryRecoverAllItems-31]
	_ = x[opInventorySplitStack-32]
	_ = x[opInventorySplitStackInto-33]
	_ = x[opGetClusterData-34]
	_ = x[opChangeCluster-35]
	_ = x[opConsoleCommand-36]
	_ = x[opChatMessage-37]
	_ = x[opReportClientError-38]
	_ = x[opRegisterToObject-39]
	_ = x[opUnRegisterFromObject-40]
	_ = x[opCraftBuildingChangeSettings-41]
	_ = x[opCraftBuildingTakeMoney-42]
	_ = x[opRepairBuildingChangeSettings-43]
	_ = x[opRepairBuildingTakeMoney-44]
	_ = x[opActionBuildingChangeSettings-45]
	_ = x[opHarvestStart-46]
	_ = x[opHarvestCancel-47]
	_ = x[opTakeSilver-48]
	_ = x[opActionOnBuildingStart-49]
	_ = x[opActionOnBuildingCancel-50]
	_ = x[opItemRerollQualityStart-51]
	_ = x[opItemRerollQualityCancel-52]
	_ = x[opInstallResourceStart-53]
	_ = x[opInstallResourceCancel-54]
	_ = x[opInstallSilver-55]
	_ = x[opBuildingFillNutrition-56]
	_ = x[opBuildingChangeRenovationState-57]
	_ = x[opBuildingBuySkin-58]
	_ = x[opBuildingClaim-59]
	_ = x[opBuildingGiveup-60]
	_ = x[opBuildingNutritionSilverStorageDeposit-61]
	_ = x[opBuildingNutritionSilverStorageWithdraw-62]
	_ = x[opBuildingNutritionSilverRewardSet-63]
	_ = x[opConstructionSiteCreate-64]
	_ = x[opPlaceableObjectPlace-65]
	_ = x[opPlaceableObjectPlaceCancel-66]
	_ = x[opPlaceableObjectPickup-67]
	_ = x[opFurnitureObjectUse-68]
	_ = x[opFarmableHarvest-69]
	_ = x[opFarmableFinishGrownItem-70]
	_ = x[opFarmableDestroy-71]
	_ = x[opFarmableGetProduct-72]
	_ = x[opFarmableFill-73]
	_ = x[opTearDownConstructionSite-74]
	_ = x[opCastleGateUse-75]
	_ = x[opAuctionCreateOffer-76]
	_ = x[opAuctionCreateRequest-77]
	_ = x[opAuctionGetOffers-78]
	_ = x[opAuctionGetRequests-79]
	_ = x[opAuctionBuyOffer-80]
	_ = x[opAuctionAbortAuction-81]
	_ = x[opAuctionModifyAuction-82]
	_ = x[opAuctionAbortOffer-83]
	_ = x[opAuctionAbortRequest-84]
	_ = x[opAuctionSellRequest-85]
	_ = x[opAuctionGetFinishedAuctions-86]
	_ = x[opAuctionGetFinishedAuctionsCount-87]
	_ = x[opAuctionFetchAuction-88]
	_ = x[opAuctionGetMyOpenOffers-89]
	_ = x[opAuctionGetMyOpenRequests-90]
	_ = x[opAuctionGetMyOpenAuctions-91]
	_ = x[opAuctionGetItemAverageStats-92]
	_ = x[opAuctionGetItemAverageValue-93]
	_ = x[opContainerOpen-94]
	_ = x[opContainerClose-95]
	_ = x[opContainerManageSubContainer-96]
	_ = x[opRespawn-97]
	_ = x[opSuicide-98]
	_ = x[opJoinGuild-99]
	_ = x[opLeaveGuild-100]
	_ = x[opCreateGuild-101]
	_ = x[opInviteToGuild-102]
	_ = x[opDeclineGuildInvitation-103]
	_ = x[opKickFromGuild-104]
	_ = x[opDuellingChallengePlayer-105]
	_ = x[opDuellingAcceptChallenge-106]
	_ = x[opDuellingDenyChallenge-107]
	_ = x[opChangeClusterTax-108]
	_ = x[opClaimTerritory-109]
	_ = x[opGiveUpTerritory-110]
	_ = x[opChangeTerritoryAccessRights-111]
	_ = x[opGetMonolithInfo-112]
	_ = x[opGetClaimInfo-113]
	_ = x[opGetAttackInfo-114]
	_ = x[opGetTerritorySeasonPoints-115]
	_ = x[opGetAttackSchedule-116]
	_ = x[opScheduleAttack-117]
	_ = x[opGetMatches-118]
	_ = x[opGetMatchDetails-119]
	_ = x[opJoinMatch-120]
	_ = x[opLeaveMatch-121]
	_ = x[opChangeChatSettings-122]
	_ = x[opLogoutStart-123]
	_ = x[opLogoutCancel-124]
	_ = x[opClaimOrbStart-125]
	_ = x[opClaimOrbCancel-126]
	_ = x[opMatchLootChestOpeningStart-127]
	_ = x[opMatchLootChestOpeningCancel-128]
	_ = x[opDepositToGuildAccount-129]
	_ = x[opWithdrawalFromAccount-130]
	_ = x[opChangeGuildPayUpkeepFlag-131]
	_ = x[opChangeGuildTax-132]
	_ = x[opGetMyTerritories-133]
	_ = x[opMorganaCommand-134]
	_ = x[opGetServerInfo-135]
	_ = x[opInviteMercenaryToMatch-136]
	_ = x[opSubscribeToCluster-137]
	_ = x[opAnswerMercenaryInvitation-138]
	_ = x[opGetCharacterEquipment-139]
	_ = x[opGetCharacterSteamAchievements-140]
	_ = x[opGetCharacterStats-141]
	_ = x[opGetKillHistoryDetails-142]
	_ = x[opLearnMasteryLevel-143]
	_ = x[opReSpecAchievement-144]
	_ = x[opChangeAvatar-145]
	_ = x[opGetRankings-146]
	_ = x[opGetRank-147]
	_ = x[opGetGvgSeasonRankings-148]
	_ = x[opGetGvgSeasonRank-149]
	_ = x[opGetGvgSeasonHistoryRankings-150]
	_ = x[opGetGvgSeasonGuildMemberHistory-151]
	_ = x[opKickFromGvGMatch-152]
	_ = x[opGetChestLogs-153]
	_ = x[opGetAccessRightLogs-154]
	_ = x[opGetGuildAccountLogs-155]
	_ = x[opGetGuildAccountLogsLargeAmount-156]
	_ = x[opInviteToPlayerTrade-157]
	_ = x[opPlayerTradeCancel-158]
	_ = x[opPlayerTradeInvitationAccept-159]
	_ = x[opPlayerTradeAddItem-160]
	_ = x[opPlayerTradeRemoveItem-161]
	_ = x[opPlayerTradeAcceptTrade-162]
	_ = x[opPlayerTradeSetSilverOrGold-163]
	_ = x[opSendMiniMapPing-164]
	_ = x[opStuck-165]
	_ = x[opBuyRealEstate-166]
	_ = x[opClaimRealEstate-167]
	_ = x[opGiveUpRealEstate-168]
	_ = x[opChangeRealEstateOutline-169]
	_ = x[opGetMailInfos-170]
	_ = x[opGetMailCount-171]
	_ = x[opReadMail-172]
	_ = x[opSendNewMail-173]
	_ = x[opDeleteMail-174]
	_ = x[opMarkMailUnread-175]
	_ = x[opClaimAttachmentFromMail-176]
	_ = x[opUpdateLfgInfo-177]
	_ = x[opGetLfgInfos-178]
	_ = x[opGetMyGuildLfgInfo-179]
	_ = x[opGetLfgDescriptionText-180]
	_ = x[opLfgApplyToGuild-181]
	_ = x[opAnswerLfgGuildApplication-182]
	_ = x[opRegisterChatPeer-183]
	_ = x[opSendChatMessage-184]
	_ = x[opJoinChatChannel-185]
	_ = x[opLeaveChatChannel-186]
	_ = x[opSendWhisperMessage-187]
	_ = x[opSay-188]
	_ = x[opPlayEmote-189]
	_ = x[opStopEmote-190]
	_ = x[opGetClusterMapInfo-191]
	_ = x[opAccessRightsChangeSettings-192]
	_ = x[opMount-193]
	_ = x[opMountCancel-194]
	_ = x[opBuyJourney-195]
	_ = x[opSetSaleStatusForEstate-196]
	_ = x[opResolveGuildOrPlayerName-197]
	_ = x[opGetRespawnInfos-198]
	_ = x[opMakeHome-199]
	_ = x[opLeaveHome-200]
	_ = x[opResurrectionReply-201]
	_ = x[opAllianceCreate-202]
	_ = x[opAllianceDisband-203]
	_ = x[opAllianceGetMemberInfos-204]
	_ = x[opAllianceInvite-205]
	_ = x[opAllianceAnswerInvitation-206]
	_ = x[opAllianceCancelInvitation-207]
	_ = x[opAllianceKickGuild-208]
	_ = x[opAllianceLeave-209]
	_ = x[opAllianceChangeGoldPaymentFlag-210]
	_ = x[opAllianceGetDetailInfo-211]
	_ = x[opGetIslandInfos-212]
	_ = x[opAbandonMyIsland-213]
	_ = x[opBuyMyIsland-214]
	_ = x[opBuyGuildIsland-215]
	_ = x[opAbandonGuildIsland-216]
	_ = x[opUpgradeMyIsland-217]
	_ = x[opUpgradeGuildIsland-218]
	_ = x[opMoveMyIsland-219]
	_ = x[opMoveGuildIsland-220]
	_ = x[opTerritoryFillNutrition-221]
	_ = x[opTeleportBack-222]
	_ = x[opPartyInvitePlayer-223]
	_ = x[opPartyAnswerInvitation-224]
	_ = x[opPartyLeave-225]
	_ = x[opPartyKickPlayer-226]
	_ = x[opPartyMakeLeader-227]
	_ = x[opPartyChangeLootSetting-228]
	_ = x[opPartyMarkObject-229]
	_ = x[opPartySetRole-230]
	_ = x[opGetGuildMOTD-231]
	_ = x[opSetGuildMOTD-232]
	_ = x[opExitEnterStart-233]
	_ = x[opExitEnterCancel-234]
	_ = x[opQuestGiverRequest-235]
	_ = x[opGoldMarketGetBuyOffer-236]
	_ = x[opGoldMarketGetBuyOfferFromSilver-237]
	_ = x[opGoldMarketGetSellOffer-238]
	_ = x[opGoldMarketGetSellOfferFromSilver-239]
	_ = x[opGoldMarketBuyGold-240]
	_ = x[opGoldMarketSellGold-241]
	_ = x[opGoldMarketCreateSellOrder-242]
	_ = x[opGoldMarketCreateBuyOrder-243]
	_ = x[opGoldMarketGetInfos-244]
	_ = x[opGoldMarketCancelOrder-245]
	_ = x[opGoldMarketGetAverageInfo-246]
	_ = x[opSiegeCampClaimStart-247]
	_ = x[opSiegeCampClaimCancel-248]
	_ = x[opTreasureChestUsingStart-249]
	_ = x[opTreasureChestUsingCancel-250]
	_ = x[opUseLootChest-251]
	_ = x[opUseShrine-252]
	_ = x[opLaborerStartJob-253]
	_ = x[opLaborerTakeJobLoot-254]
	_ = x[opLaborerDismiss-255]
	_ = x[opLaborerMove-256]
	_ = x[opLaborerBuyItem-257]
	_ = x[opLaborerUpgrade-258]
	_ = x[opBuyPremium-259]
	_ = x[opBuyTrial-260]
	_ = x[opRealEstateGetAuctionData-261]
	_ = x[opRealEstateBidOnAuction-262]
	_ = x[opGetSiegeCampCooldown-263]
	_ = x[opFriendInvite-264]
	_ = x[opFriendAnswerInvitation-265]
	_ = x[opFriendCancelnvitation-266]
	_ = x[opFriendRemove-267]
	_ = x[opInventoryStack-268]
	_ = x[opInventorySort-269]
	_ = x[opEquipmentItemChangeSpell-270]
	_ = x[opExpeditionRegister-271]
	_ = x[opExpeditionRegisterCancel-272]
	_ = x[opJoinExpedition-273]
	_ = x[opDeclineExpeditionInvitation-274]
	_ = x[opVoteStart-275]
	_ = x[opVoteDoVote-276]
	_ = x[opRatingDoRate-277]
	_ = x[opEnteringExpeditionStart-278]
	_ = x[opEnteringExpeditionCancel-279]
	_ = x[opActivateExpeditionCheckPoint-280]
	_ = x[opArenaRegister-281]
	_ = x[opArenaRegisterCancel-282]
	_ = x[opArenaLeave-283]
	_ = x[opJoinArenaMatch-284]
	_ = x[opDeclineArenaInvitation-285]
	_ = x[opEnteringArenaStart-286]
	_ = x[opEnteringArenaCancel-287]
	_ = x[opArenaCustomMatch-288]
	_ = x[opArenaCustomMatchCreate-289]
	_ = x[opUpdateCharacterStatement-290]
	_ = x[opBoostFarmable-291]
	_ = x[opGetStrikeHistory-292]
	_ = x[opUseFunction-293]
	_ = x[opUsePortalEntrance-294]
	_ = x[opResetPortalBinding-295]
	_ = x[opQueryPortalBinding-296]
	_ = x[opClaimPaymentTransaction-297]
	_ = x[opChangeUseFlag-298]
	_ = x[opClientPerformanceStats-299]
	_ = x[opExtendedHardwareStats-300]
	_ = x[opClientLowMemoryWarning-301]
	_ = x[opTerritoryClaimStart-302]
	_ = x[opTerritoryClaimCancel-303]
	_ = x[opRequestAppStoreProducts-304]
	_ = x[opVerifyProductPurchase-305]
	_ = x[opQueryGuildPlayerStats-306]
	_ = x[opQueryAllianceGuildStats-307]
	_ = x[opTrackAchievements-308]
	_ = x[opSetAchievementsAutoLearn-309]
	_ = x[opDepositItemToGuildCurrency-310]
	_ = x[opWithdrawalItemFromGuildCurrency-311]
	_ = x[opAuctionSellSpecificItemRequest-312]
	_ = x[opFishingStart-313]
	_ = x[opFishingCasting-314]
	_ = x[opFishingCast-315]
	_ = x[opFishingCatch-316]
	_ = x[opFishingPull-317]
	_ = x[opFishingGiveLine-318]
	_ = x[opFishingFinish-319]
	_ = x[opFishingCancel-320]
	_ = x[opCreateGuildAccessTag-321]
	_ = x[opDeleteGuildAccessTag-322]
	_ = x[opRenameGuildAccessTag-323]
	_ = x[opFlagGuildAccessTagGuildPermission-324]
	_ = x[opAssignGuildAccessTag-325]
	_ = x[opRemoveGuildAccessTagFromPlayer-326]
	_ = x[opModifyGuildAccessTagEditors-327]
	_ = x[opRequestPublicAccessTags-328]
	_ = x[opChangeAccessTagPublicFlag-329]
	_ = x[opUpdateGuildAccessTag-330]
	_ = x[opSteamStartMicrotransaction-331]
	_ = x[opSteamFinishMicrotransaction-332]
	_ = x[opSteamIdHasActiveAccount-333]
	_ = x[opCheckEmailAccountState-334]
	_ = x[opLinkAccountToSteamId-335]
	_ = x[opBuyGvgSeasonBooster-336]
	_ = x[opChangeFlaggingPrepare-337]
	_ = x[opOverCharge-338]
	_ = x[opOverChargeEnd-339]
	_ = x[opRequestTrusted-340]
	_ = x[opChangeGuildLogo-341]
	_ = x[opPartyFinderRegisterForUpdates-342]
	_ = x[opPartyFinderUnregisterForUpdates-343]
	_ = x[opPartyFinderEnlistNewPartySearch-344]
	_ = x[opPartyFinderDeletePartySearch-345]
	_ = x[opPartyFinderChangePartySearch-346]
	_ = x[opPartyFinderChangeRole-347]
	_ = x[opPartyFinderApplyForGroup-348]
	_ = x[opPartyFinderAcceptOrDeclineApplyForGroup-349]
	_ = x[opPartyFinderGetEquipmentSnapshot-350]
	_ = x[opPartyFinderRegisterApplicants-351]
	_ = x[opPartyFinderUnregisterApplicants-352]
	_ = x[opPartyFinderFulltextSearch-353]
	_ = x[opPartyFinderRequestEquipmentSnapshot-354]
	_ = x[opGetPersonalSeasonTrackerData-355]
	_ = x[opUseConsumableFromInventory-356]
	_ = x[opClaimPersonalSeasonReward-357]
	_ = x[opEasyAntiCheatMessageToServer-358]
	_ = x[opSetNextTutorialState-359]
	_ = x[opAddPlayerToMuteList-360]
	_ = x[opRemovePlayerFromMuteList-361]
	_ = x[opProductShopUserEvent-362]
	_ = x[opGetVanityUnlocks-363]
	_ = x[opBuyVanityUnlocks-364]
	_ = x[opGetMountSkins-365]
	_ = x[opSetMountSkin-366]
	_ = x[opSetWardrobe-367]
	_ = x[opChangeCustomization-368]
	_ = x[opSetFavoriteIsland-369]
	_ = x[opGetGuildChallengePoints-370]
	_ = x[opTravelToHideout-371]
	_ = x[opSmartQueueJoin-372]
	_ = x[opSmartQueueLeave-373]
	_ = x[opSmartQueueSelectSpawnCluster-374]
	_ = x[opUpgradeHideout-375]
	_ = x[opInitHideoutAttackStart-376]
	_ = x[opInitHideoutAttackCancel-377]
	_ = x[opHideoutFillNutrition-378]
	_ = x[opHideoutGetInfo-379]
	_ = x[opHideoutGetOwnerInfo-380]
	_ = x[opHideoutSetTribute-381]
	_ = x[opOpenWorldAttackScheduleStart-382]
	_ = x[opOpenWorldAttackScheduleCancel-383]
	_ = x[opOpenWorldAttackConquerStart-384]
	_ = x[opOpenWorldAttackConquerCancel-385]
	_ = x[opGetOpenWorldAttackDetails-386]
	_ = x[opGetNextOpenWorldAttackScheduleTime-387]
	_ = x[opRecoverVaultFromHideout-388]
	_ = x[opGetGuildEnergyDrainInfo-389]
	_ = x[opChannelingUpdate-390]
}

const _OperationType_name = "opUnusedopPingopJoinopCreateAccountopLoginopSendCrashLogopSendTraceRouteopSendVfxStatsopSendGamePingInfoopCreateCharacteropDeleteCharacteropSelectCharacteropRedeemKeycodeopGetGameServerByClusteropGetActiveSubscriptionopGetShopPurchaseUrlopGetBuyTrialDetailsopGetReferralSeasonDetailsopGetReferralLinkopGetAvailableTrialKeysopGetShopTilesForCategoryopMoveopAttackStartopCastStartopCastCancelopTerminateToggleSpellopChannelingCancelopAttackBuildingStartopInventoryDestroyItemopInventoryMoveItemopInventoryRecoverItemopInventoryRecoverAllItemsopInventorySplitStackopInventorySplitStackIntoopGetClusterDataopChangeClusteropConsoleCommandopChatMessageopReportClientErroropRegisterToObjectopUnRegisterFromObjectopCraftBuildingChangeSettingsopCraftBuildingTakeMoneyopRepairBuildingChangeSettingsopRepairBuildingTakeMoneyopActionBuildingChangeSettingsopHarvestStartopHarvestCancelopTakeSilveropActionOnBuildingStartopActionOnBuildingCancelopItemRerollQualityStartopItemRerollQualityCancelopInstallResourceStartopInstallResourceCancelopInstallSilveropBuildingFillNutritionopBuildingChangeRenovationStateopBuildingBuySkinopBuildingClaimopBuildingGiveupopBuildingNutritionSilverStorageDepositopBuildingNutritionSilverStorageWithdrawopBuildingNutritionSilverRewardSetopConstructionSiteCreateopPlaceableObjectPlaceopPlaceableObjectPlaceCancelopPlaceableObjectPickupopFurnitureObjectUseopFarmableHarvestopFarmableFinishGrownItemopFarmableDestroyopFarmableGetProductopFarmableFillopTearDownConstructionSiteopCastleGateUseopAuctionCreateOfferopAuctionCreateRequestopAuctionGetOffersopAuctionGetRequestsopAuctionBuyOfferopAuctionAbortAuctionopAuctionModifyAuctionopAuctionAbortOfferopAuctionAbortRequestopAuctionSellRequestopAuctionGetFinishedAuctionsopAuctionGetFinishedAuctionsCountopAuctionFetchAuctionopAuctionGetMyOpenOffersopAuctionGetMyOpenRequestsopAuctionGetMyOpenAuctionsopAuctionGetItemAverageStatsopAuctionGetItemAverageValueopContainerOpenopContainerCloseopContainerManageSubContaineropRespawnopSuicideopJoinGuildopLeaveGuildopCreateGuildopInviteToGuildopDeclineGuildInvitationopKickFromGuildopDuellingChallengePlayeropDuellingAcceptChallengeopDuellingDenyChallengeopChangeClusterTaxopClaimTerritoryopGiveUpTerritoryopChangeTerritoryAccessRightsopGetMonolithInfoopGetClaimInfoopGetAttackInfoopGetTerritorySeasonPointsopGetAttackScheduleopScheduleAttackopGetMatchesopGetMatchDetailsopJoinMatchopLeaveMatchopChangeChatSettingsopLogoutStartopLogoutCancelopClaimOrbStartopClaimOrbCancelopMatchLootChestOpeningStartopMatchLootChestOpeningCancelopDepositToGuildAccountopWithdrawalFromAccountopChangeGuildPayUpkeepFlagopChangeGuildTaxopGetMyTerritoriesopMorganaCommandopGetServerInfoopInviteMercenaryToMatchopSubscribeToClusteropAnswerMercenaryInvitationopGetCharacterEquipmentopGetCharacterSteamAchievementsopGetCharacterStatsopGetKillHistoryDetailsopLearnMasteryLevelopReSpecAchievementopChangeAvataropGetRankingsopGetRankopGetGvgSeasonRankingsopGetGvgSeasonRankopGetGvgSeasonHistoryRankingsopGetGvgSeasonGuildMemberHistoryopKickFromGvGMatchopGetChestLogsopGetAccessRightLogsopGetGuildAccountLogsopGetGuildAccountLogsLargeAmountopInviteToPlayerTradeopPlayerTradeCancelopPlayerTradeInvitationAcceptopPlayerTradeAddItemopPlayerTradeRemoveItemopPlayerTradeAcceptTradeopPlayerTradeSetSilverOrGoldopSendMiniMapPingopStuckopBuyRealEstateopClaimRealEstateopGiveUpRealEstateopChangeRealEstateOutlineopGetMailInfosopGetMailCountopReadMailopSendNewMailopDeleteMailopMarkMailUnreadopClaimAttachmentFromMailopUpdateLfgInfoopGetLfgInfosopGetMyGuildLfgInfoopGetLfgDescriptionTextopLfgApplyToGuildopAnswerLfgGuildApplicationopRegisterChatPeeropSendChatMessageopJoinChatChannelopLeaveChatChannelopSendWhisperMessageopSayopPlayEmoteopStopEmoteopGetClusterMapInfoopAccessRightsChangeSettingsopMountopMountCancelopBuyJourneyopSetSaleStatusForEstateopResolveGuildOrPlayerNameopGetRespawnInfosopMakeHomeopLeaveHomeopResurrectionReplyopAllianceCreateopAllianceDisbandopAllianceGetMemberInfosopAllianceInviteopAllianceAnswerInvitationopAllianceCancelInvitationopAllianceKickGuildopAllianceLeaveopAllianceChangeGoldPaymentFlagopAllianceGetDetailInfoopGetIslandInfosopAbandonMyIslandopBuyMyIslandopBuyGuildIslandopAbandonGuildIslandopUpgradeMyIslandopUpgradeGuildIslandopMoveMyIslandopMoveGuildIslandopTerritoryFillNutritionopTeleportBackopPartyInvitePlayeropPartyAnswerInvitationopPartyLeaveopPartyKickPlayeropPartyMakeLeaderopPartyChangeLootSettingopPartyMarkObjectopPartySetRoleopGetGuildMOTDopSetGuildMOTDopExitEnterStartopExitEnterCancelopQuestGiverRequestopGoldMarketGetBuyOfferopGoldMarketGetBuyOfferFromSilveropGoldMarketGetSellOfferopGoldMarketGetSellOfferFromSilveropGoldMarketBuyGoldopGoldMarketSellGoldopGoldMarketCreateSellOrderopGoldMarketCreateBuyOrderopGoldMarketGetInfosopGoldMarketCancelOrderopGoldMarketGetAverageInfoopSiegeCampClaimStartopSiegeCampClaimCancelopTreasureChestUsingStartopTreasureChestUsingCancelopUseLootChestopUseShrineopLaborerStartJobopLaborerTakeJobLootopLaborerDismissopLaborerMoveopLaborerBuyItemopLaborerUpgradeopBuyPremiumopBuyTrialopRealEstateGetAuctionDataopRealEstateBidOnAuctionopGetSiegeCampCooldownopFriendInviteopFriendAnswerInvitationopFriendCancelnvitationopFriendRemoveopInventoryStackopInventorySortopEquipmentItemChangeSpellopExpeditionRegisteropExpeditionRegisterCancelopJoinExpeditionopDeclineExpeditionInvitationopVoteStartopVoteDoVoteopRatingDoRateopEnteringExpeditionStartopEnteringExpeditionCancelopActivateExpeditionCheckPointopArenaRegisteropArenaRegisterCancelopArenaLeaveopJoinArenaMatchopDeclineArenaInvitationopEnteringArenaStartopEnteringArenaCancelopArenaCustomMatchopArenaCustomMatchCreateopUpdateCharacterStatementopBoostFarmableopGetStrikeHistoryopUseFunctionopUsePortalEntranceopResetPortalBindingopQueryPortalBindingopClaimPaymentTransactionopChangeUseFlagopClientPerformanceStatsopExtendedHardwareStatsopClientLowMemoryWarningopTerritoryClaimStartopTerritoryClaimCancelopRequestAppStoreProductsopVerifyProductPurchaseopQueryGuildPlayerStatsopQueryAllianceGuildStatsopTrackAchievementsopSetAchievementsAutoLearnopDepositItemToGuildCurrencyopWithdrawalItemFromGuildCurrencyopAuctionSellSpecificItemRequestopFishingStartopFishingCastingopFishingCastopFishingCatchopFishingPullopFishingGiveLineopFishingFinishopFishingCancelopCreateGuildAccessTagopDeleteGuildAccessTagopRenameGuildAccessTagopFlagGuildAccessTagGuildPermissionopAssignGuildAccessTagopRemoveGuildAccessTagFromPlayeropModifyGuildAccessTagEditorsopRequestPublicAccessTagsopChangeAccessTagPublicFlagopUpdateGuildAccessTagopSteamStartMicrotransactionopSteamFinishMicrotransactionopSteamIdHasActiveAccountopCheckEmailAccountStateopLinkAccountToSteamIdopBuyGvgSeasonBoosteropChangeFlaggingPrepareopOverChargeopOverChargeEndopRequestTrustedopChangeGuildLogoopPartyFinderRegisterForUpdatesopPartyFinderUnregisterForUpdatesopPartyFinderEnlistNewPartySearchopPartyFinderDeletePartySearchopPartyFinderChangePartySearchopPartyFinderChangeRoleopPartyFinderApplyForGroupopPartyFinderAcceptOrDeclineApplyForGroupopPartyFinderGetEquipmentSnapshotopPartyFinderRegisterApplicantsopPartyFinderUnregisterApplicantsopPartyFinderFulltextSearchopPartyFinderRequestEquipmentSnapshotopGetPersonalSeasonTrackerDataopUseConsumableFromInventoryopClaimPersonalSeasonRewardopEasyAntiCheatMessageToServeropSetNextTutorialStateopAddPlayerToMuteListopRemovePlayerFromMuteListopProductShopUserEventopGetVanityUnlocksopBuyVanityUnlocksopGetMountSkinsopSetMountSkinopSetWardrobeopChangeCustomizationopSetFavoriteIslandopGetGuildChallengePointsopTravelToHideoutopSmartQueueJoinopSmartQueueLeaveopSmartQueueSelectSpawnClusteropUpgradeHideoutopInitHideoutAttackStartopInitHideoutAttackCancelopHideoutFillNutritionopHideoutGetInfoopHideoutGetOwnerInfoopHideoutSetTributeopOpenWorldAttackScheduleStartopOpenWorldAttackScheduleCancelopOpenWorldAttackConquerStartopOpenWorldAttackConquerCancelopGetOpenWorldAttackDetailsopGetNextOpenWorldAttackScheduleTimeopRecoverVaultFromHideoutopGetGuildEnergyDrainInfoopChannelingUpdate"

var _OperationType_index = [...]uint16{0, 8, 14, 20, 35, 42, 56, 72, 86, 104, 121, 138, 155, 170, 194, 217, 237, 257, 283, 300, 323, 348, 354, 367, 378, 390, 412, 430, 451, 473, 492, 514, 540, 561, 586, 602, 617, 633, 646, 665, 683, 705, 734, 758, 788, 813, 843, 857, 872, 884, 907, 931, 955, 980, 1002, 1025, 1040, 1063, 1094, 1111, 1126, 1142, 1181, 1221, 1255, 1279, 1301, 1329, 1352, 1372, 1389, 1414, 1431, 1451, 1465, 1491, 1506, 1526, 1548, 1566, 1586, 1603, 1624, 1646, 1665, 1686, 1706, 1734, 1767, 1788, 1812, 1838, 1864, 1892, 1920, 1935, 1951, 1980, 1989, 1998, 2009, 2021, 2034, 2049, 2073, 2088, 2113, 2138, 2161, 2179, 2195, 2212, 2241, 2258, 2272, 2287, 2313, 2332, 2348, 2360, 2377, 2388, 2400, 2420, 2433, 2447, 2462, 2478, 2506, 2535, 2558, 2581, 2607, 2623, 2641, 2657, 2672, 2696, 2716, 2743, 2766, 2797, 2816, 2839, 2858, 2877, 2891, 2904, 2913, 2935, 2953, 2982, 3014, 3032, 3046, 3066, 3087, 3119, 3140, 3159, 3188, 3208, 3231, 3255, 3283, 3300, 3307, 3322, 3339, 3357, 3382, 3396, 3410, 3420, 3433, 3445, 3461, 3486, 3501, 3514, 3533, 3556, 3573, 3600, 3618, 3635, 3652, 3670, 3690, 3695, 3706, 3717, 3736, 3764, 3771, 3784, 3796, 3820, 3846, 3863, 3873, 3884, 3903, 3919, 3936, 3960, 3976, 4002, 4028, 4047, 4062, 4093, 4116, 4132, 4149, 4162, 4178, 4198, 4215, 4235, 4249, 4266, 4290, 4304, 4323, 4346, 4358, 4375, 4392, 4416, 4433, 4447, 4461, 4475, 4491, 4508, 4527, 4550, 4583, 4607, 4641, 4660, 4680, 4707, 4733, 4753, 4776, 4802, 4823, 4845, 4870, 4896, 4910, 4921, 4938, 4958, 4974, 4987, 5003, 5019, 5031, 5041, 5067, 5091, 5113, 5127, 5151, 5174, 5188, 5204, 5219, 5245, 5265, 5291, 5307, 5336, 5347, 5359, 5373, 5398, 5424, 5454, 5469, 5490, 5502, 5518, 5542, 5562, 5583, 5601, 5625, 5651, 5666, 5684, 5697, 5716, 5736, 5756, 5781, 5796, 5820, 5843, 5867, 5888, 5910, 5935, 5958, 5981, 6006, 6025, 6051, 6079, 6112, 6144, 6158, 6174, 6187, 6201, 6214, 6231, 6246, 6261, 6283, 6305, 6327, 6362, 6384, 6416, 6445, 6470, 6497, 6519, 6547, 6576, 6601, 6625, 6647, 6668, 6691, 6703, 6718, 6734, 6751, 6782, 6815, 6848, 6878, 6908, 6931, 6957, 6998, 7031, 7062, 7095, 7122, 7159, 7189, 7217, 7244, 7274, 7296, 7317, 7343, 7365, 7383, 7401, 7416, 7430, 7443, 7464, 7483, 7508, 7525, 7541, 7558, 7588, 7604, 7628, 7653, 7675, 7691, 7712, 7731, 7761, 7792, 7821, 7851, 7878, 7914, 7939, 7964, 7982}

func (i OperationType) String() string {
	if i >= OperationType(len(_OperationType_index)-1) {
		return "OperationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperationType_name[_OperationType_index[i]:_OperationType_index[i+1]]
}
