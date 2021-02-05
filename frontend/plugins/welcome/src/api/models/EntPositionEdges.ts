/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPositionEdges
 */
export interface EntPositionEdges {
    /**
     * Users holds the value of the users edge.
     * @type {Array<EntUser>}
     * @memberof EntPositionEdges
     */
    users?: Array<EntUser>;
}

export function EntPositionEdgesFromJSON(json: any): EntPositionEdges {
    return EntPositionEdgesFromJSONTyped(json, false);
}

export function EntPositionEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPositionEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'users': !exists(json, 'users') ? undefined : ((json['users'] as Array<any>).map(EntUserFromJSON)),
    };
}

export function EntPositionEdgesToJSON(value?: EntPositionEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'users': value.users === undefined ? undefined : ((value.users as Array<any>).map(EntUserToJSON)),
    };
}


