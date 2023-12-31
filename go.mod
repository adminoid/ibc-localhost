module bitbucket.org/decimalteam/go-smart-node

go 1.21

toolchain go1.21.2

require (
	////bitbucket.org/decimalteam/go-smart-node v0.2.1
	//cosmossdk.io/errors v1.0.0
	//cosmossdk.io/math v1.1.3-rc.1
	////cosmossdk.io/math v1.1.2
	//cosmossdk.io/core v0.6.1
	//cosmossdk.io/api v0.3.1
	//cosmossdk.io/simapp v0.0.0-20230608160436-666c345ad23d
	////cosmossdk.io/tools/rosetta v0.2.1

	cosmossdk.io/api v0.3.1
	cosmossdk.io/errors v1.0.0
	cosmossdk.io/math v1.0.1
	//cosmossdk.io/simapp v0.0.0-20230608160436-666c345ad23d
	//cosmossdk.io/tools/rosetta v0.2.1

	//github.com/coinbase/rosetta-sdk-go v0.10.0 // indirect
	//github.com/cometbft/cometbft v0.38.0 // indirect
	github.com/cometbft/cometbft v0.38.0 // todo -- previous version for evmos update compatibility
	github.com/cometbft/cometbft-db v0.8.0
	//cosmossdk.io/tools/rosetta v0.2.1
	github.com/cosmos/btcutil v1.0.5
	github.com/cosmos/cosmos-proto v1.0.0-beta.3
	// todo -- it was github.com/cosmos/cosmos-sdk v0.46.6
	//github.com/cosmos/cosmos-sdk v0.47.5
	//github.com/cosmos/cosmos-sdk v0.50.0-rc.1
	//cosmossdk.io/simapp v0.0.0-20230608160436-666c345ad23d
	//github.com/cosmos/cosmos-sdk/simapp v0.47.5
	//cosmossdk.io/simapp v0.0.0-20230608160436-666c345ad23d
	github.com/cosmos/go-bip39 v1.0.0

	//github.com/cosmos/ibc-go/modules/capability v1.0.0-rc6
	// todo -- replace from api/cosmos/capability

	// todo -- check it out
	//pgregory.net/rapid v0.5.5 // indirect

	//github.com/cosmos/cosmos-sdk v0.50.0-rc.1
	//github.com/cosmos/cosmos-sdk v0.50.0-rc.1
	//github.com/cosmos/cosmos-sdk/client/grpc/cmtservice v0.50.0-rc.1

	//github.com/cosmos/cosmos-sdk/client/grpc/tmservice v0.50.0-rc.1
	//github.com/cosmos/cosmos-sdk/client/grpc/cmtservice v0.50.0-rc.1

	//cosmossdk.io/simapp v0.0.0-20230608160436-666c345ad23d

	github.com/cosmos/ibc-go/v7 v7.3.0
	github.com/dustin/go-humanize v1.0.1
	//github.com/ethereum/go-ethereum v1.11.5
	//github.com/ethereum/go-ethereum v1.10.26-evmos-rc2
	github.com/go-resty/resty/v2 v2.7.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/joho/godotenv v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.5.1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.17.0
	github.com/status-im/keycard-go v0.2.0
	github.com/stretchr/testify v1.8.4
	go.opencensus.io v0.24.0
	golang.org/x/crypto v0.14.0
	google.golang.org/genproto/googleapis/api v0.0.0-20231012201019-e917dd12ba7a
	google.golang.org/grpc v1.58.3
	google.golang.org/protobuf v1.31.0
	gopkg.in/ini.v1 v1.67.0
)

require (
	//github.com/cosmos/ibc-go/modules/capability v1.0.0-rc6
	//cosmossdk.io/api v0.3.1
	//cosmossdk.io/api v0.7.1 // indirect
	//cosmossdk.io/collections v0.4.0 // indirect
	//cosmossdk.io/core v0.12.0 // indirect
	//cosmossdk.io/depinject v1.0.0-alpha.4 // indirect
	//cosmossdk.io/log v1.2.1 // indirect
	//cosmossdk.io/simapp v0.0.0-20230608160436-666c345ad23d // indirect
	//cosmossdk.io/store v1.0.0-rc.0 // indirect
	//cosmossdk.io/x/tx v0.10.0 // indirect
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.2 // indirect
	github.com/DataDog/zstd v1.5.5 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/Workiva/go-datastructures v1.0.53 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/speakeasy v0.1.1-0.20220910012023-760eaf8b6816 // indirect
	github.com/bits-and-blooms/bitset v1.9.0 // indirect
	//github.com/btcsuite/btcd v0.23.4 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/cockroachdb/apd/v2 v2.0.2 // indirect
	github.com/cockroachdb/errors v1.11.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v0.0.0-20231011191824-ede31f1a8e4b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/confio/ics23/go v0.9.1 // indirect
	github.com/cosmos/cosmos-db v1.0.0 // indirect
	//github.com/cosmos/cosmos-sdk/x/nft v0.1.0-alpha1 // indirect
	//github.com/cosmos/cosmos-db v1.0.0 // indirect
	github.com/cosmos/gogogateway v1.2.0 // indirect
	github.com/cosmos/gogoproto v1.4.11 // indirect
	github.com/cosmos/gorocksdb v1.2.0 // indirect
	github.com/cosmos/iavl v1.0.0-rc.1 // indirect
	github.com/cosmos/ibc-go/v5 v5.1.0 // indirect
	github.com/cosmos/ics23/go v0.10.0 // indirect
	github.com/cosmos/ledger-cosmos-go v0.13.1 // indirect
	github.com/creachadair/taskgroup v0.4.2 // indirect
	github.com/danieljoos/wincred v1.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgraph-io/badger/v2 v2.2007.4 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/emicklei/dot v1.6.0 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/getsentry/sentry-go v0.25.0 // indirect
	github.com/go-kit/kit v0.13.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/golang/glog v1.1.2 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/orderedcode v0.0.1 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-metrics v0.5.1 // indirect
	github.com/hashicorp/go-plugin v1.4.10 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/hdevalence/ed25519consensus v0.1.0 // indirect
	github.com/holiman/uint256 v1.2.3 // indirect
	github.com/huandu/skiplist v1.2.0 // indirect
	github.com/huin/goupnp v1.3.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/improbable-eng/grpc-web v0.15.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/linxGnu/grocksdb v1.8.4 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/manifoldco/promptui v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20210601165009-122bf33a46e0 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/oasisprotocol/curve25519-voi v0.0.0-20230904125328-1f23a7beb09a // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.27.6 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/petermattis/goid v0.0.0-20230904192822-1876fd5063bc // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/rs/cors v1.8.3 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	github.com/sagikazarmark/locafero v0.3.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.10.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/syndtr/goleveldb v1.0.1 // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tendermint/tendermint v0.37.0-rc2 // indirect
	github.com/tendermint/tm-db v0.6.7 // indirect
	github.com/tidwall/btree v1.7.0 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/zondax/hid v0.9.1 // indirect
	github.com/zondax/ledger-go v0.14.2 // indirect
	go.etcd.io/bbolt v1.3.7 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231012201019-e917dd12ba7a // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.0 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
	pgregory.net/rapid v1.1.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	// use cosmos fork of keyring
	github.com/99designs/keyring => github.com/cosmos/keyring v1.2.0
	// use Cosmos-SDK fork to enable Ledger functionality
	//github.com/cosmos/cosmos-sdk => github.com/adminoid/cosmos-sdk v0.47.4-evmos
	// also replace cosmossdk.io
	github.com/adminoid/cosmos-sdk => github.com/adminoid/cosmos-sdk v0.47.4-evmos
	// use Evmos geth fork
	github.com/ethereum/go-ethereum => github.com/evmos/go-ethereum v1.10.26-evmos-rc2
	// replace broken goleveldb
	github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7

	// use Cosmos-SDK fork to enable Ledger functionality
	//github.com/cosmos/cosmos-sdk => github.com/adminoid/cosmos-sdk v0.47.4-evmos
	//github.com/cosmos/cosmos-sdk/simapp => cosmossdk.io/simapp v0.0.0-20230608160436-666c345ad23d
	//github.com/cosmos/cosmos-sdk/store => cosmossdk.io/store v0.50.0-rc.1

	//github.com/cosmos/cosmos-sdk/store => github.com/adminoid/cosmos-sdk/store v0.47.4-evmos

	//github.com/cosmos/cosmos-sdk/store => github.com/adminoid/cosmos-sdk/store v0.47.4-evmos.2
	//github.com/cosmos/cosmos-sdk/store => github.com/adminoid/cosmos-sdk/store v0.47.5-evmos
	//github.com/cosmos/cosmos-sdk/x/evidence => cosmossdk.io/x/evidence
	//cosmossdk.io/x/evidence => github.com/cosmos/cosmos-sdk/x/evidence v0.50.0-rc.1

	//github.com/cosmos/cosmos-sdk/x/evidence => cosmossdk.io/x/evidence v0.47.5
	//github.com/cosmos/cosmos-sdk/x/feegrant => cosmossdk.io/x/feegrant v0.47.5
	//github.com/cosmos/cosmos-sdk/x/upgrade => cosmossdk.io/x/upgrade v0.47.5
	//github.com/cosmos/cosmos-sdk/x/api => cosmossdk.io/x/upgrade v0.47.5
	//

	// use Evmos geth fork
	//github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.10.26-evmos-rc2
	//github.com/gogo/protobuf => github.com/protocolbuffers/protobuf-go v1.13.0
	//github.com/gogo/protobuf => github.com/protocolbuffers/protobuf-go v1.13.0
	//github.com/golang/protobuf => github.com/protocolbuffers/protobuf-go v1.5.2
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

	github.com/tendermint/tendermint => github.com/cometbft/cometbft v0.34.27
//github.com/cosmos/cosmos-sdk => github.com/adminoid/cosmos-sdk v0.47.4-evmos.2
//github.com/cosmos/cosmos-sdk/client/grpc/tmservice => github.com/adminoid/cosmos-sdk/client/grpc/tmservice v0.47.5-evmos.2
//github.com/cosmos/cosmos-sdk/x/distribution/client =>

	//github.com/cosmos/cosmos-sdk/store/types =>

	//cosmossdk.io/api/cosmos/capability => github.com/cosmos/cosmos-sdk/x/capability v1
	//https://github.com/adminoid/cosmos-sdk/tree/main/ --- api/cosmos/capability/module/v1
	//cosmossdk.io/api/cosmos/capability/module/v1 => https://github.com/adminoid/cosmos-sdk/api/cosmos/capability/module/v1
	//cosmossdk.io/api/cosmos/capability => https://github.com/adminoid/cosmos-sdk/api/cosmos/capability
	//github.com/cosmos/cosmos-sdk => github.com/adminoid/cosmos-sdk v0.47.4-evmos
	//github.com/cosmos/cosmos-sdk => github.com/adminoid/cosmos-sdk v0.47.4-evmos.2

	bitbucket.org/decimalteam/go-smart-node => ./
)
