- 📁 account<br/>
  - 📃 keystore_geth.go<br/>
    - ⬜ Todo:<br/>
      *On line*              : 3 <br/>
      *Description*          : TODO: Make independent version for Nimbus<br/>
- 📁 api<br/>
  - 📃 backend_test.go<br/>
    - ⬜ Todo:<br/>
      *On line*              : 115 <br/>
      *Description*          : TODO(adam): add concurrent tests for ResetChainData()<br/>
    - ⬜ Todo:<br/>
      *On line*              : 378 <br/>
      *Description*          : TODO(adam): add concurrent tests for: SendTransaction<br/>
  - 📃 geth_backend.go<br/>
    - ⬜ Todo:<br/>
      *On line*              : 255 <br/>
      *Description*          : TODO: we should use a proper struct with optional values instead of duplicating the regular functions<br/>
       with small variants for keycard, this created too many bugs<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (b *GethStatusBackend) AppStateChange(state string)`<br/>
      *On line*              : 912 <br/>
      *Description*          : TODO: put node in low-power mode if the app is in background (or inactive)<br/>
       and normal mode if the app is in foreground.<br/>
  - 📃 nimbus_backend.go<br/>
    - ⬜ Todo:<br/>
      *On line*              : 191 <br/>
      *Description*          : TODO: we should use a proper struct with optional values instead of duplicating the regular functions<br/>
       with small variants for keycard, this created too many bugs<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (b *nimbusStatusBackend) AppStateChange(state string)`<br/>
      *On line*              : 780 <br/>
      *Description*          : TODO: put node in low-power mode if the app is in background (or inactive)<br/>
       and normal mode if the app is in foreground.<br/>
- 📁 bridge<br/>
  - 📃 bridge_test.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func TestBridgeWhisperToWaku(t *testing.T)`<br/>
      *On line*              : 84 <br/>
      *Description*          : TODO: replace with a bridge event which should not be sent by Waku<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func TestBridgeWakuToWhisper(t *testing.T)`<br/>
      *On line*              : 138 <br/>
      *Description*          : TODO: replace with a bridge event which should not be sent by Waku<br/>
- 📁 db<br/>
  - 📃 history.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (req *HistoryRequest) Histories() []TopicHistory`<br/>
      *On line*              : 144 <br/>
      *Description*          : TODO Lazy load from database on first access<br/>
  - 📃 history_store.go<br/>
    - ⬜ Todo:<br/>
      *On line*              : 76 <br/>
      *Description*          : TODO explain<br/>
- 📁 discovery<br/>
  - 📃 muxer_test.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func TestMuxerDiscovery(t *testing.T)`<br/>
      *On line*              : 235 <br/>
      *Description*          : TODO test period channel<br/>
- 📁 eth-node<br/>
  - 📁 bridge<br/>
    - 📁 nimbus<br/>
      - 📃 public_whisper_api.go<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (w *nimbusPublicWhisperAPIWrapper) NewMessageFilter(req types.Criteria) (string, error)`<br/>
          *On line*              : 119 <br/>
          *Description*          : TODO<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (w *nimbusPublicWhisperAPIWrapper) Post(ctx context.Context, req types.NewMessage) ([]byte, error)`<br/>
          *On line*              : 195 <br/>
          *Description*          : TODO: return envelope hash once nimbus_post is improved to return it<br/>
      - 📃 whisper.go<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (w *nimbusWhisperWrapper) SubscribeEnvelopeEvents(eventsProxy chan<- types.EnvelopeEvent) types.Subscription`<br/>
          *On line*              : 90 <br/>
          *Description*          : TODO: when mailserver support is implemented<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (w *nimbusWhisperWrapper) Subscribe(opts *types.SubscriptionOptions) (string, error)`<br/>
          *On line*              : 309 <br/>
          *Description*          : TODO: Check if this is done too late (race condition with onMessageHandler)<br/>
  - 📁 keystore<br/>
    - 📃 passphrase.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (m Message) CheckNonce() bool   `<br/>
        *On line*              : 273 <br/>
        *Description*          : TODO: can we do without this when unmarshalling dynamic JSON?<br/>
         why do integers in KDF params end up as float64 and not int after<br/>
         unmarshal?<br/>
- 📁 extkeys<br/>
  - 📃 hdkey.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (err decError) Error() string`<br/>
      *On line*              : 42 <br/>
      *Description*          : TODO make sure we're doing this ^^^^ !!!!!!<br/>
  - 📃 mnemonic.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (m *Mnemonic) MnemonicPhrase(strength EntropyStrength, language Language) (string, error)`<br/>
      *On line*              : 188 <br/>
      *Description*          : TODO simplify?<br/>
- 📁 mailserver<br/>
  - 📃 mailserver.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func newMailServer(cfg Config, adapter adapter, service service) (*mailServer, error)`<br/>
      *On line*              : 630 <br/>
      *Description*          : TODO: move out<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (s *mailServer) createIterator(req MessagesRequestPayload) (Iterator, error)`<br/>
      *On line*              : 1026 <br/>
      *Description*          : TODO(adam): this is invalid code. If the limit is 1000,<br/>
       it will only send 999 items and send a cursor.<br/>
- 📁 multiaccounts<br/>
  - 📁 accounts<br/>
    - 📃 database.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func TestPrepareJSONResponseErrorWithResult(t *testing.T)`<br/>
        *On line*              : 125 <br/>
        *Description*          : TODO remove photoPath from settings<br/>
- 📁 node<br/>
  - 📃 get_status_node.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (n *StatusNode) startDiscovery() error`<br/>
      *On line*              : 345 <br/>
      *Description*          : TODO(dshulyak) consider adding a flag to define this behaviour<br/>
  - 📃 geth_node.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func activateShhService(stack *node.Node, config *params.NodeConfig, db *leveldb.DB) (err error)`<br/>
      *On line*              : 357 <br/>
      *Description*          : TODO(dshulyak) add a config option to enable it by default, but disable if app is started from statusd<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func activateWakuService(stack *node.Node, config *params.NodeConfig, db *leveldb.DB) (err error)`<br/>
      *On line*              : 397 <br/>
      *Description*          : TODO(dshulyak) add a config option to enable it by default, but disable if app is started from statusd<br/>
  - 📃 nimbus_node.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (n *NimbusStatusNode) activateShhService(config *params.NodeConfig, db *leveldb.DB) (err error)`<br/>
      *On line*              : 217 <br/>
      *Description*          : TODO(dshulyak) add a config option to enable it by default, but disable if app is started from statusd<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (n *NimbusStatusNode) activateWakuService(config *params.NodeConfig, db *leveldb.DB) (err error)`<br/>
      *On line*              : 242 <br/>
      *Description*          : TODO(dshulyak) add a config option to enable it by default, but disable if app is started from statusd<br/>
       return n.Register(func(ctx *nimbussvc.ServiceContext) (nimbussvc.Service, error) {<br/>
       	var ethnode *nodebridge.NodeService<br/>
       	if err := ctx.Service(&ethnode); err != nil {<br/>
       		return nil, err<br/>
       	}<br/>
       	return shhext.New(ethnode.Node, ctx, "wakuext", shhext.EnvelopeSignalHandler{}, db, config.ShhextConfig), nil<br/>
       })<br/>
  - 📃 nimbus_status_node.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (n *NimbusStatusNode) startDiscovery() error`<br/>
      *On line*              : 332 <br/>
      *Description*          : TODO(dshulyak) consider adding a flag to define this behaviour<br/>
       	options.AllowStop = len(n.config.RegisterTopics) == 0<br/>
       	options.TrustedMailServers = parseNodesToNodeID(n.config.ClusterConfig.TrustedMailServers)<br/>
- 📁 params<br/>
  - 📃 config.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (c *NodeConfig) updatePeerLimits()`<br/>
      *On line*              : 723 <br/>
      *Description*          : TODO(dshulyak) register mailserver limits when we will change how they are handled.<br/>
- 📁 peers<br/>
  - 📃 peerpool.go<br/>
    - ⬜ Todo:<br/>
      *On line*              : 249 <br/>
      *Description*          : TODO(adam): split it into peers and discovery management loops. This should<br/>
       simplify the whole logic and allow to remove `timeout` field from `PeerPool`.<br/>
  - 📃 topicpool.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (t *TopicPool) AddPeerFromTable(server *p2p.Server) *discv5.Node`<br/>
      *On line*              : 401 <br/>
      *Description*          : TODO(adam): investigate if it's worth to keep the peer in the queue<br/>
       until the server confirms it is added and in the meanwhile only adjust its priority.<br/>
- 📁 protocol<br/>
  - 📁 communities<br/>
    - 📃 community.go<br/>
      - ⬜ Todo:<br/>
        *On line*              : 300 <br/>
        *Description*          : TODO: this should accept a request from a user to join and perform any validation<br/>
      - ⬜ Todo:<br/>
        *On line*              : 306 <br/>
        *Description*          : TODO: this should decline a request from a user to join<br/>
    - 📃 community_test.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (s *CommunitySuite) TestHandleRequestJoin()`<br/>
        *On line*              : 306 <br/>
        *Description*          : TODO -> Org is invitation only, chat is read-write for members<br/>
         INVITATION_ONLY - INVITATION_ONLY -> Error -> Org is invitation only, chat is invitation only<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (s *CommunitySuite) TestHandleRequestJoin()`<br/>
        *On line*              : 308 <br/>
        *Description*          : TODO -> Error -> Org is invitation only, member of the org need to request access for chat<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (s *CommunitySuite) TestHandleRequestJoin()`<br/>
        *On line*              : 309 <br/>
        *Description*          : TODO -> Error -> Org is on request, chat is read write for members<br/>
         ON_REQUEST - INVITATION_ONLY -> Error -> Org is on request, chat is invitation only for members<br/>
         ON_REQUEST - ON_REQUEST -> Fine -> Org is on request, chat is on request<br/>
    - 📃 manager.go<br/>
      - ⬜ Todo:<br/>
        *On line*              : 225 <br/>
        *Description*          : TODO: Finish implementing this<br/>
  - 📁 encryption<br/>
    - 📃 persistence_test.go<br/>
      - ⬜ Todo:<br/>
        *On line*              : 342 <br/>
        *Description*          : TODO: Add test for MarkBundleExpired<br/>
    - 📃 protocol.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (p *Protocol) ProcessPublicBundle(myIdentityKey *ecdsa.PrivateKey, bundle *Bundle) ([]*multidevice.Installation, error)`<br/>
        *On line*              : 266 <br/>
        *Description*          : TODO(adam): why do we add installations using identity obtained from GetIdentity()<br/>
         instead of the output of crypto.CompressPubkey()? I tried the second option<br/>
         and the unit tests TestTopic and TestMaxDevices fail.<br/>
    - 📁 publisher<br/>
      - 📃 publisher_test.go<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (*ProtocolMessage) ProtoMessage()   `<br/>
          *On line*              : 31 <br/>
          *Description*          : TODO(adam): provide more tests<br/>
    - 📁 sharedsecret<br/>
      - 📃 sharedsecret.go<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (*ProtocolMessage) ProtoMessage()   `<br/>
          *On line*              : 25 <br/>
          *Description*          : TODO: make compression of public keys a responsibility  of sqlitePersistence instead of SharedSecret.<br/>
  - 📃 message_handler.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (m *MessageHandler) matchChatEntity(chatEntity common.ChatEntity, chats map[string]*Chat, timesource common.TimeSource) (*Chat, error)`<br/>
      *On line*              : 742 <br/>
      *Description*          : TODO: this should be a three-word name used in the mobile client<br/>
  - 📃 message_validator_test.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (s *MessageValidatorSuite) TestValidatePlainTextMessage()`<br/>
      *On line*              : 287 <br/>
      *Description*          : TODO: FIX ME<br/>
  - 📃 messenger.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (m *Messenger) createChatIdentity(context chatContext) (*protobuf.ChatIdentity, error)`<br/>
      *On line*              : 631 <br/>
      *Description*          : TODO add ENS name handling to dedicate PR<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (m *Messenger) adaptIdentityImageToProtobuf(img *userimage.IdentityImage) *protobuf.IdentityImage`<br/>
      *On line*              : 678 <br/>
      *Description*          : TODO add ENS avatar handling to dedicated PR<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (m *Messenger) Init() error`<br/>
      *On line*              : 950 <br/>
      *Description*          : TODO: Get only active chats by the query.<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (m *Messenger) dispatchMessage(ctx context.Context, spec common.RawMessage) ([]byte, error)`<br/>
      *On line*              : 2314 <br/>
      *Description*          : TODO: add grant<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (m *Messenger) SendContactUpdates(ctx context.Context, ensName, profileImage string) error`<br/>
      *On line*              : 2545 <br/>
      *Description*          : TODO: This should not be sending paired messages, as we do it above<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (*ProtocolMessage) ProtoMessage()   `<br/>
      *On line*              : 2635 <br/>
      *Description*          : TODO remove use of photoPath in contacts<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (m *Messenger) encodeChatEntity(chat *Chat, message common.ChatEntity) ([]byte, error)`<br/>
      *On line*              : 4614 <br/>
      *Description*          : TODO: add grant<br/>
  - 📁 pushnotificationclient<br/>
    - 📃 client.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (c *Client) RemovePushNotificationServer(publicKey *ecdsa.PublicKey) error`<br/>
        *On line*              : 556 <br/>
        *Description*          : TODO: this needs implementing. It requires unregistering from the server and<br/>
         likely invalidate the device token of the user<br/>
  - 📁 pushnotificationserver<br/>
    - 📃 persistence.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (p *SQLitePersistence) GetPushNotificationRegistrationByPublicKeys(publicKeys [][]byte) ([]*PushNotificationIDAndRegistration, error)`<br/>
        *On line*              : 84 <br/>
        *Description*          : TODO: check for a max number of keys<br/>
  - 📁 transport<br/>
    - 📃 filter.go<br/>
      - ⬜ Todo:<br/>
        *On line*              : 5 <br/>
        *Description*          : TODO: revise fields encoding/decoding. Some are encoded using hexutil and some using encoding/hex.<br/>
    - 📁 waku<br/>
      - 📃 waku_service.go<br/>
        - ⬜ Todo:<br/>
          *On line*              : 76 <br/>
          *Description*          : TODO: leaving a chat should verify that for a given public key<br/>
                 there are no other chats. It may happen that we leave a private chat<br/>
                 but still have a public chat for a given public key.<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (a *Transport) SendPrivateOnDiscovery(ctx context.Context, newMessage *types.NewMessage, publicKey *ecdsa.PublicKey) ([]byte, error)`<br/>
          *On line*              : 295 <br/>
          *Description*          : TODO: change this anyway, it should be explicit<br/>
           and idempotent.<br/>
    - 📁 whisper<br/>
      - 📃 whisper_service.go<br/>
        - ⬜ Todo:<br/>
          *On line*              : 76 <br/>
          *Description*          : TODO: leaving a chat should verify that for a given public key<br/>
                 there are no other chats. It may happen that we leave a private chat<br/>
                 but still have a public chat for a given public key.<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (a *Transport) SendPrivateOnDiscovery(ctx context.Context, newMessage *types.NewMessage, publicKey *ecdsa.PublicKey) ([]byte, error)`<br/>
          *On line*              : 346 <br/>
          *Description*          : TODO: change this anyway, it should be explicit<br/>
           and idempotent.<br/>
  - 📁 v1<br/>
    - 📃 membership_update_message.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (g Group) validateChatID(chatID string) bool`<br/>
        *On line*              : 372 <br/>
        *Description*          : TODO: It does not verify that the prefix is a valid UUID.<br/>
               Improve it so that the prefix follows UUIDv4 spec.<br/>
- 📁 rpc<br/>
  - 📃 call_raw.go<br/>
    - ⬜ Todo:<br/>
      *On line*              : 62 <br/>
      *Description*          : TODO(divan): this function exists for compatibility and uses default<br/>
       go-ethereum's RPC client under the hood. It adds some unnecessary overhead<br/>
       by first marshalling JSON string into object to use with normal Call,<br/>
       which is then umarshalled back to the same JSON. The same goes with response.<br/>
       This is waste of CPU and memory and should be avoided if possible,<br/>
       either by changing exported API (provide only Call, not CallRaw) or<br/>
       refactoring go-ethereum's client to allow using raw JSON directly.<br/>
  - 📃 client.go<br/>
    - ⬜ Todo:<br/>
      *On line*              : 179 <br/>
      *Description*          : TODO(divan): use cancellation via context here?<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func setResultFromRPCResponse(result, response interface{}) (err error)`<br/>
      *On line*              : 219 <br/>
      *Description*          : TODO(divan): add additional checks for result underlying value, if needed:<br/>
       some example: https://golang.org/src/encoding/json/decode.go#L596<br/>
- 📁 services<br/>
  - 📁 ext<br/>
    - 📃 api.go<br/>
      - ⬜ Todo:<br/>
        *On line*              : 191 <br/>
        *Description*          : TODO: this is broken now as it requires dedup ID while a message hash should be used.<br/>
    - 📁 mailservers<br/>
      - 📃 connmanager.go<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (ps *ConnectionManager) Start()`<br/>
          *On line*              : 111 <br/>
          *Description*          : TODO treat failed requests the same way as expired<br/>
    - 📃 rpc.go<br/>
      - ⬜ Todo:<br/>
        *On line*              : 1 <br/>
        *Description*          : TODO: These types should be defined using protobuf, but protoc can only emit []byte instead of types.HexBytes,<br/>
         which causes issues when marshaling to JSON on the react side. Let's do that once the chat protocol is moved to the go repo.<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `type SendPublicMessageRPC struct`<br/>
        *On line*              : 15 <br/>
        *Description*          : TODO: remove<br/>
      - ⬜ Todo:<br/>
        *On line*              : 20 <br/>
        *Description*          : TODO: implement with accordance to https://github.com/status-im/status-go/protocol/issues/28.<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `type SendDirectMessageRPC struct`<br/>
        *On line*              : 29 <br/>
        *Description*          : TODO: remove<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `type SendDirectMessageRPC struct`<br/>
        *On line*              : 33 <br/>
        *Description*          : TODO: make sure to remove safely<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (m SendPublicMessageRPC) PublicKey() *ecdsa.PublicKey`<br/>
        *On line*              : 36 <br/>
        *Description*          : TODO: implement with accordance to https://github.com/status-im/status-go/protocol/issues/28.<br/>
    - 📃 signal.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (h PublisherSignalHandler) FilterAdded(filters []*signal.Filter)`<br/>
        *On line*              : 44 <br/>
        *Description*          : TODO(waku): change the name of the filter to generic one.<br/>
  - 📁 shhext<br/>
    - 📃 api_geth.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func NewPublicAPI(s *Service) *PublicAPI`<br/>
        *On line*              : 57 <br/>
        *Description*          : TODO: replace with an types.Envelope creator passed to the API struct<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (api *PublicAPI) RequestMessagesSync(conf ext.RetryConfig, r ext.MessagesRequest) (ext.MessagesResponse, error)`<br/>
        *On line*              : 161 <br/>
        *Description*          : FIXME this weird conversion is required because MessagesRequest expects seconds but defines time.Duration<br/>
    - 📃 api_nimbus.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (api *NimbusPublicAPI) RequestMessagesSync(conf RetryConfig, r MessagesRequest) (MessagesResponse, error)`<br/>
        *On line*              : 69 <br/>
        *Description*          : FIXME this weird conversion is required because MessagesRequest expects seconds but defines time.Duration<br/>
         		r.Timeout = time.Duration(int(r.Timeout.Seconds()))<br/>
         		requestID, err = api.RequestMessages(context.Background(), r)<br/>
         		if err != nil {<br/>
         			sub.Unsubscribe()<br/>
         			return resp, err<br/>
         		}<br/>
         		mailServerResp, err := waitForExpiredOrCompleted(types.BytesToHash(requestID), events, timeout)<br/>
         		sub.Unsubscribe()<br/>
         		if err == nil {<br/>
         			resp.Cursor = hex.EncodeToString(mailServerResp.Cursor)<br/>
         			resp.Error = mailServerResp.Error<br/>
         			return resp, nil<br/>
         		}<br/>
         		retries++<br/>
         		api.log.Error("[RequestMessagesSync] failed", "err", err, "retries", retries)<br/>
         	}<br/>
         	return resp, fmt.Errorf("failed to request messages after %d retries", retries)<br/>
         }<br/>
      - ⬜ Todo:<br/>
        *On line*              : 250 <br/>
        *Description*          : TODO: this is broken now as it requires dedup ID while a message hash should be used.<br/>
  - 📁 wakuext<br/>
    - 📃 api.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (api *PublicAPI) RequestMessagesSync(conf ext.RetryConfig, r ext.MessagesRequest) (ext.MessagesResponse, error)`<br/>
        *On line*              : 155 <br/>
        *Description*          : FIXME this weird conversion is required because MessagesRequest expects seconds but defines time.Duration<br/>
  - 📁 wallet<br/>
    - 📃 commands.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (c *controlCommand) verifyLastSynced(parent context.Context, last *DBHeader, head *types.Header) error`<br/>
        *On line*              : 623 <br/>
        *Description*          : TODO(dshulyak) make a standalone command that<br/>
         doesn't manage transfers and has an upper limit<br/>
    - 📃 database.go<br/>
      - ⬜ Todo:<br/>
        *On line*              : 49 <br/>
        *Description*          : FIXME(dshulyak) SQL big int is max 64 bits. Maybe store as bytes in big endian and hope<br/>
         that lexographical sorting will work.<br/>
    - 📃 downloader.go<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (d *ETHTransferDownloader) GetTransfers(ctx context.Context, header *DBHeader) (rst []Transfer, err error)`<br/>
        *On line*              : 65 <br/>
        *Description*          : TODO(dshulyak) consider caching balance and reset it on reorg<br/>
      - ⬜ Todo:<br/>
        *Related func or type* : `func (d *ETHTransferDownloader) getTransfersInBlock(ctx context.Context, blk *types.Block, accounts []common.Address) (rst []Transfer, err error)`<br/>
        *On line*              : 141 <br/>
        *Description*          : TODO(dshulyak) test that balance difference was covered by transactions<br/>
- 📁 t<br/>
  - 📁 e2e<br/>
    - 📁 api<br/>
      - 📃 api_test.go<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (s *APITestSuite) TestCHTUpdate()`<br/>
          *On line*              : 42 <br/>
          *Description*          : TODO(tiabc): Test that CHT is really updated.<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (s *APITestSuite) TestRaceConditions()`<br/>
          *On line*              : 74 <br/>
          *Description*          : TODO(adam): quarantined until it uses a different datadir<br/>
           as otherwise it wipes out cached blockchain data.<br/>
           func(config *params.NodeConfig) {<br/>
           s.T().Logf("async call to ResetChainData() for network: %d", config.NetworkID)<br/>
          	_, err := s.api.ResetChainData()<br/>
          	progress <- struct{}{}<br/>
           },<br/>
      - 📃 backend_test.go<br/>
        - ⬜ Todo:<br/>
          *On line*              : 24 <br/>
          *Description*          : FIXME(tiabc): There's also a test with the same name in node/manager_test.go<br/>
           so this test should only check StatusBackend logic with a mocked version of the underlying StatusNode.<br/>
        - ⬜ Todo:<br/>
          *On line*              : 83 <br/>
          *Description*          : FIXME(tiabc): There's also a test with the same name in node/manager_test.go<br/>
           so this test should only check StatusBackend logic with a mocked version of the underlying StatusNode.<br/>
    - 📁 node<br/>
      - 📃 manager_test.go<br/>
        - ⬜ Todo:<br/>
          *On line*              : 242 <br/>
          *Description*          : TODO(adam): race conditions should be tested with -race flag and unit tests, if possible.<br/>
        - ⬜ Todo:<br/>
          *On line*              : 243 <br/>
          *Description*          : TODO(boris): going via https://github.com/status-im/status-go/pull/433#issuecomment-342232645 . Testing should be with -race flag<br/>
           Research if it's possible to do the same with unit tests.<br/>
          func (s *ManagerTestSuite) TestRaceConditions() {<br/>
          	cnt := 25<br/>
          	progress := make(chan struct{}, cnt)<br/>
          	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))<br/>
        - ⬜ Todo:<br/>
          *Related func or type* : `func (s *ManagerTestSuite) TestRaceConditions()`<br/>
          *On line*              : 281 <br/>
          *Description*          : TODO(adam): quarantined until it uses a different datadir<br/>
          		// as otherwise it wipes out cached blockchain data.<br/>
          		// func(config *params.NodeConfig) {<br/>
          		//	log.Info("ResetChainData()")<br/>
          		//	_, err := s.StatusNode.ResetChainData()<br/>
          		//	s.T().Logf("ResetChainData(), error: %v", err)<br/>
          		//	progress <- struct{}{}<br/>
          		// },<br/>
          		func(config *params.NodeConfig) {<br/>
          			log.Info("RestartNode()")<br/>
          			_, err := s.StatusNode.RestartNode()<br/>
          			s.T().Logf("RestartNode(), error: %v", err)<br/>
          			progress <- struct{}{}<br/>
          		},<br/>
          		func(config *params.NodeConfig) {<br/>
          			log.Info("Config()")<br/>
          			_, err := s.StatusNode.Config()<br/>
          			s.T().Logf("Config(), error: %v", err)<br/>
          			progress <- struct{}{}<br/>
          		},<br/>
          		func(config *params.NodeConfig) {<br/>
          			log.Info("LightEthereumService()")<br/>
          			_, err := s.StatusNode.LightEthereumService()<br/>
          			s.T().Logf("LightEthereumService(), error: %v", err)<br/>
          			progress <- struct{}{}<br/>
          		},<br/>
          		func(config *params.NodeConfig) {<br/>
          			log.Info("WhisperService()")<br/>
          			_, err := s.StatusNode.WhisperService()<br/>
          			s.T().Logf("WhisperService(), error: %v", err)<br/>
          			progress <- struct{}{}<br/>
          		},<br/>
          		func(config *params.NodeConfig) {<br/>
          			log.Info("AccountManager()")<br/>
          			_, err := s.StatusNode.AccountManager()<br/>
          			s.T().Logf("AccountManager(), error: %v", err)<br/>
          			progress <- struct{}{}<br/>
          		},<br/>
          		func(config *params.NodeConfig) {<br/>
          			log.Info("AccountKeyStore()")<br/>
          			_, err := s.StatusNode.AccountKeyStore()<br/>
          			s.T().Logf("AccountKeyStore(), error: %v", err)<br/>
          			progress <- struct{}{}<br/>
          		},<br/>
          		func(config *params.NodeConfig) {<br/>
          			log.Info("RPCClient()")<br/>
          			s.StatusNode.RPCClient()<br/>
          			progress <- struct{}{}<br/>
          		},<br/>
          	}<br/>
- 📁 waku<br/>
  - 📃 waku.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (m *mockRateLimiterHandler) ExceedIPLimit() error  `<br/>
      *On line*              : 1022 <br/>
      *Description*          : TODO: This does not seem to update the bloom filter, nor topic-interest<br/>
       Note that the filter/topic-interest needs to take into account that there<br/>
       might be filters with duplicated topics, so it's not just a matter of removing<br/>
       from the map, in the topic-interest case, while the bloom filter might need to<br/>
       be rebuilt from scratch<br/>
- 📁 whisper<br/>
  - 📃 whisper.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (whisper *Whisper) runMessageLoop(p *Peer, rw p2p.MsgReadWriter) error`<br/>
      *On line*              : 1166 <br/>
      *Description*          : TODO(adam): should we limit who can send this request?<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (whisper *Whisper) runMessageLoop(p *Peer, rw p2p.MsgReadWriter) error`<br/>
      *On line*              : 1192 <br/>
      *Description*          : TODO(adam): currently, there is no feedback when a sync response<br/>
       is received. An idea to fix this:<br/>
         1. Sending a request contains an ID,<br/>
         2. Each sync response contains this ID,<br/>
         3. There is a way to call whisper.SyncMessages() and wait for the response.Final to be received for that particular request ID.<br/>
         4. If Cursor is not empty, another p2pSyncRequestCode should be sent.<br/>
