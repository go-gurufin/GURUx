package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	"github.com/gurufin2021/GURUx/app"
	ibctesting "github.com/gurufin2021/GURUx/ibc/testing"
	"github.com/gurufin2021/GURUx/testutil"
	teststypes "github.com/gurufin2021/GURUx/types/tests"
	claimstypes "github.com/gurufin2021/GURUx/x/claims/types"
	"github.com/gurufin2021/GURUx/x/recovery/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Recovery: Performing an IBC Transfer", Ordered, func() {
	coinGurux := sdk.NewCoin("agurux", sdk.NewInt(10000))
	coinOsmo := sdk.NewCoin("uosmo", sdk.NewInt(10))
	coinAtom := sdk.NewCoin("uatom", sdk.NewInt(10))

	var (
		sender, receiver       string
		senderAcc, receiverAcc sdk.AccAddress
		timeout                uint64
		claim                  claimstypes.ClaimsRecord
	)

	BeforeEach(func() {
		s.SetupTest()
	})

	Describe("from a non-authorized chain", func() {
		BeforeEach(func() {
			params := claimstypes.DefaultParams()
			params.AuthorizedChannels = []string{}
			err := s.GuruxChain.App.(*app.Gurux).ClaimsKeeper.SetParams(s.GuruxChain.GetContext(), params)
			s.Require().NoError(err)

			sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
			receiver = s.GuruxChain.SenderAccount.GetAddress().String()
			senderAcc = sdk.MustAccAddressFromBech32(sender)
			receiverAcc = sdk.MustAccAddressFromBech32(receiver)
		})
		It("should transfer and not recover tokens", func() {
			s.SendAndReceiveMessage(s.pathOsmosisGurux, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 1)

			nativeGurux := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, "agurux")
			Expect(nativeGurux).To(Equal(coinGurux))
			ibcOsmo := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
			Expect(ibcOsmo).To(Equal(sdk.NewCoin(teststypes.UosmoIbcdenom, coinOsmo.Amount)))
		})
	})

	Describe("from an authorized, non-EVM chain (e.g. Osmosis)", func() {
		Describe("to a different account on Gurux (sender != recipient)", func() {
			BeforeEach(func() {
				sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
				receiver = s.GuruxChain.SenderAccount.GetAddress().String()
				senderAcc = sdk.MustAccAddressFromBech32(sender)
				receiverAcc = sdk.MustAccAddressFromBech32(receiver)
			})

			It("should transfer and not recover tokens", func() {
				s.SendAndReceiveMessage(s.pathOsmosisGurux, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 1)

				nativeGurux := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, "agurux")
				Expect(nativeGurux).To(Equal(coinGurux))
				ibcOsmo := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
				Expect(ibcOsmo).To(Equal(sdk.NewCoin(teststypes.UosmoIbcdenom, coinOsmo.Amount)))
			})
		})

		Describe("to the sender's own eth_secp256k1 account on Gurux (sender == recipient)", func() {
			BeforeEach(func() {
				sender = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
				receiver = s.IBCOsmosisChain.SenderAccount.GetAddress().String()
				senderAcc = sdk.MustAccAddressFromBech32(sender)
				receiverAcc = sdk.MustAccAddressFromBech32(receiver)
			})

			Context("with disabled recovery parameter", func() {
				BeforeEach(func() {
					params := types.DefaultParams()
					params.EnableRecovery = false
					s.GuruxChain.App.(*app.Gurux).RecoveryKeeper.SetParams(s.GuruxChain.GetContext(), params) //nolint:errcheck
				})

				It("should not transfer or recover tokens", func() {
					s.SendAndReceiveMessage(s.pathOsmosisGurux, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
					nativeGurux := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, "agurux")
					Expect(nativeGurux).To(Equal(coinGurux))
					ibcOsmo := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
					Expect(ibcOsmo).To(Equal(sdk.NewCoin(teststypes.UosmoIbcdenom, coinOsmo.Amount)))
				})
			})

			Context("with a sender's claims record", func() {
				Context("without completed actions", func() {
					BeforeEach(func() {
						amt := sdk.NewInt(int64(100))
						coins := sdk.NewCoins(sdk.NewCoin("agurux", amt))
						claim = claimstypes.NewClaimsRecord(amt)
						s.GuruxChain.App.(*app.Gurux).ClaimsKeeper.SetClaimsRecord(s.GuruxChain.GetContext(), senderAcc, claim)

						// update the escrowed account balance to maintain the invariant
						err := testutil.FundModuleAccount(s.GuruxChain.GetContext(), s.GuruxChain.App.(*app.Gurux).BankKeeper, claimstypes.ModuleName, coins)
						s.Require().NoError(err)
					})

					It("should not transfer or recover tokens", func() {
						// Prevent further funds from getting stuck
						s.SendAndReceiveMessage(s.pathOsmosisGurux, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
						nativeGurux := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, "agurux")
						Expect(nativeGurux).To(Equal(coinGurux))
						ibcOsmo := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
					})
				})

				Context("with completed actions", func() {
					// Already has stuck funds
					BeforeEach(func() {
						amt := sdk.NewInt(int64(100))
						coins := sdk.NewCoins(sdk.NewCoin("agurux", sdk.NewInt(int64(75))))
						claim = claimstypes.NewClaimsRecord(amt)
						claim.MarkClaimed(claimstypes.ActionIBCTransfer)
						s.GuruxChain.App.(*app.Gurux).ClaimsKeeper.SetClaimsRecord(s.GuruxChain.GetContext(), senderAcc, claim)

						// update the escrowed account balance to maintain the invariant
						err := testutil.FundModuleAccount(s.GuruxChain.GetContext(), s.GuruxChain.App.(*app.Gurux).BankKeeper, claimstypes.ModuleName, coins)
						s.Require().NoError(err)

						// agurux & ibc tokens that originated from the sender's chain
						s.SendAndReceiveMessage(s.pathOsmosisGurux, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
						timeout = uint64(s.GuruxChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
					})

					It("should transfer tokens to the recipient and perform recovery", func() {
						// Escrow before relaying packets
						balanceEscrow := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), transfertypes.GetEscrowAddress("transfer", "channel-0"), "agurux")
						Expect(balanceEscrow).To(Equal(coinGurux))
						ibcOsmo := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())

						// Relay both packets that were sent in the ibc_callback
						err := s.pathOsmosisGurux.RelayPacket(CreatePacket("10000", "agurux", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisGurux.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						// Check that the agurux were recovered
						nativeGurux := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, "agurux")
						Expect(nativeGurux.IsZero()).To(BeTrue())
						ibcGurux := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AguruxIbcdenom)
						Expect(ibcGurux).To(Equal(sdk.NewCoin(teststypes.AguruxIbcdenom, coinGurux.Amount)))

						// Check that the uosmo were recovered
						ibcOsmo = s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))
					})

					It("should not claim/migrate/merge claims records", func() {
						// Relay both packets that were sent in the ibc_callback
						err := s.pathOsmosisGurux.RelayPacket(CreatePacket("10000", "agurux", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisGurux.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						claimAfter, _ := s.GuruxChain.App.(*app.Gurux).ClaimsKeeper.GetClaimsRecord(s.GuruxChain.GetContext(), senderAcc)
						Expect(claim).To(Equal(claimAfter))
					})
				})
			})

			Context("without a sender's claims record", func() {
				When("recipient has no ibc vouchers that originated from other chains", func() {
					It("should transfer and recover tokens", func() {
						// agurux & ibc tokens that originated from the sender's chain
						s.SendAndReceiveMessage(s.pathOsmosisGurux, s.IBCOsmosisChain, coinOsmo.Denom, coinOsmo.Amount.Int64(), sender, receiver, 1)
						timeout = uint64(s.GuruxChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())

						// Escrow before relaying packets
						balanceEscrow := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), transfertypes.GetEscrowAddress("transfer", "channel-0"), "agurux")
						Expect(balanceEscrow).To(Equal(coinGurux))
						ibcOsmo := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())

						// Relay both packets that were sent in the ibc_callback
						err := s.pathOsmosisGurux.RelayPacket(CreatePacket("10000", "agurux", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisGurux.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						// Check that the agurux were recovered
						nativeGurux := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, "agurux")
						Expect(nativeGurux.IsZero()).To(BeTrue())
						ibcGurux := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AguruxIbcdenom)
						Expect(ibcGurux).To(Equal(sdk.NewCoin(teststypes.AguruxIbcdenom, coinGurux.Amount)))

						// Check that the uosmo were recovered
						ibcOsmo = s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))
					})
				})

				// Do not recover uatom sent from Cosmos when performing recovery through IBC transfer from Osmosis
				When("recipient has additional ibc vouchers that originated from other chains", func() {
					BeforeEach(func() {
						params := types.DefaultParams()
						params.EnableRecovery = false
						s.GuruxChain.App.(*app.Gurux).RecoveryKeeper.SetParams(s.GuruxChain.GetContext(), params) //nolint:errcheck

						// Send uatom from Cosmos to Gurux
						s.SendAndReceiveMessage(s.pathCosmosGurux, s.IBCCosmosChain, coinAtom.Denom, coinAtom.Amount.Int64(), s.IBCCosmosChain.SenderAccount.GetAddress().String(), receiver, 1)

						params.EnableRecovery = true
						s.GuruxChain.App.(*app.Gurux).RecoveryKeeper.SetParams(s.GuruxChain.GetContext(), params) //nolint:errcheck
					})
					It("should not recover tokens that originated from other chains", func() {
						// Send uosmo from Osmosis to Gurux
						s.SendAndReceiveMessage(s.pathOsmosisGurux, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 1)

						// Relay both packets that were sent in the ibc_callback
						timeout := uint64(s.GuruxChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
						err := s.pathOsmosisGurux.RelayPacket(CreatePacket("10000", "agurux", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisGurux.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)

						// Agurux was recovered from user address
						nativeGurux := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, "agurux")
						Expect(nativeGurux.IsZero()).To(BeTrue())
						ibcGurux := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AguruxIbcdenom)
						Expect(ibcGurux).To(Equal(sdk.NewCoin(teststypes.AguruxIbcdenom, coinGurux.Amount)))

						// Check that the uosmo were retrieved
						ibcOsmo := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))

						// Check that the atoms were not retrieved
						ibcAtom := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, teststypes.UatomIbcdenom)
						Expect(ibcAtom).To(Equal(sdk.NewCoin(teststypes.UatomIbcdenom, coinAtom.Amount)))

						// Repeat transaction from Osmosis to Gurux
						s.SendAndReceiveMessage(s.pathOsmosisGurux, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 2)

						timeout = uint64(s.GuruxChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
						err = s.pathOsmosisGurux.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 3, timeout))
						s.Require().NoError(err)

						// No further tokens recovered
						nativeGurux = s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, "agurux")
						Expect(nativeGurux.IsZero()).To(BeTrue())
						ibcGurux = s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AguruxIbcdenom)
						Expect(ibcGurux).To(Equal(sdk.NewCoin(teststypes.AguruxIbcdenom, coinGurux.Amount)))

						ibcOsmo = s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo = s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))

						ibcAtom = s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, teststypes.UatomIbcdenom)
						Expect(ibcAtom).To(Equal(sdk.NewCoin(teststypes.UatomIbcdenom, coinAtom.Amount)))
					})
				})

				// Recover ibc/uatom that was sent from Osmosis back to Osmosis
				When("recipient has additional non-native ibc vouchers that originated from senders chains", func() {
					BeforeEach(func() {
						params := types.DefaultParams()
						params.EnableRecovery = false
						err := s.GuruxChain.App.(*app.Gurux).RecoveryKeeper.SetParams(s.GuruxChain.GetContext(), params)
						s.Require().NoError(err)
						s.SendAndReceiveMessage(s.pathOsmosisCosmos, s.IBCCosmosChain, coinAtom.Denom, coinAtom.Amount.Int64(), s.IBCCosmosChain.SenderAccount.GetAddress().String(), receiver, 1)

						// Send IBC transaction of 10 ibc/uatom
						transferMsg := transfertypes.NewMsgTransfer(s.pathOsmosisGurux.EndpointA.ChannelConfig.PortID, s.pathOsmosisGurux.EndpointA.ChannelID, sdk.NewCoin(teststypes.UatomIbcdenom, sdk.NewInt(10)), sender, receiver, timeoutHeight, 0, "")
						_, err = ibctesting.SendMsgs(s.IBCOsmosisChain, ibctesting.DefaultFeeAmt, transferMsg)
						s.Require().NoError(err) // message committed
						transfer := transfertypes.NewFungibleTokenPacketData("transfer/channel-1/uatom", "10", sender, receiver, "")
						packet := channeltypes.NewPacket(transfer.GetBytes(), 1, s.pathOsmosisGurux.EndpointA.ChannelConfig.PortID, s.pathOsmosisGurux.EndpointA.ChannelID, s.pathOsmosisGurux.EndpointB.ChannelConfig.PortID, s.pathOsmosisGurux.EndpointB.ChannelID, timeoutHeight, 0)
						// Receive message on the gurux side, and send ack
						err = s.pathOsmosisGurux.RelayPacket(packet)
						s.Require().NoError(err)

						// Check that the ibc/uatom are available
						osmoIBCAtom := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UatomOsmoIbcdenom)
						s.Require().Equal(osmoIBCAtom.Amount, coinAtom.Amount)

						params.EnableRecovery = true
						s.GuruxChain.App.(*app.Gurux).RecoveryKeeper.SetParams(s.GuruxChain.GetContext(), params) //nolint:errcheck
					})
					It("should not recover tokens that originated from other chains", func() {
						s.SendAndReceiveMessage(s.pathOsmosisGurux, s.IBCOsmosisChain, "uosmo", 10, sender, receiver, 2)

						// Relay packets that were sent in the ibc_callback
						timeout := uint64(s.GuruxChain.GetContext().BlockTime().Add(time.Hour * 4).Add(time.Second * -20).UnixNano())
						err := s.pathOsmosisGurux.RelayPacket(CreatePacket("10000", "agurux", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 1, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisGurux.RelayPacket(CreatePacket("10", "transfer/channel-0/transfer/channel-1/uatom", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 2, timeout))
						s.Require().NoError(err)
						err = s.pathOsmosisGurux.RelayPacket(CreatePacket("10", "transfer/channel-0/uosmo", sender, receiver, "transfer", "channel-0", "transfer", "channel-0", 3, timeout))
						s.Require().NoError(err)

						// Agurux was recovered from user address
						nativeGurux := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), senderAcc, "agurux")
						Expect(nativeGurux.IsZero()).To(BeTrue())
						ibcGurux := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, teststypes.AguruxIbcdenom)
						Expect(ibcGurux).To(Equal(sdk.NewCoin(teststypes.AguruxIbcdenom, coinGurux.Amount)))

						// Check that the uosmo were recovered
						ibcOsmo := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UosmoIbcdenom)
						Expect(ibcOsmo.IsZero()).To(BeTrue())
						nativeOsmo := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), receiverAcc, "uosmo")
						Expect(nativeOsmo).To(Equal(coinOsmo))

						// Check that the ibc/uatom were retrieved
						osmoIBCAtom := s.GuruxChain.App.(*app.Gurux).BankKeeper.GetBalance(s.GuruxChain.GetContext(), receiverAcc, teststypes.UatomOsmoIbcdenom)
						Expect(osmoIBCAtom.IsZero()).To(BeTrue())
						ibcAtom := s.IBCOsmosisChain.GetSimApp().BankKeeper.GetBalance(s.IBCOsmosisChain.GetContext(), senderAcc, teststypes.UatomIbcdenom)
						Expect(ibcAtom).To(Equal(sdk.NewCoin(teststypes.UatomIbcdenom, sdk.NewInt(10))))
					})
				})
			})
		})
	})
})
