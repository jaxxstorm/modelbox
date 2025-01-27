# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
from . import service_pb2 as service__pb2


class ModelStoreStub(object):
    """*
    ModelStore is the service exposed to upload trained models and training
    checkpoints, and manage metadata around them.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateModel = channel.unary_unary(
                '/modelbox.ModelStore/CreateModel',
                request_serializer=service__pb2.CreateModelRequest.SerializeToString,
                response_deserializer=service__pb2.CreateModelResponse.FromString,
                )
        self.ListModels = channel.unary_unary(
                '/modelbox.ModelStore/ListModels',
                request_serializer=service__pb2.ListModelsRequest.SerializeToString,
                response_deserializer=service__pb2.ListModelsResponse.FromString,
                )
        self.CreateModelVersion = channel.unary_unary(
                '/modelbox.ModelStore/CreateModelVersion',
                request_serializer=service__pb2.CreateModelVersionRequest.SerializeToString,
                response_deserializer=service__pb2.CreateModelVersionResponse.FromString,
                )
        self.ListModelVersions = channel.unary_unary(
                '/modelbox.ModelStore/ListModelVersions',
                request_serializer=service__pb2.ListModelVersionsRequest.SerializeToString,
                response_deserializer=service__pb2.ListModelVersionsResponse.FromString,
                )
        self.CreateExperiment = channel.unary_unary(
                '/modelbox.ModelStore/CreateExperiment',
                request_serializer=service__pb2.CreateExperimentRequest.SerializeToString,
                response_deserializer=service__pb2.CreateExperimentResponse.FromString,
                )
        self.ListExperiments = channel.unary_unary(
                '/modelbox.ModelStore/ListExperiments',
                request_serializer=service__pb2.ListExperimentsRequest.SerializeToString,
                response_deserializer=service__pb2.ListExperimentsResponse.FromString,
                )
        self.CreateCheckpoint = channel.unary_unary(
                '/modelbox.ModelStore/CreateCheckpoint',
                request_serializer=service__pb2.CreateCheckpointRequest.SerializeToString,
                response_deserializer=service__pb2.CreateCheckpointResponse.FromString,
                )
        self.ListCheckpoints = channel.unary_unary(
                '/modelbox.ModelStore/ListCheckpoints',
                request_serializer=service__pb2.ListCheckpointsRequest.SerializeToString,
                response_deserializer=service__pb2.ListCheckpointsResponse.FromString,
                )
        self.GetCheckpoint = channel.unary_unary(
                '/modelbox.ModelStore/GetCheckpoint',
                request_serializer=service__pb2.GetCheckpointRequest.SerializeToString,
                response_deserializer=service__pb2.GetCheckpointResponse.FromString,
                )
        self.UploadBlob = channel.stream_unary(
                '/modelbox.ModelStore/UploadBlob',
                request_serializer=service__pb2.UploadBlobRequest.SerializeToString,
                response_deserializer=service__pb2.UploadBlobResponse.FromString,
                )
        self.DownloadBlob = channel.unary_stream(
                '/modelbox.ModelStore/DownloadBlob',
                request_serializer=service__pb2.DownloadBlobRequest.SerializeToString,
                response_deserializer=service__pb2.DownloadBlobResponse.FromString,
                )


class ModelStoreServicer(object):
    """*
    ModelStore is the service exposed to upload trained models and training
    checkpoints, and manage metadata around them.
    """

    def CreateModel(self, request, context):
        """Create a new Model under a namespace. If no namespace is specified, models
        are created under a default namespace.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListModels(self, request, context):
        """List Models uploaded for a namespace 
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateModelVersion(self, request, context):
        """Creates a new model version for a model
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListModelVersions(self, request, context):
        """Lists model versions for a model.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateExperiment(self, request, context):
        """Creates a new experiment
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListExperiments(self, request, context):
        """List Experiments
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateCheckpoint(self, request, context):
        """Uploads a new checkpoint for an experiment
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListCheckpoints(self, request, context):
        """Lists all the checkpoints for an experiment
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCheckpoint(self, request, context):
        """Gets a checkpoint from the modelstore for an experiment
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UploadBlob(self, request_iterator, context):
        """UploadBlob streams a blob to ModelBox and stores the binaries to the condfigured storage
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DownloadBlob(self, request, context):
        """DownloadBlob downloads a blob from configured storage
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ModelStoreServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateModel': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateModel,
                    request_deserializer=service__pb2.CreateModelRequest.FromString,
                    response_serializer=service__pb2.CreateModelResponse.SerializeToString,
            ),
            'ListModels': grpc.unary_unary_rpc_method_handler(
                    servicer.ListModels,
                    request_deserializer=service__pb2.ListModelsRequest.FromString,
                    response_serializer=service__pb2.ListModelsResponse.SerializeToString,
            ),
            'CreateModelVersion': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateModelVersion,
                    request_deserializer=service__pb2.CreateModelVersionRequest.FromString,
                    response_serializer=service__pb2.CreateModelVersionResponse.SerializeToString,
            ),
            'ListModelVersions': grpc.unary_unary_rpc_method_handler(
                    servicer.ListModelVersions,
                    request_deserializer=service__pb2.ListModelVersionsRequest.FromString,
                    response_serializer=service__pb2.ListModelVersionsResponse.SerializeToString,
            ),
            'CreateExperiment': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateExperiment,
                    request_deserializer=service__pb2.CreateExperimentRequest.FromString,
                    response_serializer=service__pb2.CreateExperimentResponse.SerializeToString,
            ),
            'ListExperiments': grpc.unary_unary_rpc_method_handler(
                    servicer.ListExperiments,
                    request_deserializer=service__pb2.ListExperimentsRequest.FromString,
                    response_serializer=service__pb2.ListExperimentsResponse.SerializeToString,
            ),
            'CreateCheckpoint': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateCheckpoint,
                    request_deserializer=service__pb2.CreateCheckpointRequest.FromString,
                    response_serializer=service__pb2.CreateCheckpointResponse.SerializeToString,
            ),
            'ListCheckpoints': grpc.unary_unary_rpc_method_handler(
                    servicer.ListCheckpoints,
                    request_deserializer=service__pb2.ListCheckpointsRequest.FromString,
                    response_serializer=service__pb2.ListCheckpointsResponse.SerializeToString,
            ),
            'GetCheckpoint': grpc.unary_unary_rpc_method_handler(
                    servicer.GetCheckpoint,
                    request_deserializer=service__pb2.GetCheckpointRequest.FromString,
                    response_serializer=service__pb2.GetCheckpointResponse.SerializeToString,
            ),
            'UploadBlob': grpc.stream_unary_rpc_method_handler(
                    servicer.UploadBlob,
                    request_deserializer=service__pb2.UploadBlobRequest.FromString,
                    response_serializer=service__pb2.UploadBlobResponse.SerializeToString,
            ),
            'DownloadBlob': grpc.unary_stream_rpc_method_handler(
                    servicer.DownloadBlob,
                    request_deserializer=service__pb2.DownloadBlobRequest.FromString,
                    response_serializer=service__pb2.DownloadBlobResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'modelbox.ModelStore', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ModelStore(object):
    """*
    ModelStore is the service exposed to upload trained models and training
    checkpoints, and manage metadata around them.
    """

    @staticmethod
    def CreateModel(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/modelbox.ModelStore/CreateModel',
            service__pb2.CreateModelRequest.SerializeToString,
            service__pb2.CreateModelResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListModels(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/modelbox.ModelStore/ListModels',
            service__pb2.ListModelsRequest.SerializeToString,
            service__pb2.ListModelsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateModelVersion(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/modelbox.ModelStore/CreateModelVersion',
            service__pb2.CreateModelVersionRequest.SerializeToString,
            service__pb2.CreateModelVersionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListModelVersions(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/modelbox.ModelStore/ListModelVersions',
            service__pb2.ListModelVersionsRequest.SerializeToString,
            service__pb2.ListModelVersionsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateExperiment(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/modelbox.ModelStore/CreateExperiment',
            service__pb2.CreateExperimentRequest.SerializeToString,
            service__pb2.CreateExperimentResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListExperiments(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/modelbox.ModelStore/ListExperiments',
            service__pb2.ListExperimentsRequest.SerializeToString,
            service__pb2.ListExperimentsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateCheckpoint(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/modelbox.ModelStore/CreateCheckpoint',
            service__pb2.CreateCheckpointRequest.SerializeToString,
            service__pb2.CreateCheckpointResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListCheckpoints(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/modelbox.ModelStore/ListCheckpoints',
            service__pb2.ListCheckpointsRequest.SerializeToString,
            service__pb2.ListCheckpointsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCheckpoint(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/modelbox.ModelStore/GetCheckpoint',
            service__pb2.GetCheckpointRequest.SerializeToString,
            service__pb2.GetCheckpointResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UploadBlob(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/modelbox.ModelStore/UploadBlob',
            service__pb2.UploadBlobRequest.SerializeToString,
            service__pb2.UploadBlobResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DownloadBlob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/modelbox.ModelStore/DownloadBlob',
            service__pb2.DownloadBlobRequest.SerializeToString,
            service__pb2.DownloadBlobResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
