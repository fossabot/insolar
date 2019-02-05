local import_params = import '../params.libsonnet';
local params = import_params.components.insolar;

local genesis = import 'genesis.libsonnet' ;
local statefull_set() = import 'insolard_statefull_set.libsonnet';
local genesis_insolard_conf() = import "insolard-genesis.libsonnet";
local insolard_conf() = import "insolard.libsonnet";
local k = import "k.libsonnet";
local pulsar() = import 'pulsar/pulsar_common.libsonnet';

local perisitant_claim() = {
  kind: "PersistentVolumeClaim",
  apiVersion: "v1",
  metadata: {
    name: "bootstrap-config",
    labels: {
      app: "bootstrap"
    }
  },
  spec: {
    accessModes: [
      "ReadWriteMany"
    ],
    resources: {
      requests: {
        storage: "2Gi"
      }
    }
  }
};

local service() = {
  apiVersion: "v1",
  kind: "Service",
  metadata: {
    name: "bootstrap",
    labels: {
      app: "bootstrap"
    }
  },
  spec: {
    ports: [
      {
        port: 8080,
        name: "prometheus"
      },
      {
        port: params.tcp_transport_port,
        name: "network",
        protocol: "TCP"
      },
      {
        port: 19191,
        name: "api",
        protocol: "TCP"
      }
    ],
    clusterIP: "None",
    selector: {
      app: "bootstrap"
    }
  }
};

local configs() = {
  apiVersion: "v1",
  kind: "ConfigMap",
  metadata: {
    name: "seed-config"
  },
  data:{
            "genesis.yaml": std.manifestYamlDoc(genesis.generate_genesis()),
            "insolar-genesis.yaml": std.manifestYamlDoc(genesis_insolard_conf()),
            "insolar.yaml": std.manifestYamlDoc(insolard_conf()),
    }

};

k.core.v1.list.new([configs(), service(), perisitant_claim(), statefull_set()])

